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
    [...addrs] = await ethers.getSigners();
    //console.log(addrs.map((s) => s.address))
    const swapContrat = '0x9e6BA6e811665849f03f56C1f22a8894AEbb3993'
    const sbBTC_Pool = '0xD60126017fDf906668Cfe5327c566C65e7f061bA'
    const barnContract = '0x009cc14ce70b2E667984C2276490d56ae3234c43'
    const SwapContractFactory = await ethers.getContractFactory("SwapContract");
    const sbBTCFactory = await ethers.getContractFactory("sbBTCPool");
    const barnFactory = await ethers.getContractFactory("LPToken");
    const swap = await SwapContractFactory.attach(swapContrat);
    const sbBTCP = await sbBTCFactory.attach(sbBTC_Pool)
    const barn = await barnFactory.attach(barnContract)

    nodes = []
    var skipped = []

    //console.log(addrs)

    for (i = 0; i < 10; i++) {
        nodes.push({
            address: addrs[i].address,
            claim: sbBTCP.connect(addrs[i]).callStatic.claim(),
            barn: barn.balanceOf(addrs[i].address)
        })
    }
    var sum = 0;
    const claimed = await Promise.allSettled(nodes.map((s) => s.claim))
    const barnBalance = await Promise.allSettled(nodes.map((s) => s.barn))
    const nodeList = []
    claimed.forEach((s, i) => {
        if (s.status == "fulfilled") {
            nodeList.push({ address: nodes[i].address, amount: (s.value / 1e8), barnBalance: barnBalance[i].value / 1e18 })
            sum += s.value / 1e8
        } else {
            skipped.push(nodes[i].address)
        }
    })
    const sorted = nodeList.sort(function (a, b) {
        return b.amount - a.amount
    })
    console.log("sum -------------------------------------- ", sum, "time:", new Date(), "contract:", sbBTCP.address)

    sorted.forEach(async (s) => {
        console.log(`address: ${s.address} amount: ${s.amount.toFixed(8)}, bal: ${s.barnBalance.toFixed(8)}`)
    })
    console.log("skipped: ", skipped)
}
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
