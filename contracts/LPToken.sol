pragma solidity =0.7.5;

import "./BurnableToken.sol";

contract LPToken is BurnableToken {
    constructor() {
        _initialize("Swingby BTC", "sbBTC", 8, 0, true);
    }
}