const { constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect, assert } = require('chai');
const { BigNumber } = require('ethers');
const { ZERO_ADDRESS } = constants;
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18


//Web3 init
//const LPToken = artifacts.require('LPToken');
//const SwapContract = artifacts.require('SwapContract');
describe("Contract: SwapFactory", () => {


    let LPTokenFactory, SwapContractFactory, SwapContract, factory, sender, receiver, accounts

    let value, btctTest, mintValue, btctDecimals


    beforeEach(async () => {
        // The bundled BN library is the same one web3 uses under the hood
        [sender, receiver, ...addrs] = await ethers.getSigners();
        //this contains all of the account objects, access address string -> accounts[n].address
        accounts = [sender, receiver, ...addrs]

        LPToken = await ethers.getContractFactory("LPToken");
        SwapContractFactory = await ethers.getContractFactory("SwapContractFactory");
        SwapContract = await ethers.getContractFactory("SwapContract");

        value = new BigNumber.from(1)

        btctTest = await LPToken.deploy(TOKEN_DECIMALS)

        btctDecimals = await btctTest.decimals()

        mintValue = new BigNumber.from(500).mul(new BigNumber.from(10).pow(new BigNumber.from(TOKEN_DECIMALS)))

        factory = await SwapContractFactory.deploy()

    })

    // You can nest describe calls to create subsections.
    describe("Deploy New Contracts", () => {
        it('checks the owner after deployment', async () => {
            const sc = await factory.connect(receiver).deployNewContracts(receiver.address, btctTest.address, TOKEN_DECIMALS, 0)
            const receipt = await sc.wait()

            //get existing contract address -- condensed to separate consts for readability 
            const lpTokenAddress = receipt.events[receipt.events.length - 1].args.lpToken
            const swapContractAddress = receipt.events[receipt.events.length - 1].args.swapContract

            const newLPToken = await LPToken.attach(lpTokenAddress) //equivilant to truffle .at() method
            const newSwapContract = await SwapContract.attach(swapContractAddress)
            

            expect(await newLPToken.getOwner()).to.equal(newSwapContract.address)
            expect(await newSwapContract.owner()).to.equal(receiver.address)


        });
    });
});
