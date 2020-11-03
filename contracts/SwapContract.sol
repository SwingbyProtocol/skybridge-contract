pragma solidity =0.7.0;

import "./interfaces/IBurnableToken.sol";
import "./interfaces/IERC20.sol";
import "./Ownable.sol";
import "./SafeMath.sol";
import "./interfaces/ISwapContract.sol";

contract SwapContract is ISwapContract, Ownable {
    using SafeMath for uint256;

    address private _lpToken;
    // Token address -> amount
    mapping(address => uint256) _collectedFeesOfTokens;
    mapping(address => uint256) _collectedFloatAmountOfTokens;
    mapping(address => uint256) _exchangeRate;

    event RedeemWithBurnLPtoken(address indexed sender, uint256 amount);

    constructor() public {}

    /**
     * Transfer part
     */

    function singleTransferERC20(
        address _token,
        address _to,
        uint256 _amount
    ) public onlyOwner {
        require(_token != address(0));
        IERC20(_token).transfer(_to, _amount);
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
        uint256[] memory _amounts
    ) public override onlyOwner returns (bool) {
        require(_contributors.length == _amounts.length, "length is mismatch");
        for (uint256 i = 0; i < _contributors.length; i++) {
            require(IERC20(token).transfer(_contributors[i], _amounts[i]));
        }
    }

    /**
     * Float part
     */

    function mintLPToken(address _dist, uint256 _amount) public onlyOwner {
        IBurnableToken(_lpToken).mint(_dist, _amount);
    }

    function getSupplyOfLPToken() public returns (uint256) {
        return IBurnableToken(_lpToken).totalSupply();
    }

    function addFloatForBTCTokens(address _token, uint256 _amount) public {
        IERC20(_token).transferFrom(_msgSender(), address(this), _amount);
        // Update float amount
        _collectedFloatAmountOfTokens[_token] = _collectedFloatAmountOfTokens[_token]
            .add(_amount);

        // Adding BTC liquidity does not change the LP exchange rate
        _updatePool(_token);
    }

    function redeemFloatForTokensAndBTC(
        address _token,
        uint256 _amountOfLPToken
    ) public onlyOwner {
        IBurnableToken(_lpToken).transferFrom(
            _msgSender(),
            address(this),
            _amountOfLPToken
        );
        IBurnableToken(_lpToken).burn(_amountOfLPToken);
        if (_token == address(0)) {
            // operates BTC-LP -> BTC transfer
        } else {
            // operates BTC-LP -> BTC stable tokens tranfer
            // update exchange rate.
            uint256 newExchangeRate = _updatePool(_token);
            uint256 floatAmount = _amountOfLPToken.mul(newExchangeRate);
            IERC20(_token).transfer(_msgSender(), floatAmount);
        }
    }

    function _updatePool(address _token)
        internal
        returns (uint256 newExchangeRate)
    {
        // for reduce gas cost.
        uint256 collectedFess = _collectedFeesOfTokens[_token];
        uint256 totalIssued = _collectedFloatAmountOfTokens[_token];
        // for getting LP token supply
        uint256 balanceOfLPToken = getSupplyOfLPToken();
        uint256 newQuantityBTCTokens = totalIssued.add(collectedFess);
        // WIP: have to support multiple coins
        newExchangeRate = newQuantityBTCTokens.div(balanceOfLPToken);
        _exchangeRate[_token] = newExchangeRate;
        return newExchangeRate;
    }

    function distributeNodeRewards() public override {}

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
