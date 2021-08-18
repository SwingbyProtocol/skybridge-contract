const { ethers } = require("hardhat");

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
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 8
async function main() {
    const [deployer] = await ethers.getSigners();
    let BTCT_ADDR = "0x449268b65BAf7251f83fd0a4b182DbC4C20985Fd"

    console.log("Deploying contracts with the account:", deployer.address);

    console.log("Account balance:", (await deployer.getBalance()).toString());

    const LPTokenFactory = await ethers.getContractFactory("LPToken");
    const lpToken = await LPTokenFactory.deploy(TOKEN_DECIMALS);

    await lpToken.deployed();
    console.log("LPToken address:", lpToken.address);

    const SwapContractFactory = await ethers.getContractFactory("SwapContract");
    const swap = await SwapContractFactory.deploy(lpToken.address, BTCT_ADDR, 0);

    await swap.deployed();
    console.log("SwapContract address:", swap.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });