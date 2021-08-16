const { constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;
const { BigNumber } = require('ethers');
const LPToken = artifacts.require('LPToken');
const SwapContract = artifacts.require('SwapContract');

const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18

describe('Contract: SwapContract', async () => {
    //const [sender, receiver] = accounts;

    let LPTokenFactory, SwapContractFactory, lptoken, swap, sender, receiver, accounts

    let lptDecimals, btctTest, btctDecimals, convertScale, mint500wBTCs, depositFeesBPS, withdrawalFeeBPS, swapFeesBPS, zeroFees, minerFees, floatAmount, swapAmount, swapFees, totalSwapped, incomingAmount, initialPriceLP, sampleTxs, redeemedFloatTxIds

    beforeEach(async () => {
        [sender, receiver, ...addrs] = await ethers.getSigners();
        accounts = [sender, receiver, ...addrs]
        LPTokenFactory = await ethers.getContractFactory("LPToken");
        SwapContractFactory = await ethers.getContractFactory("SwapContract");

        lpToken = await LPTokenFactory.deploy(8);

        lptDecimals = await lpToken.decimals()

        btctTest = await LPTokenFactory.deploy(TOKEN_DECIMALS)

        btctDecimals = await btctTest.decimals()

        btctDecimals = new BigNumber.from(btctDecimals)

        convertScale = new BigNumber.from(10).pow(btctDecimals.sub(lptDecimals))

        mint500wBTCs = new BigNumber.from(500).mul(new BigNumber.from(10).pow(btctDecimals))

        swap = await SwapContractFactory.deploy(lpToken.address, btctTest.address, 0);

        depositFeesBPS = new BigNumber.from(50)

        withdrawalFeeBPS = new BigNumber.from(20);

        swapFeesBPS = new BigNumber.from(20);

        zeroFees = false

        minerFees = new BigNumber.from(30000)

        floatAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(btctDecimals))

        swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(btctDecimals))

        swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))

        totalSwapped = swapAmount

        incomingAmount = swapAmount.add(swapFees)

        initialPriceLP = new BigNumber.from(10).pow(lptDecimals)

        sampleTxs = [
            "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204",
            "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e",
            "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c",
            "0x000000000000000000033d05abcee8adbd9897cdcf184e135191dc06b095534b",
            "0xc33e7f89ed85bbad177f4238608360490a0accfdb1d7440b2bcd4a10db085c91"
        ]

        redeemedFloatTxIds = [
            "0x13e8785fe862e60f2caa4f838146ff9d4f4bd4a02dd6fb1f513b0a9be3452b62",
            "0xce66450451e62b9b4c406d0a83b90a5036039673d2682d4ec292f375ae571382"
        ]

        await lpToken.transferOwnership(swap.address)
    });

    describe('Test for swap actions', async () => {
        it('checks owner of LPToken contract', async () => {
            expect(await lpToken.getOwner()).to.equal(swap.address)
            expect(await lpToken.owner()).to.equal(swap.address)
        })

        it('multi send test', async () => {

            await btctTest.mint(sender.address, mint500wBTCs)
            await btctTest.connect(sender).transfer(swap.address, mint500wBTCs)

            let tokens100 = new BigNumber.from(100).mul(new BigNumber.from(10).pow(lptDecimals))
            let tokens400 = new BigNumber.from(400).mul(new BigNumber.from(10).pow(lptDecimals))

            let sendTx1 = web3.utils.padLeft(tokens100._hex + sender.address.slice(2), 64)
            let sendbackTx2 = web3.utils.padLeft(tokens400._hex + receiver.address.slice(2), 64)
            let rewards = tokens100.add(tokens400).mul(swapFeesBPS).div(new BigNumber.from(10000))

            const txs = [
                sendTx1, sendbackTx2
            ]
            await swap.multiTransferERC20TightlyPacked(btctTest.address, txs, new BigNumber.from(0), rewards, redeemedFloatTxIds)

            expect(await btctTest.balanceOf(swap.address)).equal(new BigNumber.from(0))
            expect(await btctTest.balanceOf(sender.address)).equal(tokens100.mul(new BigNumber.from(10).pow(new BigNumber.from(10))))
            expect(await btctTest.balanceOf(receiver.address)).equal(tokens400.mul(new BigNumber.from(10).pow(new BigNumber.from(10))))
        })

        it('test collectSwapFeesForBTC', async () => {
            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + sender.address.slice(2), 64)
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])
            let price = await swap.getCurrentPriceLP()
            price = new BigNumber.from(price.toString())

            expect(price).equal(new BigNumber.from(initialPriceLP))

            let rewardsAmount = swapFees.sub(minerFees)
            await swap.collectSwapFeesForBTC(ZERO_ADDRESS, incomingAmount, minerFees, rewardsAmount)
        })

        it('deposit BTC float', async () => {
            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + sender.address.slice(2), 64)
            // BTC == address(0)
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])
            let price = await swap.getCurrentPriceLP()
            price = new BigNumber.from(price.toString())
            const LP1 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price)
            const depositFees = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))

            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees))

            const tx2 = await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[1])
            let price2 = await swap.getCurrentPriceLP()
            price2 = new BigNumber.from(price2.toString())
            // LP token price is not changed.
            expect(price2).equal(initialPriceLP)

            const LP2 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price2)
            const depositFees2 = LP2.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees).add(LP2).sub(depositFees2))
        })

        it('deposit BTC float with fees collection', async () => {
            /**
            * Add float 1 BTC + miner fees
            */

            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + sender.address.slice(2), 64)
            // Add float 1 BTC + miner fees 
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])
            let price = await swap.getCurrentPriceLP()
            price = new BigNumber.from(price.toString())
            expect(price).equal(initialPriceLP)

            const LP1 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price)
            const depositFees = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))

            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees))

            /**
             * Swap 1 WBTC -> BTC 
             */

            // Swap fees are collected. (0.002 WBTC)
            await swap.collectSwapFeesForBTC(ZERO_ADDRESS, incomingAmount, minerFees, swapFees)

            // Get updated LP price = Float + Swap fees / LPT supply + 66% for Nodes (100030000 + 200000) / (100030000 + 132000) = 1.00067890
            // Get price of LP token -> 1.00067890 BTC/WBTC
            let price2 = await swap.getCurrentPriceLP()
            price2 = new BigNumber.from(price2.toString())
            // console.log(price2.toString()) 

            /**
             * Add float 1 BTC + miner fees (zero deposit fees)
             */

            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[1])

            // Get price of LP token -> 1.00067890 BTC/WBTC
            let price3 = await swap.getCurrentPriceLP()
            price3 = new BigNumber.from(price3.toString())

            const LP2 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price3)
            // const depositFees2 = LP2.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees).add(LP2))


            /**
             * Swap 1 WBTC -> BTC 
             */

            await swap.collectSwapFeesForBTC(ZERO_ADDRESS, incomingAmount, minerFees, swapFees)

            /**
             * Add float 1 BTC + miner fees (zero deposit fees)
             */

            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[2])
            // Get price of LP token -> 1.00101847 BTC/WBTC
            let price4 = await swap.getCurrentPriceLP()
            price4 = new BigNumber.from(price4.toString())
            // console.log(price4.toString())
            const LP3 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price4)
            // const depositFees3 = LP3.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees).add(LP2).add(LP3))
        })

        it('deposit WBTC float', async () => {
            await btctTest.mint(sender.address, floatAmount)
            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount._hex + sender.address.slice(2), 64)
            // WBTC address
            await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

            let price = await swap.getCurrentPriceLP()
            price = new BigNumber.from(price.toString())
            expect(price).equal(initialPriceLP)
            const LP1 = floatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price)
            const depositFees = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees))
        })

        it('deposit WBTC float after fees are collected', async () => {

            /**
             * Add float 1 BTC
             */

            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount._hex + sender.address.slice(2), 64)
            // BTC == address(0)
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])
            let price = await swap.getCurrentPriceLP()
            price = new BigNumber.from(price.toString())
            expect(price).equal(initialPriceLP)

            const LP1 = floatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price)
            const depositFees = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))

            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees))

            /**
             * Swap 1 BTC -> WBTC
             */

            await btctTest.mint(swap.address, swapAmount)

            let swapTx = web3.utils.padLeft(swapAmount.div(new BigNumber.from(10).pow(new BigNumber.from(10)))._hex + sender.address.slice(2), 64)
            // fees are collected. (0.002 WBTC)
            let rewardsAmount = web3.utils.padLeft(swapFees._hex, 64)
            await swap.multiTransferERC20TightlyPacked(btctTest.address, [swapTx], new BigNumber.from(0), rewardsAmount, redeemedFloatTxIds)

            // Get updated LP price = Float + Swap fees / LPT supply + 66% for Nodes (100030000 + 200000) / (100030000 + 132000) = 1.00067890
            // Get price of LP token -> 1.00067890 BTC/WBTC
            let price2 = await swap.getCurrentPriceLP()
            price2 = new BigNumber.from(price2.toString())
            // console.log(price2.toString()) 

            /**
             * Add float 1 BTC (zero deposit fees)
             */
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[1])

            let price3 = await swap.getCurrentPriceLP()
            price3 = new BigNumber.from(price3.toString())
            // Get price of LP token -> 1.00067890 BTC/WBTC
            const LP2 = floatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price3)
            const depositFees2 = LP2.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees).add(LP2).sub(depositFees2))

            // Mint again
            await btctTest.mint(swap.address, swapAmount)

            await swap.multiTransferERC20TightlyPacked(btctTest.address, [swapTx], new BigNumber.from(0), rewardsAmount, redeemedFloatTxIds)

            /**
             * Add float 1 BTC (zero deposit fees)
             */
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[2])
            // Get price of LP token -> 1.00101847 BTC/WBTC
            let price4 = await swap.getCurrentPriceLP()
            price4 = new BigNumber.from(price4.toString())
            //console.log(price4.toString())
            const LP3 = floatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price4)
            const depositFees3 = LP3.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFees).add(LP2).sub(depositFees2).add(LP3).sub(depositFees3))
        })

        it('withdraw BTC float', async () => {
            /**
             * Add float 1 BTC + miner fees
             */
            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + sender.address.slice(2), 64)
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

            let price1 = await swap.getCurrentPriceLP()
            price1 = new BigNumber.from(price1.toString())
            expect(price1).equal(initialPriceLP)

            const LP1 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price1)
            const depositFeesLP1 = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))

            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFeesLP1))
            // Send LP token to swap contract.
            await lpToken.transfer(swap.address, LP1.sub(depositFeesLP1))
            // Send from TSS address
            let addressesAndAmountOfLP = web3.utils.padLeft(LP1.sub(depositFeesLP1)._hex + sender.address.slice(2), 64)
            // BTC == address(0)
            await swap.recordOutcomingFloat(ZERO_ADDRESS, addressesAndAmountOfLP, minerFees, sampleTxs[1])
            // Mint LP token
            let price2 = await swap.getCurrentPriceLP()
            price2 = new BigNumber.from(price2.toString())
            expect(price2).equal(initialPriceLP)
            expect(await lpToken.balanceOf(sender.address)).equal(0)
        })

        it('withdraw BTC float after fees are collected', async () => {

            /**
             * Add float 1 BTC + miner fees
             */

            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + sender.address.slice(2), 64)
            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

            let price1 = await swap.getCurrentPriceLP()
            price1 = new BigNumber.from(price1.toString())
            expect(price1).equal(initialPriceLP)

            const LP1 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price1)
            const depositFeesLP1 = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))

            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFeesLP1))

            /**
             * Swap 1 WBTC -> BTC executed
             */

            // fees are collected. (0.002 WBTC)
            await swap.collectSwapFeesForBTC(ZERO_ADDRESS, incomingAmount, minerFees, swapFees)

            /**
             * Add float 1 BTC + miner fees (zero fees)
             */

            await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[1])
            // Get price of LP token -> 1.00033977 BTC/WBTC
            let price2 = await swap.getCurrentPriceLP()
            price2 = new BigNumber.from(price2.toString())
            // Calculate amount of LP token
            const LP2 = floatAmount.add(minerFees).mul(new BigNumber.from(10).pow(lptDecimals)).div(price2)
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFeesLP1).add(LP2))

            await lpToken.transfer(swap.address, LP1.sub(depositFeesLP1).add(LP2))

            const floats = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            // console.log(floats[0].toString(), floats[1].toString())
            // float 100060000 + 100000000

            // Get required amount and updated price.
            // result = await swap.getMinimumAmountOfLPTokens(minerFees)
            // //console.log(result[0].toString())
            // const AmountLP = web3.utils.padLeft(result[0]._hex + sender.address.slice(2), 64)

            const AmountLP = web3.utils.padLeft(LP2._hex + sender.address.slice(2), 64)
            await swap.recordOutcomingFloat(ZERO_ADDRESS, AmountLP, minerFees, sampleTxs[2])
            let price5 = await swap.getCurrentPriceLP()
            price5 = new BigNumber.from(price5.toString())
            // LP price is 1.00135713 BTC/WBTC
            // console.log(price5.toString())

            let floats2 = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            //console.log(floats2[0].toString(), floats2[1].toString())

            // console.log('Hold LPT', LP1.sub(depositFeesLP1).add(LP2).sub(result[0]).toString())

            const amountMAXWithdrawBTCfloatLP = new BigNumber.from(floats2[0].toString()).mul(new BigNumber.from(10).pow(lptDecimals)).div(price5)

            //console.log('max', amountMAXBTCfloatLP.toString())
            const AmountLP2 = web3.utils.padLeft(amountMAXWithdrawBTCfloatLP._hex + sender.address.slice(2), 64)
            await swap.recordOutcomingFloat(ZERO_ADDRESS, AmountLP2, minerFees, sampleTxs[3])
            let price6 = await swap.getCurrentPriceLP()
            price6 = new BigNumber.from(price6.toString())
            // LP price is 1.00135871 BTC/WBTC
            const amt = await (new BigNumber.from(floats2[0].toString())).mul(withdrawalFeeBPS).div(new BigNumber.from(10000))
            // console.log(amt.toString())
            const floats3 = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            const remain = amt.add(new BigNumber.from(1))
            expect(new BigNumber.from(floats3[0].toString())).equal(remain)
        })

        it('withdraw WBTC float', async () => {

            // Set miner fees = 0
            minerFees = new BigNumber.from(0)

            let inputFloatAmount = floatAmount.div(new BigNumber.from(10).pow(new BigNumber.from(10)))

            let addressesAndAmountOfFloat = web3.utils.padLeft(inputFloatAmount._hex + sender.address.slice(2), 64)
            // BTC == address(0)
            await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

            let price1 = await swap.getCurrentPriceLP()
            price1 = new BigNumber.from(price1.toString())
            expect(price1).equal(initialPriceLP)
            const LP1 = inputFloatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price1)
            const depositFeesLP1 = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFeesLP1))

            // Send LP token to swap contract.
            await lpToken.transfer(swap.address, LP1.sub(depositFeesLP1))
            // Send from TSS address
            // BTC == address(0), floatAmountOfWBTC == amountOfLPtoken
            let addressesAndAmountLPToken = web3.utils.padLeft(LP1.sub(depositFeesLP1)._hex + sender.address.slice(2), 64)

            await btctTest.mint(swap.address, floatAmount)

            await swap.recordOutcomingFloat(btctTest.address, addressesAndAmountLPToken, minerFees, sampleTxs[1])

            let price2 = await swap.getCurrentPriceLP()
            price2 = new BigNumber.from(price2.toString())

            //this constant was overloaded, I changed the name to remove ambiguity 
            const FA = await LP1.sub(depositFeesLP1).mul(new BigNumber.from(10).pow(lptDecimals)).div(price2)
            //console.log(FA)          

            const withdrawalFees = await FA.mul(withdrawalFeeBPS).div(new BigNumber.from(10000))

            expect(price2).equal(initialPriceLP)
            expect(await lpToken.balanceOf(sender.address)).equal(0)
            expect(await lpToken.balanceOf(swap.address)).equal(0)
            expect(await btctTest.balanceOf(sender.address)).equal(FA.sub(withdrawalFees).mul(new BigNumber.from(10).pow(new BigNumber.from(10))))

        })

        it('withdraw WBTC float after fees are collected', async () => {
            /**
         * Add float 1 WBTC
         */

            // Set miner fees = 0
            minerFees = new BigNumber.from(0)

            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount._hex + sender.address.slice(2), 64)
            await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

            let price1 = await swap.getCurrentPriceLP()
            price1 = new BigNumber.from(price1.toString())
            expect(price1).equal(initialPriceLP)

            const LP1 = floatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price1)
            const depositFeesLP1 = LP1.mul(depositFeesBPS).div(new BigNumber.from(10000))

            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFeesLP1))

            /**
             * Swap 1 BTC -> WBTC executed
             */

            // fees are collected. (0.002 WBTC)
            let swapTx = web3.utils.padLeft(swapAmount.div(new BigNumber.from(10).pow(new BigNumber.from(10)))._hex + receiver.address.slice(2), 64)
            let rewardsAmount = swapFees  // only BTC to WBTC

            await btctTest.mint(swap.address, swapAmount)

            await swap.multiTransferERC20TightlyPacked(btctTest.address, [swapTx], swapAmount, rewardsAmount, redeemedFloatTxIds)

            /**
             * Add float 1 WBTC (zero fees)
             */

            await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, zeroFees, sampleTxs[1])
            // Get price of LP token -> 1.00067910 BTC/WBTC
            let price2 = await swap.getCurrentPriceLP()
            price2 = new BigNumber.from(price2.toString())
            // console.log(price2.toString())
            // Calculate amount of LP token
            const LP2 = floatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price2)
            //const depositFeesLP2 = LP2.mul(depositFeesBPS).div(new BigNumber.from(10000))
            expect(await lpToken.balanceOf(sender.address)).equal(LP1.sub(depositFeesLP1).add(LP2))

            // Send all LPT
            await lpToken.transfer(swap.address, LP1.sub(depositFeesLP1).add(LP2))

            const floats = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            // console.log(floats[0].toString(), floats[1].toString())
            // float 100000000 + 100000000

            const requiredAmountOfWBTC = LP2.mul(price2).div(new BigNumber.from(10).pow(lptDecimals))
            const withdrawalFees = await requiredAmountOfWBTC.mul(withdrawalFeeBPS).div(new BigNumber.from(10000))
            // console.log(requiredAmountOfWBTC.toString())

            await btctTest.mint(swap.address, requiredAmountOfWBTC.mul(new BigNumber.from(10).pow(new BigNumber.from(10))))

            const AmountLP = web3.utils.padLeft(LP2._hex + sender.address.slice(2), 64)

            /**
             * Remove float 1 WBTC
             */

            await swap.recordOutcomingFloat(btctTest.address, AmountLP, minerFees, sampleTxs[2])
            let price3 = await swap.getCurrentPriceLP()
            price3 = new BigNumber.from(price3.toString())
            // LP price is 1.00135732 BTC/WBTC
            // console.log(price3.toString())

            expect(await btctTest.balanceOf(sender.address)).equal(requiredAmountOfWBTC.sub(withdrawalFees).mul(convertScale))

            // console.log('Hold LPT', LP1.sub(depositFeesLP1).toString())

            /**
             * Add float 1 WBTC
             */

            await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, zeroFees, sampleTxs[3])

            const floats2 = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            // console.log(floats2[0].toString(), floats2[1].toString())


            /**
             * Remove All WBTC Float 
             */

            let price4 = await swap.getCurrentPriceLP()
            price4 = new BigNumber.from(price4.toString())
            const LP3 = floatAmount.mul(new BigNumber.from(10).pow(lptDecimals)).div(price4)

            const amountMAXWithdrawWBTCfloatLP = new BigNumber.from(floats2[1].toString()).mul(new BigNumber.from(10).pow(lptDecimals)).div(price4)
            // console.log('max', amountMAXWithdrawWBTCfloatLP.toString())


            // Send all LPT
            await lpToken.transfer(swap.address, LP3)

            const AmountLP2 = web3.utils.padLeft(amountMAXWithdrawWBTCfloatLP._hex + sender.address.slice(2), 64)

            await btctTest.mint(swap.address, new BigNumber.from(floats2[1].toString()).mul(convertScale))

            await swap.recordOutcomingFloat(btctTest.address, AmountLP2, minerFees, sampleTxs[4])
            let price5 = await swap.getCurrentPriceLP()
            price5 = new BigNumber.from(price5.toString())
            // LP price is 1.00203465 BTC/WBTC
            const floats3 = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)

            const withdrawFees = await new BigNumber.from(floats2[1].toString()).mul(withdrawalFeeBPS).div(new BigNumber.from(10000))
            //console.log(withdrawFees.toString())
            expect(await lpToken.balanceOf(sender.address)).equal(0)
            expect(new BigNumber.from(floats3[1].toString())).equal(withdrawFees)
        })

        it('updates churn address and stakes', async () => {
            let rewardAddressAndAmounts = []
            let isRemoved = []
            let churnedInCount = 25
            let tssThreshold = 16
            let nodeRewardsRatio = 66
            for (i = 0; i < 1; i++) {
                let staked = new BigNumber.from(3000000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + sender.address.slice(2), 64)
                rewardAddressAndAmounts.push(addressesAndAmountStaked)
                isRemoved.push(false)
            }
            const tx1 = await swap.churn(receiver.address, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, withdrawalFeeBPS, {
                value: 0,
                gasPrice: 2 * 10 ** 6
            })
            // console.log(tx1.receipt.gasUsed)
            // Gas cost 125954 gas
            rewardAddressAndAmounts = []
            isRemoved = []
            for (i = 0; i < 1; i++) {
                let staked = new BigNumber.from(2000000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + receiver.address.slice(2), 64)
                rewardAddressAndAmounts.push(addressesAndAmountStaked)
                isRemoved.push(false)
            }
            const tx2 = await swap.connect(receiver).churn(receiver.address, rewardAddressAndAmounts, isRemoved, churnedInCount + 1, tssThreshold + 1, nodeRewardsRatio + 1, withdrawalFeeBPS, {
                value: 0,
                gasPrice: 2 * 10 ** 6
            })
            // console.log(tx2.receipt.gasUsed)
            // Gas cost 106754 gas
        })
    })
})