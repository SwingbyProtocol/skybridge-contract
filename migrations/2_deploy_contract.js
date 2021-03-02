const SwapContract = artifacts.require("SwapContract");
const LPToken = artifacts.require("LPToken");

let WBTC_ADDR = "0x449268b65BAf7251f83fd0a4b182DbC4C20985Fd"

module.exports = async function (deployer, net) {
  if (net == "development") {
    return
  }
  if (net == "goerli") {
    WBTC_ADDR = "0xEb47a21C1fC00D1E863019906df1771b80DBE182"
  }
  if (net == "mainnet") {
    WBTC_ADDR = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
  }
  await deployer.deploy(LPToken)
  const lpToken = await LPToken.deployed()
  await deployer.deploy(SwapContract, lpToken.address, WBTC_ADDR, 0);
  const swapContract = await SwapContract.deployed()
  await lpToken.transferOwnership(swapContract.address)
};
