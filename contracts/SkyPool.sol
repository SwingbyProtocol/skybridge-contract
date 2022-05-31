// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <0.8.0;

library SkyPool {

  event Swap(address from, address to, uint256 amount);

  /// @dev _swap collects swap amount to change float.
  /// @param _sourceFloat The address of source token
  /// @param _destFloat The address of target token.
  /// @param _swapAmount The amount of swap.
  function swap(
    address _sourceFloat,
    address _destFloat,
    uint256 _swapAmount
  ) internal returns (uint256 sourceFloat, uint256 destFloat) {
      destFloat = _destFloat.sub(
          _swapAmount,
          "_swap: float amount insufficient"
      );
      sourceFloat = _sourceFloat.add(
          _swapAmount
      );
      emit Swap(_sourceToken, _destToken, _swapAmount);
  }
}