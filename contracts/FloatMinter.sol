pragma solidity =0.7.0;

import "./ERC20.sol";
import "./Ownable.sol";
import "./SafeMath.sol";

// The TSS address should have ownership of this contract.
contract FloatMinter is ERC20Token {
    using SafeMath for uint256;
    address public swap;

    function initialize(address _swap) public {
        _initialize("BTC LP Token", "BTC-LP", 18, 0, true);
        swap = _swap;
    }

    // Accepts WBTC/tBTC/renBTC and redirect to swap contract
    function addFloat(address _token, uint256 _amount) public {
        IERC20(_token).transferFrom(msg.sender, swap, _amount);
        _mint(msg.sender, _amount);
    }

    // Accepts BTC-LP burn
    function redeemFloat(address _receiveToken, uint256 _amountToBurn) public {
        _burn(msg.sender, _amountToBurn);
        require(IERC20(_receiveToken).transfer(msg.sender, _amountToBurn));
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
