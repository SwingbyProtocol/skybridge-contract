// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;

interface ISwapRewards {
    function pullRewards(address _dest, address[] memory _receiver, uint256[] memory _swapped) external returns (uint256);
}
