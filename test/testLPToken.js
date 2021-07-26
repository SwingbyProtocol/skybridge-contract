const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect, assert } = require('chai');
const { BigNumber } = require('ethers');
const { ZERO_ADDRESS } = constants;
const TOKEN_DECIMALS = process.env.TOKEN_DECIMALS || 18


//Web3 init
//const LPToken = artifacts.require('LPToken');
//const SwapContract = artifacts.require('SwapContract');
describe("Contract: LPToken", () => {


    let LPTokenFactory, sender, receiver, accounts

    let value, token, mintValue

    const name = "Swingby BTC LP Token";
    const symbol = 'sbBTC';

    beforeEach(async () => {
        // The bundled BN library is the same one web3 uses under the hood
        [sender, receiver, ...addrs] = await ethers.getSigners();
        //this contains all of the account objects, access address string -> accounts[n].address
        accounts = [sender, receiver, ...addrs]

        LPTokenFactory = await ethers.getContractFactory("LPToken");

        value = new BigNumber.from(1)

        token = await LPTokenFactory.deploy(TOKEN_DECIMALS);

        mintValue = new BigNumber.from(500).mul(new BigNumber.from(10).pow(new BigNumber.from(TOKEN_DECIMALS)))

        await token.mint(sender.address, mintValue)
    })

    // You can nest describe calls to create subsections.
    describe("LPToken functions", () => {


        it('has a name', async () => {
            expect(await token.name()).to.equal(name)
        });
        it('has a symbol', async () => {
            expect(await token.symbol()).to.equal(symbol);
        });
        it('has the right decimals', async () => {
            decimals = await token.decimals()
            expect(new BigNumber.from(decimals)).equal(new BigNumber.from(TOKEN_DECIMALS));
        });
        it('reverts when minting tokens from not owner', async () => {
            // Conditions that trigger a require statement can be precisely tested\
            await expectRevert(
                token.connect(receiver).mint(receiver.address, value),
                'Ownable: caller is not the owner',
            );
        });
        it('reverts when transferring tokens to the zero address', async () => {
            // Conditions that trigger a require statement can be precisely tested
            await expectRevert(
                token.connect(sender).transfer(constants.ZERO_ADDRESS, value),
                'ERC20: transfer to the zero address',
            );
        });
        it('emits a Transfer event on successful transfers', async () => {
            let tx = await token.connect(sender).transfer(receiver.address, value);
            let receipt = await tx.wait();
            const data = receipt.events[0].args
            //console.log(data);//data emitted from Transfer Event
            
            assert.equal(receipt.events[0].event, 'Transfer', "Emits correct event")
            assert.equal(data.from, sender.address)
            assert.equal(data.to, receiver.address)
            assert.equal(data.value._hex, value._hex)


        });
        it('updates balances on successful transfers', async () => {
            await token.connect(sender).transfer(receiver.address, value)

            balanceOfReceiver = await token.balanceOf(receiver.address)
            expect(new BigNumber.from(balanceOfReceiver.sub(value))).equal('0')
        });
        it('emits a Burn event on successful burning', async () => {
            let tx = await token.connect(sender).burn(mintValue);
            let receipt = await tx.wait();
            const data = receipt.events[0].args
            //console.log(data);//data emitted from Transfer Event

            assert.equal(receipt.events[0].event, 'Transfer', "Emits correct event")
            assert.equal(data.from, sender.address)
            assert.equal(data.to, ZERO_ADDRESS)
            assert.equal(data.value._hex, mintValue._hex)
        });




    });
});