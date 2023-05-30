export default [
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_lpToken",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_btct",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_sbBTCPool",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_swapRewards",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_buybackAddress",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_initBTCFloat",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_initWBTCFloat",
        "type": "uint256"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "token",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amountOfLP",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amountOfFloat",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "currentPriceLP",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "withdrawal",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "txid",
        "type": "bytes32"
      }
    ],
    "name": "BurnLPTokensForFloat",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amountOfFloat",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amountOfLP",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "currentPriceLP",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "depositFees",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "txid",
        "type": "bytes32"
      }
    ],
    "name": "IssueLPTokensForFloat",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "previousOwner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "OwnershipTransferred",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "feesToken",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "rewards",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "rewardsLPTTotal",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "currentPriceLP",
        "type": "uint256"
      }
    ],
    "name": "RewardsCollection",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "Swap",
    "type": "event"
  },
  {
    "stateMutability": "nonpayable",
    "type": "fallback"
  },
  {
    "inputs": [],
    "name": "BTCT_ADDR",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "activeNodeCount",
    "outputs": [
      {
        "internalType": "uint8",
        "name": "",
        "type": "uint8"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "buybackAddress",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "buybackRewardsRatio",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_newOwner",
        "type": "address"
      },
      {
        "internalType": "address[]",
        "name": "_nodes",
        "type": "address[]"
      },
      {
        "internalType": "bool[]",
        "name": "_isRemoved",
        "type": "bool[]"
      },
      {
        "internalType": "uint8",
        "name": "_churnedInCount",
        "type": "uint8"
      },
      {
        "internalType": "uint8",
        "name": "_tssThreshold",
        "type": "uint8"
      }
    ],
    "name": "churn",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "churnedInCount",
    "outputs": [
      {
        "internalType": "uint8",
        "name": "",
        "type": "uint8"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_incomingAmount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_minerFee",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_rewardsAmount",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "_spenders",
        "type": "address[]"
      },
      {
        "internalType": "uint256[]",
        "name": "_swapAmounts",
        "type": "uint256[]"
      }
    ],
    "name": "collectSwapFeesForBTC",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getActiveNodes",
    "outputs": [
      {
        "internalType": "address[]",
        "name": "",
        "type": "address[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCurrentPriceLP",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "nowPrice",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_tokenA",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_tokenB",
        "type": "address"
      }
    ],
    "name": "getFloatReserve",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "reserveA",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "reserveB",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_user",
        "type": "address"
      }
    ],
    "name": "isNodeStake",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes32",
        "name": "_txid",
        "type": "bytes32"
      }
    ],
    "name": "isTxUsed",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "lpToken",
    "outputs": [
      {
        "internalType": "contract IBurnableToken",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_destToken",
        "type": "address"
      },
      {
        "internalType": "bytes32[]",
        "name": "_addressesAndAmounts",
        "type": "bytes32[]"
      },
      {
        "internalType": "uint256",
        "name": "_totalSwapped",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_rewardsAmount",
        "type": "uint256"
      },
      {
        "internalType": "bytes32[]",
        "name": "_redeemedFloatTxIds",
        "type": "bytes32[]"
      }
    ],
    "name": "multiTransferERC20TightlyPacked",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "nodeRewardsRatio",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_token",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "_addressesAndAmountOfFloat",
        "type": "bytes32"
      },
      {
        "internalType": "bytes32",
        "name": "_txid",
        "type": "bytes32"
      }
    ],
    "name": "recordIncomingFloat",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_token",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "_addressesAndAmountOfLPtoken",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "_minerFee",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "_txid",
        "type": "bytes32"
      }
    ],
    "name": "recordOutcomingFloat",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_minerFee",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "_txid",
        "type": "bytes32"
      }
    ],
    "name": "recordUTXOSweepMinerFee",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "renounceOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "sbBTCPool",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_destToken",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_to",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_totalSwapped",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_rewardsAmount",
        "type": "uint256"
      },
      {
        "internalType": "bytes32[]",
        "name": "_redeemedFloatTxIds",
        "type": "bytes32[]"
      }
    ],
    "name": "singleTransferERC20",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "sw",
    "outputs": [
      {
        "internalType": "contract ISwapRewards",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "transferOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "tssThreshold",
    "outputs": [
      {
        "internalType": "uint8",
        "name": "",
        "type": "uint8"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_sbBTCPool",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_buybackAddress",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_withdrawalFeeBPS",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_nodeRewardsRatio",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_buybackRewardsRatio",
        "type": "uint256"
      }
    ],
    "name": "updateParams",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "whitelist",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "withdrawalFeeBPS",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
];
