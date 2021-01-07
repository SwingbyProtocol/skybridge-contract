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
    uint8 public tssThreshold;
    uint8 public nodeRewardsRatio;
    uint8 public depositFeesBPS;
    uint8 public withdrawalFeeBPS;

    uint256 public lockedLPTokensForNode;

    uint256 private priceDecimals;
    uint256 private currentExchangeRate;
    uint256 private lpDecimals;
    // Support tokens
    mapping(address => bool) public whitelist;

    // Nodes
    mapping(address => bytes32) private nodes;
    mapping(address => bool) private isInList;
    address[] private nodeAddrs;
    // Token address -> amount
    mapping(address => uint256) private totalRewards;
    mapping(address => uint256) private floatAmountOf;
    mapping(bytes32 => bool) private used;

    /**
     * Events
     */

    event IssueLPTokensForFloat(
        address to,
        uint256 amountOfFloat,
        uint256 amountOfLP,
        bytes32 txid
    );

    event BurnLPTokensForFloat(
        address token,
        uint256 amountOfLP,
        uint256 amountOfFloat,
        bytes32 txid
    );

    modifier priceCheck() {
        uint256 beforePrice = getCurrentPriceLP();
        _;
        require(getCurrentPriceLP() >= beforePrice, "Invalid  LP price change");
    }

    constructor(
        address _lpToken,
        address _wbtc,
        uint256 _existingBTCFloat
    ) public {
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
        // Set withdrawalFeeBPS
        withdrawalFeeBPS = 20;
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
        floatAmountOf[address(0)] = _existingBTCFloat;
    }

    /**
     * Transfer part
     */

    /// @dev singleTransferERC20 function sends tokens from contract.
    /// @param _destToken Address of token.
    /// @param _to Recevier address.
    /// @param _amount The amount of tokens.
    /// @param _totalSwapped the amount of swapped amount which is for send.
    /// @param _rewardsAmount Value that should be paid as fees.
    /// @param _redeemedFloatTxIds the txs which is for records txids.
    function singleTransferERC20(
        address _destToken,
        address _to,
        uint256 _amount,
        uint256 _totalSwapped,
        uint256 _rewardsAmount,
        bytes32[] memory _redeemedFloatTxIds
    ) external override onlyOwner returns (bool) {
        require(whitelist[_destToken], "_destToken is not whitelisted");
        require(
            _destToken != address(0),
            "_destToken should not be address(0)"
        );
        if (_destToken == WBTC_ADDR && _totalSwapped > 0) {
            _swap(address(0), WBTC_ADDR, _totalSwapped);
        }
        _rewardsCollection(_destToken, _rewardsAmount);
        _addTxidUsed(_redeemedFloatTxIds);
        require(IERC20(_destToken).transfer(_to, _amount));
        return true;
    }

    /// @dev multiTransferERC20TightlyPacked function sends tokens from contract.
    /// @param _destToken Address of token.
    /// @param _addressesAndAmounts Recevier address and amounts.
    /// @param _totalSwapped the amount of swapped amount which is for send.
    /// @param _rewardsAmount Value that should be paid as fees.
    /// @param _redeemedFloatTxIds the txs which is for records txids.
    function multiTransferERC20TightlyPacked(
        address _destToken,
        bytes32[] memory _addressesAndAmounts,
        uint256 _totalSwapped,
        uint256 _rewardsAmount,
        bytes32[] memory _redeemedFloatTxIds
    ) external override onlyOwner returns (bool) {
        require(whitelist[_destToken], "_destToken is not whitelisted");
        require(
            _destToken != address(0),
            "_destToken should not be address(0)"
        );
        if (_destToken == WBTC_ADDR && _totalSwapped > 0) {
            _swap(address(0), WBTC_ADDR, _totalSwapped);
        }
        _rewardsCollection(_destToken, _rewardsAmount);
        _addTxidUsed(_redeemedFloatTxIds);
        for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
            require(
                IERC20(_destToken).transfer(
                    address(uint160(uint256(_addressesAndAmounts[i]))),
                    uint256(uint96(bytes12(_addressesAndAmounts[i])))
                ),
                "Batch transfer error"
            );
        }
        return true;
    }

    /// @dev collectSwapFeesForBTC function collectes fees on BTC.
    /// @param _destToken Address of token.
    /// @param _incomingAmount spent amount of BTC.
    /// @param _rewardsAmount Value that should be paid as fees.
    function collectSwapFeesForBTC(
        address _destToken,
        uint256 _incomingAmount,
        uint256 _rewardsAmount
    ) external override onlyOwner returns (bool) {
        require(_destToken == address(0), "_destToken should be address(0)");
        _swap(WBTC_ADDR, address(0), _incomingAmount);
        _rewardsCollection(_destToken, _rewardsAmount);
        return true;
    }

    /**
     * Float part
     */

    /// @dev recordIncomingFloat function mint LP token.
    /// @param _token Address of target token.
    /// @param _addressesAndAmountOfFloat Recevier address and amounts.
    /// @param _zerofee The flag of accept.
    /// @param _txid the txs which is for records txids.
    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfFloat,
        bool _zerofee,
        bytes32 _txid
    ) external override onlyOwner priceCheck returns (bool) {
        require(whitelist[_token], "_token is invalid");
        require(
            _issueLPTokensForFloat(
                _token,
                _addressesAndAmountOfFloat,
                _zerofee,
                _txid
            )
        );
        return true;
    }

    /// @dev recordOutcomingFloat function burn LP token.
    /// @param _token Address of target token.
    /// @param _addressesAndAmountOfLPtoken Sender address and amounts.
    /// @param _txid the txs which is for records txids.
    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfLPtoken,
        uint256 _minerFee,
        bytes32 _txid
    ) external override onlyOwner priceCheck returns (bool) {
        require(whitelist[_token], "_token is invalid");
        require(
            _burnLPTokensForFloat(
                _token,
                _addressesAndAmountOfLPtoken,
                withdrawalFeeBPS,
                _minerFee,
                _txid
            )
        );
        return true;
    }

    /// @dev distributeNodeRewards function sends rewards for Nodes.
    function distributeNodeRewards() external override returns (bool) {
        // Reduce Gas
        uint256 rewardLPsForNodes = lockedLPTokensForNode;
        require(rewardLPsForNodes > 0, "totalRewardLPsForNode is not positive");
        bytes32[] memory nodeList = getActiveNodes();
        uint256 totalStaked = 0;
        for (uint256 i = 0; i < nodeList.length; i++) {
            totalStaked = totalStaked.add(
                uint256(uint96(bytes12(nodeList[i])))
            );
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

    /// @dev churn function transfer contract ownership and set variables.
    /// @param _newOwner Address of new Owner.
    /// @param _rewardAddressAndAmounts Staker addresses and amounts.
    /// @param _isRemoved The flags for remove node.
    /// @param _churnedInCount The number of next N count.
    /// @param _tssThreshold The number of next T.
    /// @param _nodeRewardsRatio The number of next node rewards ratio.
    /// @param _withdrawalFeeBPS The amount of wthdrawal fees.
    function churn(
        address _newOwner,
        bytes32[] memory _rewardAddressAndAmounts,
        bool[] memory _isRemoved,
        uint8 _churnedInCount,
        uint8 _tssThreshold,
        uint8 _nodeRewardsRatio,
        uint8 _withdrawalFeeBPS
    ) external override onlyOwner returns (bool) {
        require(
            _tssThreshold >= tssThreshold && _tssThreshold <= 2**8 - 1,
            "_tssThreshold should be >= tssThreshold"
        );
        require(
            _churnedInCount >= _tssThreshold + uint8(1),
            "n should be >= t+1"
        );
        require(
            _nodeRewardsRatio >= 0 && _nodeRewardsRatio <= 100,
            "_nodeRewardsRatio is not valid"
        );
        require(
            _withdrawalFeeBPS >= 0 && _withdrawalFeeBPS <= 100,
            "_withdrawalFeeBPS is invalid"
        );
        require(
            _rewardAddressAndAmounts.length == _isRemoved.length,
            "_rewardAddressAndAmounts and _isRemoved length is not match"
        );
        transferOwnership(_newOwner);
        // Update active node list
        for (uint256 i = 0; i < _rewardAddressAndAmounts.length; i++) {
            (address newNode, ) = _splitToValues(_rewardAddressAndAmounts[i]);
            _addNode(newNode, _rewardAddressAndAmounts[i], _isRemoved[i]);
        }
        bytes32[] memory nodeList = getActiveNodes();
        if (nodeList.length > 100) {
            revert("node size should be <= 100");
        }
        churnedInCount = _churnedInCount;
        tssThreshold = _tssThreshold;
        nodeRewardsRatio = _nodeRewardsRatio;
        withdrawalFeeBPS = _withdrawalFeeBPS;
        return true;
    }

    /// @dev isTxUsed function sends rewards for Nodes.
    /// @param _txid txid of incoming tx.
    function isTxUsed(bytes32 _txid) public override view returns (bool) {
        return used[_txid];
    }

    /// @dev getCurrentPriceLP function returns exchange rate of LP token.
    function getCurrentPriceLP() public override view returns (uint256) {
        return currentExchangeRate;
    }

    /// @dev getDepositFeeRate function returns deposit fees rate
    /// @param _token The address of target token.
    /// @param _amountOfFloat The amount of float.
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

    /// @dev getMinimumAmountOfLPTokens function returns the minimum amount of LP Token.
    /// @param _minerFees The amount of miner Fees (BTC).
    function getMinimumAmountOfLPTokens(uint256 _minerFees)
        public
        override
        view
        returns (uint256, uint256)
    {
        (uint256 reserveA, uint256 reserveB) = getFloatReserve(
            address(0),
            WBTC_ADDR
        );
        uint256 totalLPs = IBurnableToken(lpToken).totalSupply();
        // decimals of totalReserved == 8, lpDecimals == 8, decimals of rate == 8
        uint256 nowPrice = totalLPs == 0
            ? currentExchangeRate
            : (reserveA.add(reserveB)).mul(lpDecimals).div(
                totalLPs.add(lockedLPTokensForNode)
            );
        uint256 requiredFloat = _minerFees.mul(10000).div(withdrawalFeeBPS);
        uint256 amountOfLPTokens = requiredFloat.add(10).mul(priceDecimals).div(
            nowPrice
        );
        return (amountOfLPTokens, nowPrice);
    }

    /// @dev getFloatReserve function returns float reserves not current balances.
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

    /// @dev getActiveNodes function returns active nodes list (stakes and amount)
    function getActiveNodes() public override view returns (bytes32[] memory) {
        uint256 nodeCount = 0;
        uint256 count = 0;
        // Seek all nodes
        for (uint256 i = 0; i < nodeAddrs.length; i++) {
            if (nodes[nodeAddrs[i]] != 0x0) {
                nodeCount = nodeCount.add(1);
            }
        }
        bytes32[] memory _nodes = new bytes32[](nodeCount);
        for (uint256 i = 0; i < nodeAddrs.length; i++) {
            if (nodes[nodeAddrs[i]] != 0x0) {
                _nodes[count] = nodes[nodeAddrs[i]];
                count = count.add(1);
            }
        }
        return _nodes;
    }

    /// @dev _issueLPTokensForFloat
    /// @param _token Address of target token.
    /// @param _transaction Recevier address and amounts.
    /// @param _zerofee The flag of accept.
    /// @param _txid the txs which is for records txids.
    function _issueLPTokensForFloat(
        address _token,
        bytes32 _transaction,
        bool _zerofee,
        bytes32 _txid
    ) internal returns (bool) {
        require(!isTxUsed(_txid), "The txid is already used");
        // (address token, bytes32 transaction) = _loadTx(_txid);
        require(_transaction != 0x0, "The transaction is not found");
        // Define target address which is recorded bottom 20bytes on tx data
        // Define amountLP which is recorded top 12bytes on tx data
        (address to, uint256 amountOfFloat) = _splitToValues(_transaction);
        // LP token price per BTC/WBTC changed
        uint256 nowPrice = _updateFloatPool(address(0), WBTC_ADDR);
        // Calculate amount of LP token
        uint256 amountOfLP = amountOfFloat.mul(priceDecimals).div(nowPrice);
        uint256 depositFeeRate = getDepositFeeRate(_token, amountOfFloat);
        uint256 depositFees = depositFeeRate != 0
            ? amountOfLP.mul(depositFeeRate).div(10000)
            : 0;

        if (_zerofee && depositFees != 0) {
            revert();
        }
        //Send LP tokens to LP
        IBurnableToken(lpToken).mint(to, amountOfLP.sub(depositFees));
        // Add deposit fees
        lockedLPTokensForNode = lockedLPTokensForNode.add(depositFees);
        // Add float amount
        _addFloat(_token, amountOfFloat);
        used[_txid] = true;
        emit IssueLPTokensForFloat(to, amountOfFloat, amountOfLP, _txid);
        return true;
    }

    /// @dev _burnLPTokensForFloat
    /// @param _token Address of target token.
    /// @param _transaction Sender address and amounts.
    /// @param _withdrawalFeeBPS The amount of withdrawal fees.
    /// @param _txid the txs which is for records txids.
    function _burnLPTokensForFloat(
        address _token,
        bytes32 _transaction,
        uint256 _withdrawalFeeBPS,
        uint256 _minerFee,
        bytes32 _txid
    ) internal returns (bool) {
        require(!isTxUsed(_txid), "The txid is already used");
        // _token should be address(0) or WBTC_ADDR
        // (address token, bytes32 transaction) = _loadTx(_txid);
        require(_transaction != 0x0, "The transaction is not found");
        // Define target address which is recorded bottom 20bytes on tx data
        // Define amountLP which is recorded top 12bytes on tx data
        (address to, uint256 amountOfLP) = _splitToValues(_transaction);
        // Calculate amountOfLP
        uint256 nowPrice = _updateFloatPool(address(0), WBTC_ADDR);
        // Calculate amountOfFloat
        uint256 amountOfFloat = amountOfLP.mul(nowPrice).div(priceDecimals);
        uint256 amountOfFees = amountOfFloat.mul(_withdrawalFeeBPS).div(10000);
        require(
            floatAmountOf[_token] >= amountOfFloat,
            "Pool balance insufficient."
        );
        require(
            _minerFee <= amountOfFees,
            "amountOfFees.sub(_minerFee) is negative"
        );
        // Burn LP tokens
        require(IBurnableToken(lpToken).burn(amountOfLP));
        // Remove float amount
        _removeFloat(_token, amountOfFloat);
        // Collect fees
        _rewardsCollection(_token, amountOfFees.sub(_minerFee));
        used[_txid] = true;
        // WBTC transfer if token address is WBTC_ADDR
        if (_token == WBTC_ADDR) {
            require(
                IERC20(_token).transfer(
                    to,
                    amountOfFloat.sub(amountOfFees).sub(_minerFee)
                ),
                "WBTC balance insufficient"
            );
        }
        emit BurnLPTokensForFloat(to, amountOfLP,amountOfFloat, _txid);
        return true;
    }

    /// @dev _checkFlips checks whether the fees are activated.
    /// @param _token Address of target token.
    /// @param _amountOfFloat The amount of float.
    function _checkFlips(address _token, uint256 _amountOfFloat)
        internal
        view
        returns (uint8)
    {
        (uint256 reserveA, uint256 reserveB) = getFloatReserve(
            address(0),
            WBTC_ADDR
        );
        uint256 threshold = reserveA
            .add(reserveB)
            .add(_amountOfFloat)
            .mul(2)
            .div(3);
        if (_token == WBTC_ADDR) {
            if (reserveA.add(_amountOfFloat) >= threshold) {
                return 1; // BTC float insufficient
            }
        } else if (_token == address(0)) {
            if (reserveB.add(_amountOfFloat) >= threshold) {
                return 2; // WBTC float insufficient
            }
        }
        return 0;
    }

    /// @dev _updateFloatPool updates float balances.
    /// @param _tokenA Address of target tokenA.
    /// @param _tokenB Address of target tokenB.
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

    /// @dev _addFloat updates Float.
    /// @param _token The address of target token.
    /// @param _amount The amount of float.
    function _addFloat(address _token, uint256 _amount) internal {
        floatAmountOf[_token] = floatAmountOf[_token].add(_amount);
    }

    /// @dev _removeFloat remove Float.
    /// @param _token The address of target token.
    /// @param _amount The amount of float.
    function _removeFloat(address _token, uint256 _amount) internal {
        floatAmountOf[_token] = floatAmountOf[_token].sub(
            _amount,
            "float amount insufficient"
        );
    }

    /// @dev _swap collects swap amount to change float.
    /// @param _sourceToken The address of source token
    /// @param _destToken The address of target token.
    /// @param _swapAmount The amount of swap.
    function _swap(
        address _sourceToken,
        address _destToken,
        uint256 _swapAmount
    ) internal {
        floatAmountOf[_destToken] = floatAmountOf[_destToken].sub(
            _swapAmount,
            "float amount insufficient"
        );
        floatAmountOf[_sourceToken] = floatAmountOf[_sourceToken].add(
            _swapAmount
        );
    }

    /// @dev _rewardsCollection collects rewards.
    /// @param _destToken The address of target token.
    /// @param _rewardsAmount The amount of rewards.
    function _rewardsCollection(address _destToken, uint256 _rewardsAmount)
        internal
    {
        if (_destToken == lpToken) return;
        if (_rewardsAmount == 0) return;
        // The fee is always collected in the source token (it's left in the float on the origin chain).
        address _feesToken = _destToken == WBTC_ADDR ? address(0) : WBTC_ADDR;
        // Add all fees into pool
        totalRewards[_feesToken] = totalRewards[_feesToken].add(_rewardsAmount);
        uint256 amountForNodes = _rewardsAmount.mul(nodeRewardsRatio).div(100);
        // Alloc LP tokens for nodes as fees
        uint256 amountLPForNode = amountForNodes.mul(priceDecimals).div(
            getCurrentPriceLP()
        );
        // Add minted LP tokens for Nodes
        lockedLPTokensForNode = lockedLPTokensForNode.add(amountLPForNode);
    }

    /// @dev _addTxidUsed updates a spent txhash.
    /// @param _txs The array of txid.
    function _addTxidUsed(bytes32[] memory _txs) internal {
        for (uint256 i = 0; i < _txs.length; i++) {
            used[_txs[i]] = true;
        }
    }

    /// @dev _addNode updates a Staker.
    /// @param _addr The address of staker.
    /// @param _data The data of staker.
    /// @param _remove The flag for remove.
    function _addNode(
        address _addr,
        bytes32 _data,
        bool _remove
    ) internal returns (bool) {
        if (_remove) {
            delete nodes[_addr];
            return true;
        }
        if (!isInList[_addr]) {
            nodeAddrs.push(_addr);
            isInList[_addr] = true;
        }
        if (nodes[_addr] == 0x0) {
            nodes[_addr] = _data;
        }
        return true;
    }

    /// @dev _splitToValues returns address and amount of stakes
    /// @param _data The data of staker.
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

    /// @dev The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
