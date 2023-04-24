// SPDX-License-Identifier: AGPL-3.0

pragma solidity >=0.6.0 <=0.8.19;

interface IWETH {
    function deposit() external payable;

    function transfer(address to, uint256 value) external returns (bool);

    function withdraw(uint256) external;

    function approve(address spender, uint256 amount) external returns (bool);
}
