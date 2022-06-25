export default [
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_dest",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "_receiver",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "_swapped",
        "type": "uint256"
      }
    ],
    "name": "pullRewards",
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
        "name": "_dest",
        "type": "address"
      },
      {
        "internalType": "address[]",
        "name": "_receiver",
        "type": "address[]"
      },
      {
        "internalType": "uint256[]",
        "name": "_swapped",
        "type": "uint256[]"
      }
    ],
    "name": "pullRewardsMulti",
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
        "name": "_pricePerBTC",
        "type": "uint256"
      }
    ],
    "name": "setSWINGBYPrice",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
];
