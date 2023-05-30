# Security Audit Report

* Date: 31-05-2023
* Test tools
  * Slither v0.9.0
  * Solc 0.8.19+commit.7dd6d404.Darwin.appleclang

## Summary
 - [unchecked-transfer](#unchecked-transfer) (4 results) (High)
 - [incorrect-equality](#incorrect-equality) (3 results) (Medium)
 - [reentrancy-no-eth](#reentrancy-no-eth) (3 results) (Medium)
 - [calls-loop](#calls-loop) (3 results) (Low)
 - [reentrancy-benign](#reentrancy-benign) (8 results) (Low)
 - [reentrancy-events](#reentrancy-events) (10 results) (Low)
 - [timestamp](#timestamp) (3 results) (Low)
 - [dead-code](#dead-code) (78 results) (Informational)
 - [solc-version](#solc-version) (56 results) (Informational)
 - [naming-convention](#naming-convention) (67 results) (Informational)
 - [tautology](#tautology) (5 results) (Medium)
 - [missing-zero-check](#missing-zero-check) (7 results) (Low)
 - [divide-before-multiply](#divide-before-multiply) (3 results) (Medium)
 - [unused-return](#unused-return) (6 results) (Medium)
 - [events-maths](#events-maths) (1 results) (Low)
 - [assembly](#assembly) (2 results) (Informational)
 - [costly-loop](#costly-loop) (2 results) (Informational)
 - [low-level-calls](#low-level-calls) (4 results) (Informational)
 - [shadowing-local](#shadowing-local) (7 results) (Low)

## Contracts
```
SafeMath:
+------+----+
| Name | ID |
+------+----+
+------+----+

IBarn:
+------------------------------+------------+
|             Name             |     ID     |
+------------------------------+------------+
|      balanceOf(address)      | 0x70a08231 |
| balanceAtTs(address,uint256) | 0x417edd4d |
+------------------------------+------------+

IBurnableToken:
+---------------------------------------+------------+
|                  Name                 |     ID     |
+---------------------------------------+------------+
|             totalSupply()             | 0x18160ddd |
|               decimals()              | 0x313ce567 |
|                symbol()               | 0x95d89b41 |
|                 name()                | 0x06fdde03 |
|               getOwner()              | 0x893d20e8 |
|           balanceOf(address)          | 0x70a08231 |
|       transfer(address,uint256)       | 0xa9059cbb |
|       allowance(address,address)      | 0xdd62ed3e |
|        approve(address,uint256)       | 0x095ea7b3 |
| transferFrom(address,address,uint256) | 0x23b872dd |
|         mint(address,uint256)         | 0x40c10f19 |
|             burn(uint256)             | 0x42966c68 |
|               mintable()              | 0x4bf365df |
+---------------------------------------+------------+

ISwapContract:
+------------------------------------------------------------------------------+------------+
|                                     Name                                     |     ID     |
+------------------------------------------------------------------------------+------------+
|                                 BTCT_ADDR()                                  | 0x0f909486 |
|                                  lpToken()                                   | 0x5fcbd285 |
|    singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])    | 0x0d63aca7 |
| multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[]) | 0xad289e76 |
|      collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])      | 0x2adf9f87 |
|                 recordIncomingFloat(address,bytes32,bytes32)                 | 0xcf10b16b |
|            recordOutcomingFloat(address,bytes32,uint256,bytes32)             | 0x2586c562 |
|                   recordUTXOSweepMinerFee(uint256,bytes32)                   | 0xc810a539 |
|                 churn(address,address[],bool[],uint8,uint8)                  | 0x6845a025 |
|            updateParams(address,address,uint256,uint256,uint256)             | 0xe880afd4 |
|                              isTxUsed(bytes32)                               | 0xe6ca2084 |
|                             getCurrentPriceLP()                              | 0x45137e27 |
|                       getFloatReserve(address,address)                       | 0xec482729 |
|                               getActiveNodes()                               | 0x6b51e919 |
|                             isNodeStake(address)                             | 0xa742329d |
+------------------------------------------------------------------------------+------------+

sbBTCPool:
+---------------------------------+------------+
|               Name              |     ID     |
+---------------------------------+------------+
|             owner()             | 0x8da5cb5b |
|       renounceOwnership()       | 0x715018a6 |
|    transferOwnership(address)   | 0xf2fde38b |
| setBarnAndSwap(address,address) | 0x1ed64040 |
|        updateAll(uint256)       | 0xa616f345 |
|    resetUnstakedNode(address)   | 0x58411c09 |
|             claim()             | 0x4e71d92d |
|            ackFunds()           | 0xacfd9325 |
|       emergencyWithdraw()       | 0xdb2e21bc |
|         balanceBefore()         | 0x94b5798a |
|       currentMultiplier()       | 0x6fbaaa1e |
|        totalNodeStaked()        | 0x8ca0a7e5 |
|     userMultiplier(address)     | 0xb1a03b6b |
|          owed(address)          | 0xdf18e047 |
|              barn()             | 0x194f480e |
|          rewardToken()          | 0xf7c618c1 |
|          swapContract()         | 0x8ea83031 |
+---------------------------------+------------+


Params:
+--------------------------------------+------------+
|                 Name                 |     ID     |
+--------------------------------------+------------+
|      minimumSwapAmountForWBTC()      | 0x1411c5b7 |
|           expirationTime()           | 0xda284dcc |
|          paraswapAddress()           | 0xf10d7c35 |
|          nodeRewardsRatio()          | 0x0b68134d |
|           depositFeesBPS()           | 0x42419255 |
|          withdrawalFeeBPS()          | 0xb6268e5d |
|             loopCount()              | 0xe91675b8 |
|               owner()                | 0x8da5cb5b |
|         renounceOwnership()          | 0x715018a6 |
|      transferOwnership(address)      | 0xf2fde38b |
|            constructor()             | 0x90fa17bb |
| setMinimumSwapAmountForWBTC(uint256) | 0x5f79ec58 |
|      setExpirationTime(uint256)      | 0xc0cc365d |
|     setParaswapAddress(address)      | 0xfb278493 |
|      setNodeRewardsRatio(uint8)      | 0xa0011cd4 |
|      setWithdrawalFeeBPS(uint8)      | 0xdd62a515 |
|       setDepositFeesBPS(uint8)       | 0x6ce81f80 |
|         setLoopCount(uint8)          | 0x56151c06 |
|      minimumSwapAmountForWBTC()      | 0x1411c5b7 |
|           expirationTime()           | 0xda284dcc |
|          paraswapAddress()           | 0xf10d7c35 |
|          nodeRewardsRatio()          | 0x0b68134d |
|           depositFeesBPS()           | 0x42419255 |
|          withdrawalFeeBPS()          | 0xb6268e5d |
|             loopCount()              | 0xe91675b8 |
+--------------------------------------+------------+


SafeMath:
+------+----+
| Name | ID |
+------+----+
+------+----+

SwapContract:
+------------------------------------------------------------------------------+------------+
|                                     Name                                     |     ID     |
+------------------------------------------------------------------------------+------------+
|                                   owner()                                    | 0x8da5cb5b |
|                             renounceOwnership()                              | 0x715018a6 |
|                          transferOwnership(address)                          | 0xf2fde38b |
|                                 BTCT_ADDR()                                  | 0x0f909486 |
|                                  lpToken()                                   | 0x5fcbd285 |
|     constructor(address,address,address,address,address,uint256,uint256)     | 0x94a103a9 |
|    singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])    | 0x0d63aca7 |
| multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[]) | 0xad289e76 |
|      collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])      | 0x2adf9f87 |
|                 recordIncomingFloat(address,bytes32,bytes32)                 | 0xcf10b16b |
|            recordOutcomingFloat(address,bytes32,uint256,bytes32)             | 0x2586c562 |
|                                  fallback()                                  | 0x552079dc |
|                   recordUTXOSweepMinerFee(uint256,bytes32)                   | 0xc810a539 |
|                 churn(address,address[],bool[],uint8,uint8)                  | 0x6845a025 |
|            updateParams(address,address,uint256,uint256,uint256)             | 0xe880afd4 |
|                              isTxUsed(bytes32)                               | 0xe6ca2084 |
|                             getCurrentPriceLP()                              | 0x45137e27 |
|                       getFloatReserve(address,address)                       | 0xec482729 |
|                               getActiveNodes()                               | 0x6b51e919 |
|                             isNodeStake(address)                             | 0xa742329d |
|                                  lpToken()                                   | 0x5fcbd285 |
|                                     sw()                                     | 0x00e5cee4 |
|                              whitelist(address)                              | 0x9b19251a |
|                                 BTCT_ADDR()                                  | 0x0f909486 |
|                               buybackAddress()                               | 0xcc2fbd66 |
|                                 sbBTCPool()                                  | 0x0085aea1 |
|                              withdrawalFeeBPS()                              | 0xb6268e5d |
|                              nodeRewardsRatio()                              | 0x0b68134d |
|                            buybackRewardsRatio()                             | 0x0b513f1d |
|                              activeNodeCount()                               | 0x75340815 |
|                               churnedInCount()                               | 0x0089356f |
|                                tssThreshold()                                | 0x12d1441e |
+------------------------------------------------------------------------------+------------+

IBurnableToken:
+---------------------------------------+------------+
|                  Name                 |     ID     |
+---------------------------------------+------------+
|             totalSupply()             | 0x18160ddd |
|               decimals()              | 0x313ce567 |
|                symbol()               | 0x95d89b41 |
|                 name()                | 0x06fdde03 |
|               getOwner()              | 0x893d20e8 |
|           balanceOf(address)          | 0x70a08231 |
|       transfer(address,uint256)       | 0xa9059cbb |
|       allowance(address,address)      | 0xdd62ed3e |
|        approve(address,uint256)       | 0x095ea7b3 |
| transferFrom(address,address,uint256) | 0x23b872dd |
|         mint(address,uint256)         | 0x40c10f19 |
|             burn(uint256)             | 0x42966c68 |
|               mintable()              | 0x4bf365df |
+---------------------------------------+------------+

ISwapRewards:
+-----------------------------------------------+------------+
|                      Name                     |     ID     |
+-----------------------------------------------+------------+
|            setSWINGBYPrice(uint256)           | 0xfabe44a8 |
|      pullRewards(address,address,uint256)     | 0xff320aae |
| pullRewardsMulti(address,address[],uint256[]) | 0x49e031a5 |
+-----------------------------------------------+------------+

Address:
+------+----+
| Name | ID |
+------+----+
+------+----+

SafeERC20:
+------+----+
| Name | ID |
+------+----+
+------+----+


SafeMath:
+------+----+
| Name | ID |
+------+----+
+------+----+

SwapRewards:
+-----------------------------------------------+------------+
|                      Name                     |     ID     |
+-----------------------------------------------+------------+
|                    owner()                    | 0x8da5cb5b |
|              renounceOwnership()              | 0x715018a6 |
|           transferOwnership(address)          | 0xf2fde38b |
|      constructor(address,address,uint256)     | 0x3bdb4e02 |
|            setSWINGBYPrice(uint256)           | 0xfabe44a8 |
|        setSwap(address,uint256,uint256)       | 0xa9d6e083 |
|      pullRewards(address,address,uint256)     | 0xff320aae |
| pullRewardsMulti(address,address[],uint256[]) | 0x49e031a5 |
|                 rewardToken()                 | 0xf7c618c1 |
|                 swapContract()                | 0x8ea83031 |
|                  rebateRate()                 | 0xb681a1e7 |
|                thresholdRatio()               | 0xc0324da1 |
|                 pricePerBTC()                 | 0x4017fe77 |
+-----------------------------------------------+------------+

IBurnableToken:
+---------------------------------------+------------+
|                  Name                 |     ID     |
+---------------------------------------+------------+
|             totalSupply()             | 0x18160ddd |
|               decimals()              | 0x313ce567 |
|                symbol()               | 0x95d89b41 |
|                 name()                | 0x06fdde03 |
|               getOwner()              | 0x893d20e8 |
|           balanceOf(address)          | 0x70a08231 |
|       transfer(address,uint256)       | 0xa9059cbb |
|       allowance(address,address)      | 0xdd62ed3e |
|        approve(address,uint256)       | 0x095ea7b3 |
| transferFrom(address,address,uint256) | 0x23b872dd |
|         mint(address,uint256)         | 0x40c10f19 |
|             burn(uint256)             | 0x42966c68 |
|               mintable()              | 0x4bf365df |
+---------------------------------------+------------+

ISwapContract:
+------------------------------------------------------------------------------+------------+
|                                     Name                                     |     ID     |
+------------------------------------------------------------------------------+------------+
|                                 BTCT_ADDR()                                  | 0x0f909486 |
|                                  lpToken()                                   | 0x5fcbd285 |
|    singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])    | 0x0d63aca7 |
| multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[]) | 0xad289e76 |
|      collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])      | 0x2adf9f87 |
|                 recordIncomingFloat(address,bytes32,bytes32)                 | 0xcf10b16b |
|            recordOutcomingFloat(address,bytes32,uint256,bytes32)             | 0x2586c562 |
|                   recordUTXOSweepMinerFee(uint256,bytes32)                   | 0xc810a539 |
|                 churn(address,address[],bool[],uint8,uint8)                  | 0x6845a025 |
|            updateParams(address,address,uint256,uint256,uint256)             | 0xe880afd4 |
|                              isTxUsed(bytes32)                               | 0xe6ca2084 |
|                             getCurrentPriceLP()                              | 0x45137e27 |
|                       getFloatReserve(address,address)                       | 0xec482729 |
|                               getActiveNodes()                               | 0x6b51e919 |
|                             isNodeStake(address)                             | 0xa742329d |
+------------------------------------------------------------------------------+------------+


SafeMath:
+------+----+
| Name | ID |
+------+----+
+------+----+

BurnableToken:
+---------------------------------------+------------+
|                  Name                 |     ID     |
+---------------------------------------+------------+
|                owner()                | 0x8da5cb5b |
|          renounceOwnership()          | 0x715018a6 |
|       transferOwnership(address)      | 0xf2fde38b |
|               mintable()              | 0x4bf365df |
|               decimals()              | 0x313ce567 |
|                symbol()               | 0x95d89b41 |
|                 name()                | 0x06fdde03 |
|               getOwner()              | 0x893d20e8 |
|             totalSupply()             | 0x18160ddd |
|           balanceOf(address)          | 0x70a08231 |
|       transfer(address,uint256)       | 0xa9059cbb |
|       allowance(address,address)      | 0xdd62ed3e |
|        approve(address,uint256)       | 0x095ea7b3 |
| transferFrom(address,address,uint256) | 0x23b872dd |
|   increaseAllowance(address,uint256)  | 0x39509351 |
|   decreaseAllowance(address,uint256)  | 0xa457c2d7 |
|         mint(address,uint256)         | 0x40c10f19 |
|             burn(uint256)             | 0x42966c68 |
+---------------------------------------+------------+


SafeMath:
+------+----+
| Name | ID |
+------+----+
+------+----+

LPToken:
+---------------------------------------+------------+
|                  Name                 |     ID     |
+---------------------------------------+------------+
|               mintable()              | 0x4bf365df |
|               decimals()              | 0x313ce567 |
|                symbol()               | 0x95d89b41 |
|                 name()                | 0x06fdde03 |
|               getOwner()              | 0x893d20e8 |
|             totalSupply()             | 0x18160ddd |
|           balanceOf(address)          | 0x70a08231 |
|       transfer(address,uint256)       | 0xa9059cbb |
|       allowance(address,address)      | 0xdd62ed3e |
|        approve(address,uint256)       | 0x095ea7b3 |
| transferFrom(address,address,uint256) | 0x23b872dd |
|   increaseAllowance(address,uint256)  | 0x39509351 |
|   decreaseAllowance(address,uint256)  | 0xa457c2d7 |
|         mint(address,uint256)         | 0x40c10f19 |
|             burn(uint256)             | 0x42966c68 |
|                owner()                | 0x8da5cb5b |
|          renounceOwnership()          | 0x715018a6 |
|       transferOwnership(address)      | 0xf2fde38b |
|           constructor(uint8)          | 0x9f217ad9 |
|            setOld(address)            | 0x242d81f0 |
|       convertTo(address,uint256)      | 0x9bfa5181 |
|       mintFrom(address,uint256)       | 0x1cc74859 |
|                 old()                 | 0xb83f8663 |
+---------------------------------------+------------+


SafeMath:
+------+----+
| Name | ID |
+------+----+
+------+----+

SwingbyToken:
+---------------------------------------+------------+
|                  Name                 |     ID     |
+---------------------------------------+------------+
|               mintable()              | 0x4bf365df |
|               decimals()              | 0x313ce567 |
|                symbol()               | 0x95d89b41 |
|                 name()                | 0x06fdde03 |
|               getOwner()              | 0x893d20e8 |
|             totalSupply()             | 0x18160ddd |
|           balanceOf(address)          | 0x70a08231 |
|       transfer(address,uint256)       | 0xa9059cbb |
|       allowance(address,address)      | 0xdd62ed3e |
|        approve(address,uint256)       | 0x095ea7b3 |
| transferFrom(address,address,uint256) | 0x23b872dd |
|   increaseAllowance(address,uint256)  | 0x39509351 |
|   decreaseAllowance(address,uint256)  | 0xa457c2d7 |
|         mint(address,uint256)         | 0x40c10f19 |
|             burn(uint256)             | 0x42966c68 |
|                owner()                | 0x8da5cb5b |
|          renounceOwnership()          | 0x715018a6 |
|       transferOwnership(address)      | 0xf2fde38b |
|             constructor()             | 0x90fa17bb |
+---------------------------------------+------------+
```


## unchecked-transfer
Impact: High
Confidence: Medium
 - [ ] ID-0
[sbBTCPool.claim()](contracts/sbBTCPool.sol#L73-L95) ignores return value by [rewardToken.transfer(msg.sender,amount)](contracts/sbBTCPool.sol#L87)

contracts/sbBTCPool.sol#L73-L95


 - [ ] ID-1
[sbBTCPool.emergencyWithdraw()](contracts/sbBTCPool.sol#L120-L124) ignores return value by [rewardToken.transfer(msg.sender,rewardToken.balanceOf(address(this)))](contracts/sbBTCPool.sol#L122)

contracts/sbBTCPool.sol#L120-L124


 - [ ] ID-2
[SwapRewards.pullRewardsMulti(address,address[],uint256[])](contracts/SwapRewards.sol#L86-L116) ignores return value by [rewardToken.transfer(_receiver[i],amount)](contracts/SwapRewards.sol#L111)

contracts/SwapRewards.sol#L86-L116


 - [ ] ID-3
[SwapRewards.pullRewards(address,address,uint256)](contracts/SwapRewards.sol#L59-L83) ignores return value by [rewardToken.transfer(_receiver,amount)](contracts/SwapRewards.sol#L79)

contracts/SwapRewards.sol#L59-L83


## incorrect-equality
Impact: Medium
Confidence: High
 - [ ] ID-4
[sbBTCPool.updateStakes(uint256)](contracts/sbBTCPool.sol#L50-L65) uses a dangerous strict equality:
	- [userMultiplier[nodes[i]] == 0](contracts/sbBTCPool.sol#L57)

contracts/sbBTCPool.sol#L50-L65


 - [ ] ID-5
[sbBTCPool.ackFunds()](contracts/sbBTCPool.sol#L100-L118) uses a dangerous strict equality:
	- [balanceNow == 0 || balanceNow <= balanceBefore](contracts/sbBTCPool.sol#L103)

contracts/sbBTCPool.sol#L100-L118


 - [ ] ID-6
[sbBTCPool.ackFunds()](contracts/sbBTCPool.sol#L100-L118) uses a dangerous strict equality:
	- [totalNodeStaked == 0](contracts/sbBTCPool.sol#L107)

contracts/sbBTCPool.sol#L100-L118


## reentrancy-no-eth
Impact: Medium
Confidence: Medium
 - [ ] ID-7
Reentrancy in [sbBTCPool.claim()](contracts/sbBTCPool.sol#L73-L95):
	External calls:
	- [require(bool,string)(swapContract.isNodeStake(msg.sender),caller is not node)](contracts/sbBTCPool.sol#L74)
	- [rewardToken.transfer(msg.sender,amount)](contracts/sbBTCPool.sol#L87)
	State variables written after the call(s):
	- [ackFunds()](contracts/sbBTCPool.sol#L90)
		- [balanceBefore = balanceNow](contracts/sbBTCPool.sol#L104)
		- [balanceBefore = balanceNow](contracts/sbBTCPool.sol#L116)
	- [ackFunds()](contracts/sbBTCPool.sol#L90)
		- [currentMultiplier = multiplier](contracts/sbBTCPool.sol#L117)

contracts/sbBTCPool.sol#L73-L95


 - [ ] ID-8
Reentrancy in [SwapContract._burnLPTokensForFloat(address,bytes32,uint256,bytes32)](contracts/SwapContract.sol#L439-L497):
	External calls:
	- [_rewardsCollection(_token,withdrawalFees)](contracts/SwapContract.sol#L476)
		- [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)
		- [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)
	State variables written after the call(s):
	- [_removeFloat(_token,amountOfFloat)](contracts/SwapContract.sol#L478)
		- [floatAmountOf[_token] = floatAmountOf[_token].sub(_amount,10)](contracts/SwapContract.sol#L510-L513)
	- [_addUsedTx(_txid)](contracts/SwapContract.sol#L480)
		- [used[_txid] = true](contracts/SwapContract.sol#L591)

contracts/SwapContract.sol#L439-L497


 - [ ] ID-9
Reentrancy in [SwapContract._issueLPTokensForFloat(address,bytes32,bytes32)](contracts/SwapContract.sol#L404-L432):
	External calls:
	- [lpToken.mint(to,amountOfLP)](contracts/SwapContract.sol#L418)
	State variables written after the call(s):
	- [_addFloat(_token,amountOfFloat)](contracts/SwapContract.sol#L420)
		- [floatAmountOf[_token] = floatAmountOf[_token].add(_amount)](contracts/SwapContract.sol#L503)
	- [_addUsedTx(_txid)](contracts/SwapContract.sol#L421)
		- [used[_txid] = true](contracts/SwapContract.sol#L591)

contracts/SwapContract.sol#L404-L432


## calls-loop
Impact: Low
Confidence: Medium
 - [ ] ID-10
[sbBTCPool._userPendingReward(address,uint256)](contracts/sbBTCPool.sol#L137-L143) has external calls inside a loop: [barn.balanceAtTs(user,timestamp).mul(multiplier).div(divisor)](contracts/sbBTCPool.sol#L142)

contracts/sbBTCPool.sol#L137-L143


 - [ ] ID-11
[Address.functionCallWithValue(address,bytes,uint256,string)](contracts/interfaces/lib/Address.sol#L122-L133) has external calls inside a loop: [(success,returndata) = target.call{value: value}(data)](contracts/interfaces/lib/Address.sol#L131)

contracts/interfaces/lib/Address.sol#L122-L133


 - [ ] ID-12
[SwapRewards.pullRewardsMulti(address,address[],uint256[])](contracts/SwapRewards.sol#L86-L116) has external calls inside a loop: [rewardToken.transfer(_receiver[i],amount)](contracts/SwapRewards.sol#L111)

contracts/SwapRewards.sol#L86-L116


## reentrancy-benign
Impact: Low
Confidence: Medium
 - [ ] ID-13
Reentrancy in [sbBTCPool.emergencyWithdraw()](contracts/sbBTCPool.sol#L120-L124):
	External calls:
	- [rewardToken.transfer(msg.sender,rewardToken.balanceOf(address(this)))](contracts/sbBTCPool.sol#L122)
	State variables written after the call(s):
	- [swapContract = ISwapContract(address(0))](contracts/sbBTCPool.sol#L123)

contracts/sbBTCPool.sol#L120-L124


 - [ ] ID-14
Reentrancy in [sbBTCPool.resetUnstakedNode(address)](contracts/sbBTCPool.sol#L67-L70):
	External calls:
	- [require(bool,string)(! swapContract.isNodeStake(_node),node is staker)](contracts/sbBTCPool.sol#L68)
	State variables written after the call(s):
	- [userMultiplier[_node] = 0](contracts/sbBTCPool.sol#L69)

contracts/sbBTCPool.sol#L67-L70


 - [ ] ID-15
Reentrancy in [sbBTCPool.setBarnAndSwap(address,address)](contracts/sbBTCPool.sol#L28-L35):
	External calls:
	- [rewardToken = IERC20(swapContract.lpToken())](contracts/sbBTCPool.sol#L33)
	State variables written after the call(s):
	- [barn = IBarn(_barn)](contracts/sbBTCPool.sol#L34)

contracts/sbBTCPool.sol#L28-L35


 - [ ] ID-16
Reentrancy in [sbBTCPool.claim()](contracts/sbBTCPool.sol#L73-L95):
	External calls:
	- [require(bool,string)(swapContract.isNodeStake(msg.sender),caller is not node)](contracts/sbBTCPool.sol#L74)
	State variables written after the call(s):
	- [ackFunds()](contracts/sbBTCPool.sol#L78)
		- [balanceBefore = balanceNow](contracts/sbBTCPool.sol#L104)
		- [balanceBefore = balanceNow](contracts/sbBTCPool.sol#L116)
	- [ackFunds()](contracts/sbBTCPool.sol#L78)
		- [currentMultiplier = multiplier](contracts/sbBTCPool.sol#L117)
	- [_updateOwed(msg.sender,block.timestamp)](contracts/sbBTCPool.sol#L80)
		- [owed[user] = owed[user].add(reward)](contracts/sbBTCPool.sol#L131)
	- [owed[msg.sender] = 0](contracts/sbBTCPool.sol#L85)
	- [updateStakes(block.timestamp)](contracts/sbBTCPool.sol#L76)
		- [totalNodeStaked = newTotalNodeStaked](contracts/sbBTCPool.sol#L63)
	- [updateStakes(block.timestamp)](contracts/sbBTCPool.sol#L76)
		- [userMultiplier[nodes[i]] = currentMultiplier](contracts/sbBTCPool.sol#L58)
	- [_updateOwed(msg.sender,block.timestamp)](contracts/sbBTCPool.sol#L80)
		- [userMultiplier[user] = currentMultiplier](contracts/sbBTCPool.sol#L132)

contracts/sbBTCPool.sol#L73-L95


 - [ ] ID-17
Reentrancy in [SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])](contracts/SwapContract.sol#L190-L207):
	External calls:
	- [sw.pullRewardsMulti(address(0),_spenders,_swapAmounts)](contracts/SwapContract.sol#L200)
	State variables written after the call(s):
	- [_swap(BTCT_ADDR,address(0),swapAmount)](contracts/SwapContract.sol#L201)
		- [floatAmountOf[_destToken] = floatAmountOf[_destToken].sub(_swapAmount,11)](contracts/SwapContract.sol#L525-L528)
		- [floatAmountOf[_sourceToken] = floatAmountOf[_sourceToken].add(_swapAmount)](contracts/SwapContract.sol#L529-L531)

contracts/SwapContract.sol#L190-L207


 - [ ] ID-18
Reentrancy in [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])](contracts/SwapContract.sol#L120-L146):
	External calls:
	- [sw.pullRewards(_destToken,_to,_totalSwapped)](contracts/SwapContract.sol#L135)
	State variables written after the call(s):
	- [_swap(address(0),BTCT_ADDR,_totalSwapped)](contracts/SwapContract.sol#L136)
		- [floatAmountOf[_destToken] = floatAmountOf[_destToken].sub(_swapAmount,11)](contracts/SwapContract.sol#L525-L528)
		- [floatAmountOf[_sourceToken] = floatAmountOf[_sourceToken].add(_swapAmount)](contracts/SwapContract.sol#L529-L531)

contracts/SwapContract.sol#L120-L146


 - [ ] ID-19
Reentrancy in [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])](contracts/SwapContract.sol#L120-L146):
	External calls:
	- [sw.pullRewards(_destToken,_to,_totalSwapped)](contracts/SwapContract.sol#L135)
	- [_rewardsCollection(_feesToken,_rewardsAmount)](contracts/SwapContract.sol#L142)
		- [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)
		- [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)
	State variables written after the call(s):
	- [_addUsedTxs(_redeemedFloatTxIds)](contracts/SwapContract.sol#L143)
		- [used[_txids[i]] = true](contracts/SwapContract.sol#L598)

contracts/SwapContract.sol#L120-L146


 - [ ] ID-20
Reentrancy in [SwapContract.multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[])](contracts/SwapContract.sol#L154-L184):
	External calls:
	- [_rewardsCollection(_feesToken,_rewardsAmount)](contracts/SwapContract.sol#L174)
		- [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)
		- [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)
	State variables written after the call(s):
	- [_addUsedTxs(_redeemedFloatTxIds)](contracts/SwapContract.sol#L175)
		- [used[_txids[i]] = true](contracts/SwapContract.sol#L598)

contracts/SwapContract.sol#L154-L184


## reentrancy-events
Impact: Low
Confidence: Medium
 - [ ] ID-21
Reentrancy in [sbBTCPool.claim()](contracts/sbBTCPool.sol#L73-L95):
	External calls:
	- [require(bool,string)(swapContract.isNodeStake(msg.sender),caller is not node)](contracts/sbBTCPool.sol#L74)
	- [rewardToken.transfer(msg.sender,amount)](contracts/sbBTCPool.sol#L87)
	Event emitted after the call(s):
	- [Claim(msg.sender,amount)](contracts/sbBTCPool.sol#L92)

contracts/sbBTCPool.sol#L73-L95


 - [ ] ID-22
Reentrancy in [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])](contracts/SwapContract.sol#L120-L146):
	External calls:
	- [sw.pullRewards(_destToken,_to,_totalSwapped)](contracts/SwapContract.sol#L135)
	Event emitted after the call(s):
	- [Swap(_sourceToken,_destToken,_swapAmount)](contracts/SwapContract.sol#L533)
		- [_swap(address(0),BTCT_ADDR,_totalSwapped)](contracts/SwapContract.sol#L136)

contracts/SwapContract.sol#L120-L146


 - [ ] ID-23
Reentrancy in [SwapContract._burnLPTokensForFloat(address,bytes32,uint256,bytes32)](contracts/SwapContract.sol#L439-L497):
	External calls:
	- [_rewardsCollection(_token,withdrawalFees)](contracts/SwapContract.sol#L476)
		- [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)
		- [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)
	- [_safeTransfer(_token,to,withdrawal)](contracts/SwapContract.sol#L484)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](contracts/interfaces/lib/SafeERC20.sol#L81)
		- [IERC20(_token).safeTransfer(_to,_amount)](contracts/SwapContract.sol#L548)
		- [(success,returndata) = target.call{value: value}(data)](contracts/interfaces/lib/Address.sol#L131)
	- [require(bool)(lpToken.burn(amountOfLP))](contracts/SwapContract.sol#L487)
	External calls sending eth:
	- [_safeTransfer(_token,to,withdrawal)](contracts/SwapContract.sol#L484)
		- [(success,returndata) = target.call{value: value}(data)](contracts/interfaces/lib/Address.sol#L131)
	Event emitted after the call(s):
	- [BurnLPTokensForFloat(to,amountOfLP,amountOfFloat,nowPrice,withdrawal,_txid)](contracts/SwapContract.sol#L488-L495)

contracts/SwapContract.sol#L439-L497


 - [ ] ID-24
Reentrancy in [SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])](contracts/SwapContract.sol#L190-L207):
	External calls:
	- [sw.pullRewardsMulti(address(0),_spenders,_swapAmounts)](contracts/SwapContract.sol#L200)
	Event emitted after the call(s):
	- [Swap(_sourceToken,_destToken,_swapAmount)](contracts/SwapContract.sol#L533)
		- [_swap(BTCT_ADDR,address(0),swapAmount)](contracts/SwapContract.sol#L201)

contracts/SwapContract.sol#L190-L207


 - [ ] ID-25
Reentrancy in [SwapContract._rewardsCollection(address,uint256)](contracts/SwapContract.sol#L554-L586):
	External calls:
	- [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)
	- [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)
	Event emitted after the call(s):
	- [RewardsCollection(_feesToken,_rewardsAmount,feesLPTTotal,nowPrice)](contracts/SwapContract.sol#L580-L585)

contracts/SwapContract.sol#L554-L586


 - [ ] ID-26
Reentrancy in [SwapContract._issueLPTokensForFloat(address,bytes32,bytes32)](contracts/SwapContract.sol#L404-L432):
	External calls:
	- [lpToken.mint(to,amountOfLP)](contracts/SwapContract.sol#L418)
	Event emitted after the call(s):
	- [IssueLPTokensForFloat(to,amountOfFloat,amountOfLP,nowPrice,0,_txid)](contracts/SwapContract.sol#L423-L430)

contracts/SwapContract.sol#L404-L432


 - [ ] ID-27
Reentrancy in [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])](contracts/SwapContract.sol#L120-L146):
	External calls:
	- [sw.pullRewards(_destToken,_to,_totalSwapped)](contracts/SwapContract.sol#L135)
	- [_rewardsCollection(_feesToken,_rewardsAmount)](contracts/SwapContract.sol#L142)
		- [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)
		- [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)
	Event emitted after the call(s):
	- [RewardsCollection(_feesToken,_rewardsAmount,feesLPTTotal,nowPrice)](contracts/SwapContract.sol#L580-L585)
		- [_rewardsCollection(_feesToken,_rewardsAmount)](contracts/SwapContract.sol#L142)

contracts/SwapContract.sol#L120-L146


 - [ ] ID-28
Reentrancy in [SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])](contracts/SwapContract.sol#L190-L207):
	External calls:
	- [sw.pullRewardsMulti(address(0),_spenders,_swapAmounts)](contracts/SwapContract.sol#L200)
	- [_rewardsCollection(_feesToken,_rewardsAmount)](contracts/SwapContract.sol#L205)
		- [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)
		- [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)
	Event emitted after the call(s):
	- [RewardsCollection(_feesToken,_rewardsAmount,feesLPTTotal,nowPrice)](contracts/SwapContract.sol#L580-L585)
		- [_rewardsCollection(_feesToken,_rewardsAmount)](contracts/SwapContract.sol#L205)

contracts/SwapContract.sol#L190-L207


 - [ ] ID-29
Reentrancy in [SwapRewards.pullRewards(address,address,uint256)](contracts/SwapRewards.sol#L59-L83):
	External calls:
	- [tokenB = swapContract.BTCT_ADDR()](contracts/SwapRewards.sol#L68)
	- [(balA,balB) = swapContract.getFloatReserve(address(0),tokenB)](contracts/SwapRewards.sol#L69-L72)
	- [rewardToken.transfer(_receiver,amount)](contracts/SwapRewards.sol#L79)
	Event emitted after the call(s):
	- [Paid(_receiver,_swapped,amount)](contracts/SwapRewards.sol#L80)

contracts/SwapRewards.sol#L59-L83


 - [ ] ID-30
Reentrancy in [SwapRewards.pullRewardsMulti(address,address[],uint256[])](contracts/SwapRewards.sol#L86-L116):
	External calls:
	- [tokenB = swapContract.BTCT_ADDR()](contracts/SwapRewards.sol#L96)
	- [(balA,balB) = swapContract.getFloatReserve(address(0),tokenB)](contracts/SwapRewards.sol#L97-L100)
	- [rewardToken.transfer(_receiver[i],amount)](contracts/SwapRewards.sol#L111)
	Event emitted after the call(s):
	- [Paid(_receiver[i],_swapped[i],amount)](contracts/SwapRewards.sol#L112)

contracts/SwapRewards.sol#L86-L116


## timestamp
Impact: Low
Confidence: Medium
 - [ ] ID-31
[sbBTCPool.ackFunds()](contracts/sbBTCPool.sol#L100-L118) uses timestamp for comparisons
	Dangerous comparisons:
	- [totalNodeStaked == 0](contracts/sbBTCPool.sol#L107)

contracts/sbBTCPool.sol#L100-L118


 - [ ] ID-32
[sbBTCPool.claim()](contracts/sbBTCPool.sol#L73-L95) uses timestamp for comparisons
	Dangerous comparisons:
	- [require(bool,string)(amount > 0,nothing to claim)](contracts/sbBTCPool.sol#L83)

contracts/sbBTCPool.sol#L73-L95


 - [ ] ID-33
[sbBTCPool.updateStakes(uint256)](contracts/sbBTCPool.sol#L50-L65) uses timestamp for comparisons
	Dangerous comparisons:
	- [userMultiplier[nodes[i]] == 0](contracts/sbBTCPool.sol#L57)
	- [totalNodeStaked != newTotalNodeStaked](contracts/sbBTCPool.sol#L62)

contracts/sbBTCPool.sol#L50-L65


## dead-code
Impact: Informational
Confidence: Medium
 - [ ] ID-34
[SafeMath.tryDiv(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69


 - [ ] ID-35
[SafeMath.tryMod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81


 - [ ] ID-36
[SafeMath.sub(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L168-L177) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L168-L177


 - [ ] ID-37
[SafeMath.tryAdd(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28


 - [ ] ID-38
[SafeMath.mod(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226


 - [ ] ID-39
[SafeMath.div(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200


 - [ ] ID-40
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-41
[SafeMath.mod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153


 - [ ] ID-42
[SafeMath.tryMul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57


 - [ ] ID-43
[SafeMath.trySub(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40


 - [ ] ID-44
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-45
[SafeMath.tryDiv(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69


 - [ ] ID-46
[Address.sendValue(address,uint256)](contracts/interfaces/lib/Address.sol#L54-L59) is never used and should be removed

contracts/interfaces/lib/Address.sol#L54-L59


 - [ ] ID-47
[Address.functionCallWithValue(address,bytes,uint256)](contracts/interfaces/lib/Address.sol#L108-L114) is never used and should be removed

contracts/interfaces/lib/Address.sol#L108-L114


 - [ ] ID-48
[SafeMath.tryMod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81


 - [ ] ID-49
[Address.functionDelegateCall(address,bytes,string)](contracts/interfaces/lib/Address.sol#L178-L187) is never used and should be removed

contracts/interfaces/lib/Address.sol#L178-L187


 - [ ] ID-50
[Address.functionDelegateCall(address,bytes)](contracts/interfaces/lib/Address.sol#L168-L170) is never used and should be removed

contracts/interfaces/lib/Address.sol#L168-L170


 - [ ] ID-51
[SafeERC20.safeIncreaseAllowance(IERC20,address,uint256)](contracts/interfaces/lib/SafeERC20.sol#L59-L66) is never used and should be removed

contracts/interfaces/lib/SafeERC20.sol#L59-L66


 - [ ] ID-52
[SafeERC20.safeApprove(IERC20,address,uint256)](contracts/interfaces/lib/SafeERC20.sol#L44-L57) is never used and should be removed

contracts/interfaces/lib/SafeERC20.sol#L44-L57


 - [ ] ID-53
[SafeMath.tryAdd(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28


 - [ ] ID-54
[SafeMath.mod(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226


 - [ ] ID-55
[SafeERC20.safeTransferFrom(IERC20,address,address,uint256)](contracts/interfaces/lib/SafeERC20.sol#L28-L35) is never used and should be removed

contracts/interfaces/lib/SafeERC20.sol#L28-L35


 - [ ] ID-56
[SafeMath.div(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200


 - [ ] ID-57
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-58
[Address.functionStaticCall(address,bytes)](contracts/interfaces/lib/Address.sol#L141-L143) is never used and should be removed

contracts/interfaces/lib/Address.sol#L141-L143


 - [ ] ID-59
[SafeMath.mod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153


 - [ ] ID-60
[SafeMath.tryMul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57


 - [ ] ID-61
[SafeMath.trySub(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40


 - [ ] ID-62
[Address.functionStaticCall(address,bytes,string)](contracts/interfaces/lib/Address.sol#L151-L160) is never used and should be removed

contracts/interfaces/lib/Address.sol#L151-L160


 - [ ] ID-63
[Address.functionCall(address,bytes)](contracts/interfaces/lib/Address.sol#L79-L81) is never used and should be removed

contracts/interfaces/lib/Address.sol#L79-L81


 - [ ] ID-64
[SafeMath.sub(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L107-L109) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L107-L109


 - [ ] ID-65
[SafeMath.tryDiv(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69


 - [ ] ID-66
[SafeMath.tryMod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81


 - [ ] ID-67
[SafeMath.sub(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L168-L177) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L168-L177


 - [ ] ID-68
[SafeMath.tryAdd(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28


 - [ ] ID-69
[SafeMath.mod(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226


 - [ ] ID-70
[SafeMath.div(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200


 - [ ] ID-71
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-72
[SafeMath.mod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153


 - [ ] ID-73
[SafeMath.tryMul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57


 - [ ] ID-74
[SafeMath.trySub(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40


 - [ ] ID-75
[SafeMath.mul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L121-L123) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L121-L123


 - [ ] ID-76
[SafeMath.tryDiv(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69


 - [ ] ID-77
[SafeMath.tryMod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81


 - [ ] ID-78
[BurnableToken._burnFrom(address,uint256)](contracts/BurnableToken.sol#L342-L352) is never used and should be removed

contracts/BurnableToken.sol#L342-L352


 - [ ] ID-79
[SafeMath.tryAdd(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28


 - [ ] ID-80
[SafeMath.mod(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226


 - [ ] ID-81
[BurnableToken._initialize(string,string,uint8,uint256,bool)](contracts/BurnableToken.sol#L23-L35) is never used and should be removed

contracts/BurnableToken.sol#L23-L35


 - [ ] ID-82
[SafeMath.div(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200


 - [ ] ID-83
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-84
[SafeMath.mod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153


 - [ ] ID-85
[SafeMath.div(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L135-L137) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L135-L137


 - [ ] ID-86
[SafeMath.tryMul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57


 - [ ] ID-87
[SafeMath.trySub(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40


 - [ ] ID-88
[SafeMath.mul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L121-L123) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L121-L123


 - [ ] ID-89
[SafeMath.tryDiv(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69


 - [ ] ID-90
[SafeMath.tryMod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81


 - [ ] ID-91
[BurnableToken._burnFrom(address,uint256)](contracts/BurnableToken.sol#L342-L352) is never used and should be removed

contracts/BurnableToken.sol#L342-L352


 - [ ] ID-92
[SafeMath.tryAdd(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28


 - [ ] ID-93
[SafeMath.mod(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226


 - [ ] ID-94
[SafeMath.div(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200


 - [ ] ID-95
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-96
[SafeMath.mod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153


 - [ ] ID-97
[SafeMath.div(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L135-L137) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L135-L137


 - [ ] ID-98
[SafeMath.tryMul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57


 - [ ] ID-99
[SafeMath.trySub(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40


 - [ ] ID-100
[SafeMath.mul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L121-L123) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L121-L123


 - [ ] ID-101
[SafeMath.tryDiv(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L64-L69


 - [ ] ID-102
[SafeMath.tryMod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L76-L81


 - [ ] ID-103
[BurnableToken._burnFrom(address,uint256)](contracts/BurnableToken.sol#L342-L352) is never used and should be removed

contracts/BurnableToken.sol#L342-L352


 - [ ] ID-104
[SafeMath.tryAdd(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L22-L28


 - [ ] ID-105
[SafeMath.mod(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L217-L226


 - [ ] ID-106
[SafeMath.div(uint256,uint256,string)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L191-L200


 - [ ] ID-107
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-108
[SafeMath.mod(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L151-L153


 - [ ] ID-109
[SafeMath.div(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L135-L137) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L135-L137


 - [ ] ID-110
[SafeMath.tryMul(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L47-L57


 - [ ] ID-111
[SafeMath.trySub(uint256,uint256)](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L35-L40


## solc-version
Impact: Informational
Confidence: High
 - [ ] ID-112
Pragma version[^0.8.0](contracts/interfaces/IERC20.sol#L2) allows old versions

contracts/interfaces/IERC20.sol#L2


 - [ ] ID-113
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-114
Pragma version[^0.8.0](contracts/interfaces/IBarn.sol#L2) allows old versions

contracts/interfaces/IBarn.sol#L2


 - [ ] ID-115
Pragma version[^0.8.0](contracts/interfaces/ISwapContract.sol#L2) allows old versions

contracts/interfaces/ISwapContract.sol#L2


 - [ ] ID-116
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-117
Pragma version[^0.8.0](contracts/sbBTCPool.sol#L2) allows old versions

contracts/sbBTCPool.sol#L2


 - [ ] ID-118
solc-0.8.19 is not recommended for deployment

 - [ ] ID-119
Pragma version[^0.8.0](contracts/interfaces/IBurnableToken.sol#L2) allows old versions

contracts/interfaces/IBurnableToken.sol#L2


 - [ ] ID-120
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


 - [ ] ID-121
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-122
Pragma version[^0.8.0](contracts/Params.sol#L2) allows old versions

contracts/Params.sol#L2


 - [ ] ID-123
solc-0.8.19 is not recommended for deployment

 - [ ] ID-124
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


 - [ ] ID-125
Pragma version[^0.8.0](contracts/interfaces/IParams.sol#L2) allows old versions

contracts/interfaces/IParams.sol#L2


 - [ ] ID-126
Pragma version[^0.8.0](contracts/interfaces/IERC20.sol#L2) allows old versions

contracts/interfaces/IERC20.sol#L2


 - [ ] ID-127
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-128
Pragma version[^0.8.0](contracts/SwapContract.sol#L2) allows old versions

contracts/SwapContract.sol#L2


 - [ ] ID-129
Pragma version[^0.8.0](contracts/interfaces/ISwapContract.sol#L2) allows old versions

contracts/interfaces/ISwapContract.sol#L2


 - [ ] ID-130
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-131
Pragma version[^0.8.0](contracts/interfaces/lib/SafeERC20.sol#L3) allows old versions

contracts/interfaces/lib/SafeERC20.sol#L3


 - [ ] ID-132
Pragma version[^0.8.0](contracts/interfaces/lib/Address.sol#L3) allows old versions

contracts/interfaces/lib/Address.sol#L3


 - [ ] ID-133
solc-0.8.19 is not recommended for deployment

 - [ ] ID-134
Pragma version[^0.8.0](contracts/interfaces/ISwapRewards.sol#L2) allows old versions

contracts/interfaces/ISwapRewards.sol#L2


 - [ ] ID-135
Pragma version[^0.8.0](contracts/interfaces/IBurnableToken.sol#L2) allows old versions

contracts/interfaces/IBurnableToken.sol#L2


 - [ ] ID-136
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


 - [ ] ID-137
Pragma version[^0.8.0](contracts/interfaces/IERC20.sol#L2) allows old versions

contracts/interfaces/IERC20.sol#L2


 - [ ] ID-138
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-139
Pragma version[^0.8.0](contracts/interfaces/ISwapContract.sol#L2) allows old versions

contracts/interfaces/ISwapContract.sol#L2


 - [ ] ID-140
Pragma version[^0.8.0](contracts/SwapRewards.sol#L2) allows old versions

contracts/SwapRewards.sol#L2


 - [ ] ID-141
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-142
solc-0.8.19 is not recommended for deployment

 - [ ] ID-143
Pragma version[^0.8.0](contracts/interfaces/IBurnableToken.sol#L2) allows old versions

contracts/interfaces/IBurnableToken.sol#L2


 - [ ] ID-144
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


 - [ ] ID-145
Pragma version[^0.8.0](contracts/BurnableToken.sol#L2) allows old versions

contracts/BurnableToken.sol#L2


 - [ ] ID-146
Pragma version[^0.8.0](contracts/interfaces/IERC20.sol#L2) allows old versions

contracts/interfaces/IERC20.sol#L2


 - [ ] ID-147
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-148
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-149
solc-0.8.19 is not recommended for deployment

 - [ ] ID-150
Pragma version[^0.8.0](contracts/interfaces/IBurnableToken.sol#L2) allows old versions

contracts/interfaces/IBurnableToken.sol#L2


 - [ ] ID-151
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


 - [ ] ID-152
Pragma version[^0.8.0](contracts/BurnableToken.sol#L2) allows old versions

contracts/BurnableToken.sol#L2


 - [ ] ID-153
Pragma version[^0.8.0](contracts/interfaces/IERC20.sol#L2) allows old versions

contracts/interfaces/IERC20.sol#L2


 - [ ] ID-154
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-155
Pragma version[^0.8.0](contracts/LPToken.sol#L2) allows old versions

contracts/LPToken.sol#L2


 - [ ] ID-156
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-157
solc-0.8.19 is not recommended for deployment

 - [ ] ID-158
Pragma version[^0.8.0](contracts/interfaces/IBurnableToken.sol#L2) allows old versions

contracts/interfaces/IBurnableToken.sol#L2


 - [ ] ID-159
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


 - [ ] ID-160
Pragma version[^0.8.0](contracts/BurnableToken.sol#L2) allows old versions

contracts/BurnableToken.sol#L2


 - [ ] ID-161
Pragma version[^0.8.0](contracts/interfaces/IERC20.sol#L2) allows old versions

contracts/interfaces/IERC20.sol#L2


 - [ ] ID-162
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-163
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-164
Pragma version[^0.8.0](contracts/SwingbyToken.sol#L2) allows old versions

contracts/SwingbyToken.sol#L2


 - [ ] ID-165
solc-0.8.19 is not recommended for deployment

 - [ ] ID-166
Pragma version[^0.8.0](contracts/interfaces/IBurnableToken.sol#L2) allows old versions

contracts/interfaces/IBurnableToken.sol#L2


 - [ ] ID-167
Pragma version[^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4) allows old versions

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


## naming-convention
Impact: Informational
Confidence: High
 - [ ] ID-168
Parameter [sbBTCPool.updateAll(uint256)._timestamp](contracts/sbBTCPool.sol#L38) is not in mixedCase

contracts/sbBTCPool.sol#L38


 - [ ] ID-169
Function [ISwapContract.BTCT_ADDR()](contracts/interfaces/ISwapContract.sol#L8) is not in mixedCase

contracts/interfaces/ISwapContract.sol#L8


 - [ ] ID-170
Parameter [sbBTCPool.setBarnAndSwap(address,address)._swap](contracts/sbBTCPool.sol#L28) is not in mixedCase

contracts/sbBTCPool.sol#L28


 - [ ] ID-171
Parameter [sbBTCPool.setBarnAndSwap(address,address)._barn](contracts/sbBTCPool.sol#L28) is not in mixedCase

contracts/sbBTCPool.sol#L28


 - [ ] ID-172
Parameter [sbBTCPool.resetUnstakedNode(address)._node](contracts/sbBTCPool.sol#L67) is not in mixedCase

contracts/sbBTCPool.sol#L67


 - [ ] ID-173
Contract [sbBTCPool](contracts/sbBTCPool.sol#L9-L144) is not in CapWords

contracts/sbBTCPool.sol#L9-L144


 - [ ] ID-174
Parameter [sbBTCPool.updateStakes(uint256)._timestamp](contracts/sbBTCPool.sol#L50) is not in mixedCase

contracts/sbBTCPool.sol#L50


 - [ ] ID-175
Constant [sbBTCPool.divisor](contracts/sbBTCPool.sol#L12) is not in UPPER_CASE_WITH_UNDERSCORES

contracts/sbBTCPool.sol#L12


 - [ ] ID-176
Parameter [Params.setMinimumSwapAmountForWBTC(uint256)._minimumSwapAmountForWBTC](contracts/Params.sol#L33) is not in mixedCase

contracts/Params.sol#L33


 - [ ] ID-177
Parameter [Params.setDepositFeesBPS(uint8)._depositFeesBPS](contracts/Params.sol#L69) is not in mixedCase

contracts/Params.sol#L69


 - [ ] ID-178
Parameter [Params.setParaswapAddress(address)._paraswapAddress](contracts/Params.sol#L49) is not in mixedCase

contracts/Params.sol#L49


 - [ ] ID-179
Parameter [Params.setWithdrawalFeeBPS(uint8)._withdrawalFeeBPS](contracts/Params.sol#L61) is not in mixedCase

contracts/Params.sol#L61


 - [ ] ID-180
Parameter [Params.setNodeRewardsRatio(uint8)._nodeRewardsRatio](contracts/Params.sol#L53) is not in mixedCase

contracts/Params.sol#L53


 - [ ] ID-181
Parameter [Params.setLoopCount(uint8)._loopCount](contracts/Params.sol#L77) is not in mixedCase

contracts/Params.sol#L77


 - [ ] ID-182
Parameter [Params.setExpirationTime(uint256)._expirationTime](contracts/Params.sol#L44) is not in mixedCase

contracts/Params.sol#L44


 - [ ] ID-183
Parameter [SwapContract.churn(address,address[],bool[],uint8,uint8)._nodes](contracts/SwapContract.sol#L283) is not in mixedCase

contracts/SwapContract.sol#L283


 - [ ] ID-184
Parameter [SwapContract.multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[])._rewardsAmount](contracts/SwapContract.sol#L158) is not in mixedCase

contracts/SwapContract.sol#L158


 - [ ] ID-185
Parameter [SwapContract.churn(address,address[],bool[],uint8,uint8)._isRemoved](contracts/SwapContract.sol#L284) is not in mixedCase

contracts/SwapContract.sol#L284


 - [ ] ID-186
Parameter [SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])._incomingAmount](contracts/SwapContract.sol#L191) is not in mixedCase

contracts/SwapContract.sol#L191


 - [ ] ID-187
Parameter [SwapContract.multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[])._destToken](contracts/SwapContract.sol#L155) is not in mixedCase

contracts/SwapContract.sol#L155


 - [ ] ID-188
Parameter [SwapContract.churn(address,address[],bool[],uint8,uint8)._churnedInCount](contracts/SwapContract.sol#L285) is not in mixedCase

contracts/SwapContract.sol#L285


 - [ ] ID-189
Parameter [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])._totalSwapped](contracts/SwapContract.sol#L124) is not in mixedCase

contracts/SwapContract.sol#L124


 - [ ] ID-190
Parameter [SwapContract.updateParams(address,address,uint256,uint256,uint256)._withdrawalFeeBPS](contracts/SwapContract.sol#L332) is not in mixedCase

contracts/SwapContract.sol#L332


 - [ ] ID-191
Parameter [SwapContract.getFloatReserve(address,address)._tokenA](contracts/SwapContract.sol#L373) is not in mixedCase

contracts/SwapContract.sol#L373


 - [ ] ID-192
Parameter [SwapContract.recordUTXOSweepMinerFee(uint256,bytes32)._minerFee](contracts/SwapContract.sol#L263) is not in mixedCase

contracts/SwapContract.sol#L263


 - [ ] ID-193
Parameter [SwapContract.updateParams(address,address,uint256,uint256,uint256)._sbBTCPool](contracts/SwapContract.sol#L330) is not in mixedCase

contracts/SwapContract.sol#L330


 - [ ] ID-194
Parameter [SwapContract.recordOutcomingFloat(address,bytes32,uint256,bytes32)._txid](contracts/SwapContract.sol#L237) is not in mixedCase

contracts/SwapContract.sol#L237


 - [ ] ID-195
Parameter [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])._redeemedFloatTxIds](contracts/SwapContract.sol#L126) is not in mixedCase

contracts/SwapContract.sol#L126


 - [ ] ID-196
Parameter [SwapContract.recordOutcomingFloat(address,bytes32,uint256,bytes32)._addressesAndAmountOfLPtoken](contracts/SwapContract.sol#L235) is not in mixedCase

contracts/SwapContract.sol#L235


 - [ ] ID-197
Variable [SwapContract.BTCT_ADDR](contracts/SwapContract.sol#L21) is not in mixedCase

contracts/SwapContract.sol#L21


 - [ ] ID-198
Parameter [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])._rewardsAmount](contracts/SwapContract.sol#L125) is not in mixedCase

contracts/SwapContract.sol#L125


 - [ ] ID-199
Parameter [SwapContract.recordUTXOSweepMinerFee(uint256,bytes32)._txid](contracts/SwapContract.sol#L264) is not in mixedCase

contracts/SwapContract.sol#L264


 - [ ] ID-200
Parameter [SwapContract.isNodeStake(address)._user](contracts/SwapContract.sol#L393) is not in mixedCase

contracts/SwapContract.sol#L393


 - [ ] ID-201
Parameter [SwapContract.recordOutcomingFloat(address,bytes32,uint256,bytes32)._token](contracts/SwapContract.sol#L234) is not in mixedCase

contracts/SwapContract.sol#L234


 - [ ] ID-202
Parameter [SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])._swapAmounts](contracts/SwapContract.sol#L195) is not in mixedCase

contracts/SwapContract.sol#L195


 - [ ] ID-203
Parameter [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])._destToken](contracts/SwapContract.sol#L121) is not in mixedCase

contracts/SwapContract.sol#L121


 - [ ] ID-204
Parameter [SwapContract.getFloatReserve(address,address)._tokenB](contracts/SwapContract.sol#L374) is not in mixedCase

contracts/SwapContract.sol#L374


 - [ ] ID-205
Parameter [SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])._spenders](contracts/SwapContract.sol#L194) is not in mixedCase

contracts/SwapContract.sol#L194


 - [ ] ID-206
Parameter [SwapContract.churn(address,address[],bool[],uint8,uint8)._newOwner](contracts/SwapContract.sol#L282) is not in mixedCase

contracts/SwapContract.sol#L282


 - [ ] ID-207
Parameter [SwapContract.multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[])._redeemedFloatTxIds](contracts/SwapContract.sol#L159) is not in mixedCase

contracts/SwapContract.sol#L159


 - [ ] ID-208
Parameter [SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])._rewardsAmount](contracts/SwapContract.sol#L193) is not in mixedCase

contracts/SwapContract.sol#L193


 - [ ] ID-209
Function [ISwapContract.BTCT_ADDR()](contracts/interfaces/ISwapContract.sol#L8) is not in mixedCase

contracts/interfaces/ISwapContract.sol#L8


 - [ ] ID-210
Parameter [SwapContract.isTxUsed(bytes32)._txid](contracts/SwapContract.sol#L346) is not in mixedCase

contracts/SwapContract.sol#L346


 - [ ] ID-211
Parameter [SwapContract.churn(address,address[],bool[],uint8,uint8)._tssThreshold](contracts/SwapContract.sol#L286) is not in mixedCase

contracts/SwapContract.sol#L286


 - [ ] ID-212
Parameter [SwapContract.multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[])._totalSwapped](contracts/SwapContract.sol#L157) is not in mixedCase

contracts/SwapContract.sol#L157


 - [ ] ID-213
Parameter [SwapContract.updateParams(address,address,uint256,uint256,uint256)._buybackRewardsRatio](contracts/SwapContract.sol#L334) is not in mixedCase

contracts/SwapContract.sol#L334


 - [ ] ID-214
Parameter [SwapContract.recordOutcomingFloat(address,bytes32,uint256,bytes32)._minerFee](contracts/SwapContract.sol#L236) is not in mixedCase

contracts/SwapContract.sol#L236


 - [ ] ID-215
Parameter [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])._amount](contracts/SwapContract.sol#L123) is not in mixedCase

contracts/SwapContract.sol#L123


 - [ ] ID-216
Parameter [SwapContract.updateParams(address,address,uint256,uint256,uint256)._buybackAddress](contracts/SwapContract.sol#L331) is not in mixedCase

contracts/SwapContract.sol#L331


 - [ ] ID-217
Parameter [SwapContract.recordIncomingFloat(address,bytes32,bytes32)._txid](contracts/SwapContract.sol#L219) is not in mixedCase

contracts/SwapContract.sol#L219


 - [ ] ID-218
Parameter [SwapContract.multiTransferERC20TightlyPacked(address,bytes32[],uint256,uint256,bytes32[])._addressesAndAmounts](contracts/SwapContract.sol#L156) is not in mixedCase

contracts/SwapContract.sol#L156


 - [ ] ID-219
Parameter [SwapContract.recordIncomingFloat(address,bytes32,bytes32)._token](contracts/SwapContract.sol#L217) is not in mixedCase

contracts/SwapContract.sol#L217


 - [ ] ID-220
Parameter [SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])._to](contracts/SwapContract.sol#L122) is not in mixedCase

contracts/SwapContract.sol#L122


 - [ ] ID-221
Parameter [SwapContract.recordIncomingFloat(address,bytes32,bytes32)._addressesAndAmountOfFloat](contracts/SwapContract.sol#L218) is not in mixedCase

contracts/SwapContract.sol#L218


 - [ ] ID-222
Parameter [SwapContract.updateParams(address,address,uint256,uint256,uint256)._nodeRewardsRatio](contracts/SwapContract.sol#L333) is not in mixedCase

contracts/SwapContract.sol#L333


 - [ ] ID-223
Parameter [SwapRewards.pullRewards(address,address,uint256)._receiver](contracts/SwapRewards.sol#L61) is not in mixedCase

contracts/SwapRewards.sol#L61


 - [ ] ID-224
Parameter [SwapRewards.pullRewardsMulti(address,address[],uint256[])._receiver](contracts/SwapRewards.sol#L88) is not in mixedCase

contracts/SwapRewards.sol#L88


 - [ ] ID-225
Parameter [SwapRewards.setSwap(address,uint256,uint256)._swap](contracts/SwapRewards.sol#L40) is not in mixedCase

contracts/SwapRewards.sol#L40


 - [ ] ID-226
Parameter [SwapRewards.pullRewardsMulti(address,address[],uint256[])._swapped](contracts/SwapRewards.sol#L89) is not in mixedCase

contracts/SwapRewards.sol#L89


 - [ ] ID-227
Parameter [SwapRewards.pullRewardsMulti(address,address[],uint256[])._dest](contracts/SwapRewards.sol#L87) is not in mixedCase

contracts/SwapRewards.sol#L87


 - [ ] ID-228
Function [ISwapContract.BTCT_ADDR()](contracts/interfaces/ISwapContract.sol#L8) is not in mixedCase

contracts/interfaces/ISwapContract.sol#L8


 - [ ] ID-229
Parameter [SwapRewards.pullRewards(address,address,uint256)._swapped](contracts/SwapRewards.sol#L62) is not in mixedCase

contracts/SwapRewards.sol#L62


 - [ ] ID-230
Parameter [SwapRewards.setSwap(address,uint256,uint256)._thresholdRatio](contracts/SwapRewards.sol#L42) is not in mixedCase

contracts/SwapRewards.sol#L42


 - [ ] ID-231
Parameter [SwapRewards.setSWINGBYPrice(uint256)._pricePerBTC](contracts/SwapRewards.sol#L34) is not in mixedCase

contracts/SwapRewards.sol#L34


 - [ ] ID-232
Parameter [SwapRewards.setSwap(address,uint256,uint256)._newRebateRate](contracts/SwapRewards.sol#L41) is not in mixedCase

contracts/SwapRewards.sol#L41


 - [ ] ID-233
Parameter [SwapRewards.pullRewards(address,address,uint256)._dest](contracts/SwapRewards.sol#L60) is not in mixedCase

contracts/SwapRewards.sol#L60


 - [ ] ID-234
Parameter [LPToken.setOld(address)._old](contracts/LPToken.sol#L13) is not in mixedCase

contracts/LPToken.sol#L13


## tautology
Impact: Medium
Confidence: High
 - [ ] ID-235
[Params.setNodeRewardsRatio(uint8)](contracts/Params.sol#L53-L59) contains a tautology or contradiction:
	- [require(bool,string)(_nodeRewardsRatio >= 0 && _nodeRewardsRatio <= 100,_nodeRewardsRatio is not valid)](contracts/Params.sol#L54-L57)

contracts/Params.sol#L53-L59


 - [ ] ID-236
[Params.setExpirationTime(uint256)](contracts/Params.sol#L44-L47) contains a tautology or contradiction:
	- [require(bool,string)(_expirationTime >= 0,_expirationTime can not be 0)](contracts/Params.sol#L45)

contracts/Params.sol#L44-L47


 - [ ] ID-237
[Params.setDepositFeesBPS(uint8)](contracts/Params.sol#L69-L75) contains a tautology or contradiction:
	- [require(bool,string)(_depositFeesBPS >= 0 && _depositFeesBPS <= 100,_depositFeesBPS is invalid)](contracts/Params.sol#L70-L73)

contracts/Params.sol#L69-L75


 - [ ] ID-238
[Params.setWithdrawalFeeBPS(uint8)](contracts/Params.sol#L61-L67) contains a tautology or contradiction:
	- [require(bool,string)(_withdrawalFeeBPS >= 0 && _withdrawalFeeBPS <= 100,_withdrawalFeeBPS is invalid)](contracts/Params.sol#L62-L65)

contracts/Params.sol#L61-L67


 - [ ] ID-239
[SwapRewards.setSwap(address,uint256,uint256)](contracts/SwapRewards.sol#L39-L56) contains a tautology or contradiction:
	- [require(bool,string)(_newRebateRate >= 0 && _newRebateRate <= 100,_newRebateRate is not valid)](contracts/SwapRewards.sol#L45-L48)

contracts/SwapRewards.sol#L39-L56


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-240
[Params.setParaswapAddress(address)._paraswapAddress](contracts/Params.sol#L49) lacks a zero-check on :
		- [paraswapAddress = _paraswapAddress](contracts/Params.sol#L50)

contracts/Params.sol#L49


 - [ ] ID-241
[SwapContract.constructor(address,address,address,address,address,uint256,uint256)._btct](contracts/SwapContract.sol#L79) lacks a zero-check on :
		- [BTCT_ADDR = _btct](contracts/SwapContract.sol#L95)

contracts/SwapContract.sol#L79


 - [ ] ID-242
[SwapContract.updateParams(address,address,uint256,uint256,uint256)._buybackAddress](contracts/SwapContract.sol#L331) lacks a zero-check on :
		- [buybackAddress = _buybackAddress](contracts/SwapContract.sol#L337)

contracts/SwapContract.sol#L331


 - [ ] ID-243
[SwapContract.constructor(address,address,address,address,address,uint256,uint256)._sbBTCPool](contracts/SwapContract.sol#L80) lacks a zero-check on :
		- [sbBTCPool = _sbBTCPool](contracts/SwapContract.sol#L87)

contracts/SwapContract.sol#L80


 - [ ] ID-244
[SwapContract.constructor(address,address,address,address,address,uint256,uint256)._buybackAddress](contracts/SwapContract.sol#L82) lacks a zero-check on :
		- [buybackAddress = _buybackAddress](contracts/SwapContract.sol#L104)

contracts/SwapContract.sol#L82


 - [ ] ID-245
[SwapContract.updateParams(address,address,uint256,uint256,uint256)._sbBTCPool](contracts/SwapContract.sol#L330) lacks a zero-check on :
		- [sbBTCPool = _sbBTCPool](contracts/SwapContract.sol#L336)

contracts/SwapContract.sol#L330


 - [ ] ID-246
[LPToken.setOld(address)._old](contracts/LPToken.sol#L13) lacks a zero-check on :
		- [old = _old](contracts/LPToken.sol#L14)

contracts/LPToken.sol#L13


## divide-before-multiply
Impact: Medium
Confidence: Medium
 - [ ] ID-247
[SwapContract._rewardsCollection(address,uint256)](contracts/SwapContract.sol#L554-L586) performs a multiplication on the result of a division:
	- [feesLPTTotal = feesTotal.mul(lpDivisor).div(nowPrice)](contracts/SwapContract.sol#L573)
	- [feesBuyback = feesLPTTotal.mul(buybackRewardsRatio).div(100)](contracts/SwapContract.sol#L575)

contracts/SwapContract.sol#L554-L586


 - [ ] ID-248
[SwapContract._burnLPTokensForFloat(address,bytes32,uint256,bytes32)](contracts/SwapContract.sol#L439-L497) performs a multiplication on the result of a division:
	- [amountOfFloat = amountOfLP.mul(nowPrice).div(lpDivisor)](contracts/SwapContract.sol#L453)
	- [withdrawalFees = amountOfFloat.mul(withdrawalFeeBPS).div(10000)](contracts/SwapContract.sol#L454)

contracts/SwapContract.sol#L439-L497


 - [ ] ID-249
[SwapContract._rewardsCollection(address,uint256)](contracts/SwapContract.sol#L554-L586) performs a multiplication on the result of a division:
	- [feesTotal = _rewardsAmount.mul(nodeRewardsRatio).div(100)](contracts/SwapContract.sol#L571)
	- [feesLPTTotal = feesTotal.mul(lpDivisor).div(nowPrice)](contracts/SwapContract.sol#L573)

contracts/SwapContract.sol#L554-L586


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-250
[SwapContract._rewardsCollection(address,uint256)](contracts/SwapContract.sol#L554-L586) ignores return value by [lpToken.mint(sbBTCPool,feesLPTTotal.sub(feesBuyback))](contracts/SwapContract.sol#L577)

contracts/SwapContract.sol#L554-L586


 - [ ] ID-251
[SwapContract._issueLPTokensForFloat(address,bytes32,bytes32)](contracts/SwapContract.sol#L404-L432) ignores return value by [lpToken.mint(to,amountOfLP)](contracts/SwapContract.sol#L418)

contracts/SwapContract.sol#L404-L432


 - [ ] ID-252
[SwapContract._rewardsCollection(address,uint256)](contracts/SwapContract.sol#L554-L586) ignores return value by [lpToken.mint(buybackAddress,feesBuyback)](contracts/SwapContract.sol#L578)

contracts/SwapContract.sol#L554-L586


 - [ ] ID-253
[SwapContract.singleTransferERC20(address,address,uint256,uint256,uint256,bytes32[])](contracts/SwapContract.sol#L120-L146) ignores return value by [sw.pullRewards(_destToken,_to,_totalSwapped)](contracts/SwapContract.sol#L135)

contracts/SwapContract.sol#L120-L146


 - [ ] ID-254
[SwapContract.collectSwapFeesForBTC(uint256,uint256,uint256,address[],uint256[])](contracts/SwapContract.sol#L190-L207) ignores return value by [sw.pullRewardsMulti(address(0),_spenders,_swapAmounts)](contracts/SwapContract.sol#L200)

contracts/SwapContract.sol#L190-L207


 - [ ] ID-255
[LPToken.convertTo(LPToken,uint256)](contracts/LPToken.sol#L19-L23) ignores return value by [target.mintFrom(_msgSender(),amount)](contracts/LPToken.sol#L22)

contracts/LPToken.sol#L19-L23


## events-maths
Impact: Low
Confidence: Medium
 - [ ] ID-256
[SwapContract.updateParams(address,address,uint256,uint256,uint256)](contracts/SwapContract.sol#L329-L342) should emit an event for: 
	- [withdrawalFeeBPS = _withdrawalFeeBPS](contracts/SwapContract.sol#L338) 
	- [nodeRewardsRatio = _nodeRewardsRatio](contracts/SwapContract.sol#L339) 
	- [buybackRewardsRatio = _buybackRewardsRatio](contracts/SwapContract.sol#L340) 

contracts/SwapContract.sol#L329-L342


## assembly
Impact: Informational
Confidence: High
 - [ ] ID-257
[Address.isContract(address)](contracts/interfaces/lib/Address.sol#L26-L36) uses assembly
	- [INLINE ASM](contracts/interfaces/lib/Address.sol#L32-L34)

contracts/interfaces/lib/Address.sol#L26-L36


 - [ ] ID-258
[Address.verifyCallResult(bool,bytes,string)](contracts/interfaces/lib/Address.sol#L195-L215) uses assembly
	- [INLINE ASM](contracts/interfaces/lib/Address.sol#L207-L210)

contracts/interfaces/lib/Address.sol#L195-L215


## costly-loop
Impact: Informational
Confidence: Medium
 - [ ] ID-259
[SwapContract.churn(address,address[],bool[],uint8,uint8)](contracts/SwapContract.sol#L281-L321) has costly operations inside a loop:
	- [activeNodeCount --](contracts/SwapContract.sol#L313)

contracts/SwapContract.sol#L281-L321


 - [ ] ID-260
[SwapContract.churn(address,address[],bool[],uint8,uint8)](contracts/SwapContract.sol#L281-L321) has costly operations inside a loop:
	- [activeNodeCount ++](contracts/SwapContract.sol#L309)

contracts/SwapContract.sol#L281-L321


## low-level-calls
Impact: Informational
Confidence: High
 - [ ] ID-261
Low level call in [Address.sendValue(address,uint256)](contracts/interfaces/lib/Address.sol#L54-L59):
	- [(success) = recipient.call{value: amount}()](contracts/interfaces/lib/Address.sol#L57)

contracts/interfaces/lib/Address.sol#L54-L59


 - [ ] ID-262
Low level call in [Address.functionCallWithValue(address,bytes,uint256,string)](contracts/interfaces/lib/Address.sol#L122-L133):
	- [(success,returndata) = target.call{value: value}(data)](contracts/interfaces/lib/Address.sol#L131)

contracts/interfaces/lib/Address.sol#L122-L133


 - [ ] ID-263
Low level call in [Address.functionStaticCall(address,bytes,string)](contracts/interfaces/lib/Address.sol#L151-L160):
	- [(success,returndata) = target.staticcall(data)](contracts/interfaces/lib/Address.sol#L158)

contracts/interfaces/lib/Address.sol#L151-L160


 - [ ] ID-264
Low level call in [Address.functionDelegateCall(address,bytes,string)](contracts/interfaces/lib/Address.sol#L178-L187):
	- [(success,returndata) = target.delegatecall(data)](contracts/interfaces/lib/Address.sol#L185)

contracts/interfaces/lib/Address.sol#L178-L187


## shadowing-local
Impact: Low
Confidence: High
 - [ ] ID-265
[SwapRewards.constructor(address,address,uint256)._owner](contracts/SwapRewards.sol#L21) shadows:
	- [Ownable._owner](node_modules/@openzeppelin/contracts/access/Ownable.sol#L21) (state variable)

contracts/SwapRewards.sol#L21


 - [ ] ID-266
[BurnableToken._approve(address,address,uint256).owner](contracts/BurnableToken.sol#L325) shadows:
	- [Ownable.owner()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L43-L45) (function)

contracts/BurnableToken.sol#L325


 - [ ] ID-267
[BurnableToken.allowance(address,address).owner](contracts/BurnableToken.sol#L107) shadows:
	- [Ownable.owner()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L43-L45) (function)

contracts/BurnableToken.sol#L107


 - [ ] ID-268
[BurnableToken._approve(address,address,uint256).owner](contracts/BurnableToken.sol#L325) shadows:
	- [Ownable.owner()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L43-L45) (function)

contracts/BurnableToken.sol#L325


 - [ ] ID-269
[BurnableToken.allowance(address,address).owner](contracts/BurnableToken.sol#L107) shadows:
	- [Ownable.owner()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L43-L45) (function)

contracts/BurnableToken.sol#L107


 - [ ] ID-270
[BurnableToken._approve(address,address,uint256).owner](contracts/BurnableToken.sol#L325) shadows:
	- [Ownable.owner()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L43-L45) (function)

contracts/BurnableToken.sol#L325


 - [ ] ID-271
[BurnableToken.allowance(address,address).owner](contracts/BurnableToken.sol#L107) shadows:
	- [Ownable.owner()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L43-L45) (function)

contracts/BurnableToken.sol#L107


contracts analyzed (49 contracts with 81 detectors), 272 result(s) found
