const SwapContract = artifacts.require("SwapContract");
const LPToken = artifacts.require("LPToken");
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 8

module.exports = async function (deployer, net) {
  if (net == "development") {
    return
  }
  let BTCT_ADDR = "0x449268b65BAf7251f83fd0a4b182DbC4C20985Fd"
  if (net == "goerli") {
    BTCT_ADDR = "0xEb47a21C1fC00D1E863019906df1771b80DBE182"
  }
  if (net == "mainnet") {
    BTCT_ADDR = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
  }
  if (net == "bsc_testnet") {
    BTCT_ADDR = "0xa88921dc290f888b5ee574cf2cd1599f412f1534"
  }
  if (net == "bsc_mainnet") {
    BTCT_ADDR = "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c"
  }
  await deployer.deploy(LPToken, TOKEN_DECIMALS)
  const lpToken = await LPToken.deployed()
  await deployer.deploy(SwapContract, lpToken.address, BTCT_ADDR, 0);
  const swapContract = await SwapContract.deployed()
  await lpToken.transferOwnership(swapContract.address)
};
