pragma solidity =0.7.0;

import "./interfaces/ISwapContract.sol";
import "./interfaces/IBurnableToken.sol";

contract Burner {
    address public token;
    address public swap;

    constructor(address _token, address _swap) public {
        token = _token;
        swap = _swap;
    }

    function burn() external {
        uint256 amount = IBurnableToken(token).balanceOf(address(this));
        IBurnableToken(token).burn(amount);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
