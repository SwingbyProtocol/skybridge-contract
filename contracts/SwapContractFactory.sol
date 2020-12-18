pragma solidity =0.7.5;

import "./SwapContract.sol";

contract SwapContractFactory {
    event Deployed(address wallet);

    function deployNewWallet(
        address _owner,
        address _lpToken,
        address _wbtc
    ) public returns (address) {
        SwapContract sc = new SwapContract(_lpToken, _wbtc);
        sc.transferOwnership(_owner);
        emit Deployed(address(sc));
        return address(sc);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
