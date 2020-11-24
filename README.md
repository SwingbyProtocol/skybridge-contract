# skybridge-contract

## Environment
- testrpc (ganache) - v2.4.0
- truffle - v5.1.54
- solc - v0.7.5
- web3 - v1.2.9

## Gas usage
- 2130554 gas for storing initial storage by `churn()`
- 592154 gas for updating storage by `churn()`

## Test 
You have to start testrpc with `8545` port before doing test
```
$ npm run test
or 
$ truffle test --network development
```