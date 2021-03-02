const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;
const WBTC_ADDR = "0xEb47a21C1fC00D1E863019906df1771b80DBE182"

const LPToken = artifacts.require('LPToken');
const SwapContract = artifacts.require('SwapContract');
const SwapContractFactory = artifacts.require('SwapContractFactory');

contract('SwapFactory', function (accounts) {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.wbtcTest = await LPToken.new()

        this.factory = await SwapContractFactory.new();
    });

    it('Deploy new contracts and checking the owner', async function () {
        const sc = await this.factory.deployNewContracts(receiver, this.wbtcTest.address, 0)
        const newLPToken = await LPToken.at(sc.receipt.logs[0].args.lpToken)
        const newSwapContract = await SwapContract.at(sc.receipt.logs[0].args.swapContract)

        expect(await newLPToken.getOwner()).to.equal(newSwapContract.address)
        expect(await newSwapContract.owner()).to.equal(receiver)
    })
})