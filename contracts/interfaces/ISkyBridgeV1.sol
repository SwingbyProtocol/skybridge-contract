// SPDX-License-Identifier: AGPL-3.0
pragma solidity ^0.8.4;

interface ISkyBridgeV1 {
    function setWhiteList(address _token, bool _add) external;

    function isWhiteListed(address _token) external view returns (bool);

    function addFloat(address _token, uint256 _amount) external;

    function removeFloat(address _token, uint256 _amount) external;

    function swap(
        address _sourceToken,
        address _destToken,
        uint256 _swapAmount
    ) external;

    function rewardsCollection(address _feesToken, uint256 _rewardsAmount)
        external;

    function addUsedTx(bytes32 _txid) external;

    function addUsedTxs(bytes32[] memory _txids) external;

    function safeTransfer(
        address _token,
        address _to,
        uint256 _amount
    ) external;

    function updateLimitBTCForSPFlow2() external;

    function isTxUsed(bytes32 _txid) external view returns (bool);

    function getCurrentPriceLP() external view returns (uint256 nowPrice);

    function getLpDivisor() external view returns (uint256);

    function priceCheck() external view returns (bool);

    function getFloatReserve(address _tokenA, address _tokenB)
        external
        view
        returns (uint256 reserveA, uint256 reserveB);

    function addTokenBalance(
        address _token,
        address _account,
        uint256 _amount
    ) external;

    function removeTokenBalance(
        address _token,
        address _account,
        uint256 _amount
    ) external;

    function addMultiTokenBalance(
        address _token,
        bytes32[] memory _addressesAndAmounts
    ) external;

    function checkTokenBalance(
        address _token,
        address _account,
        uint256 _amount
    ) external view returns (bool);

    function spRecordPendingTx(
        string calldata _destinationAddressForBTC,
        uint256 _btctAmount
    ) external;
}
