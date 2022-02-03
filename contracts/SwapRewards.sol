// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.0 <=0.8.9;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/ISwapContract.sol";

contract SwapRewards is Ownable {
    using SafeMath for uint256;

    IERC20 public immutable rewardToken;
    ISwapContract public swapContract;
    uint256 public rebateRate = 30;
    uint256 public thresholdRatio = 55; // diff is over 10%

    event Paid(address to, uint256 amount, uint256 rebate);

    constructor(
        address _owner,
        address _swingby,
        address _swap
    ) {
        require(_owner != address(0), "owner address is not be 0x0");
        require(_swingby != address(0), "swingby address must not be 0x0");

        transferOwnership(_owner);
        rewardToken = IERC20(_swingby);
        swapContract = ISwapContract(_swap);
    }

    function setSwap(
        address _swap,
        uint256 _newRebateRate,
        uint256 _thresholdRatio
    ) external {
        require(msg.sender == owner(), "!owner");
        require(
            _newRebateRate >= 0 && _newRebateRate <= 100,
            "_newRebateRate is not valid"
        );
        require(
            _thresholdRatio >= 20 && _thresholdRatio <= 100,
            "_thresholdRatio is not valid"
        );
        swapContract = ISwapContract(_swap);
        rebateRate = _newRebateRate;
        thresholdRatio = _thresholdRatio;
    }

    // pullRewards transfers the funds to the user
    function pullRewards(
        address _dest,
        address _receiver,
        uint256 _swapped
    ) external {
        require(
            msg.sender == address(swapContract),
            "caller is not swap contact"
        );
        address tokenB = swapContract.BTCT_ADDR();
        (uint256 balA, uint256 balB) = swapContract.getFloatReserve(
            address(0),
            tokenB
        );
        uint256 threshold = balA.add(balB).mul(thresholdRatio).div(100);
        if (
            (_dest == tokenB && balB >= threshold) ||
            (_dest == address(0) && balA >= threshold)
        ) {
            rewardToken.transfer(
                _receiver,
                _swapped.mul(rebateRate).div(100).mul(1e10)
            ); // decimals == 18 for payout
            emit Paid(
                _receiver,
                _swapped,
                _swapped.mul(rebateRate).div(100).mul(1e10)
            );
        }
    }

    // pullRewardsMulti transfers the funds to the user
    function pullRewardsMulti(
        address _dest,
        address[] memory _receiver,
        uint256[] memory _swapped
    ) external returns (bool){
        require(
            msg.sender == address(swapContract),
            "caller is not swap contact"
        );
        require(_receiver.length == _swapped.length, "array size is not match");
        address tokenB = swapContract.BTCT_ADDR();
        (uint256 balA, uint256 balB) = swapContract.getFloatReserve(
            address(0),
            tokenB
        );
        uint256 threshold = balA.add(balB).mul(thresholdRatio).div(100);
        if (
            (_dest == tokenB && balB >= threshold) ||
            (_dest == address(0) && balA >= threshold)
        ) {
            for (uint256 i = 0; i < _receiver.length; i++) {
                rewardToken.transfer(
                    _receiver[i],
                    _swapped[i].mul(rebateRate).div(100).mul(1e10)
                ); // decimals == 18 for payout
                emit Paid(
                    _receiver[i],
                    _swapped[i],
                    _swapped[i].mul(rebateRate).div(100).mul(1e10)
                );
            }
        }
        return true;
    }
}
