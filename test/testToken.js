const { BN, constants, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { ZERO_ADDRESS } = constants;

const LPToken = artifacts.require('LPToken');

contract('LPToken', function (accounts) {
    const [sender, receiver] = accounts
    const name = 'BTC-LP token test';
    const symbol = 'BTC-LP test';

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(8)))

        this.token = await LPToken.new();

        await this.token.mint(sender, this.mintValue)
    });

    it('has a name', async function () {
        expect(await this.token.name()).to.equal(name);
    });

    it('has a symbol', async function () {
        expect(await this.token.symbol()).to.equal(symbol);
    });

    it('has 8 decimals', async function () {
        decimals = await this.token.decimals()
        expect(decimals.toString()).to.equal("8");
    });

    it('reverts when minting tokens from not owner', async function () {
        // Conditions that trigger a require statement can be precisely tested
        await expectRevert(
            this.token.mint(receiver, this.value, { from: receiver }),
            'Ownable: caller is not the owner',
        );
    });

    it('reverts when transferring tokens to the zero address', async function () {
        // Conditions that trigger a require statement can be precisely tested
        await expectRevert(
            this.token.transfer(constants.ZERO_ADDRESS, this.value, { from: sender }),
            'ERC20: transfer to the zero address',
        );
    });

    it('emits a Transfer event on successful transfers', async function () {
        const receipt = await this.token.transfer(
            receiver, this.value, { from: sender }
        );

        // Event assertions can verify that the arguments are the expected ones
        expectEvent(receipt, 'Transfer', {
            from: sender,
            to: receiver,
            value: this.value,
        });
    });

    it('updates balances on successful transfers', async function () {
        this.token.transfer(receiver, this.value, { from: sender });

        // BN assertions are automatically available via chai-bn (if using Chai)
        balanceOfReceiver = await this.token.balanceOf(receiver)
        expect(balanceOfReceiver.sub(this.value)).to.bignumber.equal('0')
    });

    it('emits a Burn event on successful burning', async function () {
        const burn = await this.token.burn(
            this.mintValue, { from: sender }
        );

        // Event assertions can verify that the arguments are the expected ones
        expectEvent(burn, 'Transfer', {
            from: sender,
            to: ZERO_ADDRESS,
            value: this.mintValue,
        });
    });
})