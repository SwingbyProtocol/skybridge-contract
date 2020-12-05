
const LPToken = artifacts.require("LPToken");
const SwapContract = artifacts.require("SwapContract");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const lpToken = await LPToken.deployed()
        const swap = await SwapContract.deployed()
        await lpToken.transferOwnership(swap.address)
    } catch (err) {
        console.log(err)
    }
    done()
};
