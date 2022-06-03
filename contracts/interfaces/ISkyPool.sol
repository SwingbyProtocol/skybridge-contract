// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <=0.8.9;

import "../libraries/Utils.sol";

interface ISkyPool {
  function BTCT_ADDR() external returns (address);

    function swap(
        address _sourceToken,
        address _destToken,
        uint256 _swapAmount
    ) external;

    function addAllFees(
        address _feesToken,
        uint256 _rewardsAmount
    ) external;

    function removeFloat(
        address _token,
        uint256 _amount
    ) external;

    function addFloat(
        address _token,
        uint256 _amount
    ) external;

    function balanceOf(address _token, address _account) external view returns (uint256);
    function getFloat(address _token) external view returns (uint256);
    function updateLimitBTCForSPFlow2() external;
    function getSwapCount() external view returns (uint256);
    function recordReduceTokenTx(address _token, address _to, uint256 _amount) external;
    function reduceLimitBTCForSPFlow2(uint256 _amount) external;
    function increaseSwapCount() external;

    function collectSwapFeesForBTC(
        uint256 _incomingAmount,
        uint256 _minerFee,
        uint256 _rewardsAmount,
        address[] memory _spenders,
        uint256[] memory _swapAmounts,
        bool    _isUpdatelimitBTCForSPFlow2
    ) external returns (bool);

    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfFloat,
        bytes32 _txid
    ) external returns (bool);

    function recordOutcomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfLPtoken,
        uint256 _minerFee,
        bytes32 _txid
    ) external returns (bool);

    function recordTokenTX(
        address _token,
        address _to,
        uint256 _totalSwapped
    ) external returns (bool);

    function multiTokenTX(
        address _token,
        bytes32[] memory _addressesAndAmounts
    ) external returns (bool);

    function spFlow1SimpleSwap(Utils.SimpleData calldata _data) external;

    function spFlow1Uniswap(
        bool _fork,
        address _factory,
        bytes32 _initCode,
        uint256 _amountIn,
        uint256 _amountOutMin,
        address[] calldata _path
    ) external returns (uint256 receivedAmount);

    function spFlow2Uniswap(
        string memory _destinationAddressForBTC,
        bool _fork,
        address _factory,
        bytes32 _initCode,
        uint256 _amountIn,
        uint256 _amountOutMin,
        address[] calldata _path
    ) external returns (uint256 receivedAmount);

    function spFlow2SimpleSwap(
        string memory _destinationAddressForBTC,
        Utils.SimpleData calldata _data
    ) external returns (uint256 receivedAmount);

    function spCleanUpOldTXs() external;

    function spDeposit(address _token, uint256 _amount) external payable;

    function redeemEther(uint256 _amount) external;

    function redeemERC20Token(address _token, uint256 _amount) external;

    function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid)
        external
        returns (bool);
    function getFloatReserve(address _tokenA, address _tokenB)
        external
        view
        returns (uint256 reserveA, uint256 reserveB);
}