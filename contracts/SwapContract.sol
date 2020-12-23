// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <0.8.0;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/ISwapContract.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract SwapContract is Ownable, ISwapContract {
    using SafeMath for uint256;

    address public WBTC_ADDR;
    address public lpToken;

    uint8 public churnedInCount;
    uint8 public nodeRewardsRatio;
    uint8 public depositFeesBPS;

    uint256 public activeWBTCBalances;
    uint256 public lockedLPTokensForNode;

    uint256 private priceDecimals;
    uint256 private currentExchangeRate;
    uint256 private lpDecimals;
    uint256 private nodeSize;
    uint256 private nodeRemoved;
    // Support tokens
    mapping(address => bool) whitelist;

    // Nodes
    mapping(uint256 => bytes32) private nodes;
    mapping(address => uint256) private nodeIndex;

    // Token address -> amount
    mapping(address => uint256) private totalRewards;
    mapping(address => uint256) private floatAmountOf;
    mapping(bytes32 => bool) private used;

    /**
     * Events
     */

    event RecordIncomingFloat(
        address token,
        bytes32 addressesAndAmountOfFloat,
        bytes32 txid
    );

    event IssueLPTokensForFloat(address to, uint256 amountOfLP, bytes32 txid);

    event RecordOutcomingFloat(
        address token,
        bytes32 addressesAndAmountOfLPtoken,
        bytes32 txid
    );

    event BurnLPTokensForFloat(
        address token,
        uint256 amountOfFloat,
        bytes32 txid
    );

    constructor(address _lpToken, address _wbtc) public {
        // burner = new Burner();
        lpToken = _lpToken;
        // Set initial price of LP token per BTC/WBTC.
        lpDecimals = 10**IERC20(lpToken).decimals();
        // Set WBTC address
        WBTC_ADDR = _wbtc;
        // Set nodeRewardsRatio
        nodeRewardsRatio = 66;
        // Set depositFeesBPS
        depositFeesBPS = 50;
        // Set priceDecimals
        priceDecimals = 10**8;
        // Set currentExchangeRate
        currentExchangeRate = priceDecimals;
        // Set lockedLPTokensForNode
        lockedLPTokensForNode = 0;
        // SEt whitelist
        whitelist[WBTC_ADDR] = true;
        whitelist[lpToken] = true;
        whitelist[address(0)] = true;
    }

    /**
     * Transfer part
     */

    /// @dev singleTransferERC20 function sends tokens from contract.
    /// @param _token Address of token.
    /// @param _to Recevier address.
    /// @param _amount The amount of tokens.
    /// @param _totalSwapped the amount of swapped amount which is for send.
    /// @param _rewardsAmount Value that should be paid as fees.
    /// @param _redeemedFloatTxIds the txs which is for records txids.
    function singleTransferERC20(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _totalSwapped,
        uint256 _rewardsAmount,
        bytes32[] memory _redeemedFloatTxIds
    ) external override onlyOwner returns (bool) {
        require(whitelist[_token], "token is not whitelisted");
        if (_token == WBTC_ADDR && _totalSwapped > 0) {
            activeWBTCBalances = activeWBTCBalances.sub(
                _totalSwapped,
                "activeWBTCBalances insufficient"
            );
        }
        _rewardsCollection(_token, _rewardsAmount);
        _addTxidUsed(_redeemedFloatTxIds);
        require(IERC20(_token).transfer(_to, _amount));
        return true;
    }

    function multiTransferERC20TightlyPacked(
        address _token,
        bytes32[] memory _addressesAndAmounts,
        uint256 _totalSwapped,
        uint256 _rewardsAmount,
        bytes32[] memory _redeemedFloatTxIds
    ) external override onlyOwner returns (bool) {
        require(whitelist[_token], "token is not whitelisted");
        if (_token == WBTC_ADDR && _totalSwapped > 0) {
            activeWBTCBalances = activeWBTCBalances.sub(
                _totalSwapped,
                "activeWBTCBalances insufficient"
            );
        }
        _rewardsCollection(_token, _rewardsAmount);
        _addTxidUsed(_redeemedFloatTxIds);
        for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
            require(
                IERC20(_token).transfer(
                    address(uint160(uint256(_addressesAndAmounts[i]))),
                    uint256(uint96(bytes12(_addressesAndAmounts[i])))
                ),
                "Batch transfer error"
            );
        }
        return true;
    }

    function multiTransferERC20(
        address _token,
        address[] memory _contributors,
        uint256[] memory _amounts,
        uint256 _totalSwapped,
        uint256 _rewardsAmount,
        bytes32[] memory _redeemedFloatTxIds
    ) external override onlyOwner returns (bool) {
        require(
            _contributors.length == _amounts.length,
            "Length of inputs array is mismatch"
        );
        if (_token == WBTC_ADDR && _totalSwapped > 0) {
            activeWBTCBalances = activeWBTCBalances.sub(
                _totalSwapped,
                "activeWBTCBalances insufficient"
            );
        }
        _rewardsCollection(_token, _rewardsAmount);
        _addTxidUsed(_redeemedFloatTxIds);
        for (uint256 i = 0; i < _contributors.length; i++) {
            require(IERC20(_token).transfer(_contributors[i], _amounts[i]));
        }
        return true;
    }

    /**
     * @dev gas usage 90736 gas (initial), 58888 gas (update)
     */
    function collectSwapFeesForBTC(
        address _feeToken,
        uint256 _incomingAmount,
        uint256 _rewardsAmount
    ) external override onlyOwner returns (bool) {
        //require(!used[_txid], "txid is already used");
        activeWBTCBalances = activeWBTCBalances.add(_incomingAmount);
        // _feeToken should be address(0) == BTC
        _rewardsCollection(_feeToken, _rewardsAmount);
        // Add txid to used list.
        //used[_txid] = true;
        return true;
    }

    /**
     * Float part
     */

    /**
     * @dev gas usage 44910 gas
     */
    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfFloat,
        bytes32 _txid
    ) external override onlyOwner returns (bool) {
        require(whitelist[_token], "_token is invalid");
        // txs[_token][_txid] = _addressesAndAmountOfFloat;
        // emit RecordIncomingFloat(_token, _addressesAndAmountOfFloat, _txid);
        require(
            _issueLPTokensForFloat(_token, _addressesAndAmountOfFloat, _txid)
        );
        return true;
    }

    /**
     * @dev gas uasge 43628 gas
     */

    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfLPtoken,
        bytes32 _txid
    ) external override onlyOwner returns (bool) {
        require(whitelist[_token], "_token is invalid");
        // _token should be address(0) or WBTC_ADDR, txid should be unique
        // txs[_token][_txid] = _addressesAndAmountOfLPtoken;
        // emit RecordOutcomingFloat(_token, _addressesAndAmountOfLPtoken, _txid);
        require(
            _burnLPTokensForFloat(_token, _addressesAndAmountOfLPtoken, _txid)
        );
        return true;
    }

    /**
     * @dev gas usage 3129064 gas for 100 nodes
     */

    function distributeNodeRewards() external override returns (bool) {
        // Reduce Gas
        uint256 rewardLPsForNodes = lockedLPTokensForNode;
        require(rewardLPsForNodes > 0, "totalRewardLPsForNode is not positive");
        bytes32[] memory nodeList = _loadNodes();
        uint256 totalStaked = 0;
        for (uint256 i = 0; i < nodeList.length; i++) {
            totalStaked = totalStaked.add(uint256(uint96(bytes12(nodes[i]))));
        }
        for (uint256 i = 0; i < nodeList.length; i++) {
            IBurnableToken(lpToken).mint(
                address(uint160(uint256(nodeList[i]))),
                rewardLPsForNodes
                    .mul(uint256(uint96(bytes12(nodeList[i]))))
                    .div(totalStaked)
            );
        }
        lockedLPTokensForNode = 0;
        return true;
    }

    /**
     * @dev gas usage 4599585 gas (initial cost), 733763 gas (update cost) for 100 nodes
     */

    function churn(
        address _newOwner,
        bytes32[] memory _rewardAddressAndAmounts,
        bool[] memory _isRemoved,
        uint8 _churnedInCount,
        uint8 _nodeRewardsRatio
    ) external override onlyOwner returns (bool) {
        transferOwnership(_newOwner);
        // Update active node list
        for (uint256 i = 0; i < _rewardAddressAndAmounts.length; i++) {
            (address newNode, ) = _splitToValues(_rewardAddressAndAmounts[i]);
            _addNode(newNode, _rewardAddressAndAmounts[i], _isRemoved[i]);
        }
        churnedInCount = _churnedInCount;
        // The ratio should be 100x of actual rate.
        nodeRewardsRatio = _nodeRewardsRatio;
        return true;
    }

    function isTxUsed(bytes32 _txid) public override view returns (bool) {
        return used[_txid];
    }

    function getCurrentPriceLP() public override view returns (uint256) {
        return currentExchangeRate;
    }

    function getDepositFeeRate(address _token, uint256 _amountOfFloat)
        public
        override
        view
        returns (uint256 depositFeeRate)
    {
        uint8 isFlip = _checkFlips(_token, _amountOfFloat);
        if (isFlip == 1) {
            depositFeeRate = _token == WBTC_ADDR ? depositFeesBPS : 0;
        } else if (isFlip == 2) {
            depositFeeRate = _token == address(0) ? depositFeesBPS : 0;
        }
    }

    /**
     * @dev returns float amounts not balances.
     */
    function getFloatReserve(address _tokenA, address _tokenB)
        public
        override
        view
        returns (uint256 reserveA, uint256 reserveB)
    {
        (reserveA, reserveB) = (
            floatAmountOf[_tokenA].add(totalRewards[_tokenA]),
            floatAmountOf[_tokenB].add(totalRewards[_tokenB])
        );
    }

    /**
     * @dev gas usage 183033 gas
     */
    function _issueLPTokensForFloat(
        address token,
        bytes32 transaction,
        bytes32 _txid
    ) internal returns (bool) {
        require(!isTxUsed(_txid), "The txid is already used");
        // (address token, bytes32 transaction) = _loadTx(_txid);
        require(transaction != 0x0, "The transaction is not found");
        // Define target address which is recorded bottom 20bytes on tx data
        // Define amountLP which is recorded top 12bytes on tx data
        (address to, uint256 amountOfFloat) = _splitToValues(transaction);
        // LP token price per BTC/WBTC changed
        uint256 nowPrice = _updateFloatPool(address(0), WBTC_ADDR);
        // Calculate amount of LP token
        uint256 amountOfLP = amountOfFloat.mul(priceDecimals).div(nowPrice);
        uint256 depositFeeRate = getDepositFeeRate(token, amountOfFloat);
        uint256 depositFees = depositFeeRate != 0
            ? amountOfLP.mul(depositFeeRate).div(10000)
            : 0;

        //Send LP tokens to LP
        IBurnableToken(lpToken).mint(to, amountOfLP.sub(depositFees));
        // Add deposit fees
        lockedLPTokensForNode = lockedLPTokensForNode.add(depositFees);
        // Add float amount
        _addFloat(token, amountOfFloat);
        used[_txid] = true;
        emit IssueLPTokensForFloat(to, amountOfLP, _txid);
        return true;
    }

    /**
     * @dev gas uasge 63677 gas
     */
    function _burnLPTokensForFloat(
        address token,
        bytes32 transaction,
        bytes32 _txid
    ) internal returns (bool) {
        require(!isTxUsed(_txid), "The txid is already used");
        // _token should be address(0) or WBTC_ADDR
        // (address token, bytes32 transaction) = _loadTx(_txid);
        require(transaction != 0x0, "The transaction is not found");
        // Define target address which is recorded bottom 20bytes on tx data
        // Define amountLP which is recorded top 12bytes on tx data
        (address to, uint256 amountOfLP) = _splitToValues(transaction);
        // Calculate amountOfLP
        uint256 nowPrice = _updateFloatPool(address(0), WBTC_ADDR);
        // Calculate amountOfFloat
        uint256 amountOfFloat = amountOfLP.mul(nowPrice).div(priceDecimals);

        require(
            floatAmountOf[token] >= amountOfFloat,
            "Pool balance insufficient."
        );
        // Burn LP tokens
        require(IBurnableToken(lpToken).burn(amountOfLP));
        // Remove float amount
        _removeFloat(token, amountOfFloat);
        used[_txid] = true;
        // WBTC transfer if token address is WBTC_ADDR
        if (token == WBTC_ADDR) {
            require(
                IERC20(token).transfer(to, amountOfFloat),
                "WBTC balance insufficient"
            );
        }
        emit BurnLPTokensForFloat(to, amountOfFloat, _txid);
        return true;
    }

    function _checkFlips(address _token, uint256 _amountOfFloat)
        internal
        view
        returns (uint8 active)
    {
        (uint256 reserveA, uint256 reserveB) = getFloatReserve(
            address(0),
            WBTC_ADDR
        );
        if (_token == address(0)) {
            reserveA = reserveA.add(_amountOfFloat);
        } else if (_token == WBTC_ADDR) {
            reserveB = reserveB.add(_amountOfFloat);
        }
        if (activeWBTCBalances > reserveA.add(reserveB)) {
            return active;
        } else if (reserveA.add(reserveB) >= activeWBTCBalances) {
            // BTC balance == balance of BTC float + balance of WBTC float - balance of WBTC
            uint256 balBTC = reserveA.add(reserveB).sub(activeWBTCBalances);
            if (balBTC <= reserveA.add(reserveB).div(3)) {
                active = 1; // BTC float insufficient
            } else if (activeWBTCBalances <= reserveA.add(reserveB).div(3)) {
                active = 2; // WBTC float insufficient
            } else {
                active = 0; // zero fees
            }
        }
        return active;
    }

    function _updateFloatPool(address _tokenA, address _tokenB)
        internal
        returns (uint256)
    {
        // Reduce gas cost.
        (uint256 reserveA, uint256 reserveB) = getFloatReserve(
            _tokenA,
            _tokenB
        );
        uint256 totalLPs = IBurnableToken(lpToken).totalSupply();
        // decimals of totalReserved == 8, lpDecimals == 8, decimals of rate == 8
        currentExchangeRate = totalLPs == 0
            ? currentExchangeRate
            : (reserveA.add(reserveB)).mul(lpDecimals).div(
                totalLPs.add(lockedLPTokensForNode)
            );
        return currentExchangeRate;
    }

    function _addFloat(address _token, uint256 _amount) internal {
        floatAmountOf[_token] = floatAmountOf[_token].add(_amount);
        if (_token == WBTC_ADDR) {
            activeWBTCBalances = activeWBTCBalances.add(_amount);
        }
    }

    function _removeFloat(address _token, uint256 _amount) internal {
        floatAmountOf[_token] = floatAmountOf[_token].sub(
            _amount,
            "float amount insufficient"
        );
        if (_token == WBTC_ADDR) {
            activeWBTCBalances = activeWBTCBalances.sub(
                _amount,
                "activeWBTCBalances insufficient"
            );
        }
    }

    function _rewardsCollection(address _token, uint256 _rewardsAmount)
        internal
    {
        // Add all fees into pool
        totalRewards[_token] = totalRewards[_token].add(_rewardsAmount);
        uint256 amountForLP = _rewardsAmount.mul(nodeRewardsRatio).div(100);
        // Alloc LP tokens for nodes as fees
        uint256 amountLPForNode = _rewardsAmount
            .sub(amountForLP)
            .mul(priceDecimals)
            .div(getCurrentPriceLP());
        // Add minted LP tokens for Nodes
        lockedLPTokensForNode = lockedLPTokensForNode.add(amountLPForNode);
    }

    function _addTxidUsed(bytes32[] memory _txs) internal {
        for (uint256 i = 0; i < _txs.length; i++) {
            used[_txs[i]] = true;
        }
    }

    function _loadNodes() internal view returns (bytes32[] memory) {
        uint256 count = 0;
        bytes32[] memory _nodes = new bytes32[](
            nodeSize.sub(nodeRemoved, "nodeSize insufficient")
        );
        for (uint256 i = 1; i <= nodeSize; i++) {
            (address node, ) = _splitToValues(nodes[i]);
            uint256 index = nodeIndex[node];
            if (index != 2**256 - 1) {
                _nodes[count] = nodes[i];
                count = count.add(1);
            }
        }
        return _nodes;
    }

    function _addNode(
        address _addr,
        bytes32 _data,
        bool _remove
    ) internal returns (bool) {
        if (_remove) {
            nodeIndex[_addr] = 2**256 - 1;
            nodeRemoved = nodeRemoved.add(1);
            return true;
        }
        uint256 index = nodeIndex[_addr]; // 0 => not exist, != 0 => exist, 2 ^ 256 -1 => removed.
        if (index == 0) {
            nodeSize = nodeSize.add(1);
            nodes[nodeSize] = _data;
            nodeIndex[_addr] = nodeSize;
        } else {
            nodes[index] = _data;
        }
        return true;
    }

    function _splitToValues(bytes32 _data)
        internal
        pure
        returns (address, uint256)
    {
        return (
            address(uint160(uint256(_data))),
            uint256(uint96(bytes12(_data)))
        );
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
