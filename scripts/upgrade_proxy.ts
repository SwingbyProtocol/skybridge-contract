// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { ethers, upgrades } from "hardhat";

async function main() {
  const proxyAddress = process.env.PROXY ? process.env.PROXY : '';
  console.log(`Proxy address: ${proxyAddress}`);
  const SwapV2 = await ethers.getContractFactory("SwapContractV2");
  const swapV2 = await upgrades.upgradeProxy(proxyAddress, SwapV2);

  console.log("SwapContractV2 deployed at:", swapV2.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
