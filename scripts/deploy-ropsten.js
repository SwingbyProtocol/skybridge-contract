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
    //let BTCT_ADDR = "0x449268b65BAf7251f83fd0a4b182DbC4C20985Fd"
    //BTCT_ADDR = "0xbde8bb00a7ef67007a96945b3a3621177b615c44"
    //BTCT_ADDR = "0xaD6D458402F60fD3Bd25163575031ACDce07538D"//address for DAI as there is no liquidity for BTCt on Ropsten
    const BTCT_ADDR = "0x7cb2eac36b4bb7c36640f32e806d33e474d1d427"
    const wETH_ADDR = "0xc778417e063141139fce010982780140aa0cd5ab"
    const sbBTCPool = "0xf5329508Fdb96A1aada673dB6572109a228edB7e"
    const params = "0xad2CD8327BFE0E8C059EA5f789d064ba4C261BDb"
    const swapRewards = "0x53E79243D27DCd5E1070b4B71FDC42F8944a7848"

    console.log("Deploying contracts with the account:", deployer.address);

    console.log("Account balance:", (await deployer.getBalance()).toString());

    const LPTokenFactory = await ethers.getContractFactory("LPToken");
    const lpToken = await LPTokenFactory.deploy(TOKEN_DECIMALS);

    await lpToken.deployed();
    console.log("LPToken address:", lpToken.address);

    const SwapContractFactory = await ethers.getContractFactory("SwapContract");
    const swap = await SwapContractFactory.deploy(lpToken.address, BTCT_ADDR, wETH_ADDR, sbBTCPool, params, swapRewards, 0);

    await swap.deployed();
    console.log("SwapContract address:", swap.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });