// SPDX-License-Identifier: AGPL-3.0
pragma solidity ^0.8.0;

import "./BurnableToken.sol";

contract LPToken is BurnableToken {
    address public old;

    constructor(uint8 _decimals) {
        _initialize("Swingby BTC LP Token", "sbBTC", _decimals, 0, true);
    }

    function setOld(address _old) public onlyOwner {
        old = _old;
    }

    /// @dev convertTo allows to convert new LPT when user burnd this tokens at the same time.
    /// note: new LPT has to allow minting permission from this contract.
    function convertTo(LPToken target, uint256 amount) public {
        require(address(target) != address(this), "target != this contract");
        _burn(_msgSender(), amount);
        target.mintFrom(_msgSender(), amount);
    }

    function mintFrom(address target, uint256 amount) public returns (bool) {
        require(old != address(this), "target != this contract");
        require(msg.sender == old, "caller is not old");
        _mint(target, amount);
        return true;
    }
}
