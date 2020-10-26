pragma solidity =0.7.0;

import "./interfaces/IBurnableToken.sol";

contract Burner {
    address public token;

    constructor(address _token) public {
        token = _token;
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
