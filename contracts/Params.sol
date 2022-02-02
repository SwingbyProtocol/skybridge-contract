// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;
pragma experimental ABIEncoderV2;
import "./interfaces/IParams.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Params is Ownable, IParams {
    uint8 public nodeRewardsRatio;
    uint8 public depositFeesBPS;
    uint8 public withdrawalFeeBPS;

    constructor() {
        // Initialize nodeRewardsRatio
        nodeRewardsRatio = 66;
        // Initialize withdrawalFeeBPS
        withdrawalFeeBPS = 20;
        // Initialize depositFeesBPS
        depositFeesBPS = 0;
    }

    function setNodeRewardsRatio(uint8 _nodeRewardsRatio) external onlyOwner {
        require(
            _nodeRewardsRatio >= 0 && _nodeRewardsRatio <= 100,
            "_nodeRewardsRatio is not valid"
        );
        nodeRewardsRatio = _nodeRewardsRatio;
    }

    function setWithdrawalFeeBPS(uint8 _withdrawalFeeBPS) external onlyOwner {
        require(
            _withdrawalFeeBPS >= 0 && _withdrawalFeeBPS <= 100,
            "_withdrawalFeeBPS is invalid"
        );
        withdrawalFeeBPS = _withdrawalFeeBPS;
    }

    function setDepositFeesBPS(uint8 _depositFeesBPS) external onlyOwner {
        require(
            _depositFeesBPS >= 0 && _depositFeesBPS <= 100,
            "_depositFeesBPS is invalid"
        );
        depositFeesBPS = _depositFeesBPS;
    }
}
