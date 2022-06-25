
const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect, assert } = require('chai');
const { BigNumber } = require('ethers');
const { ZERO_ADDRESS } = constants;
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18


//Web3 init
//const LPToken = artifacts.require('LPToken');
//const SwapContract = artifacts.require('SwapContract');
describe("Testing Node Functions...", () => {


    let LPTokenFactory, SwapContractFactory, sender, receiver

    let value, lpToken, btctTest, btctDecimals, mintValue, totalSwapped, withdrawalFeeBPS, swap, redeemedFloatTxIds

    beforeEach(async () => {
        // The bundled BN library is the same one web3 uses under the hood
        [sender, receiver, ...addrs] = await ethers.getSigners();
        LPTokenFactory = await ethers.getContractFactory("LPToken");
        SwapContractFactory = await ethers.getContractFactory("SwapContract");

        value = new BigNumber.from(1)

        lpToken = await LPTokenFactory.deploy(8);

        btctTest = await LPTokenFactory.deploy(TOKEN_DECIMALS)

        btctDecimals = await btctTest.decimals()

        mintValue = new BigNumber.from(500).mul(new BigNumber.from(10).pow(btctDecimals))
        //console.log(mintValue)

        totalSwapped = new BigNumber.from(0)

        withdrawalFeeBPS = new BigNumber.from(20)

        //for some reason it would not work if I put these inline when initializing swap, said lpToken was undefined...
        let adr = lpToken.address
        let btctadr = btctTest.address

        swap = await SwapContractFactory.deploy(adr, btctadr, 0);
        //console.log(swap)

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

            it('Adds 50 nodes nodes, tests for duplicates, and then removes 6', async () => {
                for (i = 0; i < 50; i++) {
                    let staked = new BigNumber.from(1500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    isRemoved.push(false)
                }
                //the churn function from SwapContract is taking ~15 seconds....
                const tx1 = await swap.churn(
                    sender.address,
                    rewardAddressAndAmounts,
                    isRemoved,
                    churnedInCount,
                    tssThreshold,
                    nodeRewardsRatio,
                    withdrawalFeeBPS,
                    {
                        value: 0,
                        gasPrice: 2 * 10 ** 6
                    })

                // console.log(tx2.receipt.cumulativeGasUsed)
                // Gas cost 306494 => 1284578 gas
                let getNode1 = await swap.getActiveNodes()
                //Below prints the 100 node addresses
                //console.log("Node1: ", await swap.getActiveNodes())
                //expect(getNode1.length).to.equal(50)
                assert.equal(getNode1.length, 50, "Nodes have been added")

                //ensure duplicate addresses are not listed as a node twice
                rewardAddressAndAmounts = []
                isRemoved = []

                for (i = 0; i < 50; i++) {
                    let staked = new BigNumber.from(1500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    isRemoved.push(false)
                }
                const tx2 = await swap.churn(sender.address, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, withdrawalFeeBPS, new BigNumber.from(0), {
                    value: 0,
                    gasPrice: 2 * 10 ** 6
                })

                // console.log(tx2.receipt.cumulativeGasUsed)
                // Gas cost 306494 => 1284578 gas
                let getNode2 = await swap.getActiveNodes()
                //expect(getNode2.length).to.equal(50)
                assert.equal(getNode2.length, 50, "Nodes have not been duplicated")

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

                const distribute = await swap.distributeNodeRewards()
                // Gas 3999911 gas
                // console.log(distribute.receipt.cumulativeGasUsed)

                rewardAddressAndAmounts = []
                isRemoved = []

                for (i = 0; i < 6; i++) {
                    let staked = new BigNumber.from(500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    isRemoved.push(true)
                }
                //console.log("rewardAddressAndAmounts", rewardAddressAndAmounts)
                //console.log("isRemoved", isRemoved)
                const tx3 = await swap.churn(sender.address, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, withdrawalFeeBPS, {
                    value: 0,
                    gasPrice: 2 * 10 ** 6
                })

                let getNode3 = await swap.getActiveNodes()
                //console.log("getNode3: ", getNode3)//this is showing 0 nodes for some reason
                //expect(getNode3.length).to.equal(50 - 6)
                assert.equal(getNode3.length, 50 - 6, "Nodes have been removed correctly")
                await expectRevert(
                    swap.distributeNodeRewards(),
                    'totalRewardLPsForNode is not positive',
                );
                //console.log(tx3.receipt.cumulativeGasUsed)
                // Gas cost 51700 => 867293 gas

            })

            it('adds 30 nodes into swap contract then update 50 nodes', async () => {
                rewardAddressAndAmounts = []
                isRemoved = []
                for (i = 0; i < 30; i++) {
                    let staked = new BigNumber.from(3000000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    isRemoved.push(false)
                }
                const tx1 = await swap.churn(sender.address, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, withdrawalFeeBPS, new BigNumber.from(0), {
                    value: 0,
                    //gas: 9000000,//cannot override gas
                    gasPrice: 2 * 10 ** 6
                })
                //console.log(tx1.receipt.cumulativeGasUsed)
                // Gas cost 4615206 gas
                let getNode1 = await swap.getActiveNodes()
                expect(getNode1.length).to.equal(30)

                for (i = 0; i < 30; i++) {
                    let staked = new BigNumber.from(1500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    isRemoved.push(false)
                }
                const tx2 = await swap.churn(sender.address, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, withdrawalFeeBPS, {
                    value: 0,
                    gasPrice: 2 * 10 ** 6
                })

                // console.log(tx2.receipt.cumulativeGasUsed)
                // Gas cost 784326 gas
                let getNode2 = await swap.getActiveNodes()
                expect(getNode2.length).to.equal(30)


                let floatAmountOfBTC = new BigNumber.from(1).mul(new BigNumber.from(10).pow(new BigNumber.from(8)))

                await btctTest.mint(swap.address, floatAmountOfBTC.mul(new BigNumber.from(10).pow(new BigNumber.from(10))))

                let swapTx = web3.utils.padLeft(floatAmountOfBTC._hex + sender.address.slice(2), 64)
                // fees are collected. (0.1 WBTC)
                let rewardsAmount = web3.utils.padLeft(new BigNumber.from("1").mul(new BigNumber.from(10).pow(new BigNumber.from(5)))._hex, 64)
                await swap.multiTransferERC20TightlyPacked(btctTest.address, [swapTx], totalSwapped, rewardsAmount, redeemedFloatTxIds)
                // Second deposit tx

                const distribute = await swap.distributeNodeRewards()
                // Gas 2413235 gas
                const remainTokens = await swap.lockedLPTokensForNode()
                //console.log(remainTokens.toString())
                // console.log(distribute.receipt.cumulativeGasUsed)
                //expect(remainTokens).to.bignumber.equal(new BigNumber.from(0))
                expect(new BigNumber.from(remainTokens)).equal(new BigNumber.from(0))

                rewardAddressAndAmounts = []
                isRemoved = []

                for (i = 0; i < 50; i++) {
                    let staked = new BigNumber.from(500000).mul(new BigNumber.from(10).pow(new BigNumber.from(18)))
                    let addressesAndAmountStaked = web3.utils.padLeft(staked._hex + accounts[i].address.slice(2), 64)
                    rewardAddressAndAmounts.push(addressesAndAmountStaked)
                    isRemoved.push(false)
                }
                const tx3 = await swap.churn(sender.address, rewardAddressAndAmounts, isRemoved, churnedInCount, tssThreshold, nodeRewardsRatio, withdrawalFeeBPS, {
                    value: 0,
                    gasPrice: 2 * 10 ** 6
                })

                let getNode3 = await swap.getActiveNodes()
                expect(getNode3.length).to.equal(50)

                await expectRevert(
                    swap.distributeNodeRewards(),
                    'totalRewardLPsForNode is not positive',
                );
                // console.log(tx3.receipt.cumulativeGasUsed)
                // Gas cost 784326 gas

            });
        });


    });
});
