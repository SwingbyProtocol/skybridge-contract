const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;

const LPToken = artifacts.require('LPToken');
const SwapContract = artifacts.require('SwapContract');

contract('Test for swap actions', function (accounts) {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        this.lpToken = await LPToken.new()

        this.wbtcTest = await LPToken.new()

        this.mint500wBTCs = new BN(500).mul(new BN(10).pow(new BN(8)))

        this.swap = await SwapContract.new(this.lpToken.address, this.wbtcTest.address, 0);

        this.depositFeesBPS = new BN(50)

        this.withdrawalFeeBPS = new BN(20);

        this.swapFeesBPS = new BN(20);

        this.zeroFees = false

        this.minerFees = new BN(30000)

        this.floatAmount = new BN(1).mul(new BN(10).pow(new BN(8)))

        this.initialPriceLP = this.floatAmount

        this.swapAmount = new BN(1).mul(new BN(10).pow(new BN(8)))

        this.swapFees = this.swapAmount.mul(this.swapFeesBPS).div(new BN(10000))

        this.totalSwapped = this.swapAmount

        this.incomingAmount = this.swapAmount.add(this.swapFees)

        this.sampleTxs = [
            "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204",
            "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e",
            "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        ]

        this.redeemedFloatTxIds = [
            "0x13e8785fe862e60f2caa4f838146ff9d4f4bd4a02dd6fb1f513b0a9be3452b62",
            "0xce66450451e62b9b4c406d0a83b90a5036039673d2682d4ec292f375ae571382"
        ]
        this.multiTransferERC20TightlyPackedGasX1 = 144179
        this.multiTransferERC20TightlyPackedGasX2 = 170448
        this.recordIncomingFloatGas = 176401
        this.collectSwapFeesForBTCGas = 81701

        await this.lpToken.transferOwnership(this.swap.address)
    });

    it('checks owner of LPToken contract', async function () {
        expect(await this.lpToken.getOwner()).to.equal(this.swap.address)
        expect(await this.lpToken.owner()).to.equal(this.swap.address)
    })

    it('multi send test', async function () {
        // swap contract receives 500 tokens 
        await this.wbtcTest.mint(sender, this.mint500wBTCs)
        await this.wbtcTest.transfer(this.swap.address, this.mint500wBTCs, {
            from: sender
        })
        let tokens100 = new BN(100).mul(new BN(10).pow(new BN(8)))
        let tokens400 = new BN(400).mul(new BN(10).pow(new BN(8)))

        let sendTx1 = "0x" + web3.utils.padLeft(tokens100.toString('hex') + sender.slice(2), 64)
        let sendbackTx2 = "0x" + web3.utils.padLeft(tokens400.toString('hex') + receiver.slice(2), 64)
        let rewards = tokens100.add(tokens400).mul(this.swapFeesBPS).div(new BN(10000))
        // 0x00000000000000174876e800Fb4d4830eE2AfA5E5c6FD2C2cE3a080B6634ae0e
        const txs = [
            sendTx1, sendbackTx2
        ]
        const send = await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, txs, new BN(0), rewards, this.redeemedFloatTxIds)
        expect(send.receipt.gasUsed).to.equal(this.multiTransferERC20TightlyPackedGasX2)
        expect(await this.wbtcTest.balanceOf(this.swap.address)).to.bignumber.equal(new BN("0"))
        expect(await this.wbtcTest.balanceOf(sender)).to.bignumber.equal(tokens100)
        expect(await this.wbtcTest.balanceOf(receiver)).to.bignumber.equal(tokens400)
    })

    it('test collectSwapFeesForBTC', async function () {
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.add(this.minerFees).toString('hex') + sender.slice(2), 64)
        // Add 1 BTC float
        const tx = await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])
        expect(tx.receipt.gasUsed).to.equal(this.recordIncomingFloatGas)
        const price = await this.swap.getCurrentPriceLP()
        expect(price).to.bignumber.equal(this.initialPriceLP)
        let rewardsAmount = this.swapFees.sub(this.minerFees)
        const tx2 = await this.swap.collectSwapFeesForBTC(ZERO_ADDRESS, this.incomingAmount, rewardsAmount)
        expect(tx2.receipt.gasUsed).to.equal(this.collectSwapFeesForBTCGas)
    })

    it('deposit BTC float', async function () {
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.add(this.minerFees).toString('hex') + sender.slice(2), 64)
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])
        const price = await this.swap.getCurrentPriceLP()
        const LP1 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees))

        const tx2 = await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[1])
        const price2 = await this.swap.getCurrentPriceLP()
        // LP token price is not changed.
        expect(price2).to.bignumber.equal(this.initialPriceLP)

        const LP2 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price2)
        const depositFees2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2).sub(depositFees2))
    })

    it('deposit BTC float after fees are collected', async function () {
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.add(this.minerFees).toString('hex') + sender.slice(2), 64)
        // BTC == address(0)
        const tx = await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])
        expect(tx.receipt.gasUsed).to.equal(this.recordIncomingFloatGas)
        const price = await this.swap.getCurrentPriceLP()
        expect(price).to.bignumber.equal(this.initialPriceLP)

        const LP1 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees))

        /**
         * Swap BTC -> 1 WBTC executed, fees are 0.001 WBTC
         */

        await this.wbtcTest.mint(this.swap.address, this.incomingAmount)

        let swapTx = "0x" + web3.utils.padLeft(this.swapAmount.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.002 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(this.swapFees.toString('hex'), 64)
        const send = await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], new BN(0), rewardsAmount, this.redeemedFloatTxIds)
        expect(send.receipt.gasUsed).to.equal(this.multiTransferERC20TightlyPackedGasX1)
        // Get updated LP price = Float + Swap fees / LPT supply + 66% for Nodes (100030000 + 200000) / (100030000 + 132000) = 1.00067890
        // Get price of LP token -> 1.00067890 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // 1.00067890
        // Get price of LP token -> 1.00067890 BTC/WBTC
        // console.log(price2.toString()) 

        // Add float 1 BTC + miner fees 
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[1])

        const price3 = await this.swap.getCurrentPriceLP()
        // Get price of LP token -> 1.00067890 BTC/WBTC
        const LP2 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price3)
        const depositFees2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2).sub(depositFees2))

        // mint again
        await this.wbtcTest.mint(this.swap.address, this.incomingAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], new BN(0), rewardsAmount, this.redeemedFloatTxIds)

        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[2])
        // Get price of LP token -> 1.00101847 BTC/WBTC
        const price4 = await this.swap.getCurrentPriceLP()
        //console.log(price4.toString())
        const LP3 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price4)
        const depositFees3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2).sub(depositFees2).add(LP3).sub(depositFees3))
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
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, txid)

        const price = await this.swap.getCurrentPriceLP()
        const LP1 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees))
    })

    it('deposit WBTC float after fees are collected', async function () {
        let floatAmountOfWBTC = new BN("1").mul(new BN(10).pow(new BN(8)))

        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // WBTC !== address(0)
        const record = await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, txid1)
        // Mint LP token
        // const issue = await this.swap.issueLPTokensForFloat(txid1)
        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(floatAmountOfWBTC)
        const LP1 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))

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


        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.00033977 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP2 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        const depositFeesLP2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2))

        // mint again
        await this.wbtcTest.mint(this.swap.address, mintAmount)
        let redeemedFloatTxIds = [
            "0x13e8785fe862e60f2caa4f838146ff9d4f4bd4a02dd6fb1f513b0a9be3452b62",
            "0xce66450451e62b9b4c406d0a83b90a5036039673d2682d4ec292f375ae571382"
        ]

        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Third deposit tx
        let txid3 = "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, txid3)
        // await this.swap.issueLPTokensForFloat(txid3)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.00050969 BTC/WBTC
        const price3 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP3 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price3)
        const depositFeesLP3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2).add(LP3).sub(depositFeesLP3))
    })

    it('withdraw BTC float', async function () {
        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)

        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(floatAmountOfBTC)

        const LP1 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))
        // Send LP token to swap contract.
        await this.lpToken.transfer(this.swap.address, LP1.sub(depositFeesLP1))
        // Send from TSS address
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        let addressesAndAmountOfLP = "0x" + web3.utils.padLeft(LP1.sub(depositFeesLP1).toString('hex') + sender.slice(2), 64)
        // BTC == address(0)
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, addressesAndAmountOfLP, this.minerFees, txid2)
        // Mint LP token
        // await this.swap.burnLPTokensForFloat(txid2)
        const price2 = await this.swap.getCurrentPriceLP()
        expect(price2).to.bignumber.equal(floatAmountOfBTC)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal('0')
    })

    it('withdraw BTC float after fees are collected', async function () {
        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        let floatAmountOfBTC = new BN(1).mul(new BN(10).pow(new BN(8)))
        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)

        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(floatAmountOfBTC)

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
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)
        // const res = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address)
        // Get price of LP token -> 1.00033977 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP2 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        const depositFeesLP2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2))

        await this.wbtcTest.mint(this.swap.address, mintAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Third deposit tx
        let txid3 = "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, txid3)
        // await this.swap.issueLPTokensForFloat(txid3)
        // Get price of LP token -> 1.00050969 BTC/WBTC
        const price3 = await this.swap.getCurrentPriceLP()
        const LP3 = floatAmountOfBTC.mul(new BN(10).pow(new BN(8))).div(price3)
        const depositFeesLP3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2).add(LP3).sub(depositFeesLP3))
        // LP token amount
        const AmountOfLPtoken = "0x" + web3.utils.padLeft(LP2.sub(depositFeesLP2).toString('hex') + sender.toString().slice(2), 64)
        // send back LP token to swap contract
        await this.lpToken.transfer(this.swap.address, LP2.sub(depositFeesLP2))
        // burn LP Token tx
        let txid4 = "0x000000000000000000033d05abcee8adbd9897cdcf184e135191dc06b095534b"

        const price4 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00050969 BTC/WBTC 
        const amountFloat = LP2.sub(depositFeesLP2).mul(price4).div(new BN(10).pow(new BN(8)))
        // amount of float 0.99516900
        // console.log(price4.toString())
        result = await this.swap.getMinimumAmountOfLPTokens(this.minerFees)
        // Get required amount and updated price.
        // console.log(result[0].toString(), result[1].toString())
        const requireAmountLP = "0x" + web3.utils.padLeft(result[0].toString('hex') + sender.toString().slice(2), 64)
        // console.log(amountFloat.toString())
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, requireAmountLP, this.minerFees, txid4)
        // await this.swap.burnLPTokensForFloat(txid4)
        const price5 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00050969 BTC/WBTC
        //console.log(price5.toString())

        const AmountOfLPtoken2 = "0x" + web3.utils.padLeft(LP3.sub(depositFeesLP3).toString('hex') + sender.toString().slice(2), 64)
        await this.lpToken.transfer(this.swap.address, LP3.sub(depositFeesLP3))
        // burn LP Token txs
        let txid5 = "0x5e4e5cb9809a29b9dac6714a8945b44c7840cc58c012a95984f33666a6ceed52"
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, AmountOfLPtoken2, this.minerFees, txid5)
        const price6 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00067752 BTC/WBTC
        // console.log(price6.toString())
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))
    })

    it('withdraw WBTC float', async function () {
        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        let floatAmountOfWBTC = new BN(1).mul(new BN(10).pow(new BN(8)))

        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfWBTC.mul(new BN('1')).toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, this.zeroFees, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)

        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(floatAmountOfWBTC)
        const LP1 = floatAmountOfWBTC.mul(new BN('1')).mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))
        // Send LP token to swap contract.
        await this.lpToken.transfer(this.swap.address, LP1.sub(depositFeesLP1))
        // Send from TSS address
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        // BTC == address(0), floatAmountOfWBTC == amountOfLPtoken
        let addressesAndAmountLPToken = "0x" + web3.utils.padLeft(LP1.sub(depositFeesLP1).toString('hex') + sender.slice(2), 64)

        await this.wbtcTest.mint(this.swap.address, mintAmount)
        await this.swap.recordOutcomingFloat(this.WBTC_ADDR, addressesAndAmountLPToken, this.minerFees, txid2)

        const price2 = await this.swap.getCurrentPriceLP()
        expect(price2).to.bignumber.equal(floatAmountOfWBTC)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal('0')
    })

    it('withdraw WBTC float after fees are collected', async function () {
        let mintAmount = new BN(1).mul(new BN(10).pow(new BN(8)))
        let floatAmountOfWBTC = new BN(1).mul(new BN(10).pow(new BN(8)))

        // Send from TSS address
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        let txid1 = "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204"
        // BTC == address(0)
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, this.zeroFees, txid1)
        // Mint LP token
        // await this.swap.issueLPTokensForFloat(txid1)
        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(floatAmountOfWBTC)
        const LP1 = floatAmountOfWBTC.mul(new BN('1')).mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))

        /**
        * Swap BTC -> 1WBTC executed, fees are 0.001 WBTC
        */
        await this.wbtcTest.mint(this.swap.address, mintAmount)

        let swapTx = "0x" + web3.utils.padLeft(floatAmountOfWBTC.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.001 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(new BN("1").mul(new BN(10).pow(new BN(5))).toString('hex'), 64)
        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Second deposit tx
        let txid2 = "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e"
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, this.zeroFees, txid2)
        // await this.swap.issueLPTokensForFloat(txid2)
        // Get price of LP token -> 1.00033977 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP2 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price2)
        const depositFeesLP2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2))

        await this.wbtcTest.mint(this.swap.address, mintAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.WBTC_ADDR, [swapTx], this.totalSwapped, rewardsAmount, this.redeemedFloatTxIds)
        // Third deposit tx
        let txid3 = "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c"
        await this.swap.recordIncomingFloat(this.WBTC_ADDR, addressesAndAmountOfFloat, this.zeroFees, txid3)
        // await this.swap.issueLPTokensForFloat(txid3)
        // Get price of LP token -> 1.00050969 BTC/WBTC
        const price3 = await this.swap.getCurrentPriceLP()
        const LP3 = floatAmountOfWBTC.mul(new BN(10).pow(new BN(8))).div(price3)
        const depositFeesLP3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2).sub(depositFeesLP2).add(LP3).sub(depositFeesLP3))
        // LP token amount
        const AmountOfLPtoken = "0x" + web3.utils.padLeft(LP3.sub(depositFeesLP3).toString('hex') + sender.toString().slice(2), 64)
        // send back LP token to swap contract
        await this.lpToken.transfer(this.swap.address, LP3.sub(depositFeesLP3))

        this.minerFees = new BN(0)

        await this.wbtcTest.mint(this.swap.address, mintAmount)
        // burn LP Token tx
        let txid4 = "0x000000000000000000033d05abcee8adbd9897cdcf184e135191dc06b095534b"
        await this.swap.recordOutcomingFloat(this.WBTC_ADDR, AmountOfLPtoken, this.minerFees, txid4)
        // const burn = await this.swap.burnLPTokensForFloat(txid4)
        // LP price is 1.00050969 WBTC/BTC
        const price4 = await this.swap.getCurrentPriceLP()
        //console.log(price4.toString())

        const bal = await this.wbtcTest.balanceOf(this.swap.address)
        // WBTC balance 699000 sat
        //console.log(bal.toString())

        await this.wbtcTest.mint(this.swap.address, mintAmount)

        const AmountOfLPtoken2 = "0x" + web3.utils.padLeft(LP2.sub(depositFeesLP2).toString('hex') + sender.toString().slice(2), 64)
        await this.lpToken.transfer(this.swap.address, LP2.sub(depositFeesLP2))
        // burn LP Token txs
        let txid5 = "0x5e4e5cb9809a29b9dac6714a8945b44c7840cc58c012a95984f33666a6ceed52"
        await this.swap.recordOutcomingFloat(this.WBTC_ADDR, AmountOfLPtoken2, this.minerFees, txid5)
        const price5 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00084677 WBTC/BTC
        //console.log(price5.toString())

        const bal2 = await this.wbtcTest.balanceOf(this.swap.address)
        // WBTC balance 1347672 sat
        //console.log(bal2.toString())

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))
    })

    it('updates churn address and stakes', async function () {
        let rewardAddressAndAmounts = []
        let isRemoved = []
        let churnedInCount = 25
        let tssThreshold = 16
        let nodeRewardsRatio = 66
        for (i = 0; i < 1; i++) {
            let staked = new BN(3000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + sender.slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx1 = await this.swap.churn(receiver, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, this.withdrawalFeeBPS, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })
        // console.log(tx1.receipt.gasUsed)
        // Gas cost 125954 gas
        rewardAddressAndAmounts = []
        isRemoved = []
        for (i = 0; i < 1; i++) {
            let staked = new BN(2000000).mul(new BN(10).pow(new BN(18)))
            let addressesAndAmountStaked = "0x" + web3.utils.padLeft(staked.toString('hex') + receiver.slice(2), 64)
            rewardAddressAndAmounts.push(addressesAndAmountStaked)
            isRemoved.push(false)
        }
        const tx2 = await this.swap.churn(receiver, rewardAddressAndAmounts, isRemoved, churnedInCount + 1, tssThreshold + 1, nodeRewardsRatio + 1, this.withdrawalFeeBPS, {
            value: 0,
            gasPrice: 2 * 10 ** 6,
            from: receiver
        })
        // console.log(tx2.receipt.gasUsed)
        // Gas cost 106754 gas
    })
})