pragma solidity =0.7.5;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/ISwapContract.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract SwapContract is Ownable, ISwapContract {
    using SafeMath for uint256;

    address public WBTC_ADDR;

    uint8 public churnedInCount;
    uint8 public nodeRewardsRatio;
    uint8 public depositFeesBPS;

    uint256 private priceDecimals;
    uint256 private currentExchangeRate;
    address private lpToken;
    uint256 private lpDecimals;
    uint256 private nodeSize;

    uint256 private nextMintLPTokensForNode;

    bytes32[] private nodes;
    mapping(address => uint256) private nodeIndex;

    // Token address -> amount
    mapping(address => uint256) private totalRewards;
    mapping(address => uint256) private floatAmountOf;
    mapping(address => mapping(bytes32 => bytes32)) private txs;
    mapping(bytes32 => bool) private used;

    /**
     * Events
     */

    event RewardsCollection(address token, uint256 rewardsAmount);

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

    event DistributeNodeRewards(uint256 totalRewardsForNode);

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
        depositFeesBPS = 30;
        // Set priceDecimals
        priceDecimals = 10**8;
        // Set currentExchangeRate
        currentExchangeRate = priceDecimals;
    }

    /**
     * Transfer part
     */

    function singleTransferERC20(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _rewardsAmount
    ) public override onlyOwner returns (bool) {
        require(IERC20(_token).transfer(_to, _amount));
        _rewardsCollection(_token, _rewardsAmount);
        return true;
    }

    function multiTransferERC20TightlyPacked(
        address _token,
        bytes32[] memory _addressesAndAmounts,
        uint256 _rewardsAmount
    ) public override onlyOwner returns (bool) {
        require(_token != address(0));
        for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
            require(
                IERC20(_token).transfer(
                    address(uint160(uint256(_addressesAndAmounts[i]))),
                    uint256(uint96(bytes12(_addressesAndAmounts[i])))
                ),
                "Batch transfer error"
            );
        }
        _rewardsCollection(_token, _rewardsAmount);
        return true;
    }

    function multiTransferERC20(
        address _token,
        address[] memory _contributors,
        uint256[] memory _amounts,
        uint256 _rewardsAmount
    ) public override onlyOwner returns (bool) {
        require(
            _contributors.length == _amounts.length,
            "Length of inputs array is mismatch"
        );
        for (uint256 i = 0; i < _contributors.length; i++) {
            require(IERC20(_token).transfer(_contributors[i], _amounts[i]));
        }
        _rewardsCollection(_token, _rewardsAmount);
        return true;
    }

    /**
     * @dev gas usage 88888 gas (initial), 58888 gas (update)
     */

    function collectSwapFeesForBTC(
        address _feeToken,
        uint256 _rewardsAmount,
        bytes32 _txid
    ) public override onlyOwner returns (bool) {
        require(!used[_txid], "txid is already used");
        // _feeToken should be address(0) == BTC
        _rewardsCollection(_feeToken, _rewardsAmount);
        // Add txid to used list.
        used[_txid] = true;
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
    ) public override onlyOwner returns (bool) {
        require(txs[_token][_txid] == 0x0);
        txs[_token][_txid] = _addressesAndAmountOfFloat;
        emit RecordIncomingFloat(_token, _addressesAndAmountOfFloat, _txid);
        return true;
    }

    /**
     * @dev gas usage 131162 gas
     */

    function issueLPTokensForFloat(bytes32 _txid)
        public
        override
        returns (bool)
    {
        require(!isTxUsed(_txid), "The txid is already used");
        (address token, bytes32 transaction) = _loadTx(_txid);
        require(transaction != 0x0, "The transaction is not found");
        address to = address(uint160(uint256(transaction)));
        uint256 amountOfFloat = uint256(uint96(bytes12(transaction)));
        // LP token price per BTC/WBTC changed
        uint256 nowPrice = _updateFloatPool(address(0), WBTC_ADDR);
        // Calculate amount of LP token
        uint256 amountOfLP = amountOfFloat.mul(priceDecimals).div(nowPrice);
        uint256 depositFees;
        uint8 isFlip = _checkFlips();
        if (isFlip == 1) {
            depositFees = token == WBTC_ADDR
                ? 0
                : amountOfLP.mul(depositFeesBPS).div(10000);
        } else if (isFlip == 2) {
            depositFees = token == address(0)
                ? 0
                : amountOfLP.mul(depositFeesBPS).div(10000);
        }
        //Send LP tokens to LP
        IBurnableToken(lpToken).mint(to, amountOfLP.sub(depositFees));
        // Update deposit fees
        nextMintLPTokensForNode = nextMintLPTokensForNode.add(depositFees);
        // Add float amount
        _addFloat(token, amountOfFloat);
        used[_txid] = true;
        emit IssueLPTokensForFloat(to, amountOfLP, _txid);
        return true;
    }

    /**
     * @dev gas uasge 43628 gas
     */

    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfLPtoken,
        bytes32 _txid
    ) public override returns (bool) {
        require(txs[_token][_txid] == 0x0);
        // _token should be address(0) or WBTC_ADDR, txid should be unique
        txs[_token][_txid] = _addressesAndAmountOfLPtoken;
        emit RecordOutcomingFloat(_token, _addressesAndAmountOfLPtoken, _txid);
        return true;
    }

    /**
     * @dev gas uasge 82241 gas
     */
    function burnLPTokensForFloat(bytes32 _txid)
        public
        override
        returns (bool)
    {
        require(!isTxUsed(_txid), "The txid is already used");
        // _token should be address(0) or WBTC_ADDR
        (address token, bytes32 transaction) = _loadTx(_txid);
        require(transaction != 0x0, "The transaction is not found");
        address to = address(uint160(uint256(transaction)));
        uint256 amountOfLP = uint256(uint96(bytes12(transaction)));
        // Calculate amountOfLP
        uint256 nowPrice = _updateFloatPool(address(0), WBTC_ADDR);
        // Calculate amountOfFloat
        uint256 amountOfFloat = amountOfLP.mul(nowPrice).div(priceDecimals);

        require(
            floatAmountOf[token] >= amountOfFloat,
            "Pool balance insufficient."
        );
        // WBTC transfer if token address is WBTC
        if (token == WBTC_ADDR) {
            require(IERC20(token).transfer(to, amountOfFloat));
        }
        // Burn LP tokens
        require(IBurnableToken(lpToken).burn(amountOfLP));
        // Remove float amount
        _removeFloat(token, amountOfFloat);
        used[_txid] = true;
        emit BurnLPTokensForFloat(to, amountOfFloat, _txid);
        return true;
    }

    function distributeNodeRewards() public override returns (bool) {
        // Reduce Gas
        uint256 rewardLPsForNode = nextMintLPTokensForNode;
        require(rewardLPsForNode > 0, "totalRewardLPsForNode is not positive");
        uint256 totalStaked = 0;
        // for (uint256 i = 0; i < nodeList.length; i++) {
        //     totalStaked = totalStaked.add(
        //         uint256(uint96(bytes12(nodeList[i])))
        //     );
        // }
        // for (uint256 i = 0; i < nodeList.length; i++) {
        //     IBurnableToken(lpToken).mint(
        //         address(uint160(uint256(nodeList[i]))),
        //         rewardLPsForNode.mul(uint256(uint96(bytes12(nodeList[i])))).div(
        //             totalStaked
        //         )
        //     );
        // }
        nextMintLPTokensForNode = 0;
        emit DistributeNodeRewards(nextMintLPTokensForNode);
        return true;
    }

    /**
     * @dev gas usage 2115532 gas (initial), 592132 gas (updated)
     */

    function churn(
        address _newOwner,
        bytes32[] memory _rewardAddressAndAmounts,
        bool[] memory _isRemoved,
        uint8 _churnedInCount,
        uint8 _nodeRewardsRatio
    ) public override onlyOwner returns (bool) {
        transferOwnership(_newOwner);
        // Update active node list
        for (uint256 i = 0; i < _rewardAddressAndAmounts.length; i++) {
            (address newNode, ) = _splitToStakes(_rewardAddressAndAmounts[i]);
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

    /**
     * @dev returns reserves - deposit fees.
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

    function _checkFlips() public view returns (uint8 way) {
        (uint256 reserveA, uint256 reserveB) = getFloatReserve(
            address(0),
            WBTC_ADDR
        );
        // BTC bal == BTC float + WBTC float - WBTC bal
        uint256 balWBTC = IERC20(WBTC_ADDR).balanceOf(address(this));
        require(reserveA.add(reserveB) >= balWBTC, "balWBTC amount invalid");
        if (reserveA.add(reserveB).sub(balWBTC) > balWBTC) {
            way = 1;
        } else if (reserveA.add(reserveB).sub(balWBTC) < balWBTC) {
            way = 2;
        } else {
            way = 0;
        }
        return way;
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
                totalLPs.add(nextMintLPTokensForNode)
            );
        return currentExchangeRate;
    }

    function _addFloat(address _token, uint256 _amount) internal {
        floatAmountOf[_token] = floatAmountOf[_token].add(_amount);
    }

    function _removeFloat(address _token, uint256 _amount) internal {
        floatAmountOf[_token] = floatAmountOf[_token].sub(_amount);
    }

    function _rewardsCollection(address _token, uint256 _rewardsAmount)
        internal
    {
        // Add all fees into pool
        totalRewards[_token] = totalRewards[_token].add(_rewardsAmount);
        uint256 amountForLP = _rewardsAmount.mul(nodeRewardsRatio).div(100);
        // Future mint LP tokens 66% of fees
        uint256 amountLPForNode = _rewardsAmount
            .sub(amountForLP)
            .mul(priceDecimals)
            .div(getCurrentPriceLP());
        // Add minted LP tokens for Nodes
        nextMintLPTokensForNode = nextMintLPTokensForNode.add(amountLPForNode);
    }

    function _loadTx(bytes32 _txid) internal view returns (address, bytes32) {
        if (txs[address(0)][_txid] != 0x0) {
            return (address(0), txs[address(0)][_txid]);
        }
        if (txs[WBTC_ADDR][_txid] != 0x0) {
            return (WBTC_ADDR, txs[WBTC_ADDR][_txid]);
        }
        return (address(0x0), 0x0);
    }

    function _addNode(
        address _addr,
        bytes32 _data,
        bool _remove
    ) internal returns (bool) {
        if (_remove) {
            nodeIndex[_addr] = 2**256 - 1;
            return true;
        }
        uint256 index = nodeIndex[_addr]; // 0 => not exist, != 0 => exist, 2 ^ 256 -1 => removed.
        if (index == 0) {
            nodes.push(_data);
            nodeIndex[_addr] = nodes.length - 1;
        } else {
            nodes[index] = _data;
        }
        return true;
    }

    function loadNodes() public returns (bytes32[] memory) {
        bytes32[] memory _nodes;
        uint256 count = 0;
        for (uint256 i = 0; i < nodes.length; i++) {
            (address newNode, ) = _splitToStakes(nodes[i]);
            uint256 index = nodeIndex[newNode];
            if (index != 2**256 - 1) {
                _nodes[count] = nodes[i];
                count++;
            }
        }
        return _nodes;
    }

    function _splitToStakes(bytes32 _data)
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
