const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;

const LPToken = artifacts.require('LPToken');
const SwapContract = artifacts.require('SwapContract');
const WBTC_ADDR = "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"

contract('SwapContract', function (accounts) {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.lpToken = await LPToken.new()

        this.wbtcTest = await LPToken.new()

        this.swap = await SwapContract.new(this.lpToken.address, this.wbtcTest.address);

        await this.lpToken.transferOwnership(this.swap.address)
    });

    it('Checks owner of LPToken contract', async function () {
        expect(await this.lpToken.getOwner()).to.equal(this.swap.address)
        expect(await this.lpToken.owner()).to.equal(this.swap.address)
    })

    it('multi send test', async function () {
        // swap contract receives 500000000 tokens 
        let amount1 = new BN(5000).mul(new BN(10).pow(new BN(8)))
        await this.wbtcTest.mint(sender, amount1)
        await this.wbtcTest.transfer(this.swap.address, amount1, {
            from: sender
        })
        let amount2 = new BN(1000).mul(new BN(10).pow(new BN(8)))
        // swap contracct send back 1000 tokens
        let amount3 = new BN(4000).mul(new BN(10).pow(new BN(8)))

        let senbackTokens1 = "0x" + web3.utils.padLeft(amount2.toString('hex') + sender.slice(2), 64)
        let senbackTokens2 = "0x" + web3.utils.padLeft(amount3.toString('hex') + receiver.slice(2), 64)

        let rewardsAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        // 0x00000000000000174876e800Fb4d4830eE2AfA5E5c6FD2C2cE3a080B6634ae0e
        const txs = [
            senbackTokens1, senbackTokens2
        ]
        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, txs, rewardsAmount)
        expect(await this.wbtcTest.balanceOf(this.swap.address)).to.bignumber.equal(new BN("0"))
        expect(await this.wbtcTest.balanceOf(sender)).to.bignumber.equal(amount2)
        expect(await this.wbtcTest.balanceOf(receiver)).to.bignumber.equal(amount3)
    })

    it('deposit BTC float', async function () {
        let flaotAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(flaotAmountOfBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        let token = ZERO_ADDRESS
        await this.swap.recordIncomingFloat(token, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        await this.swap.issueLPTokensForFloat(txid1)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(flaotAmountOfBTC)
        // Second tx
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.recordIncomingFloat(token, addressesAndAmountOfFloat, txid2)
        await this.swap.issueLPTokensForFloat(txid2)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(flaotAmountOfBTC.mul(new BN('2')))
    })

    it('deposit WBTC float', async function () {
        const newToken = await LPToken.new()
        // swap contract receives 500000000 tokens 
        let amount1 = new BN(5000).mul(new BN(10).pow(new BN(8)))
        await newToken.mint(sender, amount1)
        let flaotAmountOfWBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        await newToken.transfer(this.swap.address, flaotAmountOfWBTC, {
            from: sender
        })
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(flaotAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        let txid = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // WBTC address
        let token = WBTC_ADDR
        await this.swap.recordIncomingFloat(token, addressesAndAmountOfFloat, txid)
    })

    it('withdraw BTC float', async function () {
        let amountofLPToken = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfLPToken = "0x" + web3.utils.padLeft(amountofLPToken.toString('hex') + sender.slice(2), 64)
        let txid = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        let token = ZERO_ADDRESS
        await this.swap.recordOutcomingFloat(token, addressesAndAmountOfLPToken, txid)
    })


    it('updates churn address and stakes', async function () {
        let payload = []
        let churnedInCount = 25
        let nodeRewardsRatio = 66
        for (i = 0; i < 100; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + sender.slice(2), 64)
            payload.push(addressesAndAmountStaked)
        }
        const tx1 = await this.swap.churn(receiver, payload, churnedInCount, nodeRewardsRatio, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })
        console.log(tx1.receipt.gasUsed)
        payload = []
        for (i = 0; i < 100; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + receiver.slice(2), 64)
            payload.push(addressesAndAmountStaked)
        }
        const tx2 = await this.swap.churn(receiver, payload, churnedInCount + 1, nodeRewardsRatio + 1, {
            value: 0,
            gasPrice: 2 * 10 ** 6,
            from: receiver
        })
        console.log(tx2.receipt.gasUsed)
    })
})