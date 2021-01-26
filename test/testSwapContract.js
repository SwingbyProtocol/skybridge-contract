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

        this.swapAmount = new BN(1).mul(new BN(10).pow(new BN(8)))

        this.swapFees = this.swapAmount.mul(this.swapFeesBPS).div(new BN(10000))

        this.totalSwapped = this.swapAmount

        this.incomingAmount = this.swapAmount.add(this.swapFees)

        this.initialPriceLP = this.floatAmount

        this.sampleTxs = [
            "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204",
            "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e",
            "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c",
            "0x000000000000000000033d05abcee8adbd9897cdcf184e135191dc06b095534b",
            "0xc33e7f89ed85bbad177f4238608360490a0accfdb1d7440b2bcd4a10db085c91"
        ]

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
        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, txs, new BN(0), rewards, this.redeemedFloatTxIds)
        expect(await this.wbtcTest.balanceOf(this.swap.address)).to.bignumber.equal(new BN("0"))
        expect(await this.wbtcTest.balanceOf(sender)).to.bignumber.equal(tokens100)
        expect(await this.wbtcTest.balanceOf(receiver)).to.bignumber.equal(tokens400)
    })

    it('test collectSwapFeesForBTC', async function () {
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.add(this.minerFees).toString('hex') + sender.slice(2), 64)
        // Add 1 BTC float
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])
        const price = await this.swap.getCurrentPriceLP()
        expect(price).to.bignumber.equal(this.initialPriceLP)
        let rewardsAmount = this.swapFees.sub(this.minerFees)
        await this.swap.collectSwapFeesForBTC(ZERO_ADDRESS, this.incomingAmount, this.minerFees, rewardsAmount)
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

    it('deposit BTC float with fees collection', async function () {

        /**
         * Add float 1 BTC + miner fees
         */

        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.add(this.minerFees).toString('hex') + sender.slice(2), 64)
        // Add float 1 BTC + miner fees 
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])
        const price = await this.swap.getCurrentPriceLP()
        expect(price).to.bignumber.equal(this.initialPriceLP)

        const LP1 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees))

		/**
		 * Swap 1 WBTC -> BTC 
		 */

        // Swap fees are collected. (0.002 WBTC)
        await this.swap.collectSwapFeesForBTC(ZERO_ADDRESS, this.incomingAmount, this.minerFees, this.swapFees)

        // Get updated LP price = Float + Swap fees / LPT supply + 66% for Nodes (100030000 + 200000) / (100030000 + 132000) = 1.00067890
        // Get price of LP token -> 1.00067890 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // console.log(price2.toString()) 

        /**
		 * Add float 1 BTC + miner fees (zero deposit fees)
		 */

        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[1])

        // Get price of LP token -> 1.00067890 BTC/WBTC
        const price3 = await this.swap.getCurrentPriceLP()

        const LP2 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price3)
        // const depositFees2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2))


		/**
		 * Swap 1 WBTC -> BTC 
		 */

        await this.swap.collectSwapFeesForBTC(ZERO_ADDRESS, this.incomingAmount, this.minerFees, this.swapFees)

        /**
         * Add float 1 BTC + miner fees (zero deposit fees)
         */

        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[2])
        // Get price of LP token -> 1.00101847 BTC/WBTC
        const price4 = await this.swap.getCurrentPriceLP()
        // console.log(price4.toString())
        const LP3 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price4)
        // const depositFees3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2).add(LP3))
    })

    it('deposit WBTC float', async function () {
        await this.wbtcTest.mint(sender, this.floatAmount)
        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.toString('hex') + sender.slice(2), 64)
        // WBTC address
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])

        const price = await this.swap.getCurrentPriceLP()
        expect(price).to.bignumber.equal(this.initialPriceLP)
        const LP1 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees))
    })

    it('deposit WBTC float after fees are collected', async function () {

		/**
         * Add float 1 BTC
		 */

        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.toString('hex') + sender.slice(2), 64)
        // BTC == address(0)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])
        const price = await this.swap.getCurrentPriceLP()
        expect(price).to.bignumber.equal(this.initialPriceLP)

        const LP1 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price)
        const depositFees = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees))

		/**
		 * Swap 1 BTC -> WBTC
		 */

        await this.wbtcTest.mint(this.swap.address, this.swapAmount)

        let swapTx = "0x" + web3.utils.padLeft(this.swapAmount.toString('hex') + sender.slice(2), 64)
        // fees are collected. (0.002 WBTC)
        let rewardsAmount = "0x" + web3.utils.padLeft(this.swapFees.toString('hex'), 64)
        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], new BN(0), rewardsAmount, this.redeemedFloatTxIds)

        // Get updated LP price = Float + Swap fees / LPT supply + 66% for Nodes (100030000 + 200000) / (100030000 + 132000) = 1.00067890
        // Get price of LP token -> 1.00067890 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // console.log(price2.toString()) 

		/**
         * Add float 1 BTC (zero deposit fees)
		 */
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[1])

        const price3 = await this.swap.getCurrentPriceLP()
        // Get price of LP token -> 1.00067890 BTC/WBTC
        const LP2 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price3)
        const depositFees2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2).sub(depositFees2))

        // Mint again
        await this.wbtcTest.mint(this.swap.address, this.swapAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], new BN(0), rewardsAmount, this.redeemedFloatTxIds)

        /**
         * Add float 1 BTC (zero deposit fees)
		 */
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[2])
        // Get price of LP token -> 1.00101847 BTC/WBTC
        const price4 = await this.swap.getCurrentPriceLP()
        //console.log(price4.toString())
        const LP3 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price4)
        const depositFees3 = LP3.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFees).add(LP2).sub(depositFees2).add(LP3).sub(depositFees3))
    })

    it('withdraw BTC float', async function () {

        /**
         * Add float 1 BTC + miner fees
         */

        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.add(this.minerFees).toString('hex') + sender.slice(2), 64)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])

        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(this.initialPriceLP)

        const LP1 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))
        // Send LP token to swap contract.
        await this.lpToken.transfer(this.swap.address, LP1.sub(depositFeesLP1))
        // Send from TSS address
        let addressesAndAmountOfLP = "0x" + web3.utils.padLeft(LP1.sub(depositFeesLP1).toString('hex') + sender.slice(2), 64)
        // BTC == address(0)
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, addressesAndAmountOfLP, this.minerFees, this.sampleTxs[1])
        // Mint LP token
        const price2 = await this.swap.getCurrentPriceLP()
        expect(price2).to.bignumber.equal(this.initialPriceLP)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal('0')
    })

    it('withdraw BTC float after fees are collected', async function () {

        /**
         * Add float 1 BTC + miner fees
         */

        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.add(this.minerFees).toString('hex') + sender.slice(2), 64)
        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])

        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(this.initialPriceLP)

        const LP1 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))

		/**
		 * Swap 1 WBTC -> BTC executed
		 */

        // fees are collected. (0.002 WBTC)
        await this.swap.collectSwapFeesForBTC(ZERO_ADDRESS, this.incomingAmount, this.minerFees, this.swapFees)

        /**
		 * Add float 1 BTC + miner fees (zero fees)
		 */

        await this.swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[1])
        // Get price of LP token -> 1.00033977 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // Calculate amount of LP token
        const LP2 = this.floatAmount.add(this.minerFees).mul(new BN(10).pow(new BN(8))).div(price2)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2))

        await this.lpToken.transfer(this.swap.address, LP1.sub(depositFeesLP1).add(LP2))

        const floats = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address, false)
        // console.log(floats[0].toString(), floats[1].toString())
        // float 100060000 + 100000000

        // Get required amount and updated price.
        // result = await this.swap.getMinimumAmountOfLPTokens(this.minerFees)
        // //console.log(result[0].toString())
        // const AmountLP = "0x" + web3.utils.padLeft(result[0].toString('hex') + sender.slice(2), 64)

        const AmountLP = "0x" + web3.utils.padLeft(LP2.toString('hex') + sender.slice(2), 64)
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, AmountLP, this.minerFees, this.sampleTxs[2])
        const price5 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00067890 BTC/WBTC
        // console.log(price5.toString())

        const floats2 = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address, false)
        // console.log(floats2[0].toString(), floats2[1].toString())

        // console.log('Hold LPT', LP1.sub(depositFeesLP1).add(LP2).sub(result[0]).toString())

        const amountMAXWithdrawBTCfloatLP = floats2[0].mul(new BN(10).pow(new BN(8))).div(price5)
        //console.log('max', amountMAXBTCfloatLP.toString())
        const AmountLP2 = "0x" + web3.utils.padLeft(amountMAXWithdrawBTCfloatLP.toString('hex') + sender.slice(2), 64)
        await this.swap.recordOutcomingFloat(ZERO_ADDRESS, AmountLP2, this.minerFees, this.sampleTxs[3])
        const price6 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00115425 BTC/WBTC
        // console.log(price6.toString())
        const floats3 = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address, false)

        expect(floats3[0]).to.bignumber.equal('1')
    })

    it('withdraw WBTC float', async function () {

        // Set miner fees = 0
        this.minerFees = new BN(0)

        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.toString('hex') + sender.slice(2), 64)
        // BTC == address(0)
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])

        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(this.initialPriceLP)
        const LP1 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))

        // Send LP token to swap contract.
        await this.lpToken.transfer(this.swap.address, LP1.sub(depositFeesLP1))
        // Send from TSS address
        // BTC == address(0), floatAmountOfWBTC == amountOfLPtoken
        let addressesAndAmountLPToken = "0x" + web3.utils.padLeft(LP1.sub(depositFeesLP1).toString('hex') + sender.slice(2), 64)

        await this.wbtcTest.mint(this.swap.address, this.floatAmount)

        await this.swap.recordOutcomingFloat(this.wbtcTest.address, addressesAndAmountLPToken, this.minerFees, this.sampleTxs[1])

        const price2 = await this.swap.getCurrentPriceLP()
        const floatAmount = await LP1.sub(depositFeesLP1).mul(new BN(10).pow(new BN(8))).div(price2)
        const withdrawalFees = await floatAmount.mul(this.withdrawalFeeBPS).div(new BN(10000))

        expect(price2).to.bignumber.equal(this.initialPriceLP)
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal('0')
        expect(await this.lpToken.balanceOf(this.swap.address)).to.bignumber.equal('0')
        expect(await this.wbtcTest.balanceOf(sender)).to.bignumber.equal(floatAmount.sub(withdrawalFees))
    })

    it('withdraw WBTC float after fees are collected', async function () {

        /**
         * Add float 1 WBTC
         */

        // Set miner fees = 0
        this.minerFees = new BN(0)

        let addressesAndAmountOfFloat = "0x" + web3.utils.padLeft(this.floatAmount.toString('hex') + sender.slice(2), 64)
        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[0])

        const price1 = await this.swap.getCurrentPriceLP()
        expect(price1).to.bignumber.equal(this.initialPriceLP)

        const LP1 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price1)
        const depositFeesLP1 = LP1.mul(this.depositFeesBPS).div(new BN(10000))

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1))

		/**
		 * Swap 1 BTC -> WBTC executed
		 */

        // fees are collected. (0.002 WBTC)
        let swapTx = "0x" + web3.utils.padLeft(this.swapAmount.toString('hex') + receiver.slice(2), 64)
        let rewardsAmount = this.swapFees  // only BTC to WBTC

        await this.wbtcTest.mint(this.swap.address, this.swapAmount)

        await this.swap.multiTransferERC20TightlyPacked(this.wbtcTest.address, [swapTx], this.swapAmount, rewardsAmount, this.redeemedFloatTxIds)

        /**
		 * Add float 1 WBTC (zero fees)
		 */

        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[1])
        // Get price of LP token -> 1.00067910 BTC/WBTC
        const price2 = await this.swap.getCurrentPriceLP()
        // console.log(price2.toString())
        // Calculate amount of LP token
        const LP2 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price2)
        //const depositFeesLP2 = LP2.mul(this.depositFeesBPS).div(new BN(10000))
        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal(LP1.sub(depositFeesLP1).add(LP2))

        // Send all LPT
        await this.lpToken.transfer(this.swap.address, LP1.sub(depositFeesLP1).add(LP2))

        const floats = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address, false)
        // console.log(floats[0].toString(), floats[1].toString())
        // float 100000000 + 100000000

        const requiredAmountOfWBTC = LP2.mul(price2).div(new BN(10).pow(new BN(8)))
        const withdrawalFees = await requiredAmountOfWBTC.mul(this.withdrawalFeeBPS).div(new BN(10000))
        // console.log(requiredAmountOfWBTC.toString())

        await this.wbtcTest.mint(this.swap.address, requiredAmountOfWBTC)

        const AmountLP = "0x" + web3.utils.padLeft(LP2.toString('hex') + sender.slice(2), 64)

        /**
		 * Remove float 1 WBTC
		 */

        await this.swap.recordOutcomingFloat(this.wbtcTest.address, AmountLP, this.minerFees, this.sampleTxs[2])
        const price3 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00135732 BTC/WBTC
        // console.log(price3.toString())

        expect(await this.wbtcTest.balanceOf(sender)).to.bignumber.equal(requiredAmountOfWBTC.sub(withdrawalFees))

        // console.log('Hold LPT', LP1.sub(depositFeesLP1).toString())

        /**
		 * Add float 1 WBTC
		 */

        await this.swap.recordIncomingFloat(this.wbtcTest.address, addressesAndAmountOfFloat, this.zeroFees, this.sampleTxs[3])

        const floats2 = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address, false)
        // console.log(floats2[0].toString(), floats2[1].toString())


        /**
		 * Remove All WBTC Float 
		 */

        const price4 = await this.swap.getCurrentPriceLP()

        const LP3 = this.floatAmount.mul(new BN(10).pow(new BN(8))).div(price4)

        const amountMAXWithdrawWBTCfloatLP = floats2[1].mul(new BN(10).pow(new BN(8))).div(price4)
        // console.log('max', amountMAXWithdrawWBTCfloatLP.toString())

        // Send all LPT
        await this.lpToken.transfer(this.swap.address, LP3)

        const AmountLP2 = "0x" + web3.utils.padLeft(amountMAXWithdrawWBTCfloatLP.toString('hex') + sender.slice(2), 64)

        await this.wbtcTest.mint(this.swap.address, floats2[1])

        await this.swap.recordOutcomingFloat(this.wbtcTest.address, AmountLP2, this.minerFees, this.sampleTxs[4])
        const price5 = await this.swap.getCurrentPriceLP()
        // LP price is 1.00203465 BTC/WBTC
        // console.log(price5.toString())
        const floats3 = await this.swap.getFloatReserve(ZERO_ADDRESS, this.wbtcTest.address, false)

        expect(await this.lpToken.balanceOf(sender)).to.bignumber.equal('0')
        expect(floats3[1]).to.bignumber.equal('1')
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