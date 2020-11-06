const { accounts, contract } = require('@openzeppelin/test-environment');

const {
    BN,           // Big Number support
    constants,    // Common constants, like the zero address and largest integers
    expectEvent,  // Assertions for emitted events
    expectRevert, // Assertions for transactions that should fail
} = require('@openzeppelin/test-helpers');

const LPToken = contract.fromArtifact('LPToken');

describe('LPToken', function () {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.erc20 = await LPToken.new();

        await this.erc20.mint(sender, this.mintValue)
    });

    it('reverts when minting tokens from not owner', async function () {
        // Conditions that trigger a require statement can be precisely tested
        await expectRevert(
            this.erc20.mint(receiver, this.value, { from: receiver }),
            'Ownable: caller is not the owner',
        );
    });

    it('reverts when transferring tokens to the zero address', async function () {
        // Conditions that trigger a require statement can be precisely tested
        await expectRevert(
            this.erc20.transfer(constants.ZERO_ADDRESS, this.value, { from: sender }),
            'ERC20: transfer to the zero address',
        );
    });

    it('emits a Transfer event on successful transfers', async function () {
        const receipt = await this.erc20.transfer(
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
        this.erc20.transfer(receiver, this.value, { from: sender });

        // BN assertions are automatically available via chai-bn (if using Chai)
        balanceOfSender = await this.erc20.balanceOf(receiver)
        expect(balanceOfSender - this.value, "balance is not passed !").to.equal(0)
    });
})