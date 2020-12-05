const SwapContract = artifacts.require("SwapContract");
const LPToken = artifacts.require("LPToken");

const WBTC_ADDR = "Sample token address"

module.exports = async function (deployer, net) {
  if (net == "deployment") {
    return
  }
  await deployer.deploy(LPToken)
  const lpToken = await LPToken.deployed()
  await deployer.deploy(SwapContract, lpToken.address, WBTC_ADDR);
  const swap = await SwapContract.deployed()
  await lpToken.transferOwnership(swap.address)
};
