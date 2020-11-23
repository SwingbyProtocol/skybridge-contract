pragma solidity =0.7.5;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/ISwapContract.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract SwapContract is Ownable, ISwapContract {
    using SafeMath for uint256;

    address public WBTC_ADDR;
    bytes32[] public nodes;

    uint8 public churnedInCount;
    uint8 public nodeRewardsRatio;

    uint256 private priceDecimals = 10**8;
    uint256 private currentExchangeRate = priceDecimals;
    address private lpToken;
    uint256 private lpDecimals;

    // Token address -> amount
    mapping(address => uint256) private totalRewardsForLPs;
    mapping(address => uint256) private totalRewardsForNodes;
    mapping(address => uint256) private floatAmountOfToken;
    mapping(address => mapping(bytes32 => bytes32)) private txs;
    mapping(bytes32 => bool) private used;

    constructor(address _lpToken, address _wbtc) public {
        //burner = new Burner();
        lpToken = _lpToken;
        // Set initial price of LP token per BTC/WBTC.
        lpDecimals = 10**IERC20(lpToken).decimals();
        // Set WBTC address
        WBTC_ADDR = _wbtc;
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

    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfFloat,
        bytes32 _txid
    ) public override onlyOwner returns (bool) {
        require(txs[_token][_txid] == 0x0);
        txs[_token][_txid] = _addressesAndAmountOfFloat;
        return true;
    }

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
        return true;
    }

    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfLPtoken,
        bytes32 _txid
    ) public override returns (bool) {
        require(txs[_token][_txid] == 0x0);
        // _token should be address(0) or WBTC_ADDR, txid should be unique
        txs[_token][_txid] = _addressesAndAmountOfLPtoken;
        return true;
    }

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
        IBurnableToken(lpToken).burn(amountOfLP);
        // Calculate amount of BTC/WBTC
        uint256 amountOfFloat = amountOfLP.mul(nowPrice).div(priceDecimals);
        // WBTC transfer if tokne address is WBTC
        if (token == WBTC_ADDR) {
            require(IERC20(token).transfer(to, amountOfFloat));
        }
        // Remove float amount
        floatAmountOfToken[token] = floatAmountOfToken[token].sub(
            amountOfFloat
        );
        used[_txid] = true;
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
        return true;
    }

    function churn(
        address _newOwner,
        bytes32[] memory _nodeRewardsAddressAndAmounts,
        uint8 _churnedInCount,
        uint8 _nodeRewardsRatio
    ) public override onlyOwner returns (bool) {
        transferOwnership(_newOwner);
        nodes = _nodeRewardsAddressAndAmounts;
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
    }

    function _updatePool(address _tokenA, address _tokenB)
        internal
        returns (uint256)
    {
        // Reduce gas cost.
        uint256 floatAmountOfTokenA = floatAmountOfToken[_tokenA];
        uint256 floatAmountOfTokenB = floatAmountOfToken[_tokenB];
        uint256 totalRewardOfTokenA = totalRewardsForLPs[_tokenA];
        uint256 totalRewardOfTokenB = totalRewardsForLPs[_tokenB];

        uint256 totalReserved = floatAmountOfTokenA
            .add(floatAmountOfTokenB)
            .add(totalRewardOfTokenA)
            .add(totalRewardOfTokenB);

        // Logic: LPP = (float amount of BTC + float amount of WBTC + LP fees) / (LP supply)
        // uint256 burned = IBurnableToken(lpToken).balanceOf(address(burner));
        uint256 totalLPs = IBurnableToken(lpToken).totalSupply();
        // decimals of totalReserved == 8, lpDecimals == 8, decimals of rate == 8
        currentExchangeRate = totalLPs == 0
            ? currentExchangeRate
            : totalReserved.mul(lpDecimals).div(totalLPs);
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

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
