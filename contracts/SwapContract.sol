pragma solidity =0.7.0;

import "./interfaces/IERC20.sol";
import "./Ownable.sol";
import "./SafeMath.sol";
import "./interfaces/ISwapContract.sol";

contract SwapContract is ISwapContract, Ownable {
    using SafeMath for uint256;

    function multiTransferERC20TightlyPacked(
        address _token,
        bytes32[] memory _addressesAndAmounts,
        uint8 _inputDecimals
    ) public override onlyOwner returns (bool) {
        require(_token != address(0));
        for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
            IERC20 token = IERC20(_token);
            address to = address(uint160(uint256(_addressesAndAmounts[i])));
            uint8 boost = token.decimals() - _inputDecimals;
            require(boost >= 0, "boost should be >= 0");
            uint256 amount;
            if (boost == uint8(0)) {
                amount = uint256(uint96(bytes12(_addressesAndAmounts[i])));
            } else {
                amount = uint256(uint96(bytes12(_addressesAndAmounts[i]))).mul(
                    10**uint256(boost)
                );
            }
            require(token.transfer(to, amount));
        }
    }

    function multiTransferERC20(
        address token,
        address[] memory _contributors,
        uint256[] memory _amounts
    ) public override onlyOwner returns (bool) {
        require(_contributors.length == _amounts.length, "length is mismatch");
        for (uint256 i = 0; i < _contributors.length; i++) {
            require(IERC20(token).transfer(_contributors[i], _amounts[i]));
        }
    }

    function distributeNodeRewards() public override {}

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
