
const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect, assert } = require('chai');
const { BigNumber } = require('ethers');
const { ZERO_ADDRESS } = constants;
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18
//Web3 init
//const LPToken = artifacts.require('LPToken');
//const SwapContract = artifacts.require('SwapContract');
describe("Testing Node Functions...", () => {

    let LPTokenFactory, SwapContractFactory, SwapRewardsFactory, ParamsFactory, params, swapRewards, swap, sender, receiver

    let value, lpToken, btctTest, btctDecimals, mintValue, totalSwapped, withdrawalFeeBPS, redeemedFloatTxIds

    const tempSwapAddr = "0xfbC22278A96299D91d41C453234d97b4F5Eb9B2d"
    const sbBTCPool = "0xbA99c822bb4dd80f046a75EE564f8295D44Da743"
    const wETH = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" //https://etherscan.io/token/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2

    beforeEach(async () => {
        // The bundled BN library is the same one web3 uses under the hood
        [sender, receiver, ...addrs] = await ethers.getSigners();
        LPTokenFactory = await ethers.getContractFactory("LPToken");
        SwapContractFactory = await ethers.getContractFactory("SwapContract");
        SwapRewardsFactory = await ethers.getContractFactory("SwapRewards")
        ParamsFactory = await ethers.getContractFactory("Params")
        params = await ParamsFactory.deploy()
        //value = new BigNumber.from(1)

        lpToken = await LPTokenFactory.deploy(8);

        btctTest = await LPTokenFactory.deploy(TOKEN_DECIMALS)

        btctDecimals = await btctTest.decimals()

        //mintValue = new BigNumber.from(500).mul(new BigNumber.from(10).pow(btctDecimals))
        //console.log(mintValue)

        totalSwapped = new BigNumber.from(0)

        withdrawalFeeBPS = new BigNumber.from(20)

        swapRewards = await SwapRewardsFactory.deploy(
            sender.address,
            lpToken.address,
            tempSwapAddr
        )

        //swap = await SwapContractFactory.deploy(adr, btctadr, 0);
        swap = await SwapContractFactory.deploy(
            lpToken.address,
            btctTest.address,
            wETH,
            sbBTCPool,
            params.address,
            swapRewards.address,
            0
        )

        await swapRewards.connect(sender).setSwap(
            swap.address,
            new BigNumber.from(30),
            new BigNumber.from(60)
        )


        redeemedFloatTxIds = [
            "0x13e8785fe862e60f2caa4f838146ff9d4f4bd4a02dd6fb1f513b0a9be3452b62",
            "0xce66450451e62b9b4c406d0a83b90a5036039673d2682d4ec292f375ae571382"
        ]

        await lpToken.transferOwnership(swap.address)



    });//beforeEach

    // You can nest describe calls to create subsections.
    describe("Node Functions", () => {
        let accounts
        let rewardAddressAndAmounts = []
        let nodes = []
        let isRemoved = []
        let churnedInCount = 25
        let tssThreshold = 16
        let nodeRewardsRatio = 66
        beforeEach(async () => {
            //this contains all of the account objects, access address string -> accounts[n].address
            accounts = [sender, receiver, ...addrs]
        })
        it('checks owner of LPToken contract', async () => {
            expect(await lpToken.getOwner()).to.equal(swap.address)
            expect(await lpToken.owner()).to.equal(swap.address)
        });

        //Reduced the total number of nodes by 1/2 to make sure to avoid timeout issues
        describe(`Test for churn and float btctDecimals=${TOKEN_DECIMALS}`, async () => {
            //hardhat network has 20 accounts
            it('Adds 20 nodes nodes, tests for duplicates, and then removes 6', async () => {

                for (i = 0; i < 20; i++) {
                    let staked = new BigNumber.from(1500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    nodes.push(accounts[i].address)
                    isRemoved.push(false)
                }
                churnedInCount = 20
                tssThreshold = 16

                await swap.churn(
                    sender.address,
                    nodes,
                    isRemoved,
                    churnedInCount,
                    tssThreshold,
                    new BigNumber.from(50), //_totalStakedAmount
                )
                // console.log(tx2.receipt.cumulativeGasUsed)
                // Gas cost 306494 => 1284578 gas
                let getNode1 = await swap.getActiveNodes()
                //Below prints the 100 node addresses
                //console.log("Node1: ", await swap.getActiveNodes())
                //expect(getNode1.length).to.equal(50)
                assert.equal(getNode1.length, 20, "Nodes have been added")

                //ensure duplicate addresses are not listed as a node twice
                rewardAddressAndAmounts = []
                nodes = []
                isRemoved = []

                for (i = 0; i < 20; i++) {
                    let staked = new BigNumber.from(1500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    nodes.push(accounts[i].address)
                    isRemoved.push(false)
                }
                
                const tx2 = await swap.churn(
                    sender.address,
                    nodes,
                    isRemoved,
                    churnedInCount,
                    tssThreshold,
                    new BigNumber.from(50), //_totalStakedAmount
                )

                // console.log(tx2.receipt.cumulativeGasUsed)
                // Gas cost 306494 => 1284578 gas
                let getNode2 = await swap.getActiveNodes()
                //expect(getNode2.length).to.equal(50)
                assert.equal(getNode2.length, 20, "Nodes have not been duplicated")

                let floatAmountOfBTC = new BigNumber.from(10).pow(new BigNumber.from(8))

                await btctTest.mint(swap.address, floatAmountOfBTC.mul(new BigNumber.from(10).pow(new BigNumber.from(10))))

                //Cannot pass parameters to BigNumber.toString(), there is a built in _hex attribute, no need to prepend "0x" as it is included
                let swapTx = web3.utils.padLeft(floatAmountOfBTC._hex + sender.address.slice(2), 64)
                //console.log("swapTx: ", swapTx)

                // fees are collected. (0.1 WBTC)\
                let number = new BigNumber.from(10).pow(new BigNumber.from(7))
                //console.log(web3.utils.padLeft(number._hex, 64))
                let rewardsAmount = web3.utils.padLeft(number._hex, 64)
                //console.log("rewardsAmount: ", rewardsAmount)

                await swap.multiTransferERC20TightlyPacked(btctTest.address, [swapTx], totalSwapped, rewardsAmount, redeemedFloatTxIds)
                // Second deposit tx

                // Gas 3999911 gas
                // console.log(distribute.receipt.cumulativeGasUsed)

                rewardAddressAndAmounts = []
                nodes = []
                isRemoved = []

                for (i = 0; i < 6; i++) {
                    let staked = new BigNumber.from(500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    nodes.push(accounts[i].address)
                    isRemoved.push(true)
                }
                const tx3 = await swap.churn(
                    sender.address,
                    nodes,
                    isRemoved,
                    churnedInCount,
                    tssThreshold,
                    new BigNumber.from(50), //_totalStakedAmount
                )

                let getNode3 = await swap.getActiveNodes()

                assert.equal(getNode3.length, 20 - 6, "Nodes have been removed correctly")
                /**
                 * //depricated function
                 await expectRevert(
                    swap.distributeNodeRewards(),
                    'totalRewardLPsForNode is not positive',
                );
                 */
                //console.log(tx3.receipt.cumulativeGasUsed)
                // Gas cost 51700 => 867293 gas

            })

            it('adds 15 nodes into swap contract then update 20 nodes', async () => {
                rewardAddressAndAmounts = []
                nodes = []
                isRemoved = []
                for (i = 0; i < 15; i++) {
                    let staked = new BigNumber.from(3000000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    nodes.push(accounts[i].address)
                    isRemoved.push(false)
                }                

                const tx1 = await swap.churn(
                    sender.address,
                    nodes,
                    isRemoved,
                    churnedInCount,
                    tssThreshold,
                    new BigNumber.from(50), //_totalStakedAmount
                )
                //console.log(tx1.receipt.cumulativeGasUsed)
                // Gas cost 4615206 gas
                let getNode1 = await swap.getActiveNodes()
                expect(getNode1.length).to.equal(15)                

                for (i = 0; i < 20; i++) {
                    let staked = new BigNumber.from(1500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    nodes.push(accounts[i].address)
                    isRemoved.push(false)
                }

                const tx2 = await swap.churn(
                    sender.address,
                    nodes,
                    isRemoved,
                    churnedInCount,
                    tssThreshold,
                    new BigNumber.from(50), //_totalStakedAmount
                )

                // console.log(tx2.receipt.cumulativeGasUsed)
                // Gas cost 784326 gas
                let getNode2 = await swap.getActiveNodes()
                expect(getNode2.length).to.equal(20)

                let floatAmountOfBTC = new BigNumber.from(1).mul(new BigNumber.from(10).pow(new BigNumber.from(8)))

                await btctTest.mint(swap.address, floatAmountOfBTC.mul(new BigNumber.from(10).pow(new BigNumber.from(10))))

                let swapTx = web3.utils.padLeft(floatAmountOfBTC._hex + sender.address.slice(2), 64)
                // fees are collected. (0.1 WBTC)
                let rewardsAmount = web3.utils.padLeft(new BigNumber.from("1").mul(new BigNumber.from(10).pow(new BigNumber.from(5)))._hex, 64)
                await swap.multiTransferERC20TightlyPacked(btctTest.address, [swapTx], totalSwapped, rewardsAmount, redeemedFloatTxIds)
                // Second deposit tx

                // Gas 2413235 gas
                //depricated function
                //const remainTokens = await swap.lockedLPTokensForNode()
                

                rewardAddressAndAmounts = []
                nodes = []
                isRemoved = []

                for (i = 0; i < 20; i++) {
                    let staked = new BigNumber.from(500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    nodes.push(accounts[i].address)
                    isRemoved.push(false)
                }
                const tx3 = await swap.churn(
                    sender.address,
                    nodes,
                    isRemoved,
                    churnedInCount,
                    tssThreshold,
                    new BigNumber.from(50), //_totalStakedAmount
                )

                let getNode3 = await swap.getActiveNodes()
                expect(getNode3.length).to.equal(20)

                /**
                 * //depricated function
                 await expectRevert(
                    swap.distributeNodeRewards(),
                    'totalRewardLPsForNode is not positive',
                );
                 */
                // console.log(tx3.receipt.cumulativeGasUsed)
                // Gas cost 784326 gas

            });
        });


    });
});
