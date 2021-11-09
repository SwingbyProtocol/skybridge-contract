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

            balance = await btctTest.balanceOf(swap.address)
            assert.equal(utils.formatEther(balance).toString(), utils.formatEther(mint500ERC20tokens).toString(), "Balance of tokens on swap contract is correct")

            balance = await btctTest.balanceOf(user1.address)
            assert.equal(balance.toNumber(), 0, "User's starting wallet balance is correct")

            //amount = 1 BTC, adjusted for 8 decimals
            let amount = new BigNumber.from(1).mul(new BigNumber.from(10).pow(lptDecimals))
            //console.log(utils.formatEther(amount.mul(convertScale)).toString())

            //perform recordSkyPoolsTX to assign user1 tokens in the contract
            await swap.recordSkyPoolsTX(user1.address, amount, 0) // TODO: need to add rewwards

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
            const getAllTokens = async () => {
                return await paraswap.getTokens(mainnetNetworkID)
            }
            //const AllTokens = getAllTokens()
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
                                blockNumber: 13220045,
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


            })//beforeEach

            it('executes paraSwap transactions: Flow 1 => BTC -> wBTC -> ERC20', async () => {
                /////////////////////////////// SET STARTING BALANCES //////////////////////////////////////////////
                //npx hardhat node --fork https://eth-mainnet.alchemyapi.io/v2/YfblHzLI_PlnIA0pxphS90J3yaA3dDi5 --fork-block-number 13220045
                //impersonated account at this block has .78678582 wBTC and 1621.283816188209230922 ETH
                await populateBalance(swap.address, wBTC, wBTC_SLOT, btcStartingAmount)

                /////////////////////////////// EXECUTE API CALLS //////////////////////////////////////////////
                //test API call
                //https://apiv4.paraswap.io/v2/prices/?from=ETH&to=DAI&amount=10000000000000000000&fromDecimals=18&toDecimals=18&side=SELL&network=137
                //get request for pricing data
                let getPrice = await paraswap.getPrice(
                    Tokens[mainnetNetworkID]['wBTC'],
                    Tokens[mainnetNetworkID]['UNI'],
                    endAmount,
                    mainnetNetworkID
                )
                //console.log(getPrice)

                const slippage = (decimals) => {
                    return new BigNumber.from(3).mul(new BigNumber.from(10).pow(decimals - 2)).toString() //Format ERC20 - 0.05
                }

                let decimals = Tokens[mainnetNetworkID]['UNI'].decimals
                let minDestAmount = new BigNumber.from(getPrice.price).sub(slippage(decimals))

                //POST request - build TX data to send to contract
                const txRequest = await paraswap.buildTransaction(
                    getPrice.payload,
                    Tokens[mainnetNetworkID]['wBTC'],
                    Tokens[mainnetNetworkID]['UNI'],
                    endAmount,
                    minDestAmount.toString(),
                    mainnetNetworkID,
                    user1.address,
                    true //only params - true for contract -> contract | false for standard transaction object
                )
                let data = txRequest.data //params to execute transaction contract -> contract       

                /////////////////////////////// EXECUTE TRANSACTIONS //////////////////////////////////////////////
                let balance

                //set float reserve
                const newFloatAmount = utils.parseEther("5") //5 tokens ERC20 - decimal 18
                let addressesAndAmountOfFloat = web3.utils.padLeft(newFloatAmount.add(minerFees)._hex + owner.address.slice(2), 64)
                await swap.recordIncomingFloat(wBTC, addressesAndAmountOfFloat, zeroFees, sampleTxs[0])

                //check float reserve
                balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                const startingFloatAmount = balance[1]
                assert.equal(utils.formatEther(balance[1]).toString(), utils.formatEther(newFloatAmount.add(minerFees)).toString(), "Float Reserve of BTCT tokens on the contract BEFORE skypools transaction is correct")

                //perform recordSkyPoolsTX to assign user1 btct tokens in the contract
                await swap.recordSkyPoolsTX(user1.address, endAmount, 0) // todo: need to add rewards

                balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                assert.equal(startingFloatAmount.sub(balance[1]).toString(), endAmount, "1 BTCt deducted from float") //1 BTCt -> decimal 8

                balance = await swap.tokens(wBTC, user1.address)
                assert.equal(balance.toString(), endAmount, "User balance in tokens[][] is correct")

                //Execute paraswap TX
                const dataArray = [
                    data[0].fromToken,
                    data[0].toToken,
                    data[0].fromAmount,
                    data[0].toAmount,
                    data[0].expectedAmount,
                    data[0].callees,
                    data[0].exchangeData,
                    data[0].startIndexes,
                    data[0].values,
                    data[0].beneficiary,
                    data[0].partner,
                    data[0].feePercent,
                    data[0].permit,
                    data[0].deadline,
                    data[0].uuid
                ]
                const result = await swap.connect(user1).spParaSwapBTC2Token(
                    dataArray
                )
                //Check ending balances
                balance = await UNI_Contract.balanceOf(user1.address)
                assert(balance > new BigNumber.from(0), "User has UNI tokens in wallet")

                balance = await swap.connect(user1).balanceOf(user1.address, UNI)
                assert.equal(balance.toString(), "0", "User balance of UNI tokens on swap contract is 0")

                /////////////////////////////// TEST FOR FAILURE //////////////////////////////////////////////
                it('rejects transactions when msg.sender does not match beneficiary nor holder of tokens in tokens[][]', async () => {
                    await swap.connect(user2).spParaSwapBTC2Token(
                        dataArray
                    ).should.be.rejectedWith("You can only execute swaps to your own address")
                })
                it('rejects transactions when the contract caller matches the token holder in tokens[][], but beneficiary does not', async () => {
                    let badDataArray = [
                        data[0].fromToken,
                        data[0].toToken,
                        data[0].fromAmount,
                        data[0].toAmount,
                        data[0].expectedAmount,
                        data[0].callees,
                        data[0].exchangeData,
                        data[0].startIndexes,
                        data[0].values,
                        user2.address, //malicious address
                        data[0].partner,
                        data[0].feePercent,
                        data[0].permit,
                        data[0].deadline,
                        data[0].uuid
                    ]
                    await swap.connect(user1).spParaSwapBTC2Token(
                        badDataArray
                    ).should.be.rejectedWith("You can only execute swaps to your own address")
                })
            })//END FLOW 1

            it('executes paraSwap transactions: Flow 2a => ETH -> wETH -> wBTC -> BTC', async () => {
                let balance, amount, result, overrides, receipt, event
                const ETHER = Tokens[mainnetNetworkID]['ETH'].address
                amount = utils.parseEther("1")
                overrides = {
                    value: amount
                }
                /////////////////////////////// TESTING DEPOSIT ETHER //////////////////////////////////////////////                
                result = await swap.connect(user1).spDeposit("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", amount, overrides)
                balance = await swap.balanceOf(wETH, user1.address)
                assert.equal(balance._hex, amount._hex, "Balance of wETH is correct after depositing ETHER")
                receipt = await result.wait()
                event = receipt.events[receipt.events.length - 1].args

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

                /////////////////////////////// EXECUTE API CALLS //////////////////////////////////////////////
                //test API call
                //https://apiv4.paraswap.io/v2/prices/?from=ETH&to=DAI&amount=10000000000000000000&fromDecimals=18&toDecimals=18&side=SELL&network=137
                //get request for pricing data

                let getPrice = await paraswap.getPrice(
                    Tokens[mainnetNetworkID]['WETH'],
                    Tokens[mainnetNetworkID]['wBTC'],
                    amount,
                    mainnetNetworkID
                )
                const slippage = (decimals) => {
                    return new BigNumber.from(3).mul(new BigNumber.from(10).pow(decimals - 2)).toString() //Format ERC20 - 0.05
                }
                let decimals = Tokens[mainnetNetworkID]['wBTC'].decimals

                let minDestAmount = new BigNumber.from(getPrice.price).sub(slippage(decimals))

                //POST request - build TX data to send to contract
                const txRequest = await paraswap.buildTransaction(
                    getPrice.payload,
                    Tokens[mainnetNetworkID]['WETH'],
                    Tokens[mainnetNetworkID]['wBTC'],
                    amount.toString(),
                    minDestAmount.toString(),
                    mainnetNetworkID,
                    swap.address,
                    true //only params - true for contract -> contract | false for standard transaction object
                )

                //parse output
                let data = txRequest.data //params to execute transaction contract -> contract
                const output = txRequest.config.data
                const { parse } = require('json-parser')
                const parsedOutput = parse(output)
                const contractMethod = parsedOutput.priceRoute.contractMethod
                //console.log("Recomended Contract Method:", contractMethod)//get contract method
                //console.log("Arguments for", contractMethod)

                /////////////////////////////// EXECUTE SWAP AND RECORD //////////////////////////////////////////////
                //const receivingBTC_Addr = "ms3xtHsLq2AZsQdtaPnLLnLbCtWSitUnZi"
                const receivingBTC_Addr = sampleTxs[0]

                const dataArray = [
                    data[0].fromToken,
                    data[0].toToken,
                    data[0].fromAmount,
                    data[0].toAmount,
                    data[0].expectedAmount,
                    data[0].callees,
                    data[0].exchangeData,
                    data[0].startIndexes,
                    data[0].values,
                    data[0].beneficiary,
                    data[0].partner,
                    data[0].feePercent,
                    data[0].permit,
                    data[0].deadline,
                    data[0].uuid
                ]
                balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                assert.equal(balance[1].toString(), "0", "Float Reserve of BTCT tokens on the contract BEFORE spParaSwapToken2BTC is 0")

                balance = await wETH_Contract.balanceOf(swap.address)
                assert.equal(balance.toString(), amount.toString(), "Balance of wETH on swap contract before swap is correct")
                
                
                
                
                const swapAndRecordResult = await swap.connect(user1).spParaSwapToken2BTC(
                    receivingBTC_Addr,
                    dataArray
                )

                //balance = await wBTC_Contract.balanceOf(swap.address)
                //console.log("Contract SATS AFTER SWAP: ", balance.toString())
                //balance = await swap.balanceOf(wBTC, swap.address)
                //console.log("BalanceOf SATS in tokens[][]: ", balance.toString())
                //console.log("EXPECTED AMOUNT (FROM API): ", data[0].expectedAmount.toString())

                // 49398 47333 95.81966881250253
                /////////////////////////////// CHECK ENDING BALANCES //////////////////////////////////////////////

                balance = await wETH_Contract.balanceOf(swap.address)
                assert.equal(balance.toString(), "0", "Balance of wETH on swap contract after swap is correct")

                balance = await swap.balanceOf(wBTC, swap.address)
                balance.toNumber().should.be.at.least(1, "Balance is present")
                const expectedSats = balance.toString() // 6964391 sats sas of pinned block


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


            it('executes paraSwap transactions: Flow 2b => ERC20 -> wBTC -> BTC', async () => {
                let balance, amount, result, overrides, receipt, event
                amount = utils.parseEther("1")
                /////////////////////////////// TESTING DEPOSIT UNI TOKENS //////////////////////////////////////////////
                await populateBalance(user1.address, UNI, UNI_SLOT, amount.mul(2))//amount refers to number of UNI tokens here
                balance = await UNI_Contract.balanceOf(user1.address)
                assert.equal(balance.toString(), amount.mul(2).toString(), "User has correct number of UNI tokens in their personal wallet")

                await UNI_Contract.connect(user1).approve(swap.address, amount)
                result = await swap.connect(user1).spDeposit(UNI, amount)

                balance = await swap.connect(user1).balanceOf(UNI, user1.address)
                assert.equal(balance._hex, amount._hex, "Balance is correct after depositing UNI")

                /////////////////////////////// EXECUTE API CALLS //////////////////////////////////////////////
                //test API call
                //https://apiv4.paraswap.io/v2/prices/?from=ETH&to=DAI&amount=10000000000000000000&fromDecimals=18&toDecimals=18&side=SELL&network=137
                //get request for pricing data
                let getPrice = await paraswap.getPrice(
                    Tokens[mainnetNetworkID]['UNI'],
                    Tokens[mainnetNetworkID]['wBTC'],
                    amount,
                    mainnetNetworkID
                )
                const slippage = (decimals) => {
                    return new BigNumber.from(3).mul(new BigNumber.from(10).pow(decimals - 2)).toString() //Format ERC20 - 0.05
                }

                let decimals = Tokens[mainnetNetworkID]['UNI'].decimals
                let minDestAmount = new BigNumber.from(getPrice.price).sub(slippage(decimals))

                //POST request - build TX data to send to contract
                const txRequest = await paraswap.buildTransaction(
                    getPrice.payload,
                    Tokens[mainnetNetworkID]['UNI'],
                    Tokens[mainnetNetworkID]['wBTC'],
                    amount.toString(),
                    minDestAmount.toString(),
                    mainnetNetworkID,
                    swap.address,
                    true //only params - true for contract -> contract | false for standard transaction object
                )

                //parse output
                let data = txRequest.data //params to execute transaction contract -> contract
                const output = txRequest.config.data
                const { parse } = require('json-parser')
                const parsedOutput = parse(output)
                const contractMethod = parsedOutput.priceRoute.contractMethod
                //console.log("Recomended Contract Method:", contractMethod)//get contract method
                //console.log("Arguments for", contractMethod)
                //console.log(data) 

                /////////////////////////////// EXECUTE SWAP AND RECORD //////////////////////////////////////////////
                //const receivingBTC_Addr = "ms3xtHsLq2AZsQdtaPnLLnLbCtWSitUnZi"
                const receivingBTC_Addr = sampleTxs[0]

                const dataArray = [
                    data[0].fromToken,
                    data[0].toToken,
                    data[0].fromAmount,
                    data[0].toAmount,
                    data[0].expectedAmount,
                    data[0].callees,
                    data[0].exchangeData,
                    data[0].startIndexes,
                    data[0].values,
                    data[0].beneficiary,
                    data[0].partner,
                    data[0].feePercent,
                    data[0].permit,
                    data[0].deadline,
                    data[0].uuid
                ]
                balance = await swap.getFloatReserve(ZERO_ADDRESS, wBTC)
                assert.equal(balance[1].toString(), "0", "Float Reserve of BTCT tokens on the contract BEFORE spParaSwapToken2BTC is 0")

                const swapAndRecordResult = await swap.connect(user1).spParaSwapToken2BTC(
                    receivingBTC_Addr,
                    dataArray
                )
                receipt = await swapAndRecordResult.wait()

                /////////////////////////////// CHECK BALANCES AND RECORDED DATA //////////////////////////////////////////////
                balance = await swap.balanceOf(wBTC, swap.address)
                balance.toNumber().should.be.at.least(1, "Balance is present")
                const expectedSats = balance.toString() // 47333 sats sas of pinned block

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


            })
        })//END PARASWAP
    })//END SKYPOOLS FUNCTIONS
})//END SKYPOOLS


//https://i.imgur.com/WlUaUXG.png