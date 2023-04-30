
//npx hardhat run scripts/deploy.js --network <network-name>
//run the tests: npx hardhat test



require("@nomiclabs/hardhat-waffle");
//Plugins to let hardhat use web3 and truffle libraries 
require("@nomiclabs/hardhat-web3");
require("@nomiclabs/hardhat-truffle5");
require("hardhat-watcher");
require('@symblox/hardhat-abi-gen');
require('@nomiclabs/hardhat-etherscan');
//const mnemonic = process.env.SEED

//This requires a private key rather than a memonic, this is a private key to a throw away account so this can compile, but it should be stored in a local secret.json file
//It could then be accessed below using secret.key 
const mnemonic = process.env.KEY || "4d777ee25c2bb753c12597e8f35a2eedb90ece9bc5682f335e0e2c2fdc8d5674"

module.exports = {
  defaultNetwork: "hardhat",
  watcher: {
    compilation: { //npx hardhat watch compilation -- auto compile on change
      tasks: ["compile"],
    },
    test: {//npx hardhat watch test -- run test when a file is saved
      tasks: [{ command: 'test', params: { testFiles: ['./test/testSkyPoolsV2.js'] } }], //test this file
      files: ['./test/testSkyPoolsV2.js'] //test when this file is saved
    }
  },
  solidity: {
    version: "0.8.19",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000
      }
    }
  },
  networks: {
    development: {
      url: "http://127.0.0.1:8545",
      gas: 5500000,
      gasPrice: 8000000000
    },
    hardhat: {
      //run tests on fork of mainnet 
      forking: {
        url: "https://eth-mainnet.alchemyapi.io/v2/YfblHzLI_PlnIA0pxphS90J3yaA3dDi5",
        blockNumber: 13578738 //13220045 //previous testing block
      }
    },
    goerli: {
      url: "https://goerli.infura.io/v3/f35c2a4f3d0941a38a3edb62ed10c847",
      accounts: [mnemonic],
      network_id: 5,       // Ropsten's id
      gas: "auto",        // Ropsten has a lower block limit than mainnet
      confirmations: 2,    // # of confs to wait between deployments. (default: 0)
      timeoutBlocks: 200,  // # of blocks before a deployment times out  (minimum/default: 50)
      skipDryRun: true,     // Skip dry run before migrations? (default: false for public nets )
      gasPrice: "auto",
      maxPriorityFeePerGas: 1000000000,
    },
    mainnet: {
      url: "https://mainnet.infura.io/v3/f35c2a4f3d0941a38a3edb62ed10c847",
      accounts: [mnemonic],
      network_id: 1,       // Ropsten's id
      gas: 3000000,        // Ropsten has a lower block limit than mainnet
      confirmations: 2,    // # of confs to wait between deployments. (default: 0)
      timeoutBlocks: 200,  // # of blocks before a deployment times out  (minimum/default: 50)
      skipDryRun: true,     // Skip dry run before migrations? (default: false for public nets )
      gasPrice: 93000000000
    },
    bsc_testnet: {
      url: "https://mainnet.infura.io/v3/f35c2a4f3d0941a38a3edb62ed10c847",
      accounts: [mnemonic],
      network_id: 97,       // Ropsten's id
      gas: 7500000,        // Ropsten has a lower block limit than mainnet
      confirmations: 2,    // # of confs to wait between deployments. (default: 0)
      timeoutBlocks: 200,  // # of blocks before a deployment times out  (minimum/default: 50)
      skipDryRun: true,     // Skip dry run before migrations? (default: false for public nets )
      gasPrice: 10000000000
    },
    bsc_mainnet: {
      url: "https://mainnet.infura.io/v3/f35c2a4f3d0941a38a3edb62ed10c847",
      accounts: [mnemonic],
      network_id: 56,       // Ropsten's id
      gas: 3000000,        // Ropsten has a lower block limit than mainnet
      confirmations: 2,    // # of confs to wait between deployments. (default: 0)
      timeoutBlocks: 200,  // # of blocks before a deployment times out  (minimum/default: 50)
      skipDryRun: true,     // Skip dry run before migrations? (default: false for public nets )
      gasPrice: 10000000000
    }
    // Useful for private networks
    // private: {
    // provider: () => new HDWalletProvider(mnemonic, `https://network.io`),
    // network_id: 2111,   // This network is yours, in the cloud.
    // production: true    // Treats this network as if it was a public net. (default: false)
    // }
  },
  // Set default mocha options here, use special reporters etc.
  mocha: {
    timeout: 100000
  },

  // Configure your compilers
  compilers: {
    solc: {
      version: "0.8.9",    // Fetch exact version from solc-bin (default: truffle's version)
      // docker: true,        // Use "0.5.1" you've installed locally with docker (default: false)
      settings: {          // See the solidity docs for advice about optimization and evmVersion
        optimizer: {
          enabled: true,
          runs: 200
        },
        //  evmVersion: "byzantium"
      }
    },
  },

  etherscan: {
    apiKey: "A7EAG5WB4FRAHIURRGWD8HSTM8CVYXZGZ4"
  },

  abiExporter: {
    path: './abi',
    clear: false,
    flat: true,
    spacing: 2
  },

  plugins: [
    "@chainsafe/truffle-plugin-abigen"
  ]
};