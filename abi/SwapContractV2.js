export default [
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
        "name": "token",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "user",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "balance",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "Timestamp",
        "type": "uint256"
      }
    ],
    "name": "Deposit",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "rewardLPTsForNodes",
        "type": "uint256"
      }
    ],
    "name": "DistributeNodeRewards",
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
        "name": "amountLPTokensForNode",
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
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "SwapID",
        "type": "bytes32"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "DestAddr",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "RefundAddr",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "AmountWBTC",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "Timestamp",
        "type": "uint256"
      }
    ],
    "name": "SwapTokensToBTC",
    "type": "event"
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
        "internalType": "address",
        "name": "user",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "balance",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "Timestamp",
        "type": "uint256"
      }
    ],
    "name": "Withdraw",
    "type": "event"
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
    "inputs": [
      {
        "internalType": "address",
        "name": "_token",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_user",
        "type": "address"
      }
    ],
    "name": "balanceOf",
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
    "name": "example",
    "outputs": [],
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
        "name": "_wETH",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_sbBTCPool",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_params",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_swapRewards",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_existingBTCFloat",
        "type": "uint256"
      }
    ],
    "name": "initialize",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "ip",
    "outputs": [
      {
        "internalType": "contract IParams",
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
    "name": "limitBTCForSPFlow2",
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
        "name": "_usedTxIds",
        "type": "bytes32[]"
      }
    ],
    "name": "multiRecordSkyPoolsTX",
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
    "name": "oldestActiveIndex",
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
    "inputs": [],
    "name": "paraswapAddress",
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
        "internalType": "address",
        "name": "_to",
        "type": "address"
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
        "name": "_usedTxIds",
        "type": "bytes32[]"
      }
    ],
    "name": "recordSkyPoolsTX",
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
        "name": "_token",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      }
    ],
    "name": "redeemERC20Token",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      }
    ],
    "name": "redeemEther",
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
    "name": "spCleanUpOldTXs",
    "outputs": [],
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
        "internalType": "uint256",
        "name": "_amount",
        "type": "uint256"
      }
    ],
    "name": "spDeposit",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "fromToken",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "toToken",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "fromAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "toAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "expectedAmount",
            "type": "uint256"
          },
          {
            "internalType": "address[]",
            "name": "callees",
            "type": "address[]"
          },
          {
            "internalType": "bytes",
            "name": "exchangeData",
            "type": "bytes"
          },
          {
            "internalType": "uint256[]",
            "name": "startIndexes",
            "type": "uint256[]"
          },
          {
            "internalType": "uint256[]",
            "name": "values",
            "type": "uint256[]"
          },
          {
            "internalType": "address payable",
            "name": "beneficiary",
            "type": "address"
          },
          {
            "internalType": "address payable",
            "name": "partner",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "feePercent",
            "type": "uint256"
          },
          {
            "internalType": "bytes",
            "name": "permit",
            "type": "bytes"
          },
          {
            "internalType": "uint256",
            "name": "deadline",
            "type": "uint256"
          },
          {
            "internalType": "bytes16",
            "name": "uuid",
            "type": "bytes16"
          }
        ],
        "internalType": "struct Utils.SimpleData",
        "name": "_data",
        "type": "tuple"
      }
    ],
    "name": "spFlow1SimpleSwap",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bool",
        "name": "_fork",
        "type": "bool"
      },
      {
        "internalType": "address",
        "name": "_factory",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "_initCode",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "_amountIn",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "_path",
        "type": "address[]"
      }
    ],
    "name": "spFlow1Uniswap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "receivedAmount",
        "type": "uint256"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "_destinationAddressForBTC",
        "type": "string"
      },
      {
        "components": [
          {
            "internalType": "address",
            "name": "fromToken",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "toToken",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "fromAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "toAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "expectedAmount",
            "type": "uint256"
          },
          {
            "internalType": "address[]",
            "name": "callees",
            "type": "address[]"
          },
          {
            "internalType": "bytes",
            "name": "exchangeData",
            "type": "bytes"
          },
          {
            "internalType": "uint256[]",
            "name": "startIndexes",
            "type": "uint256[]"
          },
          {
            "internalType": "uint256[]",
            "name": "values",
            "type": "uint256[]"
          },
          {
            "internalType": "address payable",
            "name": "beneficiary",
            "type": "address"
          },
          {
            "internalType": "address payable",
            "name": "partner",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "feePercent",
            "type": "uint256"
          },
          {
            "internalType": "bytes",
            "name": "permit",
            "type": "bytes"
          },
          {
            "internalType": "uint256",
            "name": "deadline",
            "type": "uint256"
          },
          {
            "internalType": "bytes16",
            "name": "uuid",
            "type": "bytes16"
          }
        ],
        "internalType": "struct Utils.SimpleData",
        "name": "_data",
        "type": "tuple"
      }
    ],
    "name": "spFlow2SimpleSwap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "receivedAmount",
        "type": "uint256"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "_destinationAddressForBTC",
        "type": "string"
      },
      {
        "internalType": "bool",
        "name": "_fork",
        "type": "bool"
      },
      {
        "internalType": "address",
        "name": "_factory",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "_initCode",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "_amountIn",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "_amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "_path",
        "type": "address[]"
      }
    ],
    "name": "spFlow2Uniswap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "receivedAmount",
        "type": "uint256"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "spGetPendingSwaps",
    "outputs": [
      {
        "components": [
          {
            "internalType": "bytes32",
            "name": "SwapID",
            "type": "bytes32"
          },
          {
            "internalType": "string",
            "name": "DestAddr",
            "type": "string"
          },
          {
            "internalType": "address",
            "name": "RefundAddr",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "AmountWBTC",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "Timestamp",
            "type": "uint256"
          }
        ],
        "internalType": "struct SwapContractV1.spPendingTx[]",
        "name": "data",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "spPendingTXs",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "SwapID",
        "type": "bytes32"
      },
      {
        "internalType": "string",
        "name": "DestAddr",
        "type": "string"
      },
      {
        "internalType": "address",
        "name": "RefundAddr",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "AmountWBTC",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "Timestamp",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
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
    "inputs": [],
    "name": "swapCount",
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
        "name": "",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "tokens",
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
    "inputs": [],
    "name": "wETH",
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
    "stateMutability": "payable",
    "type": "receive"
  }
];
