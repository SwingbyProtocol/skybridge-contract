// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;

import "./BurnableToken.sol";

contract SwingbyToken is BurnableToken {
    constructor() {
        _initialize("SWINGBY token", "SWINGBY", 18, 0, true);
    }
}
