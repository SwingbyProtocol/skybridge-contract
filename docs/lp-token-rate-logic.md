# LP exchange rate logic


In the first iteration of Skybridge v2, LP tokens represent ownership by the underlying assets from the pools, which are composed of **two assets meant to have the same value** (_"the bridged assets"_).

The first iteration of Skybridge v2 will focus on a **BTC/WBTC pool**. This document explains the process on how LP exchange rate fluctuates, with a focus on the following cases:

1. Users swapping assets (WBTC -> BTC and BTC -> WBTC)
2. Liquidity providers depositing assets (aka minting LP tokens)
3. Liquidity providers withdrawing assets(aka burning LP tokens)

Similar to other protocols like Compound or Uniswap, the price of the LP token is bound to increase over time to reflect the transaction fees being incorporated into the pools.

## 1. Description of the parameters and other notations

```initialExchangeRate```: LP token to BTC/WBTC (defined in the contract constructor). Part of the constructor of the LP contract, an arbitrary value must defined.
```currentExchangeRate```: 1 LP token to BTC (changes over time)
```amountReceivedByLP```: amount received owing to swap transaction fees. It is an accounting claim and is incorporated into the pools.

## 2. Case studies

### 2.1 Swap BTC --> WBTC

When a swap is handled, part of the transaction fee is allocated to the SC and defined in the contract through the use of a field called ```amountReceivedByLP``` (in BTC).

The fee calculation is handled at the node configuration scope (not at the smart contract level).

Once the amount is paid to LPs (recorded as```amountReceivedByLP```), the ```currentExchangeRate``` must be re-calculated such as:

```currentExchangeRate = (newQuantityBTC + newQuantityWBTC)/(numberOfLPTokens)```

with (1) ```newQuantityBTC = initialQuantityBTC + quantityBTCSwapped + amountReceivedByLP```
and (2) ```newQuantityWBTC = initialQuantityWBTC - quantityWBTCSwapped ```

Since ```newQuantityBTC + newQuantityWBTC > initialQuantityBTC + initialQuantityWBTC```, the new ```currentExchangeRate``` **increases after every new swap**.

### 2.2 Swap WBTC --> BTC

It follows a similar process as described in subsection 2.1.

Once the fee is allocated and amount is received to the contract (and recorded as```amountReceivedByLP```), the ```currentExchangeRate``` must be re-calculated such as:

```currentExchangeRate = (newQuantityBTC + newQuantityWBTC)/(numberOfLPTokens)```

with (1) ```newQuantityBTC = initialQuantityBTC - quantityBTCSwapped```
and (2) ```newQuantityWBTC = initialQuantityWBTC + quantityWBTCSwapped + amountReceivedByLP```

Since ```newQuantityBTC + newQuantityWBTC > initialQuantityBTC + initialQuantityWBTC```, the new ```currentExchangeRate``` **increases after every new swap**.

### 2.3 Addition of liquidity to one side of the pool

When a user wishes to add additional liquidity to one side of the pool, he/she must mint the LP token at the ```currentExchangeRate``` while depositing WBTC or BTC to their respective TSS addresses held by the protocol validators.

#### 2.3.1 BTC deposit

1. The user deposits BTC using a special tag placed in the KV store that tells the system it is a float deposit. The swap is processed as normal by the nodes but the destination token becomes BTC-LP.
3. The parameter```currentExchangeRate``` is read by the contract.
4. The contract mints a quantity of BTC-LP token (1-1 representation) based on ```currentExchangeRate``` and ```amountDeposited``` (as supplied by the nodes) and sends the minted BTC-LP token to the user on Ethereum.

**Adding BTC liquidity does not change the LP exchange rate (vs. BTC) since no fee is collected by LP providers**.

However, ```numberOfLPTokens``` increases accordingly based on the quantity newly minted. 

#### 2.3.2 WBTC deposit

1. User sends the WBTC to the contract on Ethereum.
2. The contract mints a quantity of the BTC-LP token (1-1 representation) based on ```currentExchangeRate``` and ```amountDeposited``` (as supplied by the nodes) and sends the minted BTC-LP token to the user on Ethereum.
3. The contract credits the swap pool with the deposited WBTC.

**Adding WBTC liquidity does not change the LP exchange rate (vs. BTC) since no fee is collected by LP providers**.

However, ```numberOfLPTokens``` increases accordingly based on the quantity newly minted. 

### 2.4 Removal of liquidity from one side of the pool

When a user wishes to withdraw liquidity to one side of the pool, he must burn the LP token, indicate the asset he wishes to redeem his LP token to, while the ```currentExchangeRate``` is recorded part of the user's request.

#### 2.4.1 WBTC redemption

1. User sends the BTC-LP token to the swap contract on Ethereum: `redeemFloat()`.
2. The swap contract sends the WBTC to the user on Ethereum.
3. Swap contract burns the received BTC-LP token.

**Removing WBTC does not change the LP exchange rate (to BTC) since no fee is collected to LP providers**.

#### 2.4.2 BTC redemption

1. User sends the BTC-LP token to the swap contract on Ethereum: `redeemFloat()`.
2. Swap contract burns the received BTC-LP token and sends the WBTC for swapping to BTC.
3. System sends the BTC to the user on the Bitcoin blockchain.

**Removing BTC does change the LP exchange rate (to BTC)** since it is effectively a two-step process: 
(1) WBTC redemption (no change in the ```currentExchangeRate```)
(2) a swap WBTC -> BTC (change in the ```currentExchangeRate```).

Similarly, the```numberOfLPTokens``` also decreases accordingly based on the quantity burnt. 