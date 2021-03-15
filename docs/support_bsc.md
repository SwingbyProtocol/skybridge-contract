# Swap Contract upgraded for BSC network

## TL;DR
* The swap contract has storages for decimal `8` tokens.
* To support decimal `18` BTC pegged tokens, the contract is upgraded.
* The contract supports transfer for [BTCB](https://testnet.bscscan.com/token/0x6ce8da28e2f864420840cf74474eff5fd80e65b8)
* A new function `_safeTransfer()` handles ERC20 token transfer with `boosting` amount.
* LP token decimal is kept `8`.
* The float balance states are still decimal `8` basis.
* All transfer input params are still decimal `8` basis


### Contract interface changed

The new LP Token contract has a new consturctor `_decimals` which is for set LP token decimals.
```js
constructor(uint8 _decimals) {
    _initialize("Swingby BTC LP Token", "sbBTC", _decimals, 0, true);
}
```

### A new internal function `_safeTransfer()` is added for amount boosting 

A new function added for sending decimals `18` token, `convertScale` is set `10 ** (btct.decimals - lpToken.decimals)`.
thus, it is set `10 ** 10`in transfer `18` decimal tokens case.
```js
/// @dev _safeTransfer executes tranfer erc20 tokens
/// @param _token The address of target token
/// @param _to The address of receiver.
/// @param _amount The amount of transfer.
function _safeTransfer(
    address _token,
    address _to,
    uint256 _amount
) internal {
    if (_token == BTCT_ADDR) {
        _amount = _amount.mul(convertScale);
    }
    require(IERC20(_token).transfer(_to, _amount));
}
```


### TIPS: Swingby node has to check ERC20 transfer amount of the refund case.
The swap contract has to refund the amount of the received `BTCT` even decimal is 18.
the tx input params which is for 'transfer are decimal `8` basis, so a tx input should be converted `8` from `18`.
All transfer functions are just supported boosting the transfer of decimal `8` basis amount. thus, the small amount of received tx will be left in the swap contract. (that is not collected into as float balances, it will left forever)