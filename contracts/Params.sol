// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;
pragma experimental ABIEncoderV2;
import "./interfaces/IParams.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Params is Ownable, IParams{
    uint8 public nodeRewardsRatio;
    uint8 public depositFeesBPS;
    uint8 public withdrawalFeeBPS;

    constructor(){
        // Initialize nodeRewardsRatio
        nodeRewardsRatio = 66;
        // Initialize withdrawalFeeBPS
        withdrawalFeeBPS = 20;
        // Initialize depositFeesBPS
        depositFeesBPS = 0;
    }

    function setNodeRewardsRatio(uint8 _new) external onlyOwner {
        nodeRewardsRatio = _new;
    }

    function setWithdrawalFeeBPS(uint8 _new) external onlyOwner {
        withdrawalFeeBPS = _new;
    }

    function setDepositFeesBPS(uint8 _new) external onlyOwner {
        depositFeesBPS = _new;
    }
}
