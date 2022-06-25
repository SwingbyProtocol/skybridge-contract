const SwingbyToken = artifacts.require("SwingbyToken");

module.exports = async function (done) {
    // await deployer.deploy(LPToken)
    try {
        const st = await SwingbyToken.new()
        console.log(st.address)
    } catch (err) {
        console.log(err)
    }
    done()
};
