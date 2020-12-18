pragma solidity =0.7.5;

import "./BurnableToken.sol";

contract LPToken is BurnableToken {
    constructor() {
        _initialize("Swingby BTC-LP Token", "sbBTC", 8, 0, true);
    }
}