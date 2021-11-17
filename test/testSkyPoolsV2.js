const ERC20ABI = require('../scripts/erc20ABI.js')
const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers')
const { expect, assert, should } = require('chai')
const { BigNumber, Ethers } = require('ethers')
const { ZERO_ADDRESS } = constants
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18
const utils = require('ethers').utils
const ParaSwap = require('../scripts/paraswap.js')
const paraV5 = "https://apiv5.paraswap.io"
const paraswap = new ParaSwap(paraV5, 5)
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

    //Store hard coded transactions from API PUT request
    const simpleDataFlow1 = ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
        '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
        '10000000',
        '24666424179035964748',
        '49332848358071929496',
        ['0xF9234CB08edb93c0d4a4d4c70cC3FfD070e78e07',
            '0xE592427A0AEce92De3Edee1F18E0157C05861564'],
        '0x91a32b690000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c59900000000000000000000000000000000000000000000000000000000007a12000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000004de5aa873c9da6541f13c89416c17271b4c21bf7b2d7414bf3890000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c5990000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f9840000000000000000000000000000000000000000000000000000000000000bb8000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee5700000000000000000000000000000000000000000000000000000000618c5f0300000000000000000000000000000000000000000000000000000000001e848000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000',
        [0, 228, 488],
        ['0', '0'],
        '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
        '0x000000000000000000000000536b79506F6f6c73',
        '0',
        '0x',
        '1636607315',
        '0x219c6130427b11ec944a9952ba50eb97']

    const simpleDataFlow2 = ['0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
        '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
        '1000000000000000000',
        '3554121',
        '7108241',
        ['0xF9234CB08edb93c0d4a4d4c70cC3FfD070e78e07'],
        '0x91a32b69000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000004de5bb2b8038a1640196fbe3e38816f3e67cba72d940',
        [0, 228],
        ['0'],
        '0x202CCe504e04bEd6fC0521238dDf04Bc9E8E15aB',
        '0x000000000000000000000000536b79506F6f6c73',
        '0',
        '0x',
        '1636608105',
        '0xf80afb90427c11ecb8ec7d3e271a910f']

    const simpleDataFlow2B = ['0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
        '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
        '1000000000000000000',
        '19439',
        '38878',
        ['0xF9234CB08edb93c0d4a4d4c70cC3FfD070e78e07'],
        '0x91a32b690000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f9840000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000004de4aa873c9da6541f13c89416c17271b4c21bf7b2d7',
        [0, 228],
        ['0'],
        '0x202CCe504e04bEd6fC0521238dDf04Bc9E8E15aB',
        '0x000000000000000000000000536b79506F6f6c73',
        '0',
        '0x',
        '1636608802',
        '0x978ceab0427e11ec9a8cb5f0ed4889af']
    const swapOnUniswapForkFlow1 = ['0x115934131916C8b277DD010Ee02de363c09d037c',
        '0x65d1a3b1e46c6e4f1be1ad5f99ef14dc488ae0549dc97db9b30afe2241ce1c7a',
        '10000000',
        '138090143398971911407',
        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
            '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984']]
    const swapOnUniswapFlow1 = ['10000000',
        '136971651253700618386',
        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
            '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984']]

    const swapOnUniswapForkFlow2 = ['0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f',
        '0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f',
        '1000000000000000000',
        '19439',
        ['0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
            '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599']]

    const swapOnUniswapFlow2 = ['1000000000000000000',
        '19439',
        ['0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
            '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599']]


    const endAmount = "100000000" //1 BTCt - decimal 8
    const btcStartingAmount = "500000000"
    const highValueEthAccount = "0x10c6b61DbF44a083Aec3780aCF769C77BE747E23" // ~.8 wBTC       

    //const paraAPI_URL = "https://apiv4.paraswap.io/v2"//V4 API URL
    const paraV5 = "https://apiv5.paraswap.io"
    const paraswap = new ParaSwap(paraV5, 5)
    const e = new ERC20ABI()
    const mainnetNetworkID = 1
    const Tokens = {
        [mainnetNetworkID]: {
            ETH: {
                address: '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE',
                decimals: 18,
            },
            WETH: {
                address: '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
                decimals: 18
            },
            USDC: {
                address: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
                decimals: 6,
            },
            MATIC: {
                address: '0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0',
                decimals: 18,
            },
            wBTC: {
                address: '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
                decimals: 8
            },
            UNI: {
                address: '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
                decimals: 18
            },
            BAT: {
                address: '0x0d8775f648430679a709e98d2b0cb6250d2887ef',
                decimals: 18
            }
        },
    }



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
            //console.log(utils.formatEther(amount.mul(convertScale)).toString()

            //perform recordSkyPoolsTX to assign user1 tokens in the contract
            let swapFeesBPS = new BigNumber.from(20);
            let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(btctDecimals))
            let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))

            await swap.recordSkyPoolsTX(user1.address, amount, swapFees, sampleTxs) // TODO: need to add rewwards

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
            const wBTC = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
            const UNI = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
            const wBTC_Contract = new ethers.Contract(Tokens[mainnetNetworkID]['wBTC'].address, e.erc20ABI(), ethers.provider)
            const UNI_Contract = new ethers.Contract(Tokens[mainnetNetworkID]['UNI'].address, e.erc20ABI(), ethers.provider)
            const wETH_Contract = new ethers.Contract(Tokens[mainnetNetworkID]['WETH'].address, e.erc20ABI(), ethers.provider)

            //Set starting balance for swap contract
            //get slot for wBTC balanceOf - https://kndrck.co/posts/local_erc20_bal_mani_w_hh/
            //npx slot20 balanceOf 0x2260fac5e5542a773aa44fbcfedf7c193bc2c599 0x10c6b61DbF44a083Aec3780aCF769C77BE747E23 -v
            //npx slot20 balanceOf 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984 0x4C6007e38Ce164Ed80FF8Ff94192225FcdAC68CD -v
            const wBTC_SLOT = 0
            const UNI_SLOT = 4
            const toBytes32 = (bn) => {
                return ethers.utils.hexlify(ethers.utils.hexZeroPad(bn, 32));
            };
            const setStorageAt = async (address, index, value) => {
                await ethers.provider.send("hardhat_setStorageAt", [address, index, value]);
                await ethers.provider.send("evm_mine", []); // Just mines to the next block
            };
            //seed balance with wBTC, using ethers.provider
            const populateBalance = async (userAddress, tokenAddress, slot, amount) => {
                //const btct = new ethers.Contract(wBTC, e.erc20ABI(), ethers.provider)
                const locallyManipulatedBalance = new BigNumber.from(amount)

                // Get storage slot index
                const index = ethers.utils.solidityKeccak256(
                    ["uint256", "uint256"],
                    [userAddress, slot] // key, slot
                );

                // Manipulate local balance (needs to be bytes32 string)
                await setStorageAt(
                    tokenAddress,
                    index.toString(),
                    toBytes32(locallyManipulatedBalance).toString()
                );
            }
            beforeEach(async () => {
                //reset back to starting block for repeated tests - see hardhat config for block number
                await network.provider.request({
                    method: "hardhat_reset",
                    params: [
                        {
                            forking: {
                                jsonRpcUrl: "https://eth-mainnet.alchemyapi.io/v2/YfblHzLI_PlnIA0pxphS90J3yaA3dDi5",
                                blockNumber: BLOCKNUM,
                            },
                        },
                    ],
                });
                //impersonate real mainnet account
                await hre.network.provider.request({
                    method: "hardhat_impersonateAccount",
                    params: [highValueEthAccount]
                })

                //re-deploy with real world wBTC address
                lpToken = await LPTokenFactory.deploy(8)
                swap = await SwapContractFactory.deploy(lpToken.address, wBTC, wETH, 0)

                await lpToken.transferOwnership(swap.address)
                expect(await lpToken.getOwner()).to.equal(swap.address)
                expect(await lpToken.owner()).to.equal(swap.address)

            })//beforeEach - PARASWAP

            describe('executes paraSwap transactions: Flow 1 => BTC -> wBTC -> ERC20', async () => {
                it('simpleSwap flow 1', async () => {
                    //Allocate wBTC tokens to user in tokens[][]
                    await populateBalance(swap.address, wBTC, wBTC_SLOT, btcStartingAmount)

                    let decimals = Tokens[mainnetNetworkID]['ETH'].decimals
                    //set float reserve
                    const newFloatAmount = new BigNumber.from("500000000")//5 tokens ERC20 - decimal 18
                    let addressesAndAmountOfFloat = web3.utils.padLeft(newFloatAmount.add(minerFees)._hex + owner.address.slice(2), 64)
                    await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

                    //check float reserve
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    const startingFloatAmount = balance[1]
                    assert.equal(utils.formatEther(balance[1]).toString(), utils.formatEther(newFloatAmount.add(minerFees)).toString(), "Float Reserve of BTCT tokens on the contract BEFORE skypools transaction is correct")

                    //perform recordSkyPoolsTX to assign user1 btct tokens in the contract
                    let swapFeesBPS = new BigNumber.from(20);
                    let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                    await swap.recordSkyPoolsTX(user1.address, endAmount, swapFees, sampleTxs) // todo: need to add rewards

                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(startingFloatAmount.sub(balance[1]).toString(), endAmount, "1 BTCt deducted from float") //1 BTCt -> decimal 8

                    balance = await swap.tokens(wBTC, user1.address)
                    assert.equal(balance.toString(), endAmount, "User balance in tokens[][] is correct")

                    //Execute paraswap TX
                    const result = await swap.connect(user1).spFlow1SimpleSwap(
                        simpleDataFlow1
                    )

                    //Check ending balances
                    balance = await UNI_Contract.balanceOf(user1.address)
                    assert(balance > new BigNumber.from(0), "User has UNI tokens in wallet")

                    balance = await swap.connect(user1).balanceOf(user1.address, UNI)
                    assert.equal(balance.toString(), "0", "User balance of UNI tokens on swap contract is 0")

                })//SimpleSwap flow 1

                it('swapOnUniswapFork flow 1', async () => {
                    //Allocate wBTC tokens to user in tokens[][]
                    await populateBalance(swap.address, wBTC, wBTC_SLOT, btcStartingAmount)

                    let decimals = Tokens[mainnetNetworkID]['ETH'].decimals
                    //set float reserve
                    const newFloatAmount = new BigNumber.from("500000000")//5 tokens ERC20 - decimal 18
                    let addressesAndAmountOfFloat = web3.utils.padLeft(newFloatAmount.add(minerFees)._hex + owner.address.slice(2), 64)
                    await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

                    //check float reserve
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    const startingFloatAmount = balance[1]
                    assert.equal(utils.formatEther(balance[1]).toString(), utils.formatEther(newFloatAmount.add(minerFees)).toString(), "Float Reserve of BTCT tokens on the contract BEFORE skypools transaction is correct")

                    //perform recordSkyPoolsTX to assign user1 btct tokens in the contract
                    let swapFeesBPS = new BigNumber.from(20);
                    let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                    await swap.recordSkyPoolsTX(user1.address, endAmount, swapFees, sampleTxs) // todo: need to add rewards

                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(startingFloatAmount.sub(balance[1]).toString(), endAmount, "1 BTCt deducted from float") //1 BTCt -> decimal 8

                    balance = await swap.tokens(wBTC, user1.address)
                    assert.equal(balance.toString(), endAmount, "User balance in tokens[][] is correct")

                    //Execute paraswap TX
                    await swap.connect(user1).spFlow1Uniswap(
                        true,
                        swapOnUniswapForkFlow1[0],
                        swapOnUniswapForkFlow1[1],
                        swapOnUniswapForkFlow1[2],
                        swapOnUniswapForkFlow1[3],
                        swapOnUniswapForkFlow1[4],
                    )
                    
                    //check contract balance
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address) //254568805347400382633
                    const expectedUniTokens = balance
                    assert(expectedUniTokens > new BigNumber.from(0), "User has UNI tokens in the contract")

                    //withdraw tokens
                    const result = await swap.connect(user1).redeemERC20Token(UNI, expectedUniTokens)

                    let receipt = await result.wait()

                    let event = receipt.events[receipt.events.length - 1].args

                    //verify event receipt is correct
                    event.token.toString().toUpperCase().should.equal(UNI.toString().toUpperCase(), 'token is correct')
                    event.user.toString().should.equal(user1.address, 'user address is correct')
                    event.amount.toString().should.equal(expectedUniTokens.toString(), 'amount is correct')
                    event.balance.toString().should.equal("0", 'balance is correct')
                    event.Timestamp.toString().length.should.be.at.least(1, 'timestamp is present')

                    //check wallet balance
                    balance = await UNI_Contract.balanceOf(user1.address)
                    assert.equal(balance.toString(), expectedUniTokens.toString(), "User has the correct amount of UNI tokens in wallet")
                    

                })//swapOnUniswapFork flow 1

                it('swapOnUniswap flow 1', async () => {
                    //Allocate wBTC tokens to user in tokens[][]
                    await populateBalance(swap.address, wBTC, wBTC_SLOT, btcStartingAmount)

                    let decimals = Tokens[mainnetNetworkID]['ETH'].decimals
                    //set float reserve
                    const newFloatAmount = new BigNumber.from("500000000")//5 tokens ERC20 - decimal 18
                    let addressesAndAmountOfFloat = web3.utils.padLeft(newFloatAmount.add(minerFees)._hex + owner.address.slice(2), 64)
                    await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

                    //check float reserve
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    const startingFloatAmount = balance[1]
                    assert.equal(utils.formatEther(balance[1]).toString(), utils.formatEther(newFloatAmount.add(minerFees)).toString(), "Float Reserve of BTCT tokens on the contract BEFORE skypools transaction is correct")

                    //perform recordSkyPoolsTX to assign user1 btct tokens in the contract
                    let swapFeesBPS = new BigNumber.from(20);
                    let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                    await swap.recordSkyPoolsTX(user1.address, endAmount, swapFees, sampleTxs) // todo: need to add rewards

                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(startingFloatAmount.sub(balance[1]).toString(), endAmount, "1 BTCt deducted from float") //1 BTCt -> decimal 8

                    balance = await swap.tokens(wBTC, user1.address)
                    assert.equal(balance.toString(), endAmount, "User balance in tokens[][] is correct")

                    //Execute paraswap TX
                    await swap.connect(user1).spFlow1Uniswap(
                        false,
                        swapOnUniswapForkFlow1[0],
                        swapOnUniswapForkFlow1[1],
                        swapOnUniswapFlow1[0],
                        swapOnUniswapFlow1[1],
                        swapOnUniswapFlow1[2],
                    )
                    
                    //check contract balance
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address) //254568805347400382633
                    const expectedUniTokens = balance
                    assert(expectedUniTokens > new BigNumber.from(0), "User has UNI tokens in the contract")

                    //withdraw tokens
                    const result = await swap.connect(user1).redeemERC20Token(UNI, expectedUniTokens)

                    let receipt = await result.wait()

                    let event = receipt.events[receipt.events.length - 1].args

                    //verify event receipt is correct
                    event.token.toString().toUpperCase().should.equal(UNI.toString().toUpperCase(), 'token is correct')
                    event.user.toString().should.equal(user1.address, 'user address is correct')
                    event.amount.toString().should.equal(expectedUniTokens.toString(), 'amount is correct')
                    event.balance.toString().should.equal("0", 'balance is correct')
                    event.Timestamp.toString().length.should.be.at.least(1, 'timestamp is present')

                    //check wallet balance
                    balance = await UNI_Contract.balanceOf(user1.address)
                    assert.equal(balance.toString(), expectedUniTokens.toString(), "User has the correct amount of UNI tokens in wallet")
                })//swapOnUniswap flow 1

                /////////////////////////////// TEST FOR FAILURE //////////////////////////////////////////////
                describe('Testing for flow 1 failure cases', async () => {
                    it('rejects transactions when msg.sender does not match beneficiary nor holder of tokens in tokens[][]', async () => {
                        await swap.connect(user2).spFlow1SimpleSwap(
                            simpleDataFlow1
                        ).should.be.rejectedWith("You can only execute swaps to your own address")
                    })
                    it('rejects transactions when the contract caller matches the token holder in tokens[][], but beneficiary does not', async () => {
                        let badData = simpleDataFlow1
                        badData[9] = user2.address //malicious beneficiary

                        await swap.connect(user1).spFlow1SimpleSwap(
                            badData
                        ).should.be.rejectedWith("You can only execute swaps to your own address")
                    })
                })
            })//END FLOW 1

            describe('executes paraSwap transactions: Flow 2 => ERC-20 -> wBTC -> BTC', async () => {
                it('simpleSwap Flow 2a => ETH -> wETH -> wBTC -> BTC', async () => {
                    const ETHER = Tokens[mainnetNetworkID]['ETH'].address
                    let amount = utils.parseEther("1")
                    overrides = {
                        value: amount
                    }

                    /////////////////////////////// TESTING DEPOSIT ETHER //////////////////////////////////////////////                
                    result = await swap.connect(user1).spDeposit("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", amount, overrides)
                    balance = await swap.balanceOf(wETH, user1.address)
                    assert.equal(balance._hex, amount._hex, "Balance of wETH is correct after depositing ETHER")
                    receipt = await result.wait()

                    let event = receipt.events[receipt.events.length - 1].args

                    //verify event receipt is correct
                    event.token.toString().should.equal(ETHER, 'token is correct')
                    event.user.toString().should.equal(user1.address, 'user address is correct')
                    event.amount.toString().should.equal(amount.toString(), 'amount is correct')
                    event.balance.toString().should.equal(amount.toString(), 'balance is correct')
                    event.Timestamp.toString().length.should.be.at.least(1, 'timestamp is present')

                    /////////////////////////////// TESTING WITHDRAW ETHER //////////////////////////////////////////////
                    result = await swap.connect(user1).redeemEther(amount)
                    balance = await swap.balanceOf(ETHER, user1.address)
                    assert.equal(balance.toString(), "0", "Balance is correct after depositing ETHER")
                    receipt = await result.wait()
                    event = receipt.events[receipt.events.length - 1].args

                    //verify event receipt is correct
                    event.token.toString().should.equal(ETHER, 'token is correct')
                    event.user.toString().should.equal(user1.address, 'user address is correct')
                    event.amount.toString().should.equal(amount.toString(), 'amount is correct')
                    event.balance.toString().should.equal("0", 'balance is correct')
                    event.Timestamp.toString().length.should.be.at.least(1, 'timestamp is present')

                    //Deposit ETHER again for flow 2 transaction
                    await swap.connect(user1).spDeposit("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", amount, overrides)

                    /////////////////////////////// EXECUTE SWAP AND RECORD //////////////////////////////////////////////
                    const receivingBTC_Addr = sampleTxs[0]

                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(balance[1].toString(), "0", "Float Reserve of BTCT tokens on the contract BEFORE spFlow2SimpleSwap is 0")

                    balance = await wETH_Contract.balanceOf(swap.address)
                    assert.equal(balance.toString(), amount.toString(), "Balance of wETH on swap contract before swap is correct")

                    //perform swap
                    const swapAndRecordResult = await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    /////////////////////////////// CHECK ENDING BALANCES //////////////////////////////////////////////

                    balance = await wETH_Contract.balanceOf(swap.address)
                    assert.equal(balance.toString(), "0", "Balance of wETH on swap contract after swap is correct")

                    balance = await swap.balanceOf(wBTC, swap.address)
                    balance.toNumber().should.be.at.least(1, "Balance is present")
                    const expectedSats = balance.toString()


                    receipt = await swapAndRecordResult.wait()

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, sampleTxs[0], "DestAddr for BTC is correct")
                    assert.equal(args.AmountWBTC.toString(), expectedSats, "AmountWBTC is correct amount")
                    assert.equal(args.RefundAddr, user1.address, "RefundAddr is correct")
                    args.Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 1, "1 item in data array returned")
                    assert.equal(data[0].SwapID.toString(), TX_ID, "Swap ID is correct")
                    assert.equal(data[0].DestAddr, receivingBTC_Addr, "Dest BTC addr is correct")
                    assert.equal(data[0].AmountWBTC.toString(), expectedSats, "Output amount matches float amount")
                    data[0].Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")


                    /////////////////////////////// CHECK FOR REFUND CASE //////////////////////////////////////////////

                    //set float reserve
                    decimals = Tokens[mainnetNetworkID]['wBTC'].decimals
                    floatAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                    let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + owner.address.slice(2), 64)
                    await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])
                    await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, zeroFees, sampleTxs[1])
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    //console.log("Starting BTC Balance: ", balance[0].toString())
                    //console.log("Starting wBTC Balance: ", balance[1].toString())


                    //prep to collectSwapFeesForBTC
                    let swapFeesBPS = new BigNumber.from(20);
                    let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                    let incomingAmount = swapAmount.add(swapFees) //100200000 sats = 1.002 BTC                

                    //allocate floats to wBTC ~1wBTC decimal 8
                    await swap.collectSwapFeesForBTC(ZERO_ADDRESS, incomingAmount, minerFees, swapFees)
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    //console.log("BTC Balance AFTER collectSwapFeesForBTC: ", balance[0].toString())
                    //console.log("wBTC Balance AFTER collectSwapFeesForBTC: ", balance[1].toString())

                    //init refund flow for skypools TX
                    const refundSwapID = data[0].SwapID.toString()

                    //get refund addr and amount
                    let refundAddr, refundAmount
                    for (let i = 0; i < data.length; i++) {
                        if (data[i].SwapID.toString() == refundSwapID) {
                            refundAddr = data[i].RefundAddr.toString()
                            refundAmount = data[i].AmountWBTC
                        }
                    }

                    //issue refund
                    await swap.connect(owner).singleTransferERC20(
                        wBTC,
                        refundAddr,
                        refundAmount,
                        new BigNumber.from(0),
                        new BigNumber.from(0),
                        redeemedFloatTxIds
                    )

                    balance = await wBTC_Contract.balanceOf(swap.address)
                    assert.equal(balance.toString(), "0", "Contract balance after refund is 0")

                    balance = await wBTC_Contract.balanceOf(user1.address)
                    assert.equal(balance.toString(), expectedSats, "User has received the full refund")

                })//flow 2a: ETH -> wETH -> wBTC -> BTC

                it('simpleSwap Flow 2b => ERC20 -> wBTC -> BTC', async () => {

                    let balance, amount, result, receipt, event
                    amount = utils.parseEther("1")

                    /////////////////////////////// TESTING DEPOSIT UNI TOKENS //////////////////////////////////////////////
                    await populateBalance(user1.address, UNI, UNI_SLOT, amount.mul(2))//amount refers to number of UNI tokens here
                    balance = await UNI_Contract.balanceOf(user1.address)
                    assert.equal(balance.toString(), amount.mul(2).toString(), "User has correct number of UNI tokens in their personal wallet")

                    await UNI_Contract.connect(user1).approve(swap.address, amount)
                    result = await swap.connect(user1).spDeposit(UNI, amount)

                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance._hex, amount._hex, "Balance is correct after depositing UNI")

                    const receivingBTC_Addr = sampleTxs[0]

                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(balance[1].toString(), "0", "Float Reserve of BTCT tokens on the contract BEFORE spFlow2SimpleSwap is 0")

                    const swapAndRecordResult = await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2B
                    )
                    receipt = await swapAndRecordResult.wait()


                    /////////////////////////////// CHECK BALANCES AND RECORDED DATA //////////////////////////////////////////////
                    balance = await swap.balanceOf(wBTC, swap.address)
                    balance.toNumber().should.be.at.least(1, "Balance is present")
                    const expectedSats = balance.toString() // ~38878

                    balance = await swap.balanceOf(wBTC, swap.address)
                    assert.equal(balance.toString(), expectedSats, "Balance is correct in contract's address in tokens[][]")

                    //check user balance on swap contract
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance.toString(), "0", "Uni token has been swapped and user no longer holds it in tokens[][]")

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, sampleTxs[0], "DestAddr for BTC is correct")
                    assert.equal(args.AmountWBTC.toString(), expectedSats, "AmountWBTC is correct amount")
                    assert.equal(args.RefundAddr, user1.address, "RefundAddr is correct")
                    args.Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 1, "1 item in data array returned")
                    assert.equal(data[0].SwapID.toString(), TX_ID, "Swap ID is correct")
                    assert.equal(data[0].DestAddr, receivingBTC_Addr, "Dest BTC addr is correct")
                    assert.equal(data[0].AmountWBTC.toString(), expectedSats, "Output amount matches float amount")
                    data[0].Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")

                })//Flow 2b => ERC20 -> wBTC -> BTC
                it('swapOnUniswapForkFlow2', async () => {
                    let balance, amount, result, receipt, event
                    amount = utils.parseEther("1")

                    /////////////////////////////// TESTING DEPOSIT UNI TOKENS //////////////////////////////////////////////
                    await populateBalance(user1.address, UNI, UNI_SLOT, amount.mul(2))//amount refers to number of UNI tokens here
                    balance = await UNI_Contract.balanceOf(user1.address)
                    assert.equal(balance.toString(), amount.mul(2).toString(), "User has correct number of UNI tokens in their personal wallet")

                    await UNI_Contract.connect(user1).approve(swap.address, amount)
                    result = await swap.connect(user1).spDeposit(UNI, amount)

                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance._hex, amount._hex, "Balance is correct after depositing UNI")

                    const receivingBTC_Addr = sampleTxs[0]

                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(balance[1].toString(), "0", "Float Reserve of BTCT tokens on the contract BEFORE spFlow2SimpleSwap is 0")

                    //fork = true
                    const swapAndRecordResult = await swap.connect(user1).spFlow2Uniswap(
                        receivingBTC_Addr,
                        true,
                        swapOnUniswapForkFlow2[0],
                        swapOnUniswapForkFlow2[1],
                        swapOnUniswapForkFlow2[2],
                        swapOnUniswapForkFlow2[3],
                        swapOnUniswapForkFlow2[4],
                    )
                    receipt = await swapAndRecordResult.wait()

                    /////////////////////////////// CHECK BALANCES AND RECORDED DATA //////////////////////////////////////////////
                    balance = await swap.balanceOf(wBTC, swap.address)
                    balance.toNumber().should.be.at.least(1, "Balance is present")
                    const expectedSats = balance.toString() //~38878

                    balance = await swap.balanceOf(wBTC, swap.address)
                    assert.equal(balance.toString(), expectedSats, "Balance is correct in contract's address in tokens[][]")

                    //check user balance on swap contract
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance.toString(), "0", "Uni token has been swapped and user no longer holds it in tokens[][]")

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, sampleTxs[0], "DestAddr for BTC is correct")
                    assert.equal(args.AmountWBTC.toString(), expectedSats, "AmountWBTC is correct amount")
                    assert.equal(args.RefundAddr, user1.address, "RefundAddr is correct")
                    args.Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 1, "1 item in data array returned")
                    assert.equal(data[0].SwapID.toString(), TX_ID, "Swap ID is correct")
                    assert.equal(data[0].DestAddr, receivingBTC_Addr, "Dest BTC addr is correct")
                    assert.equal(data[0].AmountWBTC.toString(), expectedSats, "Output amount matches float amount")
                    data[0].Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")

                })//new contract methods - flow 2
                it('swapOnUniswap', async () => {
                    let balance, amount, result, receipt, event
                    amount = utils.parseEther("1")

                    /////////////////////////////// TESTING DEPOSIT UNI TOKENS //////////////////////////////////////////////
                    await populateBalance(user1.address, UNI, UNI_SLOT, amount.mul(2))//amount refers to number of UNI tokens here
                    balance = await UNI_Contract.balanceOf(user1.address)
                    assert.equal(balance.toString(), amount.mul(2).toString(), "User has correct number of UNI tokens in their personal wallet")

                    await UNI_Contract.connect(user1).approve(swap.address, amount)
                    result = await swap.connect(user1).spDeposit(UNI, amount)

                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance._hex, amount._hex, "Balance is correct after depositing UNI")

                    const receivingBTC_Addr = sampleTxs[0]

                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(balance[1].toString(), "0", "Float Reserve of BTCT tokens on the contract BEFORE spFlow2SimpleSwap is 0")

                    //fork = true
                    const swapAndRecordResult = await swap.connect(user1).spFlow2Uniswap(
                        receivingBTC_Addr,
                        false,
                        swapOnUniswapForkFlow2[0],
                        swapOnUniswapForkFlow2[1],
                        swapOnUniswapFlow2[0],
                        swapOnUniswapFlow2[1],
                        swapOnUniswapFlow2[2],
                    )
                    receipt = await swapAndRecordResult.wait()

                    /////////////////////////////// CHECK BALANCES AND RECORDED DATA //////////////////////////////////////////////
                    balance = await swap.balanceOf(wBTC, swap.address)
                    balance.toNumber().should.be.at.least(1, "Balance is present")
                    const expectedSats = balance.toString() //~38878

                    balance = await swap.balanceOf(wBTC, swap.address)
                    assert.equal(balance.toString(), expectedSats, "Balance is correct in contract's address in tokens[][]")

                    //check user balance on swap contract
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance.toString(), "0", "Uni token has been swapped and user no longer holds it in tokens[][]")

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, sampleTxs[0], "DestAddr for BTC is correct")
                    assert.equal(args.AmountWBTC.toString(), expectedSats, "AmountWBTC is correct amount")
                    assert.equal(args.RefundAddr, user1.address, "RefundAddr is correct")
                    args.Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 1, "1 item in data array returned")
                    assert.equal(data[0].SwapID.toString(), TX_ID, "Swap ID is correct")
                    assert.equal(data[0].DestAddr, receivingBTC_Addr, "Dest BTC addr is correct")
                    assert.equal(data[0].AmountWBTC.toString(), expectedSats, "Output amount matches float amount")
                    data[0].Timestamp.toString().length.should.be.at.least(1, "Timestamp is present")
                })//new contract methods - flow 2

            })//END FLOW 2

        })//END PARASWAP
    })//END SKYPOOLS FUNCTIONS
})//END SKYPOOLS


//https://i.imgur.com/WlUaUXG.png