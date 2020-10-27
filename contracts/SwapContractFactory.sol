pragma solidity 0.7.0;

import "./SwapContract.sol";

contract SwapContractFactory {
    event Deployed(address wallet);

    function deployNewWallet(address _owner) public returns (address) {
        SwapContract wallet = new SwapContract();
        wallet.transferOwnership(_owner);
        emit Deployed(address(wallet));
        return address(wallet);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
