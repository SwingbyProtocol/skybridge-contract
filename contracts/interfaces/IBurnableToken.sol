pragma solidity =0.7.0;

import "./IERC20.sol";

interface IBurnableToken is IERC20 {
    function mint(uint256 amount) external returns (bool);

    function burn(uint256 amount) external returns (bool);
}
