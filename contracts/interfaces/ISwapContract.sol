pragma solidity =0.7.0;

interface ISwapContract {
    function multiTransferERC20TightlyPacked(
        address _token,
        bytes32[] memory _addressesAndAmounts,
        uint8 _inputDecimals
    ) external returns (bool);

    function multiTransferERC20(
        address token,
        address[] memory _contributors,
        uint256[] memory _amounts
    ) external returns (bool);

    function distributeNodeRewards() external;
}
