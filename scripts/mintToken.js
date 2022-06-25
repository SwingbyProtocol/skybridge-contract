const { BN } = require('@openzeppelin/test-helpers');

const SwingbyToken = artifacts.require("SwingbyToken");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const token = await SwingbyToken.at("0x423fa1c32541b3d9fc3690e15392185753d44fad")
        console.log("token:", token.address)
        let amount = new BN(1000000).mul(new BN(10).pow(new BN(18)))
        await token.mint(process.env.OWNER, amount)
    } catch (err) {
        console.log(err)
    }
    done()
};
