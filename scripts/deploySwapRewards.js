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
async function main() {
    const [deployer] = await ethers.getSigners();
    //const SWINGBY = "0x8287C7b963b405b7B8D467DB9d79eEC40625b13A"
    const SWINGBY = "0xFCd51B56e65605C33024A9E98a7aaDfF2e1A15b9" // goerli
    const SwapRewardsFactory = await ethers.getContractFactory("SwapRewards");
    const swapRewards = await SwapRewardsFactory.deploy(deployer.address, SWINGBY, 358520);

    await swapRewards.deployed();
    console.log("SwapRewards address:", swapRewards.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });