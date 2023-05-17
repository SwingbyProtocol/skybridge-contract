# Skybridge-contract-v3

## Environment
- hardhat (v2.14.0)
- solidity (v0.8.19)

## Deployment Order
1. Deploy sbBTCPool nodeRewards contracts
2. Deploy LP token and swap contract (initial liquidiy setup if need)
3. Initialize sbBTCPool and nodeRewards
4. Mint LPT to users (if need)
5. Set old LPT token for convert -- they can convert to new token with burning old token.
6. Transferring owner of new LPT to swap contract

## Build contract
```
$ npx hardhat compile
```

## Test 
```
$ npx hardhat test ./test/testSwapContract.js 
```

## Mainnet contracts
```
TBA
```

## Testnet contracts (goerli)
```
SwapContract
https://goerli.etherscan.io/address/0x9e6ba6e811665849f03f56c1f22a8894aebb3993
LPToken
https://goerli.etherscan.io/address/0xb10c6c5a6baf604206867cb229baddab02eea602
SwapRewards
https://goerli.etherscan.io/address/0xf4c381d077272295641f8a53d850d9a8125e0e94
BTCT
https://goerli.etherscan.io/address/0xeb47a21c1fc00d1e863019906df1771b80dbe182
sbBTCPool
https://goerli.etherscan.io/address/0xd60126017fdf906668cfe5327c566c65e7f061ba
Barn
https://goerli.etherscan.io/address/0x009cc14ce70b2e667984c2276490d56ae3234c43
```
