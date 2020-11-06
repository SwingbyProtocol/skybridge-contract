pragma solidity =0.7.0;

import "./BurnableToken.sol";

contract LPToken is BurnableToken {
    constructor() {
        _initialize("BTC-LP token test", "BTC-LP test", 18, 0, true);
    }
}