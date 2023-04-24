// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.19;

import "./BurnableToken.sol";

contract LPToken is BurnableToken { 
    constructor(uint8 _decimals) {
        _initialize("Swingby BTC LP Token", "sbBTC", _decimals, 0, true);
    }
    /// @dev convertTo allows to convert new LPT when user burnd this tokens at the same time. 
    /// note: new LPT has to allow minting permission from this contract.
    function convertTo(IBurnableToken target, uint256 amount) public {
        require(address(target) != address(this), "target != this contract");
        _burn(_msgSender(), amount);
        target.mint(_msgSender(), amount);
    }
}

