export default [
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
    "stateMutability": "nonpayable",
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
      },
      {
        "internalType": "bool",
        "name": "_isUpdatelimitBTCForSPFlow2",
        "type": "bool"
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
    "stateMutability": "nonpayable",
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
    "stateMutability": "nonpayable",
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
    "stateMutability": "nonpayable",
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
  }
];
