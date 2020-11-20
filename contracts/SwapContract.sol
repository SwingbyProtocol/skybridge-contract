pragma solidity =0.7.0;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/ISwapContract.sol";
import "./Ownable.sol";
import "./SafeMath.sol";
import "./Burner.sol";

contract SwapContract is Ownable, ISwapContract {
    using SafeMath for uint256;

    address public WBTC_ADDR = 0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599;

    address private lpToken;
    // Token address -> amount
    mapping(address => uint256) private totalRewardsForLPs;
    mapping(address => uint256) private totalRewardsForNodes;
    mapping(address => uint256) private floatAmountOfToken;
    mapping(address => uint256) private currentExchangeRate;
    mapping(address => mapping(bytes32 => bytes32)) private txs;
    Burner public burner;
    uint8 public churnedInCount;
    uint8 public nodeRewardsRatio;
    bytes32[] public nodes;

    struct Node {
        address addr;
        uint256 staked;
    }

    constructor(address _lpToken) public {
        burner = new Burner();
        lpToken = _lpToken;
        // Set initial price of LP token per BTC/WBTC.
        currentExchangeRate[address(0)] = 1e8;
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
        uint8 _inputDecimals,
        uint256 _rewardsAmount
    ) public override onlyOwner returns (bool) {
        require(_token != address(0));
        for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
            IERC20 token = IERC20(_token);
            address to = address(uint160(uint256(_addressesAndAmounts[i])));
            uint8 boost = token.decimals() - _inputDecimals;
            require(boost >= 0, "boost should be >= 0");
            uint256 amount;
            if (boost == uint8(0)) {
                amount = uint256(uint96(bytes12(_addressesAndAmounts[i])));
            } else {
                amount = uint256(uint96(bytes12(_addressesAndAmounts[i]))).mul(
                    10**uint256(boost)
                );
            }
            require(token.transfer(to, amount));
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
     * Float part
     */

    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmounts,
        bytes32 _txid
    ) public override onlyOwner returns (bool) {
        txs[_token][_txid] = _addressesAndAmounts;
        return true;
    }

    function issueLPTokensForFloat(bytes32 _txid) public {
        bytes32 payload;
        if (txs[address(0)][_txid] == 0x0) {
            payload = txs[address(0)][_txid];
        }
        if (txs[WBTC_ADDR][_txid] == 0x0) {
            payload = txs[WBTC_ADDR][_txid];
        }
        address to = address(uint160(uint256(payload)));
        uint256 amount = uint256(uint96(bytes12(payload)));
        // LP token price per BTC/WBTC changed.
        uint256 newPrice = _updatePool(address(0), WBTC_ADDR);
        // Mint LP tokens based on current LP price to users.
        IBurnableToken(lpToken).mint(to, amount.mul(1e8).div(newPrice));
    }

    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmounts,
        bytes32 _txid
    ) public override returns (bool) {
        // _token should be WBTC
        txs[_token][_txid] = _addressesAndAmounts;
        return true;
    }

    function burnLPTokensForFloat(bytes32 _txid)
        public
        override
        returns (bool)
    {
        bytes32 payload;
        if (txs[address(0)][_txid] == 0x0) {
            payload = txs[address(0)][_txid];
        }
        if (txs[WBTC_ADDR][_txid] == 0x0) {
            payload = txs[WBTC_ADDR][_txid];
        }
        // Transfer tokens to this contract from users.
        IERC20(_token).transferFrom(_msgSender(), address(this), _amount);
        // Record float amount
        floatAmountOfToken[_token] = floatAmountOfToken[_token].add(_amount);
        // Get exchange rate for BTC <> LP, And for reduece gas
        uint256 rate = getCurrentPriceLP(_token);
        // Mint LP tokens which is based on current LP price to users.
        IBurnableToken(lpToken).mint(_msgSender(), _amount.mul(1e8).div(rate));
        IBurnableToken(lpToken).transferFrom(
            _msgSender(),
            address(this),
            _amountOfLPToken
        );
        IBurnableToken(lpToken).burn(_amountOfLPToken);
        // Get exchange rate for BTC/LP, And for reduece gas
        uint256 rate = getCurrentPriceLP(_token);
        IERC20(_token).transferFrom(
            _msgSender(),
            address(this),
            _amountOfLPToken.div(rate)
        );
        return true;
    }

    function distributeNodeRewards(address _token)
        public
        override
        returns (bool)
    {
        // Reduce Gas
        uint256 totalRewardsForNode = totalRewardsForNodes[_token];
        uint256 totalStaked = 0;
        require(totalRewardsForNode > 0, "totalRewardsForNode amount is 0");
        for (uint256 i = 0; i < nodes.length; i++) {
            totalStaked = totalStaked.add(uint256(uint96(bytes12(nodes[i]))));
        }
        for (uint256 i = 0; i < nodes.length; i++) {
            IERC20(_token).transfer(
                address(uint160(uint256(nodes[i]))),
                totalRewardsForNode.mul(uint256(uint96(bytes12(nodes[i])))).div(
                    totalStaked
                )
            );
        }
        // Zerolize for storage, gas refunded.
        totalRewardsForNodes[_token] = 0;
        return true;
    }

    function churn(
        address _newOwner,
        bytes32[] memory _nodeRewardsAddressAndAmounts,
        uint8 _churnedInCount,
        uint8 _nodeRewardsRatio
    ) public override onlyOwner returns (bool) {
        _transferOwnership(_newOwner);
        nodes = _nodeRewardsAddressAndAmounts;
        churnedInCount = _churnedInCount;
        // The ratio should be 100x of actual rate.
        nodeRewardsRatio = _nodeRewardsRatio;
        return true;
    }

    function _rewardsCollection(address _token, uint256 _rewardsAmount)
        internal
    {
        // Reduce Gas
        uint256 totalRewardsForNode = totalRewardsForNodes[_token];
        // Reduce Gas
        uint256 totalRewardsForLP = totalRewardsForLPs[_token];
        // Updates rewards for nodes and LPs
        uint256 rewardsForNode = _rewardsAmount.mul(nodeRewardsRatio).div(100);
        uint256 rewardsForLP = _rewardsAmount.sub(rewardsForNode);
        totalRewardsForNodes[_token] = totalRewardsForNode.add(rewardsForNode);
        totalRewardsForLPs[_token] = totalRewardsForLP.add(rewardsForLP);
    }

    function _updatePool(address _tokenA, address _tokenB)
        internal
        returns (uint256 newExchangeRate)
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

        // Logic: LPP = (float amount of BTC + float amount of WBTC + LP fees) / (LP supply - LP burned)
        uint256 burned = IBurnableToken(lpToken).balanceOf(address(burner));
        uint256 totalLPs = IBurnableToken(lpToken).totalSupply().sub(burned);
        // LP decimals == 18 so,
        newExchangeRate = totalReserved.mul(1e8).div(totalLPs);
        currentExchangeRate[_tokenB] = newExchangeRate;
        return newExchangeRate;
    }

    function getCurrentPriceLP(address _token) internal view returns (uint256) {
        return currentExchangeRate[_token];
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
