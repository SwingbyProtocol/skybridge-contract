//SPDX-License-Identifier: Unlicense
pragma solidity >=0.6.0 <=0.8.9;
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import "./interfaces/IERC20.sol";
import "./interfaces/IParams.sol";
import "./libraries/SafeERC20.sol";
import "./libraries/Utils.sol";
contract SkyPool is Ownable, ReentrancyGuard {
  using SafeMath for uint256;
  using SafeERC20 for IERC20;

  address public immutable BTCT_ADDR;
  mapping(address => uint256) private floatAmountOf;
  
  /** Skypool */
  //skypools - token balance - call using tokens[token address][user address] to get uint256 balance - see function balanceOf
  mapping(address => mapping(address => uint256)) public tokens;
  //keep track of ether in tokens[][]
  address constant ETHER = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
  uint256 public limitBTCForSPFlow2;
  uint256 public swapCount;
  constructor(
      address _btct,
      address _swapRewards,
      uint256 _existingBTCFloat
  ) {
      // Set BTCT address
      BTCT_ADDR = _btct;
      floatAmountOf[address(0)] = _existingBTCFloat;
  }
  /**
    * Skypools part
    */
  /// @dev Record SkyPools TX - allocate tokens from float to user in tokens[][]
  /// @param _token The address of token
  /// @param _to The address of recipient.
  /// @param _totalSwapped The amount of swap amount.
  function recordTokenTX(
      address _token,
      address _to,
      uint256 _totalSwapped
  ) external onlyOwner returns (bool) {
      tokens[_token][_to] = tokens[_token][_to].add(_totalSwapped);
      return true;
  }

  function recordReduceTokenTx(address _token, address _to, uint256 _amount) external onlyOwner {
    tokens[_token][_to] = tokens[_token][_to].sub(_amount);
  }

  /// @dev multiRecordSkyPoolsTX - allocate tokens from float to user in tokens[][] in batches
  /// @param _token The address of token.
  /// @param _addressesAndAmounts The address of recipientand amount.
  function multiTokenTX(
      address _token,
      bytes32[] memory _addressesAndAmounts
  ) external onlyOwner returns (bool) {
      for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
          tokens[_token][
              address(uint160(uint256(_addressesAndAmounts[i])))
          ] = tokens[_token][
              address(uint160(uint256(_addressesAndAmounts[i])))
          ].add(uint256(uint96(bytes12(_addressesAndAmounts[i]))));
      }
      return true;
  }

  /// @dev getFloatReserve returns float reserves
  /// @param _tokenA The address of target tokenA.
  /// @param _tokenB The address of target tokenB.
  function getFloatReserve(address _tokenA, address _tokenB)
      external
      view
      returns (uint256 reserveA, uint256 reserveB)
  {
      (reserveA, reserveB) = (floatAmountOf[_tokenA], floatAmountOf[_tokenB]);
  }

  function getFloat (address _token) external view returns (uint256) {
      return floatAmountOf[_token];
  }

  /// @dev addFloat updates one side of the float.
  /// @param _token The address of target token.
  /// @param _amount The amount of float.
  function addFloat(address _token, uint256 _amount) external onlyOwner {
      floatAmountOf[_token] = floatAmountOf[_token].add(_amount);
  }

  /// @dev _removeFloat remove one side of the float - redone for skypools using tokens mapping
  /// @param _token The address of target token.
  /// @param _amount The amount of float.
  function removeFloat(address _token, uint256 _amount) external onlyOwner {
      floatAmountOf[_token] = floatAmountOf[_token].sub(
          _amount,
          "10" //"_removeFloat: float amount insufficient"
      );
  }

  /// @dev balanceOf - return user balance for given token and user for skypools
  /// @param _token The address of target token.
  /// @param _user The address of target user.
  function balanceOf(address _token, address _user)
      public
      view
      returns (uint256)
  {
      return tokens[_token][_user];
  }

  /// @dev _swap collects swap amount to change float.
  /// @param _sourceToken The address of source token
  /// @param _destToken The address of target token.
  /// @param _swapAmount The amount of swap.
  function swap(
      address _sourceToken,
      address _destToken,
      uint256 _swapAmount
  ) external onlyOwner {
      floatAmountOf[_destToken] = floatAmountOf[_destToken].sub(
          _swapAmount,
          "11" //"_swap: float amount insufficient"
      );
      floatAmountOf[_sourceToken] = floatAmountOf[_sourceToken].add(
          _swapAmount
      );
  }

  /// @dev add all fees into pool
  function addAllFees (address _feesToken, uint256 _rewardsAmount) external onlyOwner {
    floatAmountOf[_feesToken] = floatAmountOf[_feesToken].add(
        _rewardsAmount
    );
  }

  /// @dev _updateLimitBTCForSPFlow2 udpates limitBTCForSPFlow2
  function _updateLimitBTCForSPFlow2() internal {
    // Update limitBTCForSPFlow2 by adding BTC floats
    limitBTCForSPFlow2 = floatAmountOf[address(0)];
  }

  function updateLimitBTCForSPFlow2() external onlyOwner {
    limitBTCForSPFlow2 = floatAmountOf[address(0)];
  }

  /// @dev _reduceLimitBTCForSPFlow2 reduces limitBTCForSPFlow2 when new sp flow2 txs are coming.
  /// @param _amount The amount of BTCT, (use BTCT amount insatead of BTC amount for enough. always BTCT > BTC)
  function reduceLimitBTCForSPFlow2(uint256 _amount) external onlyOwner {
    if (limitBTCForSPFlow2 == 0) {
        // initialize when initial Flow2 tx has been called.
        _updateLimitBTCForSPFlow2();
    }
    limitBTCForSPFlow2 = limitBTCForSPFlow2.sub(
        _amount,
        "12" //"BTC float amount insufficient"
    );
  }

  function getSwapCount() external view returns (uint256) {
    return swapCount;
  }

  function increaseSwapCount() external onlyOwner {
      swapCount++;
  }
}