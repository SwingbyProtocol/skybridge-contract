
const mnemonic = process.env.SEED
const key = "4d777ee25c2bb753c12597e8f35a2eedb90ece9bc5682f335e0e2c2fdc8d5674"

module.exports = {
  defaultNetwork: "localhost",
  solidity: {
    version: "0.7.3",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  networks: {
    hardhat: {
    },
    goerli: {
      url: "https://goerli.infura.io/v3/f35c2a4f3d0941a38a3edb62ed10c847",
      accounts:[key],
      network_id: 5,       // Ropsten's id
      gas: 5500000,        // Ropsten has a lower block limit than mainnet
      confirmations: 2,    // # of confs to wait between deployments. (default: 0)
      timeoutBlocks: 200,  // # of blocks before a deployment times out  (minimum/default: 50)
      skipDryRun: true,     // Skip dry run before migrations? (default: false for public nets )
      gasPrice: 53000000000
    }
  },
};