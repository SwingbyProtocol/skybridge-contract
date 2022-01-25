export default [
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "lpToken",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "swapContract",
        "type": "address"
      }
    ],
    "name": "Deployed",
    "type": "event"
  },
  {
    "stateMutability": "nonpayable",
    "type": "fallback"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_owner",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_wbtc",
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
        "internalType": "uint8",
        "name": "_decimals",
        "type": "uint8"
      },
      {
        "internalType": "uint256",
        "name": "_existingBTCFloat",
        "type": "uint256"
      }
    ],
    "name": "deployNewContracts",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  }
];
