//SPDX-License-Identifier: Unlicense
pragma solidity >=0.6.0 <=0.8.9;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./interfaces/IERC20.sol";
import "./interfaces/IParams.sol";
import "./interfaces/ISkyPool.sol";
import "./libraries/SafeERC20.sol";
import "./LPToken.sol";

contract SkyBridge is Ownable {
  using SafeMath for uint256;
  using SafeERC20 for IERC20;

  ISkyPool public skyPool;
  address public lpToken;
  IParams public ip;

  address public immutable BTCT_ADDR;

  // node lists
  mapping(address => uint8) nodes;
  address[] nodeAddrs;
  uint8 public activeNodeCount;
  uint8 public churnedInCount;
  uint8 public tssThreshold;

  constructor (
    address _lpToken,
    address _btct,
    address _skyPool
  ) {
    skyPool = ISkyPool(skyPool);
    lpToken = _lpToken;
    BTCT_ADDR = _btct;
  }

  /**
    * user functions
  */
  function depositBTC(
    address _to,
    uint256 _amount,
    bytes32[] memory _redeemedFloatTxIds
  ) public onlyOwner {
    skyPool.singleTransferERC20(
      address(lpToken),
      msg.sender,
      _amount,
      0,
      0,
      _redeemedFloatTxIds
    );
  }
  
  function withdrawBTC(
    address _to,
    uint256 _amount,
    bytes32[] memory _redeemedFloatTxIds
  ) public onlyOwner {
    skyPool.singleTransferERC20(
      BTCT_ADDR,
      msg.sender,
      _amount,
      0,
      0,
      _redeemedFloatTxIds
    );
  }
  /// @dev churn transfers contract ownership and set variables of the next TSS validator set.
  /// @param _newOwner The address of new Owner.
  /// @param _nodes The reward addresses.
  /// @param _isRemoved The flags to remove node.
  /// @param _churnedInCount The number of next party size of TSS group.
  /// @param _tssThreshold The number of next threshold.
  function churn(
    address _newOwner,
    address[] memory _nodes,
    bool[] memory _isRemoved,
    uint8 _churnedInCount,
    uint8 _tssThreshold
  ) external onlyOwner returns (bool) {
    require(
        _tssThreshold >= tssThreshold && _tssThreshold <= 2**8 - 1,
        "01" //"_tssThreshold should be >= tssThreshold"
    );
    require(
        _churnedInCount >= _tssThreshold + uint8(1),
        "02" //"n should be >= t+1"
    );
    require(
        _nodes.length == _isRemoved.length,
        "05" //"_nodes and _isRemoved length is not match"
    );

    transferOwnership(_newOwner);
    // Update active node list
    for (uint256 i = 0; i < _nodes.length; i++) {
        if (!_isRemoved[i]) {
            if (nodes[_nodes[i]] == uint8(0)) {
                nodeAddrs.push(_nodes[i]);
            }
            if (nodes[_nodes[i]] != uint8(1)) {
                activeNodeCount++;
            }
            nodes[_nodes[i]] = uint8(1);
        } else {
            activeNodeCount--;
            nodes[_nodes[i]] = uint8(2);
        }
    }
    require(activeNodeCount <= 100, "Stored node size should be <= 100");
    churnedInCount = _churnedInCount;
    tssThreshold = _tssThreshold;
    return true;
  }

  /// @dev getActiveNodes returns active nodes list
  function getActiveNodes() public view returns (address[] memory) {
    uint256 count = 0;
    address[] memory _nodes = new address[](activeNodeCount);
    for (uint256 i = 0; i < nodeAddrs.length; i++) {
        if (nodes[nodeAddrs[i]] == uint8(1)) {
            _nodes[count] = nodeAddrs[i];
            count++;
        }
    }
    return _nodes;
  }

  /// @dev isNodeSake returns true if the node is churned in
  function isNodeStake(address _user) public view returns (bool) {
    if (nodes[_user] == uint8(1)) {
        return true;
    }
    return false;
  }

}