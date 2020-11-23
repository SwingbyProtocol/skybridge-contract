const { accounts, contract } = require('@openzeppelin/test-environment');

const {
    BN,           // Big Number support
    constants,    // Common constants, like the zero address and largest integers
    expectEvent,  // Assertions for emitted events
    expectRevert, // Assertions for transactions that should fail
} = require('@openzeppelin/test-helpers');
const { ZERO_ADDRESS } = require('@openzeppelin/test-helpers/src/constants');
const send = require('@openzeppelin/test-helpers/src/send');

const LPToken = contract.fromArtifact('LPToken')
const SwapContract = contract.fromArtifact('SwapContract');

describe('SwapContract', function (done) {
    const [sender, receiver] = accounts;

    beforeEach(async function () {
        // The bundled BN library is the same one web3 uses under the hood
        this.value = new BN(1);

        this.mintValue = new BN(500).mul(new BN(10).pow(new BN(18)))

        this.lpToken = await LPToken.new()

        this.swap = await SwapContract.new(this.lpToken.address);

        await this.lpToken.transferOwnership(this.swap.address)
    });

    it('Checks owner of LPToken contract', async function () {
        expect(await this.lpToken.getOwner()).to.equal(this.swap.address)
        expect(await this.lpToken.owner()).to.equal(this.swap.address)
    })

    it('multi send test', async function () {
        const newToken = await LPToken.new()
        // swap contract receives 500000000 tokens 
        let amount1 = new BN(5000).mul(new BN(10).pow(new BN(8)))
        await newToken.mint(sender, amount1)
        await newToken.transfer(this.swap.address, amount1, {
            from: sender
        })
        let amount2 = new BN(1000).mul(new BN(10).pow(new BN(8)))
        // swap contracct send back 1000 tokens
        let amount3 = new BN(4000).mul(new BN(10).pow(new BN(8)))

        let senbackTokens1 = "0x" + web3.utils.padLeft(amount2.toString('hex') + sender.slice(2), 64)
        let senbackTokens2 = "0x" + web3.utils.padLeft(amount3.toString('hex') + receiver.slice(2), 64)

        let rewardsAmount = new BN(100).mul(new BN(10).pow(new BN(8)))
        // 0x00000000000000174876e800Fb4d4830eE2AfA5E5c6FD2C2cE3a080B6634ae0e
        const txs = [
            senbackTokens1, senbackTokens2
        ]
        await this.swap.multiTransferERC20TightlyPacked(newToken.address, txs, rewardsAmount)
        expect(await newToken.balanceOf(this.swap.address)).to.bignumber.equal(new BN("0"))
        expect(await newToken.balanceOf(sender)).to.equal(amount2)
        expect(await newToken.balanceOf(receiver)).to.equal(amount3)
    })

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