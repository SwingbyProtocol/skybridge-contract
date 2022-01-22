// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;

import "./SwapContract.sol";
import "./LPToken.sol";

contract SwapContractFactory {
    event Deployed(address lpToken, address swapContract);

    function deployNewContracts(
        address _owner,
        address _wbtc,
        address _wETH,
        address _sbBTCPool,
        uint8   _decimals,
        uint256 _existingBTCFloat
    ) external returns (address) {
        LPToken lpToken = new LPToken(_decimals);
        SwapContract sc = new SwapContract(
            address(lpToken),
            _wbtc,
            _wETH,
            _sbBTCPool,
            _existingBTCFloat
        );
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
