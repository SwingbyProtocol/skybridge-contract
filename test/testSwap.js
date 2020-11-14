const { accounts, contract } = require('@openzeppelin/test-environment');

const {
    BN,           // Big Number support
    constants,    // Common constants, like the zero address and largest integers
    expectEvent,  // Assertions for emitted events
    expectRevert, // Assertions for transactions that should fail
} = require('@openzeppelin/test-helpers');
const { ZERO_ADDRESS } = require('@openzeppelin/test-helpers/src/constants');
const send = require('@openzeppelin/test-helpers/src/send');

const SwapContract = contract.fromArtifact('SwapContract');

describe('SwapContract', function () {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.swap = await SwapContract.new(sender);
    });

    it('updates churn address and stakes', async function () {
        addrs = []
        nodeRewardsRatio = 66
        for (i = 0; i < 100; i++) {
            addrs.push("0x0d9630bd2a4edf6c4bf88d0c85538ebea6e6285c751b256242e7e9a4d40599b7")
        }
        const tx1 = await this.swap.churn(receiver, addrs, nodeRewardsRatio, {
            value: 0,
            gasPrice: 2 * 10 ** 6
        })
        console.log(tx1.receipt.gasUsed)
        addrs = []
        for (i = 0; i < 100; i++) {
            addrs.push("0x3bf37ed365c59d059dd784c6a22c086215966a4c18dc5b20f3210670763d430d")
        }
        const tx2 = await this.swap.churn(receiver, addrs, nodeRewardsRatio + 1, {
            value: 0,
            gasPrice: 2 * 10 ** 6,
            from: receiver
        })
        console.log(tx2.receipt.gasUsed)


    })
})