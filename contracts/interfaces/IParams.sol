// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;
pragma experimental ABIEncoderV2;

interface IParams {


    function setNodeRewardsRatio(uint8 _new) external;

    function setWithdrawalFeeBPS(uint8 _new) external;

    function setDepositFeesBPS(uint8 _new) external;

}