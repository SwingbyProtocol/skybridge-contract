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
    [sender, receiver, dao, ...addrs] = await ethers.getSigners();
    console.log(sender.address)

    const swapContrat = '0x9e6BA6e811665849f03f56C1f22a8894AEbb3993' // goerli
    const to = '0x63d9f6A25ddD2c586F4441065Ce7C8412fbBB91e'
    const SwapContractFactory = await ethers.getContractFactory("SwapContract");
    const swap = await SwapContractFactory.attach(swapContrat);

    let nodes = []
    let isRemoved = []
    let churnedInCount = 25
    let tssThreshold = 16
    let nodeRewardsRatio = 66
    // for 17 nodes adding
    nodes.push(sender.address)
    isRemoved.push(false)
    nodes.push(receiver.address)
    isRemoved.push(false)
    nodes.push(dao.address)
    isRemoved.push(false)

    for (i = 0; i < 8; i++) {
        nodes.push(addrs[i].address)
        isRemoved.push(false)
    }

    const tx1 = await swap.churn(
        sender.address,
        nodes,
        isRemoved,
        churnedInCount,
        tssThreshold,
        {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })

    const receipt = await tx1.wait()
    console.log(receipt)
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });