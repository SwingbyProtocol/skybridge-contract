pragma solidity =0.7.0;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/ISwapContract.sol";
import "./Ownable.sol";
import "./SafeMath.sol";
import "./Burner.sol";

contract SwapContract is ISwapContract, Ownable {
    using SafeMath for uint256;

    address private _lpToken;
    // Token address -> amount
    mapping(address => uint256) _collectedFeesOfToken;
    mapping(address => uint256) _floatAmountOfToken;
    mapping(address => uint256) _currentExchangeRate;
    Burner private _burner;

    event RedeemWithBurnLPtoken(address indexed sender, uint256 amount);

    constructor(address _lpToken) public {
        _burner = new Burner(_lpToken, address(this));
    }

    /**
     * Transfer part
     */

    function singleTransferERC20(
        address _token,
        address _to,
        uint256 _amount
    ) public onlyOwner {
        require(IERC20(_token).transfer(_to, _amount));
    }

    function multiTransferERC20TightlyPacked(
        address _token,
        bytes32[] memory _addressesAndAmounts,
        uint8 _inputDecimals
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
    }

    function multiTransferERC20(
        address token,
        address[] memory _contributors,
        uint256[] memory _amounts,
        uint256 _feesRatioForLP
    ) public override onlyOwner returns (bool) {
        require(
            _contributors.length == _amounts.length,
            "Input length is mismatch"
        );
        for (uint256 i = 0; i < _contributors.length; i++) {
            require(IERC20(token).transfer(_contributors[i], _amounts[i]));
        }
    }

    /**
     * Float part
     */

    function mintLPToken(address _dist, uint256 _amount)
        public
        onlyOwner
        returns (bool)
    {
        return IBurnableToken(_lpToken).mint(_dist, _amount);
    }

    function addFloatForBTCToken(address _token, uint256 _amount) public {
        IERC20(_token).transferFrom(_msgSender(), address(this), _amount);
        // Update float amount
        _floatAmountOfToken[_token] = _floatAmountOfToken[_token].add(_amount);
        // Update LP token price
        _updatePool(_token);
    }

    function redeemFloatForBTCToken(address _token, uint256 _amountOfLPToken)
        public
    {
        IBurnableToken(_lpToken).transferFrom(
            _msgSender(),
            address(this),
            _amountOfLPToken
        );
        IBurnableToken(_lpToken).burn(_amountOfLPToken);
        // Update LP token price
        uint256 newExchangeRate = _updatePool(_token);
        uint256 floatAmount = _amountOfLPToken.mul(newExchangeRate).div(1e18);
        _floatAmountOfToken[_token] = _floatAmountOfToken[_token].sub(
            floatAmount
        );
        IERC20(_token).transfer(_msgSender(), floatAmount);
    }

    function distributeNodeRewards() public override {}

    function _burnLPToken(uint256 _amount) internal returns (bool) {
        return IBurnableToken(_lpToken).burn(_amount);
    }

    function _updatePool(address _token)
        internal
        returns (uint256 newExchangeRate)
    {
        // for reduce gas cost.
        uint256 collectedFess = _collectedFeesOfToken[_token];
        uint256 floatAmountOfToken = _floatAmountOfToken[_token];
        uint256 newQuantityBTCTokens = floatAmountOfToken.add(collectedFess);
        // WIP: have to support multiple coins
        // logic: rate = (fess + total float amount) / total amount of LP token
        uint256 totalActiveLPTokens = IBurnableToken(_lpToken)
            .totalSupply()
            .sub(IBurnableToken(_lpToken).balanceOf(address(_burner)));
        newExchangeRate = newQuantityBTCTokens.div(totalActiveLPTokens);
        _currentExchangeRate[_token] = newExchangeRate;
        return newExchangeRate;
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
