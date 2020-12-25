const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;

const LPToken = artifacts.require('LPToken');
const SwapContract = artifacts.require('SwapContract');

contract('Test for churn and float', function (accounts) {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.lpToken = await LPToken.new()

        this.wbtcTest = await LPToken.new()

        this.WBTC_ADDR = this.wbtcTest.address

        this.swap = await SwapContract.new(this.lpToken.address, this.wbtcTest.address, 0);

        await this.lpToken.transferOwnership(this.swap.address)
    });

    it('checks owner of LPToken contract', async function () {
        expect(await this.lpToken.getOwner()).to.equal(this.swap.address)
        expect(await this.lpToken.owner()).to.equal(this.swap.address)
    })

    // it('check float balances for WBTC', async function() {
    //     const bal = await this.swap.getFloatBalanceOf(this.wbtcTest.address, sender)
    //     expect(bal).to.bignumber.equal('0')
    // })

    // it('check float balances for BTC', async function() {
    //     const btcAddress = ZERO_ADDRESS
    //     const bal = await this.swap.getFloatBalanceOf(btcAddress, sender)
    //     expect(bal).to.bignumber.equal('0')
    // })

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
        // Gas cost 6675694 gas
        let getNode1 = await this.swap.getActiveNodes()
        expect(getNode1.length).to.equal(100)

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

        // console.log(tx2.receipt.cumulativeGasUsed)
        // Gas cost 306494 gas
        let getNode2 = await this.swap.getActiveNodes()
        expect(getNode2.length).to.equal(100)

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

        let getNode3 = await this.swap.getActiveNodes()
        expect(getNode3.length).to.equal(100 - 12)

        await expectRevert(
            this.swap.distributeNodeRewards(),
            'totalRewardLPsForNode is not positive',
        );
        // console.log(tx3.receipt.cumulativeGasUsed)
        // Gas cost 51700 gas
    })


})