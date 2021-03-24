
const { BN, constants } = require('@openzeppelin/test-helpers');
const { web3 } = require('@openzeppelin/test-helpers/src/setup');
const { ZERO_ADDRESS } = constants
const SwapContract = artifacts.require("SwapContract");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const swap = await SwapContract.at(process.env.WALLET)
        const res = await swap.getActiveNodes.call({ gas: 500000, gasPrice: 1 * 10 ** 9 })
        const val = new BN(1).mul(new BN(10).pow(new BN(8)))
        const result = res.map((stake) => {
            const amount = web3.utils.hexToNumber(stake.slice(0, 26))
            return {
                amount: amount / 1e8,
                address: "0x" + stake.slice(26, stake.length)
            }
        }).sort(function (a, b) {
            return b.amount - a.amount;
        });
        console.log(result)
    } catch (err) {
        console.log(err)
    }
    done()
};
