
const LPToken = artifacts.require("LPToken");
const SwapContract = artifacts.require("SwapContract");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const swap = await SwapContract.deployed()
        await swap.transferOwnership(process.env.TSS)
    } catch (err) {
        console.log(err)
    }
    done()
};
