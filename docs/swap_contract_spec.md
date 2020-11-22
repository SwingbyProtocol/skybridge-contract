# Swap Contract Specification 
## Swap contract sources
This documnet describe the details of 
- _Swap_contract_ -- (https://github.com/SwingbyProtocol/skybridge-contract/blob/dev/contracts/SwapContract.sol)
- _ISwapContract_ -- `interface` contract https://github.com/SwingbyProtocol/skybridge-contract/blob/dev/contracts/interfaces/ISwapContract.sol
## Assumptions
- `SwapContract` inherits `Ownable` contract.
- `SwapContract` owner is a TSS address which is managed Swingby nodes. -- current TSS address can changes new owner of the contract.
- `SwapContract` doen't allow to receive `ETH`.The contract will rejects any transfer `ETH` into contract.(`ERC20` tokens are available to receive.)

## 1. The Swap 
### 1.1 The swap functions.
The `TSS adress` (_the owner of contract_) calls multi-send function when user swap BTC <> WBTC swap in the skybridge process. 
The contract has `3` functions for transfering `ERC20` tokens from `SwapContract`
- 1. `singleTransferERC20` -- onlyOwner `returns (bool)`
- 2. `multiTransferERC20` -- onlyOwner `retursns (bool)`
- 3. `multiTransferERC20TightlyPacked` -- onlyOwner `returns (bool)`
#### `function singleTransferERC20` -- external [onlyOwner]
- This method allows to transfer `ERC20` tokens from `SwapContract`via calling _`transfer()`_ method on the `ERC20` contract. 
```
function singleTransferERC20(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _rewardsAmount
    ) external returns (bool);
```
##### params
- `_token` -- this param represents target `ERC20` token address
- `_to` -- this param represents recipient address.
- `_amount` -- this param represents amount of transfer.
- `_rewardAmount` -- This argument represents _the amount of total reward WBTCs_
##### note
- This method is `external` function.
- This method is only called by `TSS address`

#### `function multiTransferERC20` -- external [onlyOwner]
- This method allows to multiple batch trasnfer `ERC20` tokens from `SwapContract`
```
function multiTransferERC20(
        address _token,
        address[] memory _contributors,
        uint256[] memory _amounts,
        uint256 _rewardsAmount
    ) public override onlyOwner returns (bool) {
```
##### params
- `_token` -- This param represents target `ERC20` token address
- `_contributors` -- This param represents an array of recipient addresses.
- `_amounts` -- This param represents an array of amounts for each recipient.
- `_rewardAmount` -- This param represents the amount of total reward basis `_token`
##### note
- This method is `external` function.
- This method is only called by `TSS address`

#### `function multiTransferERC20TightlyPacked` -- external [onlyOwner]
- This method allows to multiple batch transfer `ERC20` tokens. 
```
   function multiTransferERC20TightlyPacked(
        address _token,
        bytes32[] memory _addressesAndAmounts,
        uint8 _inputDecimals,
        uint256 _rewardsAmount
    ) public override onlyOwner returns (bool)
```
##### params
- `_token` -- This param represents target `ERC20` token address
- `_addressesAndAmounts` -- This param represents an array of convined `bytes32` values (top `12bytes` for amount, `20bytes` for target address)
- `_rewardAmount` -- This param represents the amount of total reward basis `_token`
##### note
- This method is `external` function.
- This method is only called by `TSS address`

### 1.2 Collecting Swap fees (WBTC, BTC)
### 1.2.1 The swap case of BTC to WBTC
Above`3` Swap functions have _`_rewardsCollection()`_ method. 
- That method allows to collect rewards which is collected per single swap action on the way of BTC to WBTC swap.
#### `function _rewardsCollection(address _token, uint256 _rewardsAmount)` -- internal
```
 function _rewardsCollection(address _token, uint256 _rewardsAmount)
        internal
    {
        // Reduce Gas
        uint256 totalRewardsForNode = totalRewardsForNodes[_token];
        // Reduce Gas
        uint256 totalRewardsForLP = totalRewardsForLPs[_token];
        // Updates rewards for nodes and LPs
        uint256 rewardsForNode = _rewardsAmount.mul(nodeRewardsRatio).div(100);
        uint256 rewardsForLP = _rewardsAmount.sub(rewardsForNode);
        totalRewardsForNodes[_token] = totalRewardsForNode.add(rewardsForNode);
        totalRewardsForLPs[_token] = totalRewardsForLP.add(rewardsForLP);
    }
```
#### params 
- `_token` -- This param represents float token address. e.g. BTC address == address(0) in the contract state.
- `_rewardsAmount` -- This param represents a amount of total collected fees on a swap.
#### note
- This method is `internal` function. (called by internally.)
- This method calls `nodeRewardsRatio == 66 as default` variable on the `SwapContract`. 
- The `uint256 _rewardsAmount` will be splitted to `2` type of fees. _(rewardsForLP and rewardsForNode)_
- These `2` fees are charged will be added to the total amount storage (totalRewardsForNode and totalRewardFor LP) to record the total amount.

### 1.2.2 The swap case of WBTC to BTC
The swap case of WBTC to BTC is not handle in smart contract logic. Therefore, the TSS address have to call a special method to add swap  fees to the total amount storage  
#### `function collectSwapFeesForBTC(address _feeToken, uint256 _rewardsAmount, bytes32 _txid)` -- external [onlyOwner]
##### params
- `_feeToken` -- This param represents target currencty e.g. address(0) == BTC, or WBTC address.
- `_rewardsAmount` -- This param represents total amount of fees on a swap.
- `_txid` -- This param represents the txid of swap (incoming txid).
##### note
- This method is `external` function.
- This method is only called by `TSS address`

---

**_This method is not implemented yet_**


---


## 2. Float and LP token management.

### 2.1 Float logic 101

Users can deposit `BTC` or `WBTC` into the `address` which is controlled by `TSS address`.
Deposited tokens are used for "reserve" of the skybridge mechanism.

- LP tokens are sent to `float providers` _( == Liquidity Providers)_ if they are providing `BTC` or `WBTC` into `SwapContract`, `TSS address`.
- The LP token represents value of provided float Tokens (`BTC` + `WBTC`) + collected total rewards.

### 2.2 Implementaion of float deposit BTC/WBTC
- The `TSS address` receives float BTC by the process of same as swap process. 
- The `TSS address` receives float WBTC by the process of same as swap process. 
- The `TSS address` has to call a method _`recordIncomingFloat()`_ on the `SwapContract` to register a _"float minting request"_ with deposit txid (BTC tx and ETH tx) for preventing double sending.
- The _"float minting request"_ is once made with unique `txid`, Therefore, if there are another call exist which has same `txid`, the tx will be rejected.
- The `TSS address` has to call a method - _`issueLPTokensForFloat()`_ on the `SwapContract` to confirm a _"float minting request"_ for each deposit `txid`.
- After confired a _"float minting request"_, `txid` will be added into "used" txid list on the `SwapContract`. And the `LP tokens` are minted to LP.
- The `TSS address` has to query to know that txid already used for minting LP tokens, A Getter method -- _`isTxidUsed(txid)`_ will returns `true`. if the result of this method is still `false`, it is the means that incoming tx will be refunded.

#### `function recordIncomingFloat(address _token, address _to, uint256 _amountOfFloat, bytes32 _txid)` -- external [onlyOwner]
- This method allows to register a _"float minting request"_. before confirm minting of LP tokens to LP, the `TSS address` has to check whether the deposit tx is correct. 
```
    function recordIncomingFloat(address _token, address _to, uint256 _amountOfFloat, bytes32 _txid)
        public
        override
        onlyOwner
        returns (bool)
```
##### params
- `_token` -- This param represents a float target. e.g. BTC == address(0), WBTC == WBTC token address.
- `_to` -- This param represents LP user.
- `_amountOfFloat` -- This param represents a total amount of BTC/WBTC depositted by a tx sent from LP.
- `_txid` -- This param represents deposit txid, that is also used as _unique_ id to identify the tx.
##### note
- This method is `external` function.
- This method is only called by `TSS address`

#### `function issueLPTokensForFloat(bytes32 _txid)` -- external 
- This method allows to "confirm" the previous submitted a _"float minting request"_ to issue acutal LP tokens to LP.
```
    function issueLPTokensForFloat(bytes32 _txid)
        public
        override
        onlyOwner
        returns (bool)
```

***This function let LP price going down, because the supply of LP tokens is increased if without WBTCs injection into calculate logic***

##### params 
- `_txid` -- This param represents deposit txid that was registred. 
##### note
- This method is `external` function.
- The LP tokens are issued based on the current price of LP tokens --_`getCurrentPriceLP()`_


### 2.3 Implementaion of float withdraw (BTC/WBTC)
- The `TSS address` receives LP tokens by the process of same as swap process to send back BTC to LP. 
- The `TSS address` receives LP tokens by the process of same as swap process to send back WBTC to LP. 
- The `TSS address` has to call a method _`recordOutcomingFloat()`_ on the `SwapContract` to register a _"float redeem request"_ with deposit txid (BTC tx and ETH tx) for preventing double sending.
- The _"float redeem request"_ is once made with unique `txid`, Therefore, if there are another call exist which has same `txid`, the tx will be rejected.
- The `TSS address` has to call a method - _`burnLPTokensForFloat()`_ on the `SwapContract` to confirm a _"float redeem request"_ for each deposit `txid`.
- After confired a _"float redeem request"_, `txid` will be added into "used" txid list on the `SwapContract`. And the `LP tokens` are burned.
- The `TSS address` has to query to know that txid already used for redeeming LP tokens, A Getter method -- _`isTxidUsed(txid)`_ will returns `true`. if the result of this method is still `false`, it is the means that incoming tx will be refunded.

#### `function recordOutcomingFloat(address _token, address _to, uint256 _amountOfLPtoken, bytes32 _txid)` -- external [onlyOwner]
- This method allows to register a _"float redeem request"_. before confirm burning LP tokens, the `TSS address` has to check whether the deposit tx is correct. 
```
    function recordOutcomingFloat(address _token, address _to, uint256 _amountOfLPtoken, bytes32 _txid)
        public
        override
        onlyOwner
        returns (bool)
```
##### params
- `_token` – This param represents a float target. e.g. BTC == address(0), WBTC == WBTC token address.
- `_to` – This param represents LP user.
- `_amountOfLPtoken` – This param represents a total amount of BTC/WBTC depositted by a tx sent from LP.
- `_txid` – This param represents deposit txid, that is also used as unique id to identify the tx.

##### note
- This method is `external` function.
- This method is only called by `TSS address`

#### `function burnLPTokensForFloat(bytes32 _txid)` -- external 
- This method allows to "confirm" the previous submitted a _"float redeem request"_ to burn acutal LP tokens. And send back BTC or WBTC (contract side only handles WBTC transfer. If `_token` address is `address(0) == BTC`, contract will only burns LP tokens.).
```
    function burnLPTokensForFloat(bytes32 _txid)
        public
        override
        onlyOwner
        returns (bool)
```

***This function let LP price going up, because the supply of LP tokens is decreased if without WBTCs injection into calculate logic***

##### params 
- `_txid` -- This param represents deposit txid that was registred. 
##### note
- This method is `external` function.
- The LP tokens are burned based on the current price of LP tokens --_`getCurrentPriceLP()`_

### 2.4 Price calculation of LP Token 
- Th price of LP Token is basically rely on float balances of `2` Pairs token amount (`BTC` and `WBTC` in our case)
- The float balances are changed when user deposits `Float` for `BTC` and `WBTC`
- The float balances are changed when user withdraw `Float` for `BTC` and `WBTC`
- The calculation logic collects total BTC fees for LP
- The calculation logic collects total WBTC fees for LP 
#### `function _updatePool(address _tokenA, address _tokenB)` -- [internal]
```
    function _updatePool(address _tokenA, address _tokenB)
        internal
        returns (uint256)
    {
        // Reduce gas cost.
        uint256 floatAmountOfTokenA = floatAmountOfToken[_tokenA];
        uint256 floatAmountOfTokenB = floatAmountOfToken[_tokenB];
        uint256 totalRewardOfTokenA = totalRewardsForLPs[_tokenA];
        uint256 totalRewardOfTokenB = totalRewardsForLPs[_tokenB];

        uint256 totalReserved = floatAmountOfTokenA
            .add(floatAmountOfTokenB)
            .add(totalRewardOfTokenA)
            .add(totalRewardOfTokenB);

        // Logic: LPP = (float amount of BTC + float amount of WBTC + LP fees) / (LP supply)
        //uint256 burned = IBurnableToken(lpToken).balanceOf(address(burner));
        uint256 totalLPs = IBurnableToken(lpToken).totalSupply();
        // decimals of totalReserved == 8, lpDecimals == 8, decimals of rate == 8
        currentExchangeRate = totalReserved.mul(lpDecimals).div(totalLPs);
        return currentExchangeRate;
    }
```
##### params
- `_tokenA` -- This param represents a float token address. e.g. BTC address == address(0) in the contract state.
- `_tokenB` -- This param represents a target address which is pooled token address. e.g. `WBTC` address.
##### note 
- This method is `internal` function.


## 3. Node management

### 3.1 Update new TSS address and Nodes on the contract
#### `churn(address _newOwner, bytes32[] memory _nodeRewardsAddressAndAmounts, uint8 _churnedInCount, uint8 _nodeRewardsRatio)` -- external [onlyOwner]

```
    function churn(
        address _newOwner,
        bytes32[] memory _nodeRewardsAddressAndAmounts,
        uint8 _churnedInCount,
        uint8 _nodeRewardsRatio
    ) public override onlyOwner returns (bool) {
        _transferOwnership(_newOwner);
        nodes = _nodeRewardsAddressAndAmounts;
        churnedInCount = _churnedInCount;
        // The ratio should be 100x of actual rate.
        nodeRewardsRatio = _nodeRewardsRatio;
        return true;
    }
```
##### params
- `address _newOwner` -- The TSS address will be replaced to new TSS adress. `_newOwner`
- `bytes32[] memory _nodeRewardsAddressAndAmounts` is an array of convined `bytes32` values (top `12bytes` for amount, `20bytes` for node address)
##### note
- This method is `external` function.
- This method is only called by `TSS address`

### 3.2 Distribute Node rewards
- Reward currency is `LP Token`. 
- The reward amount is based on collected balance of WBTC fees and current price of `LP token`.

#### `distributeNodeRewards()` -- external [onlyOwner]
```
    function distributeNodeRewards() public override returns (bool) {
        // Reduce Gas
        uint256 totalRewardsForNode = totalRewardsForNodes[WBTC_ADDR];
        require(
            totalRewardsForNode > 0,
            "totalRewardsForNode amount is not positive"
        );
        uint256 totalStaked = 0;
        for (uint256 i = 0; i < nodes.length; i++) {
            totalStaked = totalStaked.add(uint256(uint96(bytes12(nodes[i]))));
        }
        // decimals of totalRewardsForNode  == 8, decimals of currentPrice == 8
        uint256 totalLPTokens = totalRewardsForNode.mul(priceDecimals).div(
            getCurrentPriceLP()
        );
        for (uint256 i = 0; i < nodes.length; i++) {
            IBurnableToken(lpToken).mint(
                address(uint160(uint256(nodes[i]))),
                totalLPTokens.mul(uint256(uint96(bytes12(nodes[i])))).div(
                    totalStaked
                )
            );
        }
        // Inject WBTC amount to float?
        // floatAmountOfToken[WBTC_ADDR] = floatAmountOfToken[WBTC_ADDR].add(
        //     totalRewardsForNode
        // );
        // Reset storage for WBTC fees.
        totalRewardsForNodes[WBTC_ADDR] = 0;
        return true;
    }
```
##### params

##### note
- This method is `external` function
- This method is only called by `TSS address`

## 4. Getters
#### `function isTxidUsed(bytes32 _txid) returns (bool)` -- external 
- This method returns whether the txid is used or not.

#### `function getCurrentPriceLP(address _token)  returns (uint256)` -- external
- This method returns current price of LP token.

## Refference
- [LP exchange rate logic](https://hackmd.io/@9AlaYtTSSgq4TJ4TPi-gfQ/ryns4XFOw)