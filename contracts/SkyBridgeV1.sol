// SPDX-License-Identifier: AGPL-3.0
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IParams.sol";
import "./interfaces/IERC20.sol";
import "./libraries/SafeERC20.sol";

contract SkyBridgeV1 is Ownable {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    struct spPendingTx {
        bytes32 SwapID; //swap hash for identification of this swap.
        string DestAddr; //destination BTC address for the swap
        address RefundAddr; //refund address on evm source chain for if the swap fails.
        uint256 AmountWBTC; //outbound amount for this swap.
        uint256 Timestamp; // block timestamp that is set by EVM
    }

    mapping(address => bool) public whitelist;
    mapping(address => uint256) private floatAmountOf;
    mapping(bytes32 => bool) private used; //used TX
    address public BTCT_ADDR;
    IBurnableToken public lpToken;
    uint256 private lpDivisor;
    address public sbBTCPool;
    IParams public ip;
    uint256 private convertScale;
    uint256 public limitBTCForSPFlow2;
    address public skyPool;
    uint256 public swapCount;
    //skypools - token balance - call using tokens[token address][user address] to get uint256 balance - see function balanceOf
    mapping(address => mapping(address => uint256)) public tokens;
    //Data and indexes for pending swap objects
    mapping(uint256 => spPendingTx) public spPendingTXs; //index => pending TX object
    uint256 public oldestActiveIndex;
    /** TSS */
    // Node lists state { 0 => not exist, 1 => exist, 2 => removed }
    mapping(address => uint8) private nodes;
    address[] private nodeAddrs;
    uint8 public activeNodeCount;
    uint8 public churnedInCount;
    uint8 public tssThreshold;
    bool isInitialized;

    // intialize
    function initialize(
        address _btct,
        address _lpToken,
        address _sbBTCPool,
        address _params,
        uint256 _existingBTCFloat
    ) public onlyOwner {
      require(!isInitialized, "You already initialized!");
        BTCT_ADDR = _btct;
        lpToken = IBurnableToken(_lpToken);
        // Set initial lpDivisor of LP token
        lpDivisor = 10**IERC20(_lpToken).decimals();
        sbBTCPool = _sbBTCPool;
        //set IParams
        ip = IParams(_params);
        // Set convertScale
        convertScale = 10**(IERC20(_btct).decimals() - 8);
        whitelist[_btct] = true;
        whitelist[_lpToken] = true;
        whitelist[address(0)] = true;
        floatAmountOf[address(0)] = _existingBTCFloat;
    }

    // modifers

    modifier onlyPool() {
        require(msg.sender == skyPool, "is not Pool!");
        _;
    }

    // ownable set functions owner
    function setWhiteList(address _token, bool _add) external onlyOwner {
        whitelist[_token] = _add;
    }

    function setPool(address _skyPool) public onlyOwner {
        skyPool = _skyPool;
    }

    // set functions for skyPool

    function addFloat(address _token, uint256 _amount) external onlyPool {
        floatAmountOf[_token] = floatAmountOf[_token].add(_amount);
    }

    function removeFloat(address _token, uint256 _amount) external onlyPool {
        floatAmountOf[_token] = floatAmountOf[_token].sub(
            _amount,
            "10" //"_removeFloat: float amount insufficient"
        );
    }

    /// @dev _addUsedTx updates txid list which is spent. (single hash)
    /// @param _txid The array of txid.
    function addUsedTx(bytes32 _txid) external onlyPool {
        used[_txid] = true;
    }

    /// @dev _addUsedTxs updates txid list which is spent. (multiple hashes)
    /// @param _txids The array of txid.
    function addUsedTxs(bytes32[] memory _txids) external onlyPool {
        for (uint256 i = 0; i < _txids.length; i++) {
            used[_txids[i]] = true;
        }
    }

    function addTokenBalance(
        address _token,
        address _account,
        uint256 _amount
    ) external onlyPool {
        tokens[_token][_account] = tokens[_token][_account].add(_amount);
    }

    function removeTokenBalance(
        address _token,
        address _account,
        uint256 _amount
    ) external onlyPool {
        tokens[_token][_account] = tokens[_token][_account].sub(_amount);
    }

    function addMultiTokenBalance(
        address _token,
        bytes32[] memory _addressesAndAmounts
    ) external onlyPool {
        for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
            tokens[_token][
                address(uint160(uint256(_addressesAndAmounts[i])))
            ] = tokens[_token][
                address(uint160(uint256(_addressesAndAmounts[i])))
            ].add(uint256(uint96(bytes12(_addressesAndAmounts[i]))));
        }
    }

    function checkTokenBalance(
        address _token,
        address _account,
        uint256 _amount
    ) external view onlyPool returns (bool) {
        return tokens[_token][_account] >= _amount;
    }

    /// @dev updateLimitBTCForSPFlow2 udpates limitBTCForSPFlow2
    function updateLimitBTCForSPFlow2() external onlyPool {
        // Update limitBTCForSPFlow2 by adding BTC floats
        _updateLimitBTCForSPFlow2();
    }

    /// @dev _updateLimitBTCForSPFlow2 udpates limitBTCForSPFlow2
    function _updateLimitBTCForSPFlow2() internal {
        // Update limitBTCForSPFlow2 by adding BTC floats
        limitBTCForSPFlow2 = floatAmountOf[address(0)];
    }

    /// @dev _reduceLimitBTCForSPFlow2 reduces limitBTCForSPFlow2 when new sp flow2 txs are coming.
    /// @param _amount The amount of BTCT, (use BTCT amount insatead of BTC amount for enough. always BTCT > BTC)
    function _reduceLimitBTCForSPFlow2(uint256 _amount) internal {
        if (limitBTCForSPFlow2 == 0) {
            // initialize when initial Flow2 tx has been called.
            _updateLimitBTCForSPFlow2();
        }
        limitBTCForSPFlow2 = limitBTCForSPFlow2.sub(
            _amount,
            "12" //"BTC float amount insufficient"
        );
    }

    /// @dev _swap collects swap amount to change float.
    /// @param _sourceToken The address of source token
    /// @param _destToken The address of target token.
    /// @param _swapAmount The amount of swap.
    function swap(
        address _sourceToken,
        address _destToken,
        uint256 _swapAmount
    ) external onlyPool {
        floatAmountOf[_destToken] = floatAmountOf[_destToken].sub(
            _swapAmount,
            "11" //"_swap: float amount insufficient"
        );
        floatAmountOf[_sourceToken] = floatAmountOf[_sourceToken].add(
            _swapAmount
        );
    }

    /// @dev _safeTransfer executes tranfer erc20 tokens
    /// @param _token The address of target token
    /// @param _to The address of receiver.
    /// @param _amount The amount of transfer.
    function safeTransfer(
        address _token,
        address _to,
        uint256 _amount
    ) external onlyPool {
        if (_token == BTCT_ADDR) {
            _amount = _amount.mul(convertScale);
        }
        IERC20(_token).safeTransfer(_to, _amount);
    }

    /// @dev _rewardsCollection collects tx rewards.
    /// @param _feesToken The token address for collection fees.
    /// @param _rewardsAmount The amount of rewards.
    function rewardsCollection(address _feesToken, uint256 _rewardsAmount)
        external
        onlyPool
    {
        if (_rewardsAmount == 0) return;
        // if (_feesToken == lpToken) {
        //     IBurnableToken(lpToken).transfer(sbBTCPool, _rewardsAmount);
        //     emit RewardsCollection(_feesToken, 0, _rewardsAmount, 0);
        //     return;
        // }

        // Get current LP token price.
        uint256 nowPrice = getCurrentPriceLP();
        // Add all fees into pool
        floatAmountOf[_feesToken] = floatAmountOf[_feesToken].add(
            _rewardsAmount
        );
        uint256 amountForNodes = _rewardsAmount.mul(ip.nodeRewardsRatio()).div(
            100
        );
        // Alloc LP tokens for nodes as fees
        uint256 amountLPTokensForNode = amountForNodes.mul(lpDivisor).div(
            nowPrice
        );
        // Mints LP tokens for Nodes
        lpToken.mint(sbBTCPool, amountLPTokensForNode);
    }

    // get functions for sky pool
    function isWhiteListed(address _token) external view returns (bool) {
        return whitelist[_token];
    }

    function priceCheck() external view onlyPool returns (bool) {
        uint256 beforePrice = getCurrentPriceLP();
        return getCurrentPriceLP() >= beforePrice;
    }

    /// @dev isTxUsed sends rewards for Nodes.
    /// @param _txid The txid which is for recording.
    function isTxUsed(bytes32 _txid) external view returns (bool) {
        return used[_txid];
    }

    function getLpDivisor() external view returns (uint256) {
        return lpDivisor;
    }

    // public get functions
    /// @dev getCurrentPriceLP returns the current exchange rate of LP token.
    function getCurrentPriceLP() public view returns (uint256 nowPrice) {
        (uint256 reserveA, uint256 reserveB) = getFloatReserve(
            address(0),
            BTCT_ADDR
        );
        uint256 totalLPs = lpToken.totalSupply();
        // decimals of totalReserved == 8, lpDivisor == 8, decimals of rate == 8
        nowPrice = totalLPs == 0
            ? lpDivisor
            : (reserveA.add(reserveB)).mul(lpDivisor).div(totalLPs);
        return nowPrice;
    }

    /// @dev getFloatReserve returns float reserves
    /// @param _tokenA The address of target tokenA.
    /// @param _tokenB The address of target tokenB.
    function getFloatReserve(address _tokenA, address _tokenB)
        public
        view
        returns (uint256 reserveA, uint256 reserveB)
    {
        (reserveA, reserveB) = (floatAmountOf[_tokenA], floatAmountOf[_tokenB]);
    }

    /// @dev _spRecordPendingTx - hash a unique swap ID, and add it to the array of pending TXs, and then emit event
    /// @param _destinationAddressForBTC The BTC address to send BTC to.
    /// @param _btctAmount amount in BTC decimal 8.
    function spRecordPendingTx(
        string calldata _destinationAddressForBTC,
        uint256 _btctAmount
    ) external onlyPool {
        //hash TX data for unique ID
        bytes32 ID = keccak256(
            abi.encodePacked(
                BTCT_ADDR, //specific to current chain
                swapCount,
                _destinationAddressForBTC,
                _btctAmount,
                block.timestamp
            )
        );

        spPendingTXs[swapCount] = spPendingTx(
            ID,
            _destinationAddressForBTC,
            msg.sender,
            _btctAmount,
            block.timestamp
        );

        //clean up expired TXs
        spCleanUpOldTXs();

        swapCount = swapCount.add(1); //increment TX count after cleaning up pending TXs to not loop over next empty index

        _reduceLimitBTCForSPFlow2(_btctAmount);
    }

    /// @dev _spCleanUpOldTXs - call when executing flow 2 swaps, cleans up expired TXs and moves the indices
    function spCleanUpOldTXs() public {
        uint256 max = oldestActiveIndex.add(ip.loopCount());

        if (max >= swapCount) {
            max = swapCount;
        }

        uint256 current = block.timestamp;
        for (uint256 i = oldestActiveIndex; i < max; i++) {
            if (spPendingTXs[i].Timestamp.add(ip.expirationTime()) < current) {
                delete spPendingTXs[i];
                oldestActiveIndex = i.add(1);
            }
        }
    }

    /// @dev churn transfers contract ownership and set variables of the next TSS validator set.
    /// @param _newOwner The address of new Owner.
    /// @param _nodes The reward addresses.
    /// @param _isRemoved The flags to remove node.
    /// @param _churnedInCount The number of next party size of TSS group.
    /// @param _tssThreshold The number of next threshold.
    function churn(
        address _newOwner,
        address[] memory _nodes,
        bool[] memory _isRemoved,
        uint8 _churnedInCount,
        uint8 _tssThreshold
    ) external onlyOwner returns (bool) {
        require(
            _tssThreshold >= tssThreshold && _tssThreshold <= 2**8 - 1,
            "01" //"_tssThreshold should be >= tssThreshold"
        );
        require(
            _churnedInCount >= _tssThreshold + uint8(1),
            "02" //"n should be >= t+1"
        );
        require(
            _nodes.length == _isRemoved.length,
            "05" //"_nodes and _isRemoved length is not match"
        );

        transferOwnership(_newOwner);
        // Update active node list
        for (uint256 i = 0; i < _nodes.length; i++) {
            if (!_isRemoved[i]) {
                if (nodes[_nodes[i]] == uint8(0)) {
                    nodeAddrs.push(_nodes[i]);
                }
                if (nodes[_nodes[i]] != uint8(1)) {
                    activeNodeCount++;
                }
                nodes[_nodes[i]] = uint8(1);
            } else {
                activeNodeCount--;
                nodes[_nodes[i]] = uint8(2);
            }
        }
        require(activeNodeCount <= 100, "Stored node size should be <= 100");
        churnedInCount = _churnedInCount;
        tssThreshold = _tssThreshold;
        return true;
    }

    /// @dev balanceOf - return user balance for given token and user for skypools
    /// @param _token The address of target token.
    /// @param _user The address of target user.
    function balanceOf(address _token, address _user)
        public
        view
        returns (uint256)
    {
        return tokens[_token][_user];
    }

    /// @dev spGetPendingSwaps - returns array of pending swaps
    /// @return data - returns array of pending swap struct objects
    function spGetPendingSwaps()
        external
        view
        returns (spPendingTx[] memory data)
    {
        //require(swapCount != 0);

        uint256 index = 0;
        data = new spPendingTx[](swapCount.sub(oldestActiveIndex));

        for (uint256 i = oldestActiveIndex.add(1); i <= swapCount; i++) {
            data[index] = spPendingTXs[index.add(oldestActiveIndex)];
            index = index.add(1);
        }

        return data;
    }

    /// @dev getActiveNodes returns active nodes list
    function getActiveNodes() public view returns (address[] memory) {
        uint256 count = 0;
        address[] memory _nodes = new address[](activeNodeCount);
        for (uint256 i = 0; i < nodeAddrs.length; i++) {
            if (nodes[nodeAddrs[i]] == uint8(1)) {
                _nodes[count] = nodeAddrs[i];
                count++;
            }
        }
        return _nodes;
    }

    /// @dev isNodeSake returns true if the node is churned in
    function isNodeStake(address _user) public view returns (bool) {
        if (nodes[_user] == uint8(1)) {
            return true;
        }
        return false;
    }
}
