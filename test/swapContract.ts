import { AddressZero } from "@ethersproject/constants";
import { expect } from "chai";
import { Contract } from "ethers";
import { ethers, upgrades } from "hardhat";
import { deployFixtures } from "./fixtures";


describe("SwapContract", function () {
  let swapV1: Contract;
  let swapV2: Contract;
  this.beforeEach( async() => {
    const fixture = await deployFixtures();
    swapV1 = fixture.swapV1;
    swapV2 = fixture.swapV2;
  })
  it("Should be same address with upgraded contract", async function () {
    expect(swapV1.address).to.eq(swapV2.address);
  });
});
