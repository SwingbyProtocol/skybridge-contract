const { ethers } = require("hardhat");
BN = require('bn.js')
const { BigNumber } = require('ethers');

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

    // send 0.1 BTC as float
    minerFees = new BigNumber.from(30000)

    //prep to collectSwapFeesForBTC
    let swapFeesBPS = new BigNumber.from(20);
    //let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))

    let swapAmount = new BigNumber.from(1000000) //.01 BTC
    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
    //let rewardsAmount = swapFees.sub(minerFees)
    let incomingAmount = swapAmount.add(swapFees) //100200000 sats = 1.002 BTC                
    let spenders = [to]
    let amounts = [incomingAmount]

    const tx = await swap.collectSwapFeesForBTC(
        swapAmount,
        minerFees,
        swapFees,
        spenders,
        amounts,
    )
    console.log("txhash:", tx);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });