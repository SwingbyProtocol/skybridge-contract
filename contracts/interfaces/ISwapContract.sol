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

    function mintLPToken(address _dist, uint256 _amount)
        external
        returns (bool);

    function addFloatForBTCToken(address _token, uint256 _amount)
        external
        returns (bool);

    function redeemFloatForBTCToken(address _token, uint256 _amountOfLPToken)
        external
        returns (bool);

    function distributeNodeRewards(address _token) external returns (bool);

    function churn(
        address _newOwner,
        bytes32[] memory _nodeRewardsAddressAndAmounts,
        uint8 _churnedInCount,
        uint8 _rewards_rate,
        uint8 _nodeRewardsRatio
    ) external returns (bool);
}
