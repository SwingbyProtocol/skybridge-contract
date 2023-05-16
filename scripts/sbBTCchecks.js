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
    const swapContrat = '0x9e6BA6e811665849f03f56C1f22a8894AEbb3993'
    const sbBTC_Pool = '0xec2946aD323f0879269910cbBB5420E8CD578a30'
    const SwapContractFactory = await ethers.getContractFactory("SwapContract");
    const sbBTCFactory = await ethers.getContractFactory("sbBTCPool");
    const swap = await SwapContractFactory.attach(swapContrat);
    const sbBTCP = sbBTCFactory.attach(sbBTC_Pool)

    //let nodes = await swap.getActiveNodes()

    array = []
    var skipped = []

    array.push(sbBTCP.connect(sender).callStatic.claim())
    array.push(sbBTCP.connect(receiver).callStatic.claim())
    array.push(sbBTCP.connect(dao).callStatic.claim())

    //console.log(addrs)

    for (i = 0; i < 10; i++) {
        array.push(sbBTCP.connect(addrs[i]).callStatic.claim())
    }
    var sum = 0;
    const result = await Promise.allSettled(array)
    var nodeList = []
    var skipped = []
    result.forEach((s, i) => {
        var address = ""
        switch (i) {
            case 0: address = sender.address;
                break;
            case 1: address = receiver.address
                break
            case 2: address = dao.address
                break;
            default: address = addrs[i-3].address
        }
        if (s.status == "fulfilled") {
            nodeList.push({ address: address, amount: (s.value / 1e8) })
            sum += s.value / 1e8
        } else {
            skipped.push(address)
        }
    })
    const sorted = nodeList.sort(function (a, b) {
        return b.amount - a.amount
    })
    console.log("sum -------------------------------------- ", sum, "time:", new Date(), "contract", sbBTCP.address,)
    sorted.forEach((s) => console.log(`address: ${s.address} amount: ${s.amount.toFixed(8)}`))
    //console.log("skipped: ", skipped)
}
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
