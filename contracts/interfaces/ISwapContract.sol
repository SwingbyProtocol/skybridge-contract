pragma solidity =0.7.0;

interface ISwapContract {
    function singleTransferERC20(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _rewardsAmount
    ) external returns (bool);

    function multiTransferERC20TightlyPacked(
        address _token,
        bytes32[] memory _addressesAndAmounts,
        uint8 _inputDecimals,
        uint256 _rewardsAmount
    ) external returns (bool);

    function multiTransferERC20(
        address token,
        address[] memory _contributors,
        uint256[] memory _amounts,
        uint256 _rewardsAmount
    ) external returns (bool);

    function collectSwapFeesForBTC(
        address _feeToken,
        uint256 _rewardsAmount,
        bytes32 _txid
    ) external returns (bool);

    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfBTC,
        bytes32 _txid
    ) external returns (bool);

    function issueLPTokensForFloat(bytes32 _txid) external returns (bool);

    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfLPtoken,
        bytes32 _txid
    ) external returns (bool);

    function burnLPTokensForFloat(bytes32 _txid) external returns (bool);

    function distributeNodeRewards() external returns (bool);

    function churn(
        address _newOwner,
        bytes32[] memory _nodeRewardsAddressAndAmounts,
        uint8 _churnedInCount,
        uint8 _nodeRewardsRatio
    ) external returns (bool);

    function isTxUsed(bytes32 _txid) external view returns (bool);

    function getCurrentPriceLP() external view returns (uint256);
}
