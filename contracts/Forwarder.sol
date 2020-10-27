pragma solidity =0.7.0;

import "./interfaces/IERC20.sol";

contract Forwarder {
    address payable private parent;

    constructor(address payable _parent) public {
        parent = _parent;
    }

    function forward(address _token) public payable returns (uint256 amount) {
        if (_token == address(0)) {
            amount = address(this).balance;
            parent.transfer(amount);
        } else {
            amount = IERC20(_token).balanceOf(address(this));
            IERC20(_token).transfer(parent, amount);
        }
    }
}
