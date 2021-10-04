// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;

import "./IERC20.sol";

interface IBurnableToken is IERC20 {
    function mint(address target, uint256 amount) external returns (bool);

    function burn(uint256 amount) external returns (bool);

    function mintable() external returns (bool);
}
