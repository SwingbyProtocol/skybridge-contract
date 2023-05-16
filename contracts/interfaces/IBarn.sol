// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.0 <=0.8.19;
pragma experimental ABIEncoderV2;

interface IBarn {
    // balanceOf returns the current BOND balance of a user (bonus not included)
    function balanceOf(address user) external view returns (uint256);

    // balanceAtTs returns the amount of BOND that the user currently staked (bonus NOT included)
    function balanceAtTs(
        address user,
        uint256 timestamp
    ) external view returns (uint256);
}
