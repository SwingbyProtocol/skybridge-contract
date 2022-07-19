import { AddressZero } from "@ethersproject/constants";
import { expect } from "chai";
import { Contract } from "ethers";
import { parseEther } from "ethers/lib/utils";
import { ethers, upgrades } from "hardhat";

interface Fixture {
  swapV1: Contract;
  swapV2: Contract;
}

export const deployFixtures = async (): Promise <Fixture> => {
  const [owner, ...otherAccounts] = await ethers.getSigners();
  const LpToken = await ethers.getContractFactory("LPToken");
  const lpToken = await LpToken.deploy(18);

  const BTCT = await ethers.getContractFactory("BTCT");
  const btct = await BTCT.deploy(parseEther("1000000"));

  const WETH = await ethers.getContractFactory("WETH9");
  const weth = await WETH.deploy();

  const Params = await ethers.getContractFactory("Params");
  const params = await Params.deploy();

  const Token = await ethers.getContractFactory("Token");
  const rewardToken = await Token.deploy(parseEther("1000000"));

  const SwapRewards = await ethers.getContractFactory("SwapRewards");
  const swapRewards = await SwapRewards.deploy(
    owner.address,
    rewardToken.address,
    10
  );

  const sbBTCPool = otherAccounts[otherAccounts.length - 1];

  const SwapContractV1 = await ethers.getContractFactory("SwapContractV1");
  const SwapContractV2 = await ethers.getContractFactory("SwapContractV2");
  const swapV1 = await upgrades.deployProxy(SwapContractV1, [
    lpToken.address,
    btct.address,
    weth.address,
    sbBTCPool.address,
    params.address,
    swapRewards.address,
    1
  ], {initializer: 'initialize'});
  const swapV2 = await upgrades.upgradeProxy(swapV1.address, SwapContractV2);

  return {
    swapV1,
    swapV2
  }
}

