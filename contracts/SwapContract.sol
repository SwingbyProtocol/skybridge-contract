pragma solidity =0.7.5;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/ISwapContract.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract SwapContract is Ownable, ISwapContract {
    using SafeMath for uint256;

    address public WBTC_ADDR;
    uint256 public DEFAULT_INDEX = 2**256 - 1;
    bytes32[] public nodes;

    uint8 public churnedInCount;
    uint8 public nodeRewardsRatio;

    uint256 private priceDecimals;
    uint256 private currentExchangeRate;
    address private lpToken;
    uint256 private lpDecimals;

    // Token address -> amount
    mapping(address => uint256) private totalRewardsForLPs;
    mapping(address => uint256) private totalRewardsForNodes;
    mapping(address => uint256) private floatAmountOfToken;
    mapping(address => mapping(bytes32 => bytes32)) private txs;
    mapping(bytes32 => bool) private used;

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
        //burner = new Burner();
        lpToken = _lpToken;
        // Set initial price of LP token per BTC/WBTC.
        lpDecimals = 10**IERC20(lpToken).decimals();
        // Set WBTC address
        WBTC_ADDR = _wbtc;
        // Set nodeRewardsRatio
        nodeRewardsRatio = 66;
        // set priceDecimals
        priceDecimals = 10**8;
        // SEt currentExchangeRate
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
        uint256 nowPrice = _updatePool(address(0), WBTC_ADDR);
        // Calculate amount of LP token
        uint256 amountOfLP = amountOfFloat.mul(priceDecimals).div(nowPrice);
        // Mint LP tokens
        IBurnableToken(lpToken).mint(to, amountOfLP);
        // Add float amount
        floatAmountOfToken[token] = floatAmountOfToken[token].add(
            amountOfFloat
        );
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
        //  _token should be address(0) or WBTC_ADDR
        (address token, bytes32 transaction) = _loadTx(_txid);
        require(transaction != 0x0, "The transaction is not found");
        address to = address(uint160(uint256(transaction)));
        uint256 amountOfLP = uint256(uint96(bytes12(transaction)));
        uint256 nowPrice = _updatePool(address(0), WBTC_ADDR);
        // Burn LP tokens
        require(IBurnableToken(lpToken).burn(amountOfLP));
        // Calculate amount of BTC/WBTC
        uint256 amountOfFloat = amountOfLP.mul(nowPrice).div(priceDecimals);
        // WBTC transfer if tokne address is WBTC
        if (token == WBTC_ADDR) {
            require(IERC20(token).transfer(to, amountOfFloat));
        }
        // Remove float amount
        require(
            floatAmountOfToken[token] >= amountOfFloat,
            "floatAmount is not enough"
        );
        floatAmountOfToken[token] = floatAmountOfToken[token].sub(
            amountOfFloat
        );
        used[_txid] = true;
        emit BurnLPTokensForFloat(to, amountOfFloat, _txid);
        return true;
    }

    function distributeNodeRewards() public override returns (bool) {
        // Reduce Gas
        uint256 totalRewardsForNode = totalRewardsForNodes[WBTC_ADDR];
        require(
            totalRewardsForNode > 0,
            "totalRewardsForNode amount is not positive"
        );
        uint256 totalStaked = 0;
        for (uint256 i = 0; i < nodes.length; i++) {
            totalStaked = totalStaked.add(uint256(uint96(bytes12(nodes[i]))));
        }
        // decimals of totalRewardsForNode  == 8, decimals of currentPrice == 8
        uint256 totalLPTokens = totalRewardsForNode.mul(priceDecimals).div(
            getCurrentPriceLP()
        );
        for (uint256 i = 0; i < nodes.length; i++) {
            IBurnableToken(lpToken).mint(
                address(uint160(uint256(nodes[i]))),
                totalLPTokens.mul(uint256(uint96(bytes12(nodes[i])))).div(
                    totalStaked
                )
            );
        }
        // Inject WBTC amount to float?
        // floatAmountOfToken[WBTC_ADDR] = floatAmountOfToken[WBTC_ADDR].add(
        //     totalRewardsForNode
        // );
        // Reset storage for WBTC fees.
        totalRewardsForNodes[WBTC_ADDR] = 0;
        emit DistributeNodeRewards(totalRewardsForNode);
        return true;
    }

    /**
     * @dev gas usage 2115532 gas (initial), 592132 gas (updated)
     */

    function churn(
        address _newOwner,
        bytes32[] memory _newRewardsAddressAndAmounts,
        address[] memory _removedRewardsAddresses,
        uint8 _churnedInCount,
        uint8 _nodeRewardsRatio
    ) public override onlyOwner returns (bool) {
        transferOwnership(_newOwner);
        // Update active node list
        for (uint256 i = 0; i < _newRewardsAddressAndAmounts.length; i++) {
            (address newNode, ) = _splitToStakes(
                _newRewardsAddressAndAmounts[i]
            );
            uint256 index = _checkMatch(newNode);
            if (index != DEFAULT_INDEX) {
                // Update stakes
                nodes[index] = _newRewardsAddressAndAmounts[i];
            } else {
                // Add stakes
                nodes.push(_newRewardsAddressAndAmounts[i]);
            }
        }
        // Remove all removed list
        for (uint256 i = 0; i < _removedRewardsAddresses.length; i++) {
            uint256 index = _checkMatch(_removedRewardsAddresses[i]);
            if (index != DEFAULT_INDEX) {
                delete (nodes[index]);
            }
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

    function getFloatReserve(address _tokenA, address _tokenB)
        public
        override
        view
        returns (uint256 reserveA, uint256 reserveB)
    {
        (reserveA, reserveB) = (
            floatAmountOfToken[_tokenA].add(totalRewardsForLPs[_tokenA]),
            floatAmountOfToken[_tokenB].add(totalRewardsForLPs[_tokenB])
        );
    }

    function _rewardsCollection(address _token, uint256 _rewardsAmount)
        internal
    {
        // Reduce Gas
        uint256 totalRewardsForNode = totalRewardsForNodes[_token];
        uint256 totalRewardsForLP = totalRewardsForLPs[_token];
        // Updates rewards for nodes and LPs
        uint256 rewardsForNode = _rewardsAmount.mul(nodeRewardsRatio).div(100);
        uint256 rewardsForLP = _rewardsAmount.sub(rewardsForNode);
        totalRewardsForNodes[_token] = totalRewardsForNode.add(rewardsForNode);
        totalRewardsForLPs[_token] = totalRewardsForLP.add(rewardsForLP);
        emit RewardsCollection(_token, _rewardsAmount);
    }

    function _updatePool(address _tokenA, address _tokenB)
        internal
        returns (uint256)
    {
        // Reduce gas cost.
        (uint256 reserveA, uint256 reserveB) = getFloatReserve(
            _tokenA,
            _tokenB
        );
        // uint256 burned = IBurnableToken(lpToken).balanceOf(address(burner));
        uint256 totalLPs = IBurnableToken(lpToken).totalSupply();
        // decimals of totalReserved == 8, lpDecimals == 8, decimals of rate == 8
        currentExchangeRate = totalLPs == 0
            ? currentExchangeRate
            : (reserveA.add(reserveB)).mul(lpDecimals).div(totalLPs);
        return currentExchangeRate;
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

    function _checkMatch(address _node) internal view returns (uint256 index) {
        index = DEFAULT_INDEX;
        for (uint256 i = 0; i < nodes.length; i++) {
            (address node, ) = _splitToStakes(nodes[i]);
            if (node == _node) {
                index = i;
                return index;
            }
        }
        return index;
    }

    function _splitToStakes(bytes32 _data)
        internal
        view
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
