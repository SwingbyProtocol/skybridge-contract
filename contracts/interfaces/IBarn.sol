// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.0 <=0.8.9;
pragma experimental ABIEncoderV2;

interface IBarn {
    // balanceOf returns the current BOND balance of a user (bonus not included)
    function balanceOf(address user) external view returns (uint256);
}
