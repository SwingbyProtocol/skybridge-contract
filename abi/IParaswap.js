export default [
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "amountInMax",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "amountOut",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "path",
        "type": "address[]"
      }
    ],
    "name": "buyOnUniswap",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "factory",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "initCode",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "amountInMax",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "amountOut",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "path",
        "type": "address[]"
      }
    ],
    "name": "buyOnUniswapFork",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "tokenIn",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amountInMax",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "amountOut",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "weth",
        "type": "address"
      },
      {
        "internalType": "uint256[]",
        "name": "pools",
        "type": "uint256[]"
      }
    ],
    "name": "buyOnUniswapV2Fork",
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
            "internalType": "address payable",
            "name": "beneficiary",
            "type": "address"
          },
          {
            "components": [
              {
                "internalType": "uint256",
                "name": "fromAmountPercent",
                "type": "uint256"
              },
              {
                "components": [
                  {
                    "internalType": "address",
                    "name": "to",
                    "type": "address"
                  },
                  {
                    "internalType": "uint256",
                    "name": "totalNetworkFee",
                    "type": "uint256"
                  },
                  {
                    "components": [
                      {
                        "internalType": "address payable",
                        "name": "adapter",
                        "type": "address"
                      },
                      {
                        "internalType": "uint256",
                        "name": "percent",
                        "type": "uint256"
                      },
                      {
                        "internalType": "uint256",
                        "name": "networkFee",
                        "type": "uint256"
                      },
                      {
                        "components": [
                          {
                            "internalType": "uint256",
                            "name": "index",
                            "type": "uint256"
                          },
                          {
                            "internalType": "address",
                            "name": "targetExchange",
                            "type": "address"
                          },
                          {
                            "internalType": "uint256",
                            "name": "percent",
                            "type": "uint256"
                          },
                          {
                            "internalType": "bytes",
                            "name": "payload",
                            "type": "bytes"
                          },
                          {
                            "internalType": "uint256",
                            "name": "networkFee",
                            "type": "uint256"
                          }
                        ],
                        "internalType": "struct Utils.Route[]",
                        "name": "route",
                        "type": "tuple[]"
                      }
                    ],
                    "internalType": "struct Utils.Adapter[]",
                    "name": "adapters",
                    "type": "tuple[]"
                  }
                ],
                "internalType": "struct Utils.Path[]",
                "name": "path",
                "type": "tuple[]"
              }
            ],
            "internalType": "struct Utils.MegaSwapPath[]",
            "name": "path",
            "type": "tuple[]"
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
        "internalType": "struct Utils.MegaSwapSellData",
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "megaSwap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
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
            "internalType": "address payable",
            "name": "beneficiary",
            "type": "address"
          },
          {
            "components": [
              {
                "internalType": "address",
                "name": "to",
                "type": "address"
              },
              {
                "internalType": "uint256",
                "name": "totalNetworkFee",
                "type": "uint256"
              },
              {
                "components": [
                  {
                    "internalType": "address payable",
                    "name": "adapter",
                    "type": "address"
                  },
                  {
                    "internalType": "uint256",
                    "name": "percent",
                    "type": "uint256"
                  },
                  {
                    "internalType": "uint256",
                    "name": "networkFee",
                    "type": "uint256"
                  },
                  {
                    "components": [
                      {
                        "internalType": "uint256",
                        "name": "index",
                        "type": "uint256"
                      },
                      {
                        "internalType": "address",
                        "name": "targetExchange",
                        "type": "address"
                      },
                      {
                        "internalType": "uint256",
                        "name": "percent",
                        "type": "uint256"
                      },
                      {
                        "internalType": "bytes",
                        "name": "payload",
                        "type": "bytes"
                      },
                      {
                        "internalType": "uint256",
                        "name": "networkFee",
                        "type": "uint256"
                      }
                    ],
                    "internalType": "struct Utils.Route[]",
                    "name": "route",
                    "type": "tuple[]"
                  }
                ],
                "internalType": "struct Utils.Adapter[]",
                "name": "adapters",
                "type": "tuple[]"
              }
            ],
            "internalType": "struct Utils.Path[]",
            "name": "path",
            "type": "tuple[]"
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
        "internalType": "struct Utils.SellData",
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "multiSwap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
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
            "internalType": "address payable",
            "name": "beneficiary",
            "type": "address"
          },
          {
            "components": [
              {
                "internalType": "uint256",
                "name": "fromAmountPercent",
                "type": "uint256"
              },
              {
                "components": [
                  {
                    "internalType": "address",
                    "name": "to",
                    "type": "address"
                  },
                  {
                    "internalType": "uint256",
                    "name": "totalNetworkFee",
                    "type": "uint256"
                  },
                  {
                    "components": [
                      {
                        "internalType": "address payable",
                        "name": "adapter",
                        "type": "address"
                      },
                      {
                        "internalType": "uint256",
                        "name": "percent",
                        "type": "uint256"
                      },
                      {
                        "internalType": "uint256",
                        "name": "networkFee",
                        "type": "uint256"
                      },
                      {
                        "components": [
                          {
                            "internalType": "uint256",
                            "name": "index",
                            "type": "uint256"
                          },
                          {
                            "internalType": "address",
                            "name": "targetExchange",
                            "type": "address"
                          },
                          {
                            "internalType": "uint256",
                            "name": "percent",
                            "type": "uint256"
                          },
                          {
                            "internalType": "bytes",
                            "name": "payload",
                            "type": "bytes"
                          },
                          {
                            "internalType": "uint256",
                            "name": "networkFee",
                            "type": "uint256"
                          }
                        ],
                        "internalType": "struct Utils.Route[]",
                        "name": "route",
                        "type": "tuple[]"
                      }
                    ],
                    "internalType": "struct Utils.Adapter[]",
                    "name": "adapters",
                    "type": "tuple[]"
                  }
                ],
                "internalType": "struct Utils.Path[]",
                "name": "path",
                "type": "tuple[]"
              }
            ],
            "internalType": "struct Utils.MegaSwapPath[]",
            "name": "path",
            "type": "tuple[]"
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
        "internalType": "struct Utils.MegaSwapSellData",
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "protectedMegaSwap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
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
            "internalType": "address payable",
            "name": "beneficiary",
            "type": "address"
          },
          {
            "components": [
              {
                "internalType": "address",
                "name": "to",
                "type": "address"
              },
              {
                "internalType": "uint256",
                "name": "totalNetworkFee",
                "type": "uint256"
              },
              {
                "components": [
                  {
                    "internalType": "address payable",
                    "name": "adapter",
                    "type": "address"
                  },
                  {
                    "internalType": "uint256",
                    "name": "percent",
                    "type": "uint256"
                  },
                  {
                    "internalType": "uint256",
                    "name": "networkFee",
                    "type": "uint256"
                  },
                  {
                    "components": [
                      {
                        "internalType": "uint256",
                        "name": "index",
                        "type": "uint256"
                      },
                      {
                        "internalType": "address",
                        "name": "targetExchange",
                        "type": "address"
                      },
                      {
                        "internalType": "uint256",
                        "name": "percent",
                        "type": "uint256"
                      },
                      {
                        "internalType": "bytes",
                        "name": "payload",
                        "type": "bytes"
                      },
                      {
                        "internalType": "uint256",
                        "name": "networkFee",
                        "type": "uint256"
                      }
                    ],
                    "internalType": "struct Utils.Route[]",
                    "name": "route",
                    "type": "tuple[]"
                  }
                ],
                "internalType": "struct Utils.Adapter[]",
                "name": "adapters",
                "type": "tuple[]"
              }
            ],
            "internalType": "struct Utils.Path[]",
            "name": "path",
            "type": "tuple[]"
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
        "internalType": "struct Utils.SellData",
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "protectedMultiSwap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
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
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "protectedSimpleBuy",
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
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "protectedSimpleSwap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "receivedAmount",
        "type": "uint256"
      }
    ],
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
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "simpleBuy",
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
        "name": "data",
        "type": "tuple"
      }
    ],
    "name": "simpleSwap",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "receivedAmount",
        "type": "uint256"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "amountIn",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "path",
        "type": "address[]"
      }
    ],
    "name": "swapOnUniswap",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "factory",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "initCode",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "amountIn",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address[]",
        "name": "path",
        "type": "address[]"
      }
    ],
    "name": "swapOnUniswapFork",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "tokenIn",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amountIn",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "weth",
        "type": "address"
      },
      {
        "internalType": "uint256[]",
        "name": "pools",
        "type": "uint256[]"
      }
    ],
    "name": "swapOnUniswapV2Fork",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "contract IERC20",
        "name": "fromToken",
        "type": "address"
      },
      {
        "internalType": "contract IERC20",
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
        "name": "amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "exchange",
        "type": "address"
      },
      {
        "internalType": "bytes",
        "name": "payload",
        "type": "bytes"
      }
    ],
    "name": "swapOnZeroXv2",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "contract IERC20",
        "name": "fromToken",
        "type": "address"
      },
      {
        "internalType": "contract IERC20",
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
        "name": "amountOutMin",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "exchange",
        "type": "address"
      },
      {
        "internalType": "bytes",
        "name": "payload",
        "type": "bytes"
      }
    ],
    "name": "swapOnZeroXv4",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  }
];
