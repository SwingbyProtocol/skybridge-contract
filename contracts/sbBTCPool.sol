// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.6.0 <=0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./interfaces/IBarn.sol";
import "./interfaces/ISwapContract.sol";

contract sbBTCPool is Ownable {
    using SafeMath for uint256;

    uint256 constant divisor = 10 ** 27;

    uint256 public balanceBefore;
    uint256 public currentMultiplier;
    uint256 public totalNodeStaked;

    mapping(address => uint256) public userMultiplier;
    mapping(address => uint256) public owed;

    IBarn public barn;
    IERC20 public rewardToken;
    ISwapContract public swapContract;

    event Claim(address indexed user, uint256 amount);

    // setBarn sets the address of the BarnBridge Barn into the state variable
    function setBarnAndSwap(address _barn, address _swap) public {
        require(_barn != address(0), "barn address must not be 0x0");
        require(_swap != address(0), "swap contract address must not be 0x0");
        require(msg.sender == owner(), "!owner");
        swapContract = ISwapContract(_swap);
        rewardToken = IERC20(swapContract.lpToken());
        barn = IBarn(_barn);
    }

    // update all node rewards.
    function updateAll(uint256 _timestamp) public {
        updateStakes(_timestamp);
        // Getting rewards
        ackFunds();
        // Update all distribution.
        address[] memory nodes = swapContract.getActiveNodes();
        for (uint256 i = 0; i < nodes.length; i++) {
            _updateOwed(nodes[i], _timestamp);
        }
    }

    // check all active nodes to calculate current stakes.
    function updateStakes(uint256 _timestamp) public {
        address[] memory nodes = swapContract.getActiveNodes();
        uint256 newTotalNodeStaked;
        for (uint256 i = 0; i < nodes.length; i++) {
            newTotalNodeStaked = newTotalNodeStaked.add(
                barn.balanceAtTs(nodes[i], _timestamp)
            );
            if (userMultiplier[nodes[i]] == 0) {
                userMultiplier[nodes[i]] = currentMultiplier;
            }
        }
        // only change when stakers had actions.
        if (totalNodeStaked != newTotalNodeStaked) {
            totalNodeStaked = newTotalNodeStaked;
        }
    }

    function resetUnstakedNode(address _node) public {
        require(!swapContract.isNodeStake(_node), "node is staker");
        userMultiplier[_node] = 0;
    }

    // claim calculates the currently owed reward and transfers the funds to the user
    function claim() public returns (uint256) {
        require(swapContract.isNodeStake(msg.sender), "caller is not node");

        updateStakes(block.timestamp);

        ackFunds();

        _updateOwed(msg.sender, block.timestamp);

        uint256 amount = owed[msg.sender];
        require(amount > 0, "nothing to claim");

        owed[msg.sender] = 0;

        rewardToken.transfer(msg.sender, amount);

        // acknowledge the amount that was transferred to the user
        ackFunds();

        emit Claim(msg.sender, amount);

        return amount;
    }

    // ackFunds checks the difference between the last known balance of `token` and the current one
    // if it goes up, the multiplier is re-calculated
    // if it goes down, it only updates the known balance
    function ackFunds() public {
        uint256 balanceNow = rewardToken.balanceOf(address(this));

        if (balanceNow == 0 || balanceNow <= balanceBefore) {
            balanceBefore = balanceNow;
            return;
        }
        if (totalNodeStaked == 0) {
            return;
        }

        uint256 diff = balanceNow.sub(balanceBefore);
        uint256 multiplier = currentMultiplier.add(
            diff.mul(divisor).div(totalNodeStaked)
        );

        balanceBefore = balanceNow;
        currentMultiplier = multiplier;
    }

    function emergencyWithdraw() public {
        require(msg.sender == owner(), "!owner");
        rewardToken.transfer(msg.sender, rewardToken.balanceOf(address(this)));
        swapContract = ISwapContract(address(0));
    }

    // _updateOwed calculates and updates the total amount that is owed to an user and updates the user's multiplier
    // to the current value
    // it automatically attempts to pull the token from the source and acknowledge the funds
    function _updateOwed(address user, uint256 timestamp) internal {
        uint256 reward = _userPendingReward(user, timestamp);
        owed[user] = owed[user].add(reward);
        userMultiplier[user] = currentMultiplier;
    }

    // _userPendingReward calculates the reward that should be based on the current multiplier / anything that's not included in the `owed[user]` value
    // it does not represent the entire reward that's due to the user unless added on top of `owed[user]`
    function _userPendingReward(
        address user,
        uint256 timestamp
    ) internal view returns (uint256) {
        uint256 multiplier = currentMultiplier.sub(userMultiplier[user]);
        return barn.balanceAtTs(user, timestamp).mul(multiplier).div(divisor);
    }
}
