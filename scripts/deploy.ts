// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { ethers, upgrades } from "hardhat";

async function main() {
  const [owner, ...otherAccounts] = await ethers.getSigners();
  // constants that should be defined
  const rewardTokenAddress = "";
  const pricePerBTC = 0;
  const sbBTCPoolAddress = "";
  const BTCT_address = "";
  const WETH_address = "";
  const existingBTCFloat = 0;

  const LPToken = await ethers.getContractFactory("LPToken");
  const lpToken = await LPToken.deploy(18);
  await lpToken.deployed();

  console.log("LPToken deployed to:", lpToken.address);

  const Params = await ethers.getContractFactory("Params");
  const params = await Params.deploy();
  await params.deployed();

  console.log("Params deployed to:", params.address);

  const SwapRewards = await ethers.getContractFactory("SwapRewards");
  const swapRewards = await SwapRewards.deploy(
    owner.address,
    rewardTokenAddress,
    pricePerBTC
  )
  await swapRewards.deployed();

  console.log("SwapRewards deployed to:", swapRewards.address);

  const SwapV1 = await ethers.getContractFactory("SwapContractV1");
  const swapV1 = await upgrades.deployProxy(SwapV1, [
    lpToken.address,
    BTCT_address,
    WETH_address,
    sbBTCPoolAddress,
    params.address,
    swapRewards.address,
    existingBTCFloat
  ], {initializer: 'intialize'});

  console.log("SwapContractProxy deployed to:", swapV1.address);

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
