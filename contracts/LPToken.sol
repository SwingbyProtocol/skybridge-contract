pragma solidity =0.7.0;

import "./BurnableToken.sol";

contract LPToken is BurnableToken {
    constructor(address _owner) {
        _initialize("BTC-LP token test", "BTC-LP", 18, 0, true, _owner);
    }
}