const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;

const LPToken = artifacts.require('LPToken');
const SwapContract = artifacts.require('SwapContract');

contract('SwapContract', function (accounts) {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.lpToken = await LPToken.new()

        this.wbtcTest = await LPToken.new()

        this.WBTC_ADDR = this.wbtcTest.address

        this.swap = await SwapContract.new(this.lpToken.address, this.wbtcTest.address);

        await this.lpToken.transferOwnership(this.swap.address)
    });

    it('checks owner of LPToken contract', async function () {
        expect(await this.lpToken.getOwner()).to.equal(this.swap.address)
        expect(await this.lpToken.owner()).to.equal(this.swap.address)
    })

    it('add 100 nodes into swap contract', async function () {
        let rewardAddressAndAmounts = []
        let isRemoved = []
        let churnedInCount = 25
        let nodeRewardsRatio = 66
        for (i = 0; i < 100; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx1 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, nodeRewardsRatio, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })
        console.log(tx1.receipt.cumulativeGasUsed)

        rewardAddressAndAmounts = []
        isRemoved = []

        for (i = 0; i < 100; i++) {
            let staked = new BN(1500000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx2 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, nodeRewardsRatio, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })

        console.log(tx2.receipt.cumulativeGasUsed)

        rewardAddressAndAmounts = []
        isRemoved = []

        for (i = 0; i < 12; i++) {
            let staked = new BN(500000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(true)
        }
        const tx3 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, nodeRewardsRatio, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })

        const dist = await this.swap.distributeNodeRewards()

        console.log(dist.receipt.cumulativeGasUsed)
    })


})