
const { BN, constants } = require('@openzeppelin/test-helpers');
const { ZERO_ADDRESS } = constants
const SwapContract = artifacts.require("SwapContract");
const LPToken = artifacts.require("LPToken");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const swap = await SwapContract.deployed()
        // send 0.1 BTC as float
        let amount = new BN(3231122121817231).mul(new BN(10).pow(new BN(1)))
        //let amount = new BN(3933663654516931)
        let to = process.env.TO
        let AmountLP2 = "0x" + web3.utils.padLeft(amount.toString('hex') + to.slice(2), 64)
        let minerFees = 0
        let txid = "0x1c12143203aa1f42cdf1b1acee5b1b1c1fedc144cb309a3bf5edcffafb0ac731"
        let btct = await swap.BTCT_ADDR()
        let lptAddr = await swap.lpToken()

        const lpt = await LPToken.at(lptAddr)
        const deposit = await lpt.transfer(swap.address, amount)
        //console.log('max', amountMAXBTCfloatLP.toString())
        await swap.recordOutcomingFloat(ZERO_ADDRESS, AmountLP2, minerFees, deposit.tx)
    } catch (err) {
        console.log(err)
    }
    done()
};
