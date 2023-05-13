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
    const BTCT_ADDR = "0xEb47a21C1fC00D1E863019906df1771b80DBE182"
    const sbBTCPool = "0xec2946aD323f0879269910cbBB5420E8CD578a30"
    const swapRewards = "0xF4c381d077272295641F8A53D850d9a8125e0e94"
    const dao = "0xA740E20712C630d602D5007b618a1d604D3f41e9"
    const initialBTC = 1111
    const initialWBTC = 1111
    console.log("Deploying contracts with the account:", deployer.address);

    console.log("Account balance:", (await deployer.getBalance() / 1e18).toFixed(8));

    const LPTokenFactory = await ethers.getContractFactory("LPToken");
    const lpToken = await LPTokenFactory.deploy(TOKEN_DECIMALS);

    await lpToken.deployed();
    console.log("LPToken address:", lpToken.address);

    const SwapContractFactory = await ethers.getContractFactory("SwapContract");
    const swap = await SwapContractFactory.deploy(
        lpToken.address,
        BTCT_ADDR,
        sbBTCPool,
        swapRewards,
        dao,
        initialBTC,
        initialWBTC);

    await swap.deployed();
    console.log("SwapContract address:", swap.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });