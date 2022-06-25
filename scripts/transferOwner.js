
const SwingbyToken = artifacts.require("SwingbyToken");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const token = await SwingbyToken.at("0xbd6c7cdf51d143b62bb5e5141313a6d41ea3dd54")
        console.log("token:", token.address)
        await token.transferOwnership(process.env.OWNER)
    } catch (err) {
        console.log(err)
    }
    done()
};
