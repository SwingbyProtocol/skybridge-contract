const { ZERO_ADDRESS } = require("@openzeppelin/test-helpers/src/constants");
const { ethers } = require("hardhat");
BN = require('bn.js')
/*********************************************************************
 * 
 * Boiler Plate script to deploy
 * 
 * July 28, 2021 
 * 
 * Deploy with:
 * npx hardhat run ./scripts/deploy.js --network development || npx hardhat run ./scripts/deploy.js --network hardhat
 * There are still some errors with the above
 */
async function main() {
    const swapContrat = '0x9e6BA6e811665849f03f56C1f22a8894AEbb3993'
    const to = '0x63d9f6A25ddD2c586F4441065Ce7C8412fbBB91e'
    const SwapContractFactory = await ethers.getContractFactory("SwapContract");
    const swap = await SwapContractFactory.attach(swapContrat);

    // send 1 BTC & 1WBTC as float
    let amount = new BN(1).mul(new BN(10).pow(new BN(8)))
    let float = "0x" + web3.utils.padLeft(amount.toString('hex') + to.slice(2), 64)
    let txid_1 = "0x1c12143203a48f42cdf7b1acee5b1b1c1fedc144cb309a3bf5edbffafc0ad734"
    let txid_2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
    let btct = await swap.BTCT_ADDR()
    await swap.recordIncomingFloat(btct.toString(), float, txid_1)
    await swap.recordIncomingFloat(ZERO_ADDRESS, float, txid_2)
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });