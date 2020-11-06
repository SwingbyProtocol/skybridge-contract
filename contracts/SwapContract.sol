pragma solidity =0.7.0;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/ISwapContract.sol";
import "./Ownable.sol";
import "./SafeMath.sol";
import "./Burner.sol";

contract SwapContract is ISwapContract, Ownable {
    using SafeMath for uint256;

    address private lpToken;
    // Token address -> amount
    mapping(address => uint256) private totalRewardsForLPs; // 0x0 space is for node operators
    mapping(address => uint256) private totalRewardsForNodes; // 0x0 space is for node operators
    mapping(address => uint256) private floatAmountOfToken;
    mapping(address => uint256) private currentExchangeRate;
    Burner public burner;
    uint256 public nodeRewardsRatio;
    address[] public nodes;

    constructor(address _lpToken) public {
        burner = new Burner(_lpToken);
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

    function mintLPToken(address _dist, uint256 _amount)
        public
        onlyOwner
        returns (bool)
    {
        return IBurnableToken(lpToken).mint(_dist, _amount);
    }

    function addFloatForBTCToken(address _token, uint256 _amount)
        public
        override
        returns (bool)
    {
        // Transfer tokens to this contract from users.
        IERC20(_token).transferFrom(_msgSender(), address(this), _amount);
        // Update float amount
        floatAmountOfToken[_token] = floatAmountOfToken[_token].add(_amount);
        // Update LP token price
        _updatePool(_token);
        return true;
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
        // Update LP token price
        uint256 newExchangeRate = _updatePool(_token);
        uint256 floatAmount = _amountOfLPToken.mul(newExchangeRate).div(1e18);
        floatAmountOfToken[_token] = floatAmountOfToken[_token].sub(
            floatAmount
        );
        require(IERC20(_token).transfer(_msgSender(), floatAmount));
        return true;
    }

    function distributeNodeRewards(address _token)
        public
        override
        returns (bool)
    {
        // Reduce Gas
        uint256 totalRewardsForNode = totalRewardsForNodes[_token];
        require(totalRewardsForNode > 0, "totalRewardsForNode amount is 0");
        for (uint256 i = 0; i < nodes.length; i++) {
            require(
                IERC20(_token).transfer(
                    nodes[i],
                    totalRewardsForNode / nodes.length
                )
            );
        }
        // Zerolize for storage, gas refunded.
        totalRewardsForNodes[_token] = 0;
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

    function _updatePool(address _token)
        internal
        returns (uint256 newExchangeRate)
    {
        // for reduce gas cost.
        uint256 totalReward = totalRewardsForLPs[_token];
        uint256 floatAmount = floatAmountOfToken[_token];
        uint256 newReservedBTCTokens = floatAmount.add(totalReward);
        // WIP: have to support multiple coins
        // logic: rate = (fess + total float amount) / total active amount of LP token
        uint256 totalActiveLPTokens = IBurnableToken(lpToken).totalSupply().sub(
            IBurnableToken(lpToken).balanceOf(address(burner))
        );
        newExchangeRate = newReservedBTCTokens.mul(100).div(
            totalActiveLPTokens
        );
        currentExchangeRate[_token] = newExchangeRate;
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
