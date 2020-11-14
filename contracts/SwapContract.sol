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
    Burner public burner;
    uint8 churned_in_count;
    uint8 public nodeRewardsRatio;
    bytes32[] public nodes;

    struct Node {
        address addr;
        uint256 staked;
    }

    constructor(address _lpToken) public {
        burner = new Burner();
        lpToken = _lpToken;
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

    function mintLPToken(address _dist, uint256 _amountOfBTC)
        public
        override
        onlyOwner
        returns (bool)
    {
        // BTC float increased.
        floatAmountOfToken[address(0)] = floatAmountOfToken[address(0)].add(
            _amountOfBTC
        );
        // LP token price per BTC/WBTC changed.
        uint256 newPrice = _updatePool(address(0), WBTC_ADDR);
        // Mint LP tokens based on current LP price to users.
        IBurnableToken(lpToken).mint(
            _dist,
            _amountOfBTC.mul(100).div(newPrice)
        );
        return true;
    }

    function addFloatForBTCToken(address _token, uint256 _amount)
        public
        override
        returns (bool)
    {
        // _token should be WBTC
        require(_token == WBTC_ADDR);
        // Transfer tokens to this contract from users.
        IERC20(_token).transferFrom(_msgSender(), address(this), _amount);
        // Get exchange rate for BTC <> LP, And for reduece gas
        uint256 rate = currentExchangeRate[_token];
        // Mint LP tokens based on current LP price to users.
        IBurnableToken(lpToken).mint(_msgSender(), _amount.mul(100).div(rate));

        floatAmountOfToken[_token] = floatAmountOfToken[_token].add(_amount);
        return true;
    }

    function redeemFloatForBTC(uint256 _amountOfLPToken) public {
        IBurnableToken(lpToken).transferFrom(
            _msgSender(),
            address(this),
            _amountOfLPToken
        );
        IBurnableToken(lpToken).burn(_amountOfLPToken);
        // Changed LP token price per WBTC
        _updatePool(address(0), WBTC_ADDR);
        // Transfer BTC from TSS address.
    }

    function redeemFloatForBTCToken(address _token, uint256 _amountOfLPToken)
        public
        override
        returns (bool)
    {
        IBurnableToken(lpToken).transferFrom(
            _msgSender(),
            address(this),
            _amountOfLPToken
        );
        IBurnableToken(lpToken).burn(_amountOfLPToken);
        // Get exchange rate for BTC/LP, And for reduece gas
        uint256 rate = currentExchangeRate[_token];
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
            uint256 staked = uint256(uint96(bytes12(nodes[i])));
            totalStaked = totalStaked.add(staked);
        }
        for (uint256 i = 0; i < nodes.length; i++) {
            address to = address(uint160(uint256(nodes[i])));
            uint256 staked = uint256(uint96(bytes12(nodes[i])));
            IERC20(_token).transfer(
                to,
                totalRewardsForNode.mul(staked).div(totalStaked)
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
        uint8 _rewardsRate,
        uint8 _nodeRewardsRatio
    ) public override onlyOwner returns (bool) {
        _transferOwnership(_newOwner);
        nodes = _nodeRewardsAddressAndAmounts;
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
        totalRewardsForNodes[_token] = totalRewardsForNode.add(rewardsForNode);
        totalRewardsForLPs[_token] = totalRewardsForLP.add(
            _rewardsAmount.sub(rewardsForNode)
        );
    }

    function _updatePool(address _floatToken, address _feeToken)
        internal
        returns (uint256 newExchangeRate)
    {
        // for reduce gas cost.
        uint256 floatAmountOfFloatToken = floatAmountOfToken[_floatToken];
        uint256 floatAmountOfFeeToken = floatAmountOfToken[_feeToken];
        uint256 totalReward = totalRewardsForLPs[_feeToken];

        uint256 newReserved = floatAmountOfFloatToken
            .add(floatAmountOfFeeToken)
            .add(totalReward);
        // WIP: have to support multiple coins
        // Logic: rate = (fess + total float amount) / (LP supply - LP burned)
        uint256 burned = IBurnableToken(lpToken).balanceOf(address(burner));
        uint256 totalLPs = IBurnableToken(lpToken).totalSupply().sub(burned);
        newExchangeRate = newReserved.mul(100).div(totalLPs);
        currentExchangeRate[_feeToken] = newExchangeRate;
        return newExchangeRate;
    }

    function _burnLPToken(uint256 _amount) internal returns (bool) {
        return IBurnableToken(lpToken).burn(_amount);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
