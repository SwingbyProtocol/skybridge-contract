// SPDX-License-Identifier: AGPL-3.0
pragma solidity >=0.6.0 <0.8.0;
import "./interfaces/IERC20.sol";

library SkyBridge {
  /// @dev _safeTransfer executes tranfer erc20 tokens
  /// @param _token The address of target token
  /// @param _to The address of receiver.
  /// @param _amount The amount of transfer.
  function safeTransfer(
    address _token,
    address _to,
    uint256 _amount
  ) internal {
    address BTCT_ADDR = 0x820A8481451e893Bc66DCe50C84d45617CaC3705;
      if (_token == BTCT_ADDR) {
          _amount = _amount.mul(convertScale);
      }
      require(IERC20(_token).transfer(_to, _amount));
  }
}