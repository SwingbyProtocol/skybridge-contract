const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;

const LPToken = artifacts.require('LPToken');
const SwapContract = artifacts.require('SwapContract');

contract('Test for swap actions', function (accounts) {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.lpToken = await LPToken.new()

        this.wbtcTest = await LPToken.new()

        this.WBTC_ADDR = this.wbtcTest.address

        this.swap = await SwapContract.new(this.lpToken.address, this.wbtcTest.address, 0);

        this.depositFeesBPS = new BN(20)

        this.totalSwapped = new BN(0)

        this.incomingAmount = new BN(20)

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

    it('multi send test', async function () {
        // swap contract receives 5000 tokens 
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
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN(1).mul(new BN(10).pow(new BN(8))).toString('hex'), 64)
        // 0x00000000000000174876e800Fb4d4830eE2AfA5E5c6FD2C2cE3a080B6634ae0e
        const txs = [
            senbackTokens1, senbackTokens2
        ]
        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, txs, this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        expect(await this.wbtcTest.balanceOf(this.swap.address)).to.bignumber.equal(new BN("0"))
        expect(await this.wbtcTest.balanceOf(sender)).to.bignumber.equal(amount2)
        expect(await this.wbtcTest.balanceOf(receiver)).to.bignumber.equal(amount3)
    })

    it('test collectSwapFeesForBTC', async function () {
        let rewardsAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        const tx = await this.swap.collectSwapFeesForBTC(ZERO_ADDRESS, this.incomingAmount, rewardsAmount)
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.collectSwapFeesForBTC(ZERO_ADDRESS, this.incomingAmount, rewardsAmount)
        // expect(await this.swap.isTxUsed(txid1)).equal(true)
        // åexpect(await this.swap.isTxUsed(txid2)).equal(true)
        // console.log(tx.receipt.gasUsed)
    })

    it('deposit BTC float', async function () {
        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        // const tx = await this.swap.issueLPTokensForFloat(txid1)
        // console.log(tx.receipt.gasUsed)
        const price = await this.swap.getCurrentPriceLP()
        const LP1 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees))

        // Second tx
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)

        const price2 = await this.swap.getCurrentPriceLP()
        const LP2 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        const depositFees2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2).sub(depositFees2))
    })

    it('deposit BTC float after fees are collected', async function () {
        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)

        const price = await this.swap.getCurrentPriceLP()
        const mintedLPs = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = mintedLPs.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(mintedLPs.sub(depositFees))

        /**
         * Swap BTC -> 1 WBTC executed, fees are 0.001 WBTC
         */

        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        // swap contract receives tokens 
        await this.wbtcTest.mint(this.swap.address, mintAmount)

        let senbackTokens1 = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.001 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN("1").mul(new BN(10).pow(new BN(5))).toString('hex'), 64)
        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [senbackTokens1], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Second deposit tx
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.06600000 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const mintedLPs2 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        const depositFees2 = mintedLPs2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(mintedLPs.sub(depositFees).add(mintedLPs2).sub(depositFees2))

        // mint again
        await this.wbtcTest.mint(this.swap.address, mintAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [senbackTokens1], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Third deposit tx
        let txid3 = "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid3)
        // await this.swap.issueLPTokensForFloat(txid3)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.05128417 BTC/WBTC
        const price3 = await this.swap.getCurrentPriceLP()
        const mintedLPs3 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price3)
        const depositFees3 = mintedLPs3.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(mintedLPs.sub(depositFees).add(mintedLPs2).sub(depositFees2).add(mintedLPs3).sub(depositFees3))
    })

    it('deposit WBTC float', async function () {
        let floatAmountOfWBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        await this.wbtcTest.mint(sender, mintAmount)
        await this.wbtcTest.transfer(this.swap.address, floatAmountOfWBTC, {
            from: sender
        })
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        let txid = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // WBTC address
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, txid)
    })

    it('deposit WBTC float after fees are collected', async function () {
        let floatAmountOfWBTC = new BN("1").mul(new BN(10).pow(new BN(8)))

        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // WBTC !== address(0)
        const record = await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        // const issue = await this.swap.issueLPTokensForFloat(txid1)

        const price1 = await this.swap.getCurrentPriceLP()
        const LP1 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price1)
        // const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1)

        /**
         * Swap BTC -> 1WBTC executed, fees are 0.001 WBTC
         */

        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        await this.wbtcTest.mint(this.swap.address, mintAmount)

        let swapTx = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.001 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN(1).mul(new BN(10).pow(new BN(5))).toString('hex'), 64)

        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Second deposit tx
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.00962888 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP2 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        // const depositFeesLP2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.add(LP2))

        // mint again
        await this.wbtcTest.mint(this.swap.address, mintAmount)
        let redeemedFloatTxIds = [
            "0x13e8785fe862e60f2caa4f838146ff9d4f4bd4a02dd6fb1f513b0a9be3452b62",
            "0xce66450451e62b9b4c406d0a83b90a5036039673d2682d4ec292f375ae571382"
        ]

        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Third deposit tx
        let txid3 = "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, txid3)
        // await this.swap.issueLPTokensForFloat(txid3)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.05128417 BTC/WBTC
        const price3 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP3 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price3)
        // const depositFeesLP3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.add(LP2).add(LP3))
    })

    it('withdraw BTC float', async function () {
        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)

        const price1 = await this.swap.getCurrentPriceLP()
        const LP1 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))
        // Send LP token to swap contract.
        await this.lpToken.transfer(this.swap.address, LP1.sub(depositFeesLP1))
        // Send from TSS address
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        let addressesAndAmountOfLP = "0x" + web3.utils.padLeft(LP1.sub(depositFeesLP1).toString('hex') + sender.slice(2), 64)
        // BTC == address(0)
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, addressesAndAmountOfLP, txid2)
        // Mint LP token
        // await this.swap.burnLPTokensForFloat(txid2)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal('0')
    })

    it('withdraw BTC float after fees are collected', async function () {
        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)

        const price1 = await this.swap.getCurrentPriceLP()
        const LP1 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))

        /**
         * Swap BTC -> 1WBTC executed, fees are 0.001 WBTC
         */

        await this.wbtcTest.mint(this.swap.address, mintAmount)

        let swapTx = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.1 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN("1").mul(new BN(10).pow(new BN(5))).toString('hex'), 64)
        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Second deposit tx
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.034 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP2 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        const depositFeesLP2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2))

        await this.wbtcTest.mint(this.swap.address, mintAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Third deposit tx
        let txid3 = "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, txid3)
        // await this.swap.issueLPTokensForFloat(txid3)
        // Get price of LP token -> 1.05128417 BTC/WBTC
        const price3 = await this.swap.getCurrentPriceLP()
        const LP3 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price3)
        const depositFeesLP3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2).add(LP3).sub(depositFeesLP3))
        // LP token amount
        const AmountOfLPtoken = "0x" + web3.utils.padLeft(LP2.toString('hex'), 64)
        // send back LP token to swap contract
        await this.lpToken.transfer(this.swap.address, AmountOfLPtoken)
        // burn LP Token tx
        let txid4 = "0x000000000000000000033d05abcee8adbd9897cdcf184e135191dc06b095534b"
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, AmountOfLPtoken, txid4)
        // await this.swap.burnLPTokensForFloat(txid4)

        const AmountOfLPtoken2 = "0x" + web3.utils.padLeft(LP3.toString('hex'), 64)
        await this.lpToken.transfer(this.swap.address, AmountOfLPtoken2)
        // burn LP Token txs
        let txid5 = "0x5e4e5cb9809a29b9dac6714a8945b44c7840cc58c012a95984f33666a6ceed52"
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, AmountOfLPtoken2, txid5)
        // await this.swap.burnLPTokensForFloat(txid5)

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).sub(depositFeesLP2).sub(depositFeesLP3))
    })

    it('withdraw WBTC float', async function () {
        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        let floatAmountOfWBTC = new BN(1).mul(new BN(10).pow(new BN(8)))

        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfWBTC.mul(new BN('1')).toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)

        const price1 = await this.swap.getCurrentPriceLP()
        const LP1 = floatAmountOfWBTC.mul(new BN('1')).mul(new BN(10).pow(new BN(8))).div(price1)
        //const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1)
        // Send LP token to swap contract.
        await this.lpToken.transfer(this.swap.address, LP1)
        // Send from TSS address
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        // BTC == address(0), floatAmountOfWBTC == amountOfLPtoken
        let addressesAndAmountLPToken = "0x" + web3.utils.padLeft(LP1.toString('hex') + sender.slice(2), 64)

        await this.wbtcTest.mint(this.swap.address, mintAmount)

        await this.swap.recordOutcomingFloat(this.WBTC_ADDR, addressesAndAmountLPToken, txid2)
        // Burn LP token
        const tx = // await this.swap.burnLPTokensForFloat(txid2)
            // console.log(tx.receipt.gasUsed)

            expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal('0')
    })

    it('withdraw WBTC float after fees are collected', async function () {
        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        let floatAmountOfWBTC = new BN(1).mul(new BN(10).pow(new BN(8)))

        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)
        const price1 = await this.swap.getCurrentPriceLP()
        const LP1 = floatAmountOfWBTC.mul(new BN('1')).mul(new BN(10).pow(new BN(8))).div(price1)
        // const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1)

        /**
        * Swap BTC -> 1WBTC executed, fees are 0.001 WBTC
        */
        await this.wbtcTest.mint(this.swap.address, mintAmount)

        let swapTx = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.1 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN("1").mul(new BN(10).pow(new BN(6))).toString('hex'), 64)
        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Second deposit tx
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)
        // Get price of LP token -> 1.034 BTC/WBTC
        const price = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP2 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price)
        // const depositFeesLP2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.add(LP2))

        await this.wbtcTest.mint(this.swap.address, mintAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Third deposit tx
        let txid3 = "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, txid3)
        // await this.swap.issueLPTokensForFloat(txid3)
        // Get price of LP token -> 1.05128417 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        const LP3 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.add(LP2).add(LP3))
        // LP token amount
        const AmountOfLPtoken = "0x" + web3.utils.padLeft(LP3.toString('hex'), 64)
        // send back LP token to swap contract
        await this.lpToken.transfer(this.swap.address, AmountOfLPtoken)
        // burn LP Token tx
        let txid4 = "0x000000000000000000033d05abcee8adbd9897cdcf184e135191dc06b095534b"
        const record = await this.swap.recordOutcomingFloat(this.WBTC_ADDR, AmountOfLPtoken, txid4)
        // const burn = await this.swap.burnLPTokensForFloat(txid4)

        const AmountOfLPtoken2 = "0x" + web3.utils.padLeft(LP2.toString('hex'), 64)
        await this.lpToken.transfer(this.swap.address, AmountOfLPtoken2)
        // burn LP Token txs
        let txid5 = "0x5e4e5cb9809a29b9dac6714a8945b44c7840cc58c012a95984f33666a6ceed52"
        await this.swap.recordOutcomingFloat(this.WBTC_ADDR, AmountOfLPtoken2, txid5)
        // await this.swap.burnLPTokensForFloat(txid5)

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1)
    })

    it('updates churn address and stakes', async function () {
        let rewardAddressAndAmounts = []
        let isRemoved = []
        let churnedInCount = 25
        let nodeRewardsRatio = 66
        for (i = 0; i < 1; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + sender.slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx1 = await this.swap.churn(receiver, rewardAddressAndAmounts, isRemoved, churnedInCount, nodeRewardsRatio, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })
        // console.log(tx1.receipt.gasUsed)
        rewardAddressAndAmounts = []
        isRemoved = []
        for (i = 0; i < 1; i++) {
            let staked = new BN(2000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + receiver.slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx2 = await this.swap.churn(receiver, rewardAddressAndAmounts, isRemoved, churnedInCount + 1, nodeRewardsRatio + 1, {
            value: 0,
            gasPrice: 2 * 10 ** 6,
            from: receiver
        })
        // console.log(tx2.receipt.gasUsed)
    })
})