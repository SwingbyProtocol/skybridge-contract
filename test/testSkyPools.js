const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect, assert } = require('chai');
const { BigNumber } = require('ethers');
const { ZERO_ADDRESS } = constants;
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18


//Web3 init
//const LPToken = artifacts.require('LPToken');
//const SwapContract = artifacts.require('SwapContract');
describe("SkyPools", () => {


    let LPTokenFactory, SwapContractFactory, lpToken, swap, sender, receiver, accounts
    let lptDecimals, btctTest, initialPriceLP
    beforeEach(async () => {
        [sender, receiver, ...addrs] = await ethers.getSigners();
        accounts = [sender, receiver, ...addrs]
        LPTokenFactory = await ethers.getContractFactory("LPToken");
        lpToken = await LPTokenFactory.deploy(8);

        SwapContractFactory = await ethers.getContractFactory("SwapContract");

        lptDecimals = await lpToken.decimals()

        btctTest = await LPTokenFactory.deploy(TOKEN_DECIMALS)

        swap = await SwapContractFactory.deploy(lpToken.address, btctTest.address, 0);    
        initialPriceLP = new BigNumber.from(10).pow(lptDecimals)
    });


    // You can nest describe calls to create subsections.
    describe("SkyPool functions", () => {

        it('records SkyPools Transactions', async () => {

        });
        it('executes paraSwap transactions', async () => {

        });
        it('executes 1Inch trades', async () => {

        });
        it('allows the redemption of wBTC tokens', async () => {
            const wBTCaddr = 0x0

        });
    });
});
