// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <0.8.0;

import "./BurnableToken.sol";

contract LPToken is BurnableToken {
    constructor() {
        _initialize("Swingby BTC LP Token", "sbBTC", 18, 0, true);
    }
}