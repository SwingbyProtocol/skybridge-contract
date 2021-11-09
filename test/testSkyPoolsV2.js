const ERC20ABI = require('../scripts/erc20ABI.js')
const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers')
const { expect, assert, should } = require('chai')
const { BigNumber, Ethers } = require('ethers')
const { ZERO_ADDRESS } = constants
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18
const utils = require('ethers').utils
const ParaSwap = require('../scripts/paraswap.js')
//Web3 init
//const LPToken = artifacts.require('LPToken')
//const SwapContract = artifacts.require('SwapContract')


require('chai')
    .use(require('chai-as-promised'))
    .should()

describe("SkyPools", () => {    
    let LPTokenFactory, SwapContractFactory, lptoken, swap, owner, receiver, accounts

    let convertScale, lptDecimals, btctTest, btctDecimals, mint500ERC20tokens, balance, zeroFees, minerFees, floatAmount, sampleTxs, redeemedFloatTxIds

    const wETH = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" //https://etherscan.io/token/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
    const BLOCKNUM = 13578738 //pinned block number for all testing data - paraswap TXs are with respect to this block

    const srcAmountETH = "1000000000000000000"//1 ETHER
    const srcAmountBTC = "10000000"//.1 wBTC

    const simpleDataFlow1 = [ '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
    '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
    '10000000',
    '702062308029837650',
    '1404124616059675300',
    [ '0xdef1c0ded9bec7f1a1670819833240f027b25eff' ],
    '0xaa77476c000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c59900000000000000000000000000000000000000000000000013aed978d540a30000000000000000000000000000000000000000000000000000000000009a1d200000000000000000000000000000006daea1723962647b7e189d311d757fb793000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57000000000000000000000000f73d63c3eb97389cb5a28c4ad5e8ac428cb164170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006189bd3b75c584835392e8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000001c1681a9e8fd6e29ffd26e9885ced9357a97c18b2116ac83a5323c61a8edd28ce213effddfd61f437136d11071b44342e6d04a99e54345763491b658e70521a55c0000000000000000000000000000000000000000000000000000000000989680',
    [ 0, 484 ],
    [ '0' ],
    '0xf73D63C3eB97389cB5A28C4aD5e8AC428cb16417', //set to beneficiary
    '0x000000000000000000000000536b79506F6f6c73',
    '0',
    '0x',
    '1636416827',
    '0xb285291040f111ecbd35cff44415087d' ]

    const simpleDataFlow2 = [ '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
    '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
    '100000000000000000',
    '355624',
    '711247',
    [ '0xF9234CB08edb93c0d4a4d4c70cC3FfD070e78e07' ],
    '0x91a32b69000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000004de5bb2b8038a1640196fbe3e38816f3e67cba72d940',
    [ 0, 228 ],
    [ '0' ],
    '0xf73D63C3eB97389cB5A28C4aD5e8AC428cb16417', //set to beneficiary
    '0x000000000000000000000000536b79506F6f6c73',
    '0',
    '0x',
    '1636437969',
    '0xd70aca2040f011ec99e22154ce3687c5' ]

    const swapOnUniswapFlow1 = [ '10000000',
    '698729514669918696',
    [ '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
      '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2' ] ]
    const swapOnUniswapFlow2 = [ '1000000000000000000',
    '3556087',
    [ '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
      '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599' ] ]

    const swapOnUniswapForkFlow1 = [ '0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac',
    '0xe18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303',
    '10000000',
    '701071243021484014',
    [ '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
      '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2' ] ]
    const swapOnUniswapForkFlow2 = [ '0x9DEB29c9a4c7A88a3C0257393b7f3335338D9A9D',
    '0x69d637e77615df9f235f642acebbdad8963ef35c5523142078c9b8f9d0ceba7e',
    '1000000000000000000',
    '3557483',
    [ '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
      '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599' ] ]





    let paraswap
    before(async () => {
        [owner, receiver, user1, user2, ...addrs] = await ethers.getSigners()
        accounts = [owner, receiver, user1, user2, ...addrs]
        LPTokenFactory = await ethers.getContractFactory("LPToken")
        SwapContractFactory = await ethers.getContractFactory("SwapContract")

        lpToken = await LPTokenFactory.deploy(8)

        lptDecimals = await lpToken.decimals()

        btctTest = await LPTokenFactory.deploy(TOKEN_DECIMALS)

        btctDecimals = await btctTest.decimals()

        btctDecimals = new BigNumber.from(btctDecimals)

        convertScale = new BigNumber.from(10).pow(btctDecimals.sub(lptDecimals)) //convert BTC decimals to ERC20

        mint500ERC20tokens = new BigNumber.from(500).mul(new BigNumber.from(10).pow(btctDecimals))

        swap = await SwapContractFactory.deploy(lpToken.address, btctTest.address, wETH, 0)

        zeroFees = false

        minerFees = new BigNumber.from(30000)

        floatAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(btctDecimals))

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
    })


    // You can nest describe calls to create subsections.
    describe("Testing all functions for SkyPools using hard coded API data from block", () => {
        it('checks owner of LPToken contract', async () => {
            expect(await lpToken.getOwner()).to.equal(swap.address)
            expect(await lpToken.owner()).to.equal(swap.address)
        })

        it('records SkyPools Transaction and allows user to redeem BTCT tokens', async () => {
            //mint tokens and assign to owner so they can be transferred from the contract
            await btctTest.mint(owner.address, mint500ERC20tokens)
            await btctTest.connect(owner).transfer(swap.address, mint500ERC20tokens)

            //set float reserve
            let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + owner.address.slice(2), 64)

            await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

            //check starting balances
            balance = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            const startingFloatAmount = balance[1]
            assert.equal(utils.formatEther(balance[1]).toString(), utils.formatEther(floatAmount.add(minerFees)).toString(), "Float Reserve of BTCT tokens on the contract BEFORE skypools transaction is correct")

            balance = await btctTest.balanceOf(swap.address)
            assert.equal(utils.formatEther(balance).toString(), utils.formatEther(mint500ERC20tokens).toString(), "Balance of tokens on swap contract is correct")

            balance = await btctTest.balanceOf(user1.address)
            assert.equal(balance.toNumber(), 0, "User's starting wallet balance is correct")

            //amount = 1 BTC, adjusted for 8 decimals
            let amount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(lptDecimals))
            //console.log(utils.formatEther(amount.mul(convertScale)).toString())

            //perform recordSkyPoolsTX to assign user1 tokens in the contract
            await swap.recordSkyPoolsTX(btctTest.address, user1.address, amount, 0, redeemedFloatTxIds)

            //check ending balances
            balance = await swap.tokens(btctTest.address, user1.address)
            assert.equal(balance.toString(), amount.toString(), "Ending user balance is correct")

            balance = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            assert.equal(balance[1].toString(), startingFloatAmount.sub(amount).toString(), "Ending swap balance is correct")


            //get users swap balance before they redeem tokens
            balance = await swap.balanceOf(btctTest.address, user1.address)
            assert.equal(balance.toString(), amount.toString(), "User's swap balance is correct - balanceOf function works correctly")

            //check float reserve after skypools transaction
            balance = await swap.getFloatReserve(ZERO_ADDRESS, btctTest.address)
            let expectedFloat = utils.formatEther(floatAmount.add(minerFees).sub(amount))
            assert.equal(utils.formatEther(balance[1]).toString(), expectedFloat, "Float Reserve of BTCT tokens on the contract AFTER skypools transaction")

            //redeem tokens
            const result = await swap.connect(user1).redeemERC20Token(btctTest.address, amount)

            //Verify result of token redemption based on event emitted 
            const receipt = await result.wait()
            const args = receipt.events[1].args

            assert.equal(receipt.events[1].event, "Withdraw", "Event name is correct")
            assert.equal(args.token, btctTest.address, "Token address is correct")
            assert.equal(args.user, user1.address, "Token address is correct")
            assert.equal(args.amount.toString(), amount.toString(), "Amount is correct")
            assert.equal(args.balance.toNumber(), 0, "User balance is correct")

            //check ending balances
            balance = await swap.balanceOf(btctTest.address, user1.address)
            assert.equal(balance.toNumber(), 0, "User's ending balance on the swap contract is correct")

            balance = await btctTest.balanceOf(user1.address)
            let endingAmountInBTC = amount.div(new BigNumber.from(10).pow(8)) //convert to BTC decimals
            assert.equal(utils.formatEther(balance).toString(), endingAmountInBTC.toNumber(), "User's wallet balance contains the tokens in the correct amount")

        })
        describe("Executing paraSwap transactions", () => {
            
            beforeEach(async () => {
                
            })//beforeEach - PARASWAP

            describe('executes paraSwap transactions: Flow 1 => BTC -> wBTC -> ERC20', async () => {
                it('executes paraSwap transactions using simpleSwap', async () => {
                   
                    /////////////////////////////// TEST FOR FAILURE //////////////////////////////////////////////
                    describe('Testing for flow 1 failure cases', async () => {
                        it('rejects transactions when msg.sender does not match beneficiary nor holder of tokens in tokens[][]', async () => {
                           
                        })
                        it('rejects transactions when the contract caller matches the token holder in tokens[][], but beneficiary does not', async () => {
                            
                        })
                    })
                })//SimpleSwap flow 1

                it('testing for other contract methods', async () => {
                   
                })//new contract methods - flow 1
            })//END FLOW 1

            it('executes paraSwap transactions: Flow 2a => ETH -> wETH -> wBTC -> BTC', async () => {
                
            })//flow 2a: ETH -> wETH -> wBTC -> BTC

            it('executes paraSwap transactions: Flow 2b => ERC20 -> wBTC -> BTC', async () => {

            })
        })//END PARASWAP
    })//END SKYPOOLS FUNCTIONS
})//END SKYPOOLS


//https://i.imgur.com/WlUaUXG.png