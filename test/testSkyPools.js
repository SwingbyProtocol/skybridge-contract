const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect, assert } = require('chai');
const { BigNumber } = require('ethers');
const { ZERO_ADDRESS } = constants;
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18
const utils = require('ethers').utils;

//Web3 init
//const LPToken = artifacts.require('LPToken');
//const SwapContract = artifacts.require('SwapContract');

describe("SkyPools", () => {


    let LPTokenFactory, SwapContractFactory, lptoken, swap, sender, receiver, accounts

    let lptDecimals, btctTest, btctDecimals, convertScale, mint500wBTCs, depositFeesBPS, withdrawalFeeBPS, swapFeesBPS, zeroFees, minerFees, floatAmount, swapAmount, swapFees, totalSwapped, incomingAmount, initialPriceLP, sampleTxs, redeemedFloatTxIds

    beforeEach(async () => {
        [sender, receiver, user1, user2, ...addrs] = await ethers.getSigners();
        accounts = [sender, receiver, user1, user2, ...addrs]
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


    // You can nest describe calls to create subsections.
    describe("SkyPool functions", () => {

        it('checks owner of LPToken contract', async () => {
            expect(await lpToken.getOwner()).to.equal(swap.address)
            expect(await lpToken.owner()).to.equal(swap.address)
        })
        
        
        it('records SkyPools Transactions', async () => {
            let balance
            //set float reserve
            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + sender.address.slice(2), 64)
            await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])
            //console.log(utils.formatEther(floatAmount).toString())               
    
            balance = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)  
            const startingFloatAmount = balance[1]  
            assert.equal(balance[1].toString(), floatAmount.add(minerFees).toString(), "Starting swap balance is correct")            
            
            balance = await swap.tokens(btctTest.address, user1.address) 
            assert.equal(balance.toNumber(), 0, "Starting user balance is correct") 

            let amount = new BigNumber.from(100).mul(new BigNumber.from(10).pow(lptDecimals))
            //console.log(utils.formatEther(amount).toString())

            
            //perform recordSkyPoolsTX
            await swap.recordSkyPoolsTX(btctTest.address, user1.address, amount, 0, redeemedFloatTxIds)


            //check ending balances
            balance = await swap.tokens(btctTest.address, user1.address) 
            assert.equal(balance.toString(), amount.toString(), "Ending user balance is correct")

            balance = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            assert.equal(balance[1].toString(), startingFloatAmount.sub(amount).toString(), "Ending swap balance is correct") 
                 
        });
        it('executes paraSwap transactions', async () => {
        
        });
        it('executes 1Inch trades', async () => {
        
        });
        it('allows the redemption of wBTC tokens', async () => {
        
        });
    });
});