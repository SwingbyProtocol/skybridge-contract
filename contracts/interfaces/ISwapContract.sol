// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <0.8.0;

interface ISwapContract {
    function singleTransferERC20(
        address _destToken,
        address _to,
        uint256 _amount,
        uint256 _totalSwapped,
        uint256 _rewardsAmount,
        bytes32[] memory _redeemedFloatTxIds
    ) external returns (bool);

    function multiTransferERC20TightlyPacked(
        address _destToken,
        bytes32[] memory _addressesAndAmounts,
        uint256 _totalSwapped,
        uint256 _rewardsAmount,
        bytes32[] memory _redeemedFloatTxIds
    ) external returns (bool);

    function collectSwapFeesForBTC(
        address _destToken,
        uint256 _incomingAmount,
        uint256 _rewardsAmount
    ) external returns (bool);

    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfFloat,
        bool _zerofee,
        bytes32 _txid
    ) external returns (bool);

    // function issueLPTokensForFloat(bytes32 _txid) external returns (bool);

    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfLPtoken,
        uint256 _minerFee,
        bytes32 _txid
    ) external returns (bool);

    // function burnLPTokensForFloat(bytes32 _txid) external returns (bool);

    function distributeNodeRewards() external returns (bool);

    function churn(
        address _newOwner,
        bytes32[] memory _rewardAddressAndAmounts,
        bool[] memory _isRemoved,
        uint8 _churnedInCount,
        uint8 _tssThreshold,
        uint8 _nodeRewardsRatio,
        uint8 _withdrawalFeeBPS
    ) external returns (bool);

    function isTxUsed(bytes32 _txid) external view returns (bool);

    function getCurrentPriceLP() external view returns (uint256);

    function getDepositFeeRate(address _token, uint256 _amountOfFloat)
        external
        view
        returns (uint256);

    function getFloatReserve(
        address _tokenA,
        address _tokenB,
        bool _mergeRewards
    ) external returns (uint256 reserveA, uint256 reserveB);

    function getActiveNodes() external returns (bytes32[] memory);

    function getMinimumAmountOfLPTokens(uint256 _minerFees)
        external
        view
        returns (uint256, uint256);
}
