const ERC20ABI = require('../scripts/erc20ABI.js')
const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers')
const { expect, assert, should } = require('chai')
const { BigNumber, Ethers } = require('ethers')
const { ZERO_ADDRESS } = constants
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18
const utils = require('ethers').utils
const ParaSwap = require('../scripts/paraswap.js')
const { web3 } = require('@openzeppelin/test-helpers/src/setup')
const paraV5 = "https://apiv5.paraswap.io"
const paraswap = new ParaSwap(paraV5, 5)
const rewardsAddr = "0xe1Bb00f0C0c36135A644b44D422Af60F9b7D490d"
const tempSwapAddr = "0xfbC22278A96299D91d41C453234d97b4F5Eb9B2d"
//Web3 init
//const LPToken = artifacts.require('LPToken')
//const SwapContract = artifacts.require('SwapContract')


require('chai')
    .use(require('chai-as-promised'))
    .should()

describe("SkyPools", () => {
    let LPTokenFactory, SwapContractFactory, SwapRewardsFactory, swapRewards, lptoken, swap, owner, receiver, accounts

    let convertScale, lptDecimals, btctTest, btctDecimals, mint500ERC20tokens, balance, zeroFees, minerFees, floatAmount, sampleTxs, redeemedFloatTxIds

    const sbBTCPool = "0xbA99c822bb4dd80f046a75EE564f8295D44Da743"
    const wETH = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" //https://etherscan.io/token/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
    const BLOCKNUM = 13578738 //pinned block number for all testing data - paraswap TXs are with respect to this block

    const srcAmountETH = "1000000000000000000"//1 ETHER
    const srcAmountBTC = "10000000"//.1 wBTC

    //for Flow 2 - hex of utf-8 string BTC addr tb1p5cyxnuxmeuwuvkwfem96lqzszd02n6xdcjrs20cac6yqjjwudpxqp3mvzv
    //const receivingBTC_Addr = "0x74623170356379786e75786d65757775766b7766656d39366c717a737a6430326e367864636a727332306361633679716a6a77756470787170336d767a76"
    const receivingBTC_Addr = "bc1pt5rdh83k9v04kjcdkuw6f6072f3yaregu087hk3grpk4uhjauujspg9w6s"

    //Store hard coded transactions from API PUT request
    const simpleDataFlow1 = [
        '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599', //fromToken
        '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984', //toToken
        '10000000', //fromAmount
        '24666424179035964748', //toAmount
        '49332848358071929496', //expectedAmount - minimum expected amount
        ['0xF9234CB08edb93c0d4a4d4c70cC3FfD070e78e07',
            '0xE592427A0AEce92De3Edee1F18E0157C05861564'], //callees
        '0x91a32b690000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c59900000000000000000000000000000000000000000000000000000000007a12000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000004de5aa873c9da6541f13c89416c17271b4c21bf7b2d7414bf3890000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c5990000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f9840000000000000000000000000000000000000000000000000000000000000bb8000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee5700000000000000000000000000000000000000000000000000000000618c5f0300000000000000000000000000000000000000000000000000000000001e848000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000', //exchangeData
        [0, 228, 488], //startIndexes
        ['0', '0'], //values
        '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC', //beneficiary
        '0x000000000000000000000000536b79506F6f6c73', //partner
        '0', //feePercent
        '0x', //permit
        '1636607315', //deadline
        '0x219c6130427b11ec944a9952ba50eb97'] //uuid

    const simpleDataFlow1ETH = ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599', //wBTC
        '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE', //naitive ETH
        '10000000',
        '690231785171252979',
        '1380463570342505957',
        ['0xdef1c0ded9bec7f1a1670819833240f027b25eff',
            '0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57'],
        '0xaa77476c000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c599000000000000000000000000000000000000000000000000135d7b001eb3b40000000000000000000000000000000000000000000000000000000000009a4494000000000000000000000000b3c839dbde6b96d37c56ee4f9dad3390d49310aa000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee570000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000619c09880000000000000000000000000000000000000000000000000000017d49853b400000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000001cdf10b297e8aea756ba184e37aeb050cd7a7336323790e7f0e447ecc54e2bef704795fe48ffcbc0d2fdf0f0c642737f7e758e8e22367e07d75a6a6a0090243dcf0000000000000000000000000000000000000000000000000000000000989680e1829cfe000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
        [0, 484, 520],
        ['0', '0'],
        '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
        '0x000000000000000000000000536b79506F6f6c73',
        '0',
        '0x',
        '1637616008',
        '0xb0e496c04bd911ecbcc39ff5b6a635aa']

    const simpleDataFlow2 = ['0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2', //wETH9 https://etherscan.io/address/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
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

    const simpleDataFlow2B = ['0x1f9840a85d5af5bf1d1762f925bdaddc4201f984', //UNI token 
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
    const swapOnUniswapForkFlow1 = [
        '0x115934131916C8b277DD010Ee02de363c09d037c', //factory
        '0x65d1a3b1e46c6e4f1be1ad5f99ef14dc488ae0549dc97db9b30afe2241ce1c7a', //initCode
        '10000000', //amountIn
        '138090143398971911407', //amountOutMin
        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
            '0x1f9840a85d5af5bf1d1762f925bdaddc4201f984']] //path - first index is fromToken, last index is ending token
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

    const swapOnUniswapFlow2 = [
        '1000000000000000000', //amountIn
        '19439', //amountOutMin
        ['0x1f9840a85d5af5bf1d1762f925bdaddc4201f984',
            '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599']] //path - first index is fromToken, last index is ending token


    const endAmount = "100000000" //1 BTCt - decimal 8
    const btcStartingAmount = "500000000" //5 BTC decimal 8
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
        SwapRewardsFactory = await ethers.getContractFactory("SwapRewards")


        lpToken = await LPTokenFactory.deploy(8)

        lptDecimals = await lpToken.decimals()

        btctTest = await LPTokenFactory.deploy(TOKEN_DECIMALS)

        btctDecimals = await btctTest.decimals()

        btctDecimals = new BigNumber.from(btctDecimals)

        convertScale = new BigNumber.from(10).pow(btctDecimals.sub(lptDecimals)) //convert BTC decimals to ERC20

        mint500ERC20tokens = new BigNumber.from(500).mul(new BigNumber.from(10).pow(btctDecimals))

        swapRewards = await SwapRewardsFactory.deploy(
            owner.address,
            lpToken.address,
            tempSwapAddr
        )

        swap = await SwapContractFactory.deploy(
            lpToken.address,
            btctTest.address,
            wETH,
            sbBTCPool,
            swapRewards.address,
            0
        )

        await swapRewards.connect(owner).setSwap(
            swap.address,
            new BigNumber.from(30),
            new BigNumber.from(60)
        )

        zeroFees = false

        minerFees = new BigNumber.from(30000)

        floatAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(btctDecimals))

        sampleTxs = [
            "0x1c12443203a48f42cdf7b1acee5b4b1c1fedc144cb909a3bf5edbffafb0cd204",
            "0x6a167c4b6750c3213320098178f913478fe50d3f75d5f0377ee7cec9a630ad9e",
            "0xbceaa7c52bcb637ddbb7dab980ec8e015f259b3aa4f8b4c4115fd1bcb4a5779c",
            "0x000000000000000000033d05abcee8adbd9897cdcf184e135191dc06b095534b",
            "0xc33e7f89ed85bbad177f4238608360490a0accfdb1d7440b2bcd4a10db085c91",
            "0x50818e936aB61377A18bCAEc0f1C32cA27E389230f1C32cA27E389230f1C32cA"
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

        describe("Testing recording of SkyPools TXs", () => {
            beforeEach(async () => {
                //mint tokens and assign to owner so they can be transferred from the contract
                await btctTest.mint(owner.address, mint500ERC20tokens)
                await btctTest.connect(owner).transfer(swap.address, mint500ERC20tokens)
            })
            it('records SkyPools Transaction and allows user to redeem BTCT tokens', async () => {
                //set float reserve
                let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + owner.address.slice(2), 64)

                await swap.recordIncomingFloat(btctTest.address, addressesAndAmountOfFloat, sampleTxs[0])

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

            it("records batched SkyPools TXs and allocates BTCT to users correctly in tokens[][]", async () => {
                //set float reserve
                let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + owner.address.slice(2), 64)

                await swap.recordIncomingFloat(
                    btctTest.address,
                    addressesAndAmountOfFloat,
                    "0xce66450451e62b9b4c406d0a83b90a5036039673d2682d4ec292f375ae571382"
                )

                //amount = 1 BTC, adjusted for 8 decimals
                let amount = new BigNumber.from(10000000) //.1 BTC

                //prepare to multi record
                let swapFeesBPS = new BigNumber.from(20);
                let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(btctDecimals))
                let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                let rewards = amount.add(amount).mul(swapFeesBPS).div(new BigNumber.from(10000))


                let addrAndAmount1 = web3.utils.padLeft(amount._hex + user1.address.slice(2), 64)
                let addrAndAmount2 = web3.utils.padLeft(amount._hex + user2.address.slice(2), 64)


                const TXs = [
                    addrAndAmount1,
                    addrAndAmount2
                ]

                const txIDs = [
                    "0x23e8785fe862e60f2caa4f738146ff9d4f4bd4a02dd6fb1f513b0a9be3452b64",
                    "0xce25470451e62b9b4a406d0a83b90a5036039673d2682d4ec292f375ae571382"
                ]
                const totalSwapped = amount.mul(2)


                balance = await swap.balanceOf(btctTest.address, user1.address)
                assert.equal(balance.toString(), 0, "User1 has 0 balance of BTCT before multi record")

                balance = await swap.balanceOf(btctTest.address, user2.address)
                assert.equal(balance.toString(), 0, "User2 has 0 balance of BTCT before multi record")

                //multiRecord
                await swap.connect(owner).multiRecordSkyPoolsTX(
                    TXs,
                    totalSwapped,
                    rewards,
                    txIDs
                )

                balance = await swap.balanceOf(btctTest.address, user1.address)
                assert.equal(balance.toString(), amount.toString(), `User1 has ${amount.toString()} sats after multiRecord`)

                balance = await swap.balanceOf(btctTest.address, user2.address)
                assert.equal(balance.toString(), amount.toString(), `User2 has ${amount.toString()} sats after multiRecord`)

            })
        })//recording SkyPools TXs

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

                swapRewards = await SwapRewardsFactory.deploy(
                    owner.address,
                    lpToken.address,
                    tempSwapAddr
                )

                swap = await SwapContractFactory.deploy(lpToken.address, wBTC, wETH, sbBTCPool, swapRewards.address, 0)


                await swapRewards.connect(owner).setSwap(
                    swap.address,
                    new BigNumber.from(30),
                    new BigNumber.from(60)
                )

                await lpToken.transferOwnership(swap.address)
                expect(await lpToken.getOwner()).to.equal(swap.address)
                expect(await lpToken.owner()).to.equal(swap.address)

            })//beforeEach - PARASWAP

            describe('executes paraSwap transactions: Flow 1 => BTC -> wBTC -> ERC20', async () => {
                beforeEach(async () => {
                    //Allocate wBTC tokens to user in tokens[][]
                    await populateBalance(swap.address, wBTC, wBTC_SLOT, btcStartingAmount)

                    let decimals = Tokens[mainnetNetworkID]['ETH'].decimals
                    //set float reserve
                    const newFloatAmount = new BigNumber.from("500000000")//5 tokens ERC20 - decimal 18
                    let addressesAndAmountOfFloat = web3.utils.padLeft(newFloatAmount.add(minerFees)._hex + owner.address.slice(2), 64)
                    await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, sampleTxs[0])

                    //perform recordSkyPoolsTX to assign user1 btct tokens in the contract
                    let swapFeesBPS = new BigNumber.from(20);
                    let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                    await swap.recordSkyPoolsTX(user1.address, endAmount, swapFees, sampleTxs) // todo: need to add rewards

                })
                it('simpleSwap flow 1', async () => {
                    //Execute paraswap TX
                    const result = await swap.connect(user1).spFlow1SimpleSwap(
                        simpleDataFlow1
                    )

                    //Check ending balances
                    balance = await UNI_Contract.balanceOf(user1.address)
                    assert(balance > new BigNumber.from(0), "User has UNI tokens in wallet")

                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance.toString(), "0", "User balance of UNI tokens on swap contract is 0")

                })//SimpleSwap flow 1

                it('simpleSwap ETH flow 1', async () => {

                    const initEthBalance = await ethers.provider.getBalance(user1.address)

                    //Execute paraswap TX
                    const result = await swap.connect(user1).spFlow1SimpleSwap(
                        simpleDataFlow1ETH
                    )

                    //Check ending balances
                    balance = await ethers.provider.getBalance(user1.address)
                    assert(balance > new BigNumber.from(initEthBalance), "User has received ETH")

                })//simpleSwap ETH flow 1

                it('swapOnUniswapFork flow 1', async () => {

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

                it('swapOnUniswap flow 1 -> wETH and withdraw naitive ETH', async () => {
                    //Use path wBTC -> wETH, user withdraws ETH to their wallet

                    const initEthBalance = await ethers.provider.getBalance(user1.address)
                    const ETHER = Tokens[mainnetNetworkID]['ETH'].address
                    const flow1ToETHUniswap = ['0x115934131916C8b277DD010Ee02de363c09d037c',
                        '0x65d1a3b1e46c6e4f1be1ad5f99ef14dc488ae0549dc97db9b30afe2241ce1c7a',
                        '10000000',
                        '659829776477491556',
                        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
                            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2']]

                    await swap.connect(user1).spFlow1Uniswap(
                        true,
                        flow1ToETHUniswap[0],
                        flow1ToETHUniswap[1],
                        flow1ToETHUniswap[2],
                        flow1ToETHUniswap[3],
                        flow1ToETHUniswap[4],
                    )

                    balance = await swap.balanceOf(wETH, user1.address)
                    const receivedAmount = balance
                    assert(receivedAmount > 0, "User received wETH to their allocation on the swap contract")

                    //withdraw and receive ETH
                    const result = await swap.connect(user1).redeemEther(receivedAmount)
                    balance = await swap.balanceOf(wETH, user1.address)
                    assert.equal(balance.toString(), "0", "Balance is correct after depositing ETHER")
                    const receipt = await result.wait()
                    const event = receipt.events[receipt.events.length - 1].args

                    //verify event receipt is correct
                    event.token.toString().should.equal(ETHER, 'token is correct')
                    event.user.toString().should.equal(user1.address, 'user address is correct')
                    event.amount.toString().should.equal(receivedAmount.toString(), 'amount is correct')
                    event.balance.toString().should.equal("0", 'balance is correct')
                    event.Timestamp.toString().length.should.be.at.least(1, 'timestamp is present')

                    //verify user balance is correct 
                    balance = await ethers.provider.getBalance(user1.address)
                    assert(balance > initEthBalance, "User received ETHER to their personal wallet")

                })


                it('allows for multiple flow 1 TXs without failure', async () => {

                    //repeat all previous TXs at once
                    const spFlow1SimpleSwapResult = await swap.connect(user1).spFlow1SimpleSwap(
                        simpleDataFlow1
                    )

                    const spFlow1SimpleSwapETHResult = await swap.connect(user1).spFlow1SimpleSwap(
                        simpleDataFlow1ETH
                    )

                    await swap.connect(user1).spFlow1Uniswap(
                        true,
                        swapOnUniswapForkFlow1[0],
                        swapOnUniswapForkFlow1[1],
                        swapOnUniswapForkFlow1[2],
                        swapOnUniswapForkFlow1[3],
                        swapOnUniswapForkFlow1[4],
                    )

                    await swap.connect(user1).spFlow1Uniswap(
                        false,
                        swapOnUniswapForkFlow1[0],
                        swapOnUniswapForkFlow1[1],
                        swapOnUniswapFlow1[0],
                        swapOnUniswapFlow1[1],
                        swapOnUniswapFlow1[2],
                    )

                    //more TXs including less tokens w/o direct swap from wBTC
                    //Contract methods were selected by paraswap API based on market conditions at the time of testing

                    //swap wBTC -> COMP via swapOnUniswapFork
                    const flow1ToCOMP = ['0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac',
                        '0xe18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303',
                        '10000000',
                        '9784939186051670608',
                        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
                            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
                            '0xc00e94cb662c3520282e6f5717214004a7f26888']]

                    await swap.connect(user1).spFlow1Uniswap(
                        true,
                        flow1ToCOMP[0],
                        flow1ToCOMP[1],
                        flow1ToCOMP[2],
                        flow1ToCOMP[3],
                        flow1ToCOMP[4],
                    )
                    //CHECK BALANCE TODO


                    //swap wBTC -> USDC via swapOnUniswapFork
                    const flow1ToUSDC = ['0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac',
                        '0xe18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303',
                        '10000000',
                        '2801097984',
                        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
                            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
                            '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48']]

                    await swap.connect(user1).spFlow1Uniswap(
                        true,
                        flow1ToUSDC[0],
                        flow1ToUSDC[1],
                        flow1ToUSDC[2],
                        flow1ToUSDC[3],
                        flow1ToUSDC[4],
                    )
                    //CHECK BALANCE TODO

                    //swap wBTC -> USDC via swapOnUniswap
                    const flow1ToMATTIC = ['10000000',
                        '1799754185989227903071',
                        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
                            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
                            '0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0']]

                    await swap.connect(user1).spFlow1Uniswap(
                        false,
                        flow1ToUSDC[0],
                        flow1ToUSDC[1],
                        flow1ToMATTIC[0],
                        flow1ToMATTIC[1],
                        flow1ToMATTIC[2]
                    )
                    //CHECK BALANCE TODO

                    //swap wBTC -> MKR via simpleSwap
                    const flow1ToMKR = ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
                        '0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2',
                        '10000000',
                        '983574506400367217',
                        '1967149012800734433',
                        ['0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0'],
                        '0x569706eb00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000098968000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000050000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c599000000000000000000000000fee7eeaa0c2f3f7c7e6301751a8de55ce4d059ec0000000000000000000000001f573d6fb3f13d689ff844b4ce37794d79a7ff1c000000000000000000000000f553e6ea4ce2f7deecbe7837e27931850ec15fab0000000000000000000000009f8f72aa9304c8b593d555f12ef6589cc3a579a2',
                        [0, 356],
                        ['0'],
                        '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',
                        '0x000000000000000000000000536b79506F6f6c73',
                        '0',
                        '0x',
                        '1637641451',
                        '0xeaae5b304be211ec9f59dd3d48aed4dd']

                    await swap.connect(user1).spFlow1SimpleSwap(
                        flow1ToMKR
                    )
                    //CHECK BALANCE TODO

                    //swap wBTC -> SWINGBY via swapOnUniswap
                    const flow1toSWINGBY = ['10000000',
                        '38608541168427523513677',
                        ['0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
                            '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
                            '0x8287c7b963b405b7b8d467db9d79eec40625b13a']]

                    await swap.connect(user1).spFlow1Uniswap(
                        false,
                        flow1ToUSDC[0],
                        flow1ToUSDC[1],
                        flow1toSWINGBY[0],
                        flow1toSWINGBY[1],
                        flow1toSWINGBY[2]
                    )
                    //CHECK BALANCE TODO

                })


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
                let balance, result, receipt, event
                const ETHER = Tokens[mainnetNetworkID]['ETH'].address
                const amount = utils.parseEther("1")
                const overrides = {
                    value: amount.mul(5)
                }

                const expectedFloats = new BigNumber.from(101032000)

                beforeEach(async () => {
                    //set float reserve
                    decimals = Tokens[mainnetNetworkID]['wBTC'].decimals
                    floatAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                    let addressesAndAmountOfFloat = web3.utils.padLeft(floatAmount.add(minerFees)._hex + owner.address.slice(2), 64)


                    await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, sampleTxs[0])
                    await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, sampleTxs[1])



                    //prep to collectSwapFeesForBTC
                    let swapFeesBPS = new BigNumber.from(20);
                    //let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))

                    let swapAmount = new BigNumber.from(1000000) //.01 BTC
                    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                    let rewardsAmount = swapFees.sub(minerFees)
                    let incomingAmount = swapAmount.add(swapFees) //100200000 sats = 1.002 BTC                
                    let spenders = [rewardsAddr]
                    let amounts = [incomingAmount]
                    //allocate floats to wBTC ~1wBTC decimal 8

                    let test = await swap.swapRewards()

                    
                     await swap.connect(owner).collectSwapFeesForBTC(
                        incomingAmount,
                        minerFees,
                        swapFees,
                        spenders,
                        amounts,
                        true
                    )
                     

                    /////////////////////////////// DEPOSIT UNI TOKENS //////////////////////////////////////////////
                    await populateBalance(user1.address, UNI, UNI_SLOT, amount.mul(5))//amount refers to number of UNI tokens here
                    await UNI_Contract.connect(user1).approve(swap.address, amount)
                    result = await swap.connect(user1).spDeposit(UNI, amount)

                    /////////////////////////////// TESTING DEPOSIT ETHER //////////////////////////////////////////////                
                    result = await swap.connect(user1).spDeposit("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", amount, overrides)
                    balance = await swap.balanceOf(wETH, user1.address)
                    assert.equal(balance.toString(), amount.mul(5).toString(), "Balance of wETH is correct after depositing ETHER")
                    receipt = await result.wait()

                    event = receipt.events[receipt.events.length - 1].args

                    //verify event receipt is correct
                    event.token.toString().should.equal(ETHER, 'token is correct')
                    event.user.toString().should.equal(user1.address, 'user address is correct')
                    event.amount.toString().should.equal(amount.mul(5).toString(), 'amount is correct')
                    event.balance.toString().should.equal(amount.mul(5).toString(), 'balance is correct')
                    event.Timestamp.toString().length.should.be.at.least(1, 'timestamp is present')




                })

                it("converts wETH to naitive ETH on withdraw", async () => {
                    const initEthBalance = await ethers.provider.getBalance(user1.address)
                    result = await swap.connect(user1).redeemEther(amount)
                    balance = await swap.balanceOf(ETHER, user1.address)
                    assert.equal(balance.toString(), "0", "Balance is correct after depositing ETHER")
                    receipt = await result.wait()
                    event = receipt.events[receipt.events.length - 1].args

                    //verify event receipt is correct
                    event.token.toString().should.equal(ETHER, 'token is correct')
                    event.user.toString().should.equal(user1.address, 'user address is correct')
                    event.amount.toString().should.equal(amount.toString(), 'amount is correct')
                    event.balance.toString().should.equal(amount.mul(4).toString(), 'balance is correct')
                    event.Timestamp.toString().length.should.be.at.least(1, 'timestamp is present')

                    //verify user balance is correct 
                    balance = await ethers.provider.getBalance(user1.address)
                    assert(balance > initEthBalance, "User received ETHER to their personal wallet")
                })

                it('simpleSwap Flow 2a => ETH -> wETH -> wBTC -> BTC', async () => {
                    /////////////////////////////// EXECUTE SWAP AND RECORD //////////////////////////////////////////////
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(balance[1].toString(), expectedFloats.toString(), "Float Reserve of BTCT tokens on the contract BEFORE spFlow2SimpleSwap is correct")

                    balance = await wETH_Contract.balanceOf(swap.address)
                    assert.equal(balance.toString(), amount.mul(5).toString(), "Balance of wETH on swap contract before swap is correct")


                    //ensure beneficiary is current swap contract address
                    simpleDataFlow2[9] = swap.address


                    //perform swap
                    const swapAndRecordResult = await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    /////////////////////////////// CHECK ENDING BALANCES //////////////////////////////////////////////

                    balance = await wETH_Contract.balanceOf(swap.address)
                    assert.equal(balance.toString(), amount.mul(4).toString(), "Balance of wETH on swap contract after swap is correct")


                    const expectedSats = '7111530'

                    receipt = await swapAndRecordResult.wait()

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, receivingBTC_Addr, "DestAddr for BTC is correct")
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
                    await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, sampleTxs[2])
                    await swap.recordIncomingFloat(ZERO_ADDRESS, addressesAndAmountOfFloat, sampleTxs[3])
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)


                    /**
                      //prep to collectSwapFeesForBTC
                     let swapFeesBPS = new BigNumber.from(20);
                     let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))
                     let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                     let incomingAmount = swapAmount.add(swapFees) //100200000 sats = 1.002 BTC                
 
                     //allocate floats to wBTC ~1wBTC decimal 8
                     await swap.collectSwapFeesForBTC(ZERO_ADDRESS, incomingAmount, minerFees, swapFees, true)
                     
                     */


                    //prep to collectSwapFeesForBTC
                    let swapFeesBPS = new BigNumber.from(20);
                    //let swapAmount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(decimals))

                    let swapAmount = new BigNumber.from(1000000) //.01 BTC
                    let swapFees = swapAmount.mul(swapFeesBPS).div(new BigNumber.from(10000))
                    let rewardsAmount = swapFees.sub(minerFees)
                    let incomingAmount = swapAmount.add(swapFees) //100200000 sats = 1.002 BTC                
                    let spenders = [rewardsAddr]
                    let swapAmounts = [new BigNumber.from(5)]
                    //allocate floats to wBTC ~1wBTC decimal 8

                    let test = await swap.swapRewards()

                    await swap.collectSwapFeesForBTC(
                        incomingAmount,
                        minerFees,
                        swapFees,
                        spenders,
                        swapAmounts,
                        true
                    )
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)

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

                    //ensure beneficiary is current swap contract address
                    simpleDataFlow2B[9] = swap.address

                    const swapAndRecordResult = await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2B
                    )
                    receipt = await swapAndRecordResult.wait()


                    /////////////////////////////// CHECK BALANCES AND RECORDED DATA //////////////////////////////////////////////

                    const expectedSats = '38878'

                    //check user balance on swap contract
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance.toString(), "0", "Uni token has been swapped and user no longer holds it in tokens[][]")

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, receivingBTC_Addr, "DestAddr for BTC is correct")
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

                    //fork = true  -  use paraswap function swapOnUniswapFork
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

                    const expectedSats = '38878'



                    //check user balance on swap contract
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance.toString(), "0", "Uni token has been swapped and user no longer holds it in tokens[][]")

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, receivingBTC_Addr, "DestAddr for BTC is correct")
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

                    const expectedSats = '38878'



                    //check user balance on swap contract
                    balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                    assert.equal(balance.toString(), "0", "Uni token has been swapped and user no longer holds it in tokens[][]")

                    //check data from event emitted 
                    event = receipt.events[receipt.events.length - 1]
                    let args = event.args
                    assert.equal(event.event, "SwapTokensToBTC", "Correct event name")
                    args.SwapID.toString().length.should.be.at.least(1, "SwapID is present")
                    const TX_ID = args.SwapID
                    assert.equal(args.DestAddr, receivingBTC_Addr, "DestAddr for BTC is correct")
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

                it('cleans up old TXs and correctly moves the indexes', async () => {
                    /////////////////////////////// EXECUTE SWAP AND RECORD //////////////////////////////////////////////
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(balance[1].toString(), expectedFloats.toString(), "Float Reserve of BTCT tokens on the contract BEFORE spFlow2SimpleSwap is correct")

                    balance = await wETH_Contract.balanceOf(swap.address)
                    assert.equal(balance.toString(), amount.mul(5).toString(), "Balance of wETH on swap contract before swap is correct")

                    //perform swap eth -> wBTC
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 1, "1 item in data array returned")

                    //perform swap eth -> wBTC
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 2, "2 items in data array returned")

                    //perform swap UNI -> wBTC
                    await swap.connect(user1).spFlow2Uniswap(
                        receivingBTC_Addr,
                        true,
                        swapOnUniswapForkFlow2[0],
                        swapOnUniswapForkFlow2[1],
                        swapOnUniswapForkFlow2[2],
                        swapOnUniswapForkFlow2[3],
                        swapOnUniswapForkFlow2[4],
                    )

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 3, "3 items in data array returned")

                    //check swapCount index before advancing time
                    let swapCount = await swap.swapCount()
                    assert.equal(swapCount, 3, "swapCount is correct")

                    let oldestActiveIndex = await swap.oldestActiveIndex()
                    assert.equal(oldestActiveIndex, 0, "oldestActiveIndex is correct")

                    const TwoDaysSeconds = 172800
                    //ADVANCE 2 DAYS
                    let currentBlock = await ethers.provider.getBlockNumber()
                    await ethers.provider.send("evm_increaseTime", [TwoDaysSeconds]) //set EVM time, will be applied to next block
                    await ethers.provider.send("evm_mine") //next block
                    currentBlock = await ethers.provider.getBlockNumber()
                    let blockObj = await ethers.provider.getBlock(currentBlock) //https://docs.ethers.io/v5/api/providers/types/#providers-Block

                    /////////////////////////////// DO ANOTHER TX - SHOULD CLEAN UP THE FIRST 3 //////////////////////////////////////////////

                    //check swapCount index before advancing time
                    swapCount = await swap.swapCount()
                    assert.equal(swapCount, 3, "swapCount is correct after elapsed time")

                    oldestActiveIndex = await swap.oldestActiveIndex()
                    assert.equal(oldestActiveIndex, 0, "oldestActiveIndex is correct after elapsed time")

                    //perform swap eth -> wBTC
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    //perform swap eth -> wBTC again 
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 2, "The first 3 pending swaps have been cleaned up, the newest 2 remain")

                    swapCount = await swap.swapCount()
                    assert.equal(swapCount, 5, "5 total swaps, swapCount index is correct")

                    oldestActiveIndex = await swap.oldestActiveIndex()
                    assert.equal(oldestActiveIndex, 3, "oldest active index is correct")


                    //ADVANCE 2 DAYS
                    currentBlock = await ethers.provider.getBlockNumber()
                    await ethers.provider.send("evm_increaseTime", [TwoDaysSeconds]) //set EVM time, will be applied to next block
                    await ethers.provider.send("evm_mine") //next block
                    currentBlock = await ethers.provider.getBlockNumber()
                    blockObj = await ethers.provider.getBlock(currentBlock) //https://docs.ethers.io/v5/api/providers/types/#providers-Block

                    /////////////////////////////// DO ANOTHER TX - SHOULD CLEAN UP PREVIOUS 2 //////////////////////////////////////////////

                    //swapOnUniswapFork wETH -> wBTC
                    const finalTX = ['0x75e48C954594d64ef9613AeEF97Ad85370F13807',
                        '0xb2b53dca60cae1d1f93f64d80703b888689f28b63c483459183f2f4271fa0308',
                        '100000000000000000',
                        '408244',
                        ['0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
                            '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599']]

                    await swap.connect(user1).spFlow2Uniswap(
                        receivingBTC_Addr,
                        true,
                        finalTX[0],
                        finalTX[1],
                        finalTX[2],
                        finalTX[3],
                        finalTX[4],
                    )

                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 1, "The first 5 pending swaps have been cleaned up, only the newest one remains")

                    swapCount = await swap.swapCount()
                    assert.equal(swapCount, 6, "6 total swaps, swapCount index is correct")

                    oldestActiveIndex = await swap.oldestActiveIndex()
                    assert.equal(oldestActiveIndex, 5, "oldest active index is correct")
                })
                it('allows for manual cleanup of old TXs after changing expiration times', async () => {
                    /////////////////////////////// EXECUTE SWAP AND RECORD //////////////////////////////////////////////
                    balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                    assert.equal(balance[1].toString(), expectedFloats.toString(), "Float Reserve of BTCT tokens on the contract BEFORE spFlow2SimpleSwap is correct")

                    balance = await wETH_Contract.balanceOf(swap.address)
                    assert.equal(balance.toString(), amount.mul(5).toString(), "Balance of wETH on swap contract before swap is correct")

                    //perform swap eth -> wBTC
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    //perform swap eth -> wBTC
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    //perform swap UNI -> wBTC
                    await swap.connect(user1).spFlow2Uniswap(
                        receivingBTC_Addr,
                        true,
                        swapOnUniswapForkFlow2[0],
                        swapOnUniswapForkFlow2[1],
                        swapOnUniswapForkFlow2[2],
                        swapOnUniswapForkFlow2[3],
                        swapOnUniswapForkFlow2[4],
                    )

                    //check spGetPendingSwaps()
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 3, "3 items in data array returned")

                    const TwoDaysSeconds = 172800
                    //ADVANCE 2 DAYS
                    let currentBlock = await ethers.provider.getBlockNumber()
                    await ethers.provider.send("evm_increaseTime", [TwoDaysSeconds]) //set EVM time, will be applied to next block
                    await ethers.provider.send("evm_mine") //next block
                    currentBlock = await ethers.provider.getBlockNumber()
                    let blockObj = await ethers.provider.getBlock(currentBlock) //https://docs.ethers.io/v5/api/providers/types/#providers-Block

                    /////////////////////////////// DO ANOTHER TX - SHOULD CLEAN UP THE FIRST 3 //////////////////////////////////////////////

                    //check swapCount index before advancing time
                    swapCount = await swap.swapCount()
                    assert.equal(swapCount, 3, "swapCount is correct after elapsed time")

                    oldestActiveIndex = await swap.oldestActiveIndex()
                    assert.equal(oldestActiveIndex, 0, "oldestActiveIndex is correct after elapsed time")

                    //perform swap eth -> wBTC
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    //perform swap eth -> wBTC again 
                    await swap.connect(user1).spFlow2SimpleSwap(
                        receivingBTC_Addr,
                        simpleDataFlow2
                    )

                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 2, "The first 3 pending swaps have been cleaned up, the newest 2 remain")

                    /////////////////////////////// CHURN AND SHORTEN EXPIRATION TIME //////////////////////////////////////////////
                    let currentExpirationTime = await swap.expirationTime()
                    assert.equal(currentExpirationTime, 172800, "Expiration time is 2 days")

                    let nodes = ["0x8bAf8b6Ed0E0ddB6557Af1A7391a86949FAFa3a8"]
                    let isRemoved = [false]
                    let churnedInCount = 25
                    let tssThreshold = 16
                    let nodeRewardsRatio = 66
                    let withdrawalFeeBPS = new BigNumber.from(20);
                    let newExpirationTime = new BigNumber.from(86400) //1 day seconds
                    /*
                    for (i = 0; i < 1; i++) {
                        let staked = new BigNumber.from(3000000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                        let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + owner.address.slice(2), 64)
                        nodes.push(addressesAndAmountStaked)
                        isRemoved.push(false)
                    }
                    */


                    await swap.churn(
                        owner.address,
                        nodes,
                        isRemoved,
                        churnedInCount,
                        tssThreshold,
                        nodeRewardsRatio,
                        withdrawalFeeBPS,
                        new BigNumber.from(50),
                        new BigNumber.from(0),
                        newExpirationTime
                    )

                    currentExpirationTime = await swap.expirationTime()
                    assert.equal(currentExpirationTime, 86400, "Expiration time is 1 day")

                    /////////////////////////////// MANUAL CLEANUP OF OLD TXS WITH HIGH LOOP COUNT //////////////////////////////////////////////

                    await swap.spCleanUpOldTXs(data.length + 5)

                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 2, "No TXs cleaned up yet, expiration time not yet elapsed")

                    //ADVANCE 2 DAYS
                    currentBlock = await ethers.provider.getBlockNumber()
                    await ethers.provider.send("evm_increaseTime", [TwoDaysSeconds]) //set EVM time, will be applied to next block
                    await ethers.provider.send("evm_mine") //next block
                    currentBlock = await ethers.provider.getBlockNumber()
                    blockObj = await ethers.provider.getBlock(currentBlock) //https://docs.ethers.io/v5/api/providers/types/#providers-Block

                    await swap.spCleanUpOldTXs(data.length + 5)
                    data = await swap.spGetPendingSwaps()
                    assert.equal(data.length, 0, "All TXs have been cleaned up")


                })
                /////////////////////////////// TEST FOR FAILURE //////////////////////////////////////////////
                describe('Testing for flow 2 failure cases', async () => {
                    it('Existing BTC floats must be sufficient', async () => {
                        //swapOnUniswap
                        const giantAmountData = ['15000000000000000000',
                            '65138277',
                            ['0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
                                '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599']]



                        const amount = utils.parseEther("1")
                        const overrides = {
                            value: amount.mul(16)
                        }


                        await swap.connect(user1).spDeposit("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", amount.mul(16), overrides)


                        await swap.connect(user1).spFlow2Uniswap(
                            receivingBTC_Addr,
                            false,
                            swapOnUniswapForkFlow1[0],
                            swapOnUniswapForkFlow1[1],
                            giantAmountData[0],
                            giantAmountData[1],
                            giantAmountData[2],
                        ).should.be.rejectedWith("reverted with reason string \'12\'")
                    })
                })//END FLOW 2 FAILURE
            })//END FLOW 2
        })//END PARASWAP
    })//END SKYPOOLS FUNCTIONS
})//END SKYPOOLS


//https://i.imgur.com/WlUaUXG.png