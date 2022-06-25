
const { BN, constants } = require('@openzeppelin/test-helpers');
const { ZERO_ADDRESS } = constants
const SwapContract = artifacts.require("SwapContract");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const swap = await SwapContract.deployed()
        // send 0.1 BTC as float
        let amount = new BN(131122121817231).mul(new BN(10).pow(new BN(1)))
        let to = process.env.TO
        let float = "0x" + web3.utils.padLeft(amount.toString('hex') + to.slice(2), 64)
        let isZerofees = false
        let txid = "0x1c12143203a48f42cdf7b1acee5b1b1c1fedc144cb309a3bf5edbffafb0ad734"
        let btct = await swap.BTCT_ADDR()
        await swap.recordIncomingFloat(btct.toString(), float, isZerofees, txid)
    } catch (err) {
        console.log(err)
    }
    done()
};
