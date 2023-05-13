const { ethers } = require("hardhat");

const _owner = '0x64496f51779e400C5E955228E56fA41563Fb4dd8';
const _sbBTC = '0x1B5B6dF2C72D7c406df1C30E640df8dBaE57d21d';
const _barn = '0x009cc14ce70b2e667984c2276490d56ae3234c43';
const _swap = '0xe8d45281d7BD836b30F9B371001676d1ed59465D'

async function main() {
    const sbBTCFactory = await ethers.getContractFactory("sbBTCPool");

    const [owner] = await ethers.getSigners();
    console.log(owner.address)
    const sbBTCPool = await sbBTCFactory.deploy();
    console.log(`sbBTCPool deployed at: ${sbBTCPool.address}`);

    // await sbBTCPool.setBarnAndSwap(_barn, _swap);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });