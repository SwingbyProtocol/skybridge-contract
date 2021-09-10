const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers')
const { expect, assert } = require('chai')
const { BigNumber, Ethers } = require('ethers')
const { ZERO_ADDRESS } = constants
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18
const utils = require('ethers').utils
const ParaSwap = require('../scripts/paraswap.js')
//Web3 init
//const LPToken = artifacts.require('LPToken')
//const SwapContract = artifacts.require('SwapContract')

describe("SkyPools", () => {
    let LPTokenFactory, SwapContractFactory, lptoken, swap, owner, receiver, accounts

    let convertScale, lptDecimals, testToken, btctTest, btctDecimals, mint500ERC20tokens, balance, zeroFees, minerFees, floatAmount, sampleTxs, redeemedFloatTxIds

    beforeEach(async () => {
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

        swap = await SwapContractFactory.deploy(lpToken.address, btctTest.address, 0)

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
    describe("SkyPool functions", () => {

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

            //console.log("Starting Float amount: ", balance[1].toString())

            balance = await btctTest.balanceOf(swap.address)
            assert.equal(utils.formatEther(balance).toString(), utils.formatEther(mint500ERC20tokens).toString(), "Balance of tokens on swap contract is correct")

            balance = await btctTest.balanceOf(user1.address)
            assert.equal(balance.toNumber(), 0, "User's starting wallet balance is correct")

            balance = await btctTest.balanceOf(swap.address)
            assert.equal(utils.formatEther(balance).toString(), utils.formatEther(mint500ERC20tokens).toString(), "Balance of BTCT tokens in the contract is correct")


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

            //console.log("Ending Float amount: ", balance[1].toString())
            //console.log("Difference: ", (startingFloatAmount.sub(balance[1]).toString()))

            //get users swap balance before they redeem tokens
            balance = await swap.connect(user1.address).balanceOf(btctTest.address)
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
            balance = await swap.connect(user1.address).balanceOf(btctTest.address)
            assert.equal(balance.toNumber(), 0, "User's ending balance on the swap contract is correct")

            balance = await btctTest.balanceOf(user1.address)
            let endingAmountInBTC = amount.div(new BigNumber.from(10).pow(8)) //convert to BTC decimals
            assert.equal(utils.formatEther(balance).toString(), endingAmountInBTC.toNumber(), "User's wallet balance contains the tokens in the correct amount")

        })
        it('executes paraSwap transactions', async () => {
            //test API call
            //https://apiv4.paraswap.io/v2/prices/?from=ETH&to=DAI&amount=10000000000000000000&fromDecimals=18&toDecimals=18&side=SELL&network=137
            const paraAPI_URL = "https://apiv4.paraswap.io/v2"
            const paraV5 = "https://apiv5.paraswap.io"
            const paraswap = new ParaSwap(paraV5, 5)

            const normalize = (amount, token) => {
                return new BigNumber.from(amount).mul(new BigNumber.from(10).pow(token.decimals))
            }
            const denormalize = (amount, token) => {
                return new BigNumber.from(amount).div(new BigNumber.from(10).pow(token.decimals))
            }

            const highValueEthAccount = "0x00000000219ab540356cBB839Cbe05303d7705Fa"
            const swingbyContractAddress = "0xbe83f11d3900f3a13d8d12fb62f5e85646cda45e"
            //impersonate real mainnet account
            await hre.network.provider.request({
                method: "hardhat_impersonateAccount",
                params: [highValueEthAccount]
            })

            const signer = await ethers.getSigner(highValueEthAccount)
            const provider = new ethers.providers.JsonRpcProvider("https://eth-mainnet.alchemyapi.io/v2/YfblHzLI_PlnIA0pxphS90J3yaA3dDi5");
            const swingby = await ethers.getSigner(swingbyContractAddress)


            const mainnetNetworkID = 1
            const REST_TIME = 5 * 1000 // 5 seconds
            const Tokens = {
                [mainnetNetworkID]: {
                    ETH: {
                        address: '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE',
                        decimals: 18,
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
                    }
                },
            }
            const minABI = [
                // balanceOf
                {
                    constant: true,

                    inputs: [{ name: "_owner", type: "address" }],

                    name: "balanceOf",

                    outputs: [{ name: "balance", type: "uint256" }],

                    type: "function",
                },

            ];




            //signer.sendTransaction()
            const wBTC_Contract = new ethers.Contract(Tokens[mainnetNetworkID]['wBTC'].address, minABI, provider)
            const swingbywBTC_Balance = await wBTC_Contract.balanceOf(swingby.address)
            const swingbyETH_Balance = await provider.getBalance(swingbyContractAddress)
            console.log(utils.formatEther(swingbyETH_Balance).toString())

            const srcAmount = normalize(
                1,
                Tokens[mainnetNetworkID]['wBTC']
            )
            let getPrice = await paraswap.getPrice(
                Tokens[mainnetNetworkID]['wBTC'],
                Tokens[mainnetNetworkID]['ETH'],
                srcAmount,
                mainnetNetworkID
            )

            //console.log(getPrice)


            const slippage = (decimals) => {
                return new BigNumber.from(3).mul(new BigNumber.from(10).pow(decimals - 2)).toString() //ERC20 - 0.03
            }


            let decimals = Tokens[mainnetNetworkID]['ETH'].decimals
            let minDestAmount = new BigNumber.from(getPrice.price).sub(slippage(decimals))



            const txRequest = await paraswap.buildTransaction(
                getPrice.payload,
                Tokens[mainnetNetworkID]['wBTC'],
                Tokens[mainnetNetworkID]['ETH'],
                srcAmount.toString(),
                minDestAmount.toString(),
                mainnetNetworkID,
                signer.address,
                true //only params - true for contract -> contract | false for standard transaction object
            )
            
             let data = txRequest.data //params to execute transaction contract -> contract
            //console.log(data)
            const output = txRequest.config.data
            const { parse } = require('json-parser')
            const parsedOutput = parse(output)
            const contractMethod = parsedOutput.priceRoute.contractMethod


            console.log("Expected Amount: ", parsedOutput.destAmount.toString())

            if (decimals == 18) {
                console.log("Formatted ERC-20 amount: ", utils.formatEther(parsedOutput.destAmount).toString())
            }

            console.log("Recomended Contract Method:", contractMethod)//get contract method

            console.log("Arguments for", contractMethod)
            console.log(data)

            const paraSwapAddress = "0x1bD435F3C054b6e901B7b108a0ab7617C808677b"


            if (contractMethod == "swapOnUniswap") {
                let abi = ["function approve(address _spender, uint256 _value) public returns (bool success)"]
                let contract = new ethers.Contract(data.path[0], abi, provider)
                await contract.connect(signer).approve(highValueEthAccount, srcAmount)

                const result = await swap.connect(signer).doParaSwapOnUniswap(
                    paraSwapAddress,
                    data.amountIn,
                    data.amountOutMin,
                    data.path,
                    data.referrer
                )
                console.log(result.receipt)
            } else if (contractMethod == "swapOnUniswapFork") {
                const result = await swap.connect(signer).doParaSwapOnUniswapFork(
                    paraSwapAddress,
                    data.factory,
                    data.initCode,
                    data.amountIn,
                    data.amountOutMin,
                    data.path,
                    data.referrer
                )
                console.log(result.receipt)
            } else if (contractMethod == "simpleSwap") {
                let abi = ["function approve(address _spender, uint256 _value) public returns (bool success)"]
                let contract = new ethers.Contract(data.fromToken, abi, provider)
                await contract.connect(signer).approve(highValueEthAccount, srcAmount)

                await swap.connect(signer).doParaSwapSimpleSwap(
                    //paraSwapAddress,
                    data.fromToken,
                    data.toToken,
                    data.fromAmount,
                    data.toAmount,
                    data.expectedAmount,
                    data.callees,
                    data.exchangeData,
                    data.startIndexes,
                    data.values,
                    data.beneficiary,
                    data.referrer,
                    data.useReduxToken
                )

            } else if (contractMethod == "multiSwap") {
                let abi = ["function approve(address _spender, uint256 _value) public returns (bool success)"]
                let contract = new ethers.Contract(data.data.fromToken, abi, provider)
                await contract.connect(signer).approve(highValueEthAccount, srcAmount)


            }
             

        })

    })
})