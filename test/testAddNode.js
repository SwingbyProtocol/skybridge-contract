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

        this.totalSwapped = new BN(0)

        this.withdrawalFeeBPS = new BN(20)

        this.swap = await SwapContract.new(this.lpToken.address, this.wbtcTest.address, 0);

        this.redeemedFloatTxIds = [
            "0x13e8785fe862e60f2caa4f838146ff9d4f4bd4a02dd6fb1f513b0a9be3452b62",
            "0xce66450451e62b9b4c406d0a83b90a5036039673d2682d4ec292f375ae571382"
        ]


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

    it('add 100 nodes into swap contract then remove 12 nodes', async function () {
        let rewardAddressAndAmounts = []
        let isRemoved = []
        let churnedInCount = 25
        let tssThreshold = 16
        let nodeRewardsRatio = 66
        for (i = 0; i < 100; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx1 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, this.withdrawalFeeBPS, {
            value: 0,
            gas: 9000000,
            gasPrice: 2 * 10 ** 6
        })
        // console.log(tx1.receipt.cumulativeGasUsed)
        // Gas cost 6676812 => 7653778 gas
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
        const tx2 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, this.withdrawalFeeBPS, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })

        // console.log(tx2.receipt.cumulativeGasUsed)
        // Gas cost 306494 => 1284578 gas
        let getNode2 = await this.swap.getActiveNodes()
        expect(getNode2.length).to.equal(100)


        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))

        await this.wbtcTest.mint(this.swap.address, floatAmountOfBTC)

        let swapTx = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.1 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN("1").mul(new BN(10).pow(new BN(5))).toString('hex'), 64)
        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Second deposit tx

        const distribute = await this.swap.distributeNodeRewards()
        // Gas 3999911 gas
        // console.log(distribute.receipt.cumulativeGasUsed)

        rewardAddressAndAmounts = []
        isRemoved = []

        for (i = 0; i < 12; i++) {
            let staked = new BN(500000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(true)
        }
        const tx3 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, this.withdrawalFeeBPS, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })

        let getNode3 = await this.swap.getActiveNodes()
        expect(getNode3.length).to.equal(100 - 12)

        await expectRevert(
            this.swap.distributeNodeRewards(),
            'totalRewardLPsForNode is not positive',
        );
        //console.log(tx3.receipt.cumulativeGasUsed)
        // Gas cost 51700 => 867293 gas
    })
    it('add 60 nodes into swap contract then update 100 nodes', async function () {
        let rewardAddressAndAmounts = []
        let isRemoved = []
        let churnedInCount = 25
        let tssThreshold = 16
        let nodeRewardsRatio = 66
        for (i = 0; i < 60; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx1 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, this.withdrawalFeeBPS, {
            value: 0,
            gas: 9000000,
            gasPrice: 2 * 10 ** 6
        })
        //console.log(tx1.receipt.cumulativeGasUsed)
        // Gas cost 4615206 gas
        let getNode1 = await this.swap.getActiveNodes()
        expect(getNode1.length).to.equal(60)

        rewardAddressAndAmounts = []
        isRemoved = []

        for (i = 0; i < 60; i++) {
            let staked = new BN(1500000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx2 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, this.withdrawalFeeBPS, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })

        // console.log(tx2.receipt.cumulativeGasUsed)
        // Gas cost 784326 gas
        let getNode2 = await this.swap.getActiveNodes()
        expect(getNode2.length).to.equal(60)


        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))

        await this.wbtcTest.mint(this.swap.address, floatAmountOfBTC)

        let swapTx = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.1 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN("1").mul(new BN(10).pow(new BN(5))).toString('hex'), 64)
        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Second deposit tx

        const distribute = await this.swap.distributeNodeRewards()
        // Gas 3999911 gas
        // console.log(distribute.receipt.cumulativeGasUsed)

        rewardAddressAndAmounts = []
        isRemoved = []

        for (i = 0; i < 100; i++) {
            let staked = new BN(500000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + accounts[i].slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx3 = await this.swap.churn(sender, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, this.withdrawalFeeBPS, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })

        let getNode3 = await this.swap.getActiveNodes()
        expect(getNode3.length).to.equal(100)

        await expectRevert(
            this.swap.distributeNodeRewards(),
            'totalRewardLPsForNode is not positive',
        );
        // console.log(tx3.receipt.cumulativeGasUsed)
        // Gas cost 784326 gas
    })
})