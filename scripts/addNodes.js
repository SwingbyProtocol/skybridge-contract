
const { BN, constants } = require('@openzeppelin/test-helpers');
const { ZERO_ADDRESS } = constants
const SwapContract = artifacts.require("SwapContract");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const swap = await SwapContract.deployed()
        // send 0.1 BTC as float
        let rewardAddressAndAmounts = []
        let isRemoved = []
        let churnedInCount = 25
        let tssThreshold = 16
        let nodeRewardsRatio = 66
        let withdrawFess = 20
        let accounts = [
            "0xde2C9714845B18706c350629787621f8e03Cb8Bf",
            "0x219B35fF0528fe11e55F68F9a63E0b1392B0a299"
        ]
        for (i = 0; i < 2; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx1 = await swap.churn(accounts[0], rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, withdrawFess)
    } catch (err) {
        console.log(err)
    }
    done()
};
