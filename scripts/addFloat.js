
const { BN, constants } = require('@openzeppelin/test-helpers');
const { ZERO_ADDRESS } = constants
const SwapContract = artifacts.require("SwapContract");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const swap = await SwapContract.deployed()
        // send 0.1 BTC as float
        let amount = new BN(1).mul(new BN(10).pow(new BN(7)))
        let to = process.env.TO
        let float = "0x" + web3.utils.padLeft(amount.toString('hex') + to.slice(2), 64)
        let txid = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        await swap.recordIncomingFloat(ZERO_ADDRESS, float, txid)
        await swap.issueLPTokensForFloat(txid)
    } catch (err) {
        console.log(err)
    }
    done()
};
