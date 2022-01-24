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

    constructor(
        address _owner,
        address _swingby,
        address _swap
    ) {
        require(_owner != address(0), "owner address is not be 0x0");
        require(_swingby != address(0), "swingby address must not be 0x0");

        transferOwnership(_owner);
        rewardToken = IERC20(_swingby);
        swapContract = IERC20(_swap);
    }

    function setSwap(address _swap, uint256 _newRebateRate) public {
        require(swapContract != address(0), "barn address must not be 0x0");
        require(msg.sender == owner(), "!owner");
        swapContract = ISwapContract(_swap);
        rebateRate = _newRebateRate;
    }

    // pullRewards transfers the funds to the user
    function pullRewards(
        address _dest,
        address[] memory _receiver,
        uint256[] memory _swapped
    ) public returns (uint256) {
        require(
            msg.sender == address(swapContract),
            "caller is not swap contact"
        );
        address tokenB = swapContract.BTCT_ADDR();
        (balA, balB) = swapContract.getFloatReserve(address(0), tokenB);
        uint256 threshold = balA.add(balB).mul(2).div(3);
        if (
            (_dest == tokenB && balA <= threshold) ||
            (_dest == address(0) && balB <= threshold)
        ) {
            for (uint256 i = 0; i < _receiver.length; i++) {
                rewardToken.transfer(
                    _receiver[i],
                    _swapped[i].mul(rebateRate).div(100).mul(1e10)
                ); // decimals == 18 for payout
            }
        }
        return rewards;
    }
}
