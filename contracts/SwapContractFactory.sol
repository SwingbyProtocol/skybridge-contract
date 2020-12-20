pragma solidity =0.7.5;

import "./SwapContract.sol";
import "./LPToken.sol";

contract SwapContractFactory {
    event Deployed(address lpToken, address swapContract);

    function deployNewContracts(address _owner, address _wbtc)
        public
        returns (address)
    {
        LPToken lpToken = new LPToken();
        SwapContract sc = new SwapContract(address(lpToken), _wbtc);
        lpToken.transferOwnership(address(sc));
        sc.transferOwnership(_owner);
        emit Deployed(address(lpToken), address(sc));
        return address(sc);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
