# Skybridge System Contracts on Ethereum

## Use Cases

### Deposit BTC

1. User deposits BTC using a special amount/tag that tells the system it’s a float. e.g. x.42424242. (Will be improved with BIP32)
2. System mints the BTC ‘LP’ token (1-1 representation).
3. System sends the minted BTC ‘LP’ token to User on Ethereum.

### Withdraw BTC

1. User initiates a normal swap via the API. (BTC ‘LP’ token on Ethereum -> BTC on Bitcoin)
2. User sends the BTC ‘LP’ token to the Bridge on Ethereum.
3. System sends the BTC to the User on Bitcoin.
4. Deposit WBTC/tBTC/renBTC:
5. User sends the WBTC/tBTC/renBTC to the Float Contract on Ethereum.
6. Float Contract mints the BTC ‘LP’ token (1-1 representation).
7. Float Contract sends the BTC ‘LP’ token to the User on Ethereum.
8. Float Contract funds the Swap Contract with the deposited WBTC/tBTC/renBTC.

### Withdraw WBTC

1. User sends the BTC ‘LP’ token to the Swap Contract on Ethereum: `redeemFloat(symbol)`.
2. Swap Contract sends the WBTC/tBTC/renBTC to the User on Ethereum.
3. Swap Contract burns the received BTC ‘LP’ token.

## Overview

### Swap Contract

- Allows transfer ownership to new TSS address.
- Accepts normal ERC20 transfers for swaps.
- Contains the ERC20 Multi-Send function. (with symbol arg)
- Redeems and burns BTC ‘LP’ token for WBTC/tBTC/renBTC.
- Allows contract owner (TSS) to forceRedeem float (burns BTC ‘LP’ token).
- Contains a pool for node rewards distributions.

### Float Contract

- Allows transfer ownership to new TSS address.
- Accepts normal ERC20 transfers for staking float.
- Mints BTC ‘LP’ token and redirects floated funds (WBTC/tBTC/renBTC) into Swap Contract.
- Allows Query of BTC ‘LP’ token holders list for fee distributions.
  Note: Can be integrated into the Multi-Send function of the Swap Contract for easy fee distributions?

## Interfaces

### Swap Contract

- `payable` (erc20)
- `addFloat(symbol)` - `payable`, called by Float Contract, _onlyOwner_
- `redeemFloat(symbol)` - payable, receives BTC ‘LP’ token
- `singleSend(address, token, amount, rewards_amount)`, _onlyOwner_
- `multiSend([]addresses, token, []amounts, rewards_amount)`, _onlyOwner_
- `distributeNodeRewards()`
- `churn(new_tss_addr, []node_addrs, []node_stake_amts, churned_in_count, node_rewards_ratio = 0.66)`, _onlyOwner_

### Float Contract

- `payable` (erc20) - mints BTC ‘LP’ token and transfers WBTC to Swap Contract
- `changeOwner(new_tss_address)`, _onlyOwner_
- `mintLPToken(amount, destination)`, _onlyOwner_
