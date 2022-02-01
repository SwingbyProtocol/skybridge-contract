// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;
pragma experimental ABIEncoderV2;

interface IParams {
    function nodeRewardsRatio() external returns (uint8);
    function depositFeesBPS() external view returns (uint8);
    function withdrawalFeeBPS() external view returns (uint8);

    function setNodeRewardsRatio(uint8 _new) external;

    function setWithdrawalFeeBPS(uint8 _new) external;

    function setDepositFeesBPS(uint8 _new) external;
}
