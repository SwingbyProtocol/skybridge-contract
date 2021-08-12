const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect, assert } = require('chai');
const { BigNumber } = require('ethers');
const { ZERO_ADDRESS } = constants;
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18


//Web3 init
//const LPToken = artifacts.require('LPToken');
//const SwapContract = artifacts.require('SwapContract');
describe("SkyPools", () => {



    beforeEach(async () => {

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

        });
    });
});
