# skybridge-contract-v2

## Environment
- testrpc (ganache-cli) - v2.4.0
- truffle - v5.1.54
- solc - v0.7.5
- nodejs - v14.8.0
- web3 - v1.2.9

## Deploy contract
There are 2 steps to setup contracts on the ethereum testnet.
1. Deploy SwapContract and LP token contract
```
$ export SEED=<your mnemonic key>
$ truffle migrate --network {development/goerli}
```
2. Change the owner of SwapContract to TSS address
```
$ TSS={The TSS wallet address} truffle exec scripts/moveSCOwner.js --network {development/goerli} 
```

## Build contract
```
$ npn run build
```

## Test 
You have to start testrpc with 100 accounts and binding port `8545` before doing test
```
$ ganache-cli -a 100 -l 9000000
```
```
$ npm run test
or 
$ truffle test --network development
```

## Mainnet contracts
```
SkypoolContract
https://etherscan.io/address/0x4A084C0D1f89793Bb57f49b97c4e3a24cA539aAA
LPToken
https://etherscan.io/address/0x44a62c7121a64691b61aef669f21c628258e7d52
PraswapRouter
https://etherscan.io/address/0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57
BTCT
https://etherscan.io/token/0x2260fac5e5542a773aa44fbcfedf7c193bc2c599
```

## Testnet contracts (ropsten)
```
SkypoolContract
https://ropsten.etherscan.io/address/0x92c95b6227a9f0b4602649bd83f83adc48dae903
LPToken
https://ropsten.etherscan.io/address/0x679199877e56b8e68fdb1ddae122e843ecaca268
PraswapRouter
https://ropsten.etherscan.io/address/0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57
BTCT
https://ropsten.etherscan.io/address/0x7cb2eac36b4bb7c36640f32e806d33e474d1d427
```