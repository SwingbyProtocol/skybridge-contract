pragma solidity =0.7.0;

import "./interfaces/IBurnableToken.sol";

contract Burner {
    function burn(address _token) external {
        uint256 amount = IBurnableToken(_token).balanceOf(address(this));
        IBurnableToken(_token).burn(amount);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
