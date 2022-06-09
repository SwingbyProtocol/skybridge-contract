//SPDX-License-Identifier: Unlicense
pragma solidity >=0.6.0 <=0.8.9;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

import "./interfaces/IERC20.sol";
import "./interfaces/IParams.sol";
import "./interfaces/ISwapRewards.sol";
import "./interfaces/IBurnableToken.sol";
import "./interfaces/ISkyPool.sol";
import "./interfaces/IWETH.sol";
import "./libraries/SafeERC20.sol";
import "./interfaces/IParaswap.sol";
import "./interfaces/IAugustusSwapper.sol";
import "./LPToken.sol";


contract SkyBridge is Ownable, ReentrancyGuard{
  using SafeMath for uint256;
  using SafeERC20 for IERC20;

  struct spPendingTx {
    bytes32 SwapID; //swap hash for identification of this swap.
    string DestAddr; //destination BTC address for the swap
    address RefundAddr; //refund address on evm source chain for if the swap fails.
    uint256 AmountWBTC; //outbound amount for this swap.
    uint256 Timestamp; // block timestamp that is set by EVM
  }

  ISkyPool public skyPool;
  IBurnableToken public lpToken;
  IParams public ip;
  ISwapRewards public sw;

  mapping(address => bool) public whitelist;
  address public BTCT_ADDR;
  address public sbBTCPool;
  address public wETH;
  uint256 private convertScale;
  uint256 private lpDivisor;

  mapping(bytes32 => bool) private used; //used TX

  // node lists
  mapping(address => uint8) nodes;
  address[] nodeAddrs;
  uint8 public activeNodeCount;
  uint8 public churnedInCount;
  uint8 public tssThreshold;

  //Data and indexes for pending swap objects
  mapping(uint256 => spPendingTx) public spPendingTXs; //index => pending TX object
  uint256 public oldestActiveIndex;
  address constant ETHER = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
  address public paraswapAddress;
  bool intialized;  

  modifier priceCheck() {
    uint256 beforePrice = getCurrentPriceLP();        
    _;
    require(getCurrentPriceLP() >= beforePrice, "Invalid LPT price");
  }

  function initialize (
    address _lpToken,
    address _btct,
    address _skyPool,
    address _weth,
    address _sbBTCPool,
    address _swapRewards,
    address _params
  ) public {
    require(!intialized, "already initialized");
    skyPool = ISkyPool(skyPool);
    // Set lpToken address
    lpToken = IBurnableToken(_lpToken);
    BTCT_ADDR = _btct;
    wETH = _weth;
    //set address for sbBTCpool
    sbBTCPool = _sbBTCPool;
    //set IParams
    ip = IParams(_params);
    //set paraswap address
    paraswapAddress = ip.paraswapAddress();
    //set ISwapRewards
    sw = ISwapRewards(_swapRewards);
    convertScale = 10**(IERC20(_btct).decimals() - 8);
    lpDivisor = 10**IERC20(_lpToken).decimals();
    whitelist[_btct] = true;
    whitelist[_lpToken] = true;
    whitelist[address(0)] = true;
    //init latest removed index and swapCount
    oldestActiveIndex = 0;
  }

  /**
     * Transfer part
     */
    /// @dev singleTransferERC20 sends tokens from contract.
    /// @param _destToken The address of target token.
    /// @param _to The address of recipient.
    /// @param _amount The amount of tokens.
    /// @param _totalSwapped The amount of swap.
    /// @param _rewardsAmount The fees that should be paid.
    /// @param _redeemedFloatTxIds The txids which is for recording.
    function singleTransferERC20(
      address _destToken,
      address _to,
      uint256 _amount,
      uint256 _totalSwapped,
      uint256 _rewardsAmount,
      bytes32[] memory _redeemedFloatTxIds
  ) external  onlyOwner returns (bool) {
      require(whitelist[_destToken], "14"); //_destToken is not whitelisted
      require(
          _destToken != address(0),
          "15" //_destToken should not be address(0)
      );
      address _feesToken = address(0);
      if (_totalSwapped > 0) {
          sw.pullRewards(_destToken, _to, _totalSwapped);
          skyPool.swap(address(0), BTCT_ADDR, _totalSwapped);
      } else {
          _feesToken = (_destToken == address(lpToken)) ? address(lpToken) : BTCT_ADDR;
      }
      _rewardsCollection(_feesToken, _rewardsAmount);
      _addUsedTxs(_redeemedFloatTxIds);
      _safeTransfer(_destToken, _to, _amount);
      return true;
  }

  /// @dev multiTransferERC20TightlyPacked sends tokens from contract.
  /// @param _destToken The address of target token.
  /// @param _addressesAndAmounts The address of recipient and amount.
  /// @param _totalSwapped The amount of swap.
  /// @param _rewardsAmount The fees that should be paid.
  /// @param _redeemedFloatTxIds The txids which is for recording.
  function multiTransferERC20TightlyPacked(
      address _destToken,
      bytes32[] memory _addressesAndAmounts,
      uint256 _totalSwapped,
      uint256 _rewardsAmount,
      bytes32[] memory _redeemedFloatTxIds
  ) external  onlyOwner returns (bool) {
      require(whitelist[_destToken], "_destToken is not whitelisted");
      require(
          _destToken != address(0),
          "_destToken should not be address(0)"
      );
      address _feesToken = address(0);
      if (_totalSwapped > 0) {
          skyPool.swap(address(0), BTCT_ADDR, _totalSwapped);
      } else {
          _feesToken = (_destToken == address(lpToken)) ? address(lpToken) : BTCT_ADDR;
      }
      _rewardsCollection(_feesToken, _rewardsAmount);
      _addUsedTxs(_redeemedFloatTxIds);
      for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
          _safeTransfer(
              _destToken,
              address(uint160(uint256(_addressesAndAmounts[i]))),
              uint256(uint96(bytes12(_addressesAndAmounts[i])))
          );
      }
      return true;
  }

  /// @dev collectSwapFeesForBTC collects fees in the case of swap BTCT to BTC.
  /// @param _incomingAmount The spent amount. (BTCT)
  /// @param _minerFee The miner fees of BTC transaction.
  /// @param _rewardsAmount The fees that should be paid.
  function collectSwapFeesForBTC(
      uint256 _incomingAmount,
      uint256 _minerFee,
      uint256 _rewardsAmount,
      address[] memory _spenders,
      uint256[] memory _swapAmounts,
      bool _isUpdatelimitBTCForSPFlow2
  ) external  onlyOwner returns (bool) {
      address _feesToken = BTCT_ADDR;
      if (_incomingAmount > 0) {
          uint256 swapAmount = _incomingAmount.sub(_rewardsAmount);
          sw.pullRewardsMulti(address(0), _spenders, _swapAmounts);
          skyPool.swap(BTCT_ADDR, address(0), swapAmount);
      } else if (_incomingAmount == 0) {
          _feesToken = address(0);
      }
      _rewardsCollection(_feesToken, _rewardsAmount);
      if (_isUpdatelimitBTCForSPFlow2) {
          skyPool.updateLimitBTCForSPFlow2();
      }
      return true;
  }

  /**
   * Float part
   */
  /// @dev recordIncomingFloat mints LP token.
  /// @param _token The address of target token.
  /// @param _addressesAndAmountOfFloat The address of recipient and amount.
  /// @param _txid The txids which is for recording.
  function recordIncomingFloat(
      address _token,
      bytes32 _addressesAndAmountOfFloat,
      bytes32 _txid
  ) external  onlyOwner priceCheck returns (bool) {
      require(whitelist[_token], "16"); //_token is invalid
      require(
          _issueLPTokensForFloat(_token, _addressesAndAmountOfFloat, _txid)
      );
      return true;
  }

  /// @dev recordOutcomingFloat burns LP token.
  /// @param _token The address of target token.
  /// @param _addressesAndAmountOfLPtoken The address of recipient and amount.
  /// @param _minerFee The miner fees of BTC transaction.
  /// @param _txid The txid which is for recording.
  function recordOutcomingFloat(
      address _token,
      bytes32 _addressesAndAmountOfLPtoken,
      uint256 _minerFee,
      bytes32 _txid
  ) external  onlyOwner priceCheck returns (bool) {
      require(whitelist[_token], "16"); //_token is invalid
      require(
          _burnLPTokensForFloat(
              _token,
              _addressesAndAmountOfLPtoken,
              _minerFee,
              _txid
          )
      );
      return true;
  }

  /**
   * Skypools part
   */
  /// @dev Record SkyPools TX - allocate tokens from float to user in tokens[][]
  /// @param _to The address of recipient.
  /// @param _totalSwapped The amount of swap amount.
  /// @param _rewardsAmount The fees that should be paid.
  /// @param _usedTxIds The txids which is for recording.
  function recordSkyPoolsTX(
      address _to,
      uint256 _totalSwapped,
      uint256 _rewardsAmount,
      bytes32[] memory _usedTxIds
  ) external onlyOwner returns (bool) {
      require(_totalSwapped != 0);
      require(_rewardsAmount != 0);

      skyPool.swap(address(0), BTCT_ADDR, _totalSwapped);
      skyPool.recordTokenTX(BTCT_ADDR, _to, _totalSwapped);
      _rewardsCollection(address(0), _rewardsAmount);
      _addUsedTxs(_usedTxIds);
      return true;
  }

  /// @dev multiRecordSkyPoolsTX - allocate tokens from float to user in tokens[][] in batches
  /// @param _addressesAndAmounts The address of recipientand amount.
  /// @param _totalSwapped The amount of swap amount.
  /// @param _rewardsAmount The fees that should be paid.
  /// @param _usedTxIds The txids which is for recording.
  function multiRecordSkyPoolsTX(
      bytes32[] memory _addressesAndAmounts,
      uint256 _totalSwapped,
      uint256 _rewardsAmount,
      bytes32[] memory _usedTxIds
  ) external onlyOwner returns (bool) {
      require(_totalSwapped != 0);
      require(_rewardsAmount != 0);

      skyPool.swap(address(0), BTCT_ADDR, _totalSwapped);

      _rewardsCollection(address(0), _rewardsAmount);

      _addUsedTxs(_usedTxIds);
      skyPool.multiTokenTX(BTCT_ADDR, _addressesAndAmounts);
      return true;
  }

  /// @dev _spCleanUpOldTXs - call when executing flow 2 swaps, cleans up expired TXs and moves the indices
  function spCleanUpOldTXs() public {
      uint256 max = oldestActiveIndex.add(ip.loopCount());
      uint256 swapCount = skyPool.getSwapCount();
      if (max >= swapCount) {
          max = swapCount;
      }

      uint256 current = block.timestamp;
      for (uint256 i = oldestActiveIndex; i < max; i++) {
          if (spPendingTXs[i].Timestamp.add(ip.expirationTime()) < current) {
              delete spPendingTXs[i];
              oldestActiveIndex = i.add(1);
          }
      }
  }

  
  /// @dev spDeposit - ERC-20 ONLY - users deposit ERC-20 tokens, balances to be stored in tokens[][]
  /// @param _token The address of the ERC-20 token contract.
  /// @param _amount amount to be deposited.
  function spDeposit(address _token, uint256 _amount)
      external
      payable
      nonReentrant
  {
      if (msg.value == 0) {
          require(_token != ETHER);
          require(_token != BTCT_ADDR);

          uint256 initBalance = IERC20(_token).balanceOf(address(this));

          IERC20(_token).safeTransferFrom(msg.sender, address(this), _amount);

          uint256 received = IERC20(_token).balanceOf(address(this)).sub(initBalance);

          skyPool.recordTokenTX(_token, msg.sender, received);
      } else {
          require(msg.value > 0);
          //swap to wETH tokens - contract now holds wETH instead of ether
          IWETH(wETH).deposit{value: msg.value}();
          skyPool.recordTokenTX(wETH, msg.sender, msg.value);
      }
  }

  /// @dev redeemEther for skypools - swap wETH for ether and send to user's wallet
  /// @param _amount amount to withdraw
  function redeemEther(uint256 _amount) external nonReentrant {
      require(skyPool.balanceOf(wETH, msg.sender) >= _amount);
      IWETH(wETH).withdraw(_amount);
      skyPool.removeFloat(wETH, _amount);
      address payable sender = payable(msg.sender);

      (
          bool success, /*bytes memory data*/

      ) = sender.call{value: _amount}("");

      require(success, "receiver rejected ETH transfer");
  }

  receive() external payable {
      assert(msg.sender == wETH); // only accept ETH via fallback from the WETH contract
  }

  /// @dev redeemERC20Token for skypools - redeem erc20 token
  /// @param _token The address of target token.
  /// @param _amount The amount to withdraw - call with BTC decimals (8) for BTC
  function redeemERC20Token(address _token, uint256 _amount)
      external
      nonReentrant
  {
      require(skyPool.balanceOf(_token, msg.sender) >= _amount, "Insufficient Balance");
      skyPool.removeFloat(_token, _amount);
      _safeTransfer(_token, msg.sender, _amount);
  }

  /**
   * Life cycle part
   */

  /// @dev recordUTXOSweepMinerFee reduces float amount by collected miner fees.
  /// @param _minerFee The miner fees of BTC transaction.
  /// @param _txid The txid which is for recording.
  function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid)
      public
      onlyOwner
      returns (bool)
  {
      require(!isTxUsed(_txid), "The txid is already used");
      skyPool.removeFloat(address(0), _minerFee);
      _addUsedTx(_txid);
      return true;
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


  /// @dev spGetPendingSwaps - returns array of pending swaps
  /// @return data - returns array of pending swap struct objects
  function spGetPendingSwaps()
      external
      view
      returns (spPendingTx[] memory data)
  {
      //require(swapCount != 0);

      uint256 index = 0;
      uint256 swapCount = skyPool.getSwapCount();
      data = new spPendingTx[](swapCount.sub(oldestActiveIndex));

      for (uint256 i = oldestActiveIndex.add(1); i <= swapCount; i++) {
          data[index] = spPendingTXs[index.add(oldestActiveIndex)];
          index = index.add(1);
      }

      return data;
  }

  /// @dev isTxUsed sends rewards for Nodes.
  /// @param _txid The txid which is for recording.
  function isTxUsed(bytes32 _txid) public view returns (bool) {
      return used[_txid];
  }

  /// @dev getCurrentPriceLP returns the current exchange rate of LP token.
  function getCurrentPriceLP()
      public
      view
      returns (uint256 nowPrice)
  {
      (uint256 reserveA, uint256 reserveB) = skyPool.getFloatReserve(
          address(0),
          BTCT_ADDR
      );
      uint256 totalLPs = lpToken.totalSupply();
      // decimals of totalReserved == 8, lpDivisor == 8, decimals of rate == 8
      nowPrice = totalLPs == 0
          ? lpDivisor
          : (reserveA.add(reserveB)).mul(lpDivisor).div(totalLPs);
      return nowPrice;
  }

  /// @dev _issueLPTokensForFloat
  /// @param _token The address of target token.
  /// @param _transaction The recevier address and amount.
  /// @param _txid The txid which is for recording.
  function _issueLPTokensForFloat(
      address _token,
      bytes32 _transaction,
      bytes32 _txid
  ) internal returns (bool) {
      require(!isTxUsed(_txid), "06"); //"The txid is already used");
      require(_transaction != 0x0, "07"); //"The transaction is not valid");
      // Define target address which is recorded on the tx data (20 bytes)
      // Define amountOfFloat which is recorded top on tx data (12 bytes)
      (address to, uint256 amountOfFloat) = _splitToValues(_transaction);
      // Calculate the amount of LP token
      uint256 nowPrice = getCurrentPriceLP();
      uint256 amountOfLP = amountOfFloat.mul(lpDivisor).div(nowPrice);
      // Send LP tokens to LP
      lpToken.mint(to, amountOfLP);
      // Add float amount
      skyPool.addFloat(_token, amountOfFloat);
      _addUsedTx(_txid);

      return true;
  }

  /// @dev _burnLPTokensForFloat
  /// @param _token The address of target token.
  /// @param _transaction The address of sender and amount.
  /// @param _minerFee The miner fees of BTC transaction.
  /// @param _txid The txid which is for recording.
  function _burnLPTokensForFloat(
      address _token,
      bytes32 _transaction,
      uint256 _minerFee,
      bytes32 _txid
  ) internal returns (bool) {
      require(!isTxUsed(_txid), "06"); //"The txid is already used");
      require(_transaction != 0x0, "07"); //"The transaction is not valid");
      // Define target address which is recorded on the tx data (20bytes)
      // Define amountLP which is recorded top on tx data (12bytes)
      (address to, uint256 amountOfLP) = _splitToValues(_transaction);
      // Calculate the amount of LP token
      uint256 nowPrice = getCurrentPriceLP();
      // Calculate the amountOfFloat
      uint256 amountOfFloat = amountOfLP.mul(nowPrice).div(lpDivisor);
      uint256 withdrawalFees = amountOfFloat.mul(ip.withdrawalFeeBPS()).div(
          10000
      );
      require(
          amountOfFloat.sub(withdrawalFees) >= _minerFee,
          "09" //"Error: amountOfFloat.sub(withdrawalFees) < _minerFee"
      );
      uint256 withdrawal = amountOfFloat.sub(withdrawalFees).sub(_minerFee);
      (uint256 reserveA, uint256 reserveB) = skyPool.getFloatReserve(
          address(0),
          BTCT_ADDR
      );
      if (_token == address(0)) {
          require(
              reserveA >= amountOfFloat.sub(withdrawalFees),
              "08" //"The float balance insufficient."
          );
      } else if (_token == BTCT_ADDR) {
          require(
              reserveB >= amountOfFloat.sub(withdrawalFees),
              "12" //"BTC float amount insufficient"
          );
      }
      // Collect fees before remove float
      _rewardsCollection(_token, withdrawalFees);
      // Remove float amount
      skyPool.removeFloat(_token, amountOfFloat);
      // Add txid for recording.
      _addUsedTx(_txid);
      // BTCT transfer if token address is BTCT_ADDR
      if (_token == BTCT_ADDR) {
          // _minerFee should be zero
          _safeTransfer(_token, to, withdrawal);
      }
      // Burn LP tokens
      require(lpToken.burn(amountOfLP));
      return true;
  }

  /// @dev _safeTransfer executes tranfer erc20 tokens
  /// @param _token The address of target token
  /// @param _to The address of receiver.
  /// @param _amount The amount of transfer.
  function _safeTransfer(
      address _token,
      address _to,
      uint256 _amount
  ) internal {
      if (_token == BTCT_ADDR) {
          _amount = _amount.mul(convertScale);
      }
      IERC20(_token).safeTransfer(_to, _amount);
  }

  /// @dev _rewardsCollection collects tx rewards.
  /// @param _feesToken The token address for collection fees.
  /// @param _rewardsAmount The amount of rewards.
  function _rewardsCollection(address _feesToken, uint256 _rewardsAmount)
      internal
  {
      if (_rewardsAmount == 0) return;
      
      // Get current LP token price.
      uint256 nowPrice = getCurrentPriceLP();
      // Add all fees into pool
      skyPool.addAllFees(_feesToken, _rewardsAmount);
      uint256 amountForNodes = _rewardsAmount.mul(ip.nodeRewardsRatio()).div(
          100
      );
      // Alloc LP tokens for nodes as fees
      uint256 amountLPTokensForNode = amountForNodes.mul(lpDivisor).div(
          nowPrice
      );
      // Mints LP tokens for Nodes
      lpToken.mint(sbBTCPool, amountLPTokensForNode);

  }

  /// @dev _addUsedTx updates txid list which is spent. (single hash)
  /// @param _txid The array of txid.
  function _addUsedTx(bytes32 _txid) internal {
      used[_txid] = true;
  }

  /// @dev _addUsedTxs updates txid list which is spent. (multiple hashes)
  /// @param _txids The array of txid.
  function _addUsedTxs(bytes32[] memory _txids) internal {
      for (uint256 i = 0; i < _txids.length; i++) {
          used[_txids[i]] = true;
      }
  }

  /// @dev _splitToValues returns address and amount of staked SWINGBYs
  /// @param _data The info of a staker.
  function _splitToValues(bytes32 _data)
      internal
      pure
      returns (address, uint256)
  {
      return (
          address(uint160(uint256(_data))),
          uint256(uint96(bytes12(_data)))
      );
  }
  
  /// @dev spFlow1SimpleSwap - FLOW 1 - execute paraswap TX using simpleSwap, ending tokens sent DIRECTLY to user's wallet
  /// @param _data A struct containing the data for simpleSwap, from the paraswap Utils lib.
  function spFlow1SimpleSwap(Utils.SimpleData calldata _data)
      external
      nonReentrant
  {
      require(_data.beneficiary == msg.sender, "beneficiary != msg.sender");

      require(
          skyPool.balanceOf(_data.fromToken, _data.beneficiary) >= _data.fromAmount,
          "Balance insufficient"
      );
      require(
          _data.fromToken == BTCT_ADDR,
          "fromToken != BTCT_ADDR"
      );

      skyPool.recordReduceTokenTx(_data.fromToken, _data.beneficiary, _data.fromAmount);

      _doSimpleSwap(_data); //no received amount, tokens to go user's wallet
  }

  /// @dev spParaSwapToken2BTC - FLOW 2 -> swap ERC20 -> wBTC
  /// @param _destinationAddressForBTC The BTC address to send BTC to.
  /// @param _data simpleData from paraswap API call, param for simpleSwap
  function spFlow2SimpleSwap(
    string calldata _destinationAddressForBTC,
    Utils.SimpleData calldata _data
  ) external nonReentrant returns (uint256 receivedAmount) {
    //bytes32 destBytes32 = _stringToBytes32(destinationAddressForBTC);
    //console.log("Converted to bytes32 and back to String:",_bytes32ToString(destBytes32));

    require(_data.fromToken != BTCT_ADDR, "Must not swap from BTC token");
    require(_data.toToken == BTCT_ADDR, "Must swap to BTC token");
    require(
        _data.beneficiary == address(this),
        "beneficiary != swap contract"
    );
    require(
        skyPool.balanceOf(_data.fromToken, msg.sender) >= _data.fromAmount,
        "Balance insufficient"
    );

    uint256 preSwapBalance = IERC20(_data.toToken).balanceOf(address(this));

    skyPool.recordReduceTokenTx(_data.fromToken, msg.sender, _data.fromAmount);
    _doSimpleSwap(_data);

    receivedAmount = IERC20(_data.toToken).balanceOf(address(this)).sub(
        preSwapBalance
    );

    require(
        receivedAmount >= _data.expectedAmount,
        "Received amount insufficient"
    );
    require(
        receivedAmount >= ip.minimumSwapAmountForWBTC(),
        "receivedAmount < minimumSwapAmountForWBTC"
    );

    _spRecordPendingTx(_destinationAddressForBTC, receivedAmount);

    return receivedAmount;
  }


  /// @dev spFlow1Uniswap - FLOW 1 - execute paraswap TX using uniswap, ending tokens sent to users allocation in tokens[][] mapping
  /// @param _fork - BOOL to determine if using swapOnUniswap or swapOnUniswapFork paraswap contract methods
  /// @param _factory - param for swapOnUniswapFork
  /// @param _initCode - param for swapOnUniswapFork
  /// @param _amountIn - param for swapOnUniswapFork or swapOnUniswap
  /// @param _amountOutMin - param for swapOnUniswapFork or swapOnUniswap
  /// @param _path - param for swapOnUniswapFork or swapOnUniswap
  function spFlow1Uniswap(
    bool _fork,
    address _factory,
    bytes32 _initCode,
    uint256 _amountIn,
    uint256 _amountOutMin,
    address[] calldata _path
  ) external nonReentrant returns (uint256 receivedAmount) {
    address fromToken = _path[0];
    address endToken = _path[_path.length - 1];

    require(
        skyPool.balanceOf(fromToken, msg.sender) >= _amountIn,
        "Balance insufficient"
    );
    require(fromToken == BTCT_ADDR, "fromToken != BTCT_ADDR");
    require(endToken != ETHER, "Use path wBTC -> wETH");

    uint256 preSwapBalance = IERC20(endToken).balanceOf(address(skyPool));

    skyPool.recordReduceTokenTx(fromToken, msg.sender, _amountIn);
    //do swap
    if (_fork) {
        _doUniswapFork(
            _factory,
            _initCode,
            _amountIn,
            _amountOutMin,
            _path
        );
    } else {
        _doUniswap(_amountIn, _amountOutMin, _path);
    }

    receivedAmount = IERC20(endToken).balanceOf(address(skyPool)).sub(
        preSwapBalance
    );

    require(
        receivedAmount >= _amountOutMin,
        "Received < minimum"
    );
    skyPool.recordTokenTX(endToken, msg.sender, receivedAmount);
    
    return receivedAmount;
  }

  /// @dev spFlow2Uniswap - FLOW 1 - execute paraswap TX using uniswap, ending tokens sent to users allocation in tokens[][] mapping
  /// @param _fork - BOOL to determine if using swapOnUniswap or swapOnUniswapFork paraswap contract methods
  /// @param _factory - param for swapOnUniswapFork
  /// @param _initCode - param for swapOnUniswapFork
  /// @param _amountIn - param for swapOnUniswapFork or swapOnUniswap
  /// @param _amountOutMin - param for swapOnUniswapFork or swapOnUniswap
  /// @param _path - param for swapOnUniswapFork or swapOnUniswap
  function spFlow2Uniswap(
    string calldata _destinationAddressForBTC,
    bool _fork,
    address _factory,
    bytes32 _initCode,
    uint256 _amountIn,
    uint256 _amountOutMin,
    address[] calldata _path
  ) external nonReentrant returns (uint256 receivedAmount) {
    address fromToken = _path[0];
    address endToken = _path[_path.length - 1];

    require(
        skyPool.balanceOf(fromToken, msg.sender) >= _amountIn,
        "Balance insufficient"
    );
    require(fromToken != ETHER, "Use path wETH -> wBTC");
    require(endToken == BTCT_ADDR, "swap => BTCT");

    uint256 preSwapBalance = IERC20(endToken).balanceOf(address(this));

    skyPool.recordReduceTokenTx(fromToken, msg.sender, _amountIn);
    //do swap
    if (_fork) {
        _doUniswapFork(
            _factory,
            _initCode,
            _amountIn,
            _amountOutMin,
            _path
        );
    } else {
        _doUniswap(_amountIn, _amountOutMin, _path);
    }

    receivedAmount = IERC20(endToken).balanceOf(address(this)).sub(
        preSwapBalance
    );

    require(
        receivedAmount >= _amountOutMin,
        "Received < minimum"
    );
    require(
        receivedAmount >= ip.minimumSwapAmountForWBTC(),
        "receivedAmount < minimumSwapAmountForWBTC"
    );

    _spRecordPendingTx(_destinationAddressForBTC, receivedAmount);

    return receivedAmount;
  }

  /// @dev _doUniswap - performs paraswap transaction - BALANCE & TOKEN CHECKS MUST OCCUR BEFORE CALLING THIS
  /// @param _amountIn - param for swapOnUniswap
  /// @param _amountOutMin - param for swapOnUniswap
  /// @param _path - param for swapOnUniswap
  function _doUniswap(
    uint256 _amountIn,
    uint256 _amountOutMin,
    address[] calldata _path
  ) internal {
    //address fromToken = _path[0];

    //address proxy = IAugustusSwapper(paraswapAddress).getTokenTransferProxy();

    IERC20(_path[0]).safeIncreaseAllowance(
        IAugustusSwapper(paraswapAddress).getTokenTransferProxy(), 
        _amountIn
    );

    IParaswap(paraswapAddress).swapOnUniswap(
        _amountIn,
        _amountOutMin,
        _path
    );
  }
  /// @dev _doUniswapFork - performs paraswap transaction - BALANCE & TOKEN CHECKS MUST OCCUR BEFORE CALLING THIS
  /// @param _factory - param for swapOnUniswapFork
  /// @param _initCode - param for swapOnUniswapFork
  /// @param _amountIn - param for swapOnUniswapFork
  /// @param _amountOutMin - param for swapOnUniswapFork
  /// @param _path - param for swapOnUniswapFork
  function _doUniswapFork(
    address _factory,
    bytes32 _initCode,
    uint256 _amountIn,
    uint256 _amountOutMin,
    address[] calldata _path
  ) internal {
    //address fromToken = _path[0];

    //address proxy = IAugustusSwapper(paraswapAddress).getTokenTransferProxy();

    IERC20(_path[0]).safeIncreaseAllowance(
        IAugustusSwapper(paraswapAddress).getTokenTransferProxy(),
        _amountIn
    );

    IParaswap(paraswapAddress).swapOnUniswapFork(
        _factory,
        _initCode,
        _amountIn,
        _amountOutMin,
        _path
    );
  }
  /// @dev _doSimpleSwap - performs paraswap transaction - BALANCE & TOKEN CHECKS MUST OCCUR BEFORE CALLING THIS
  /// @param _data data from API call that is ready to be sent to paraswap interface
  function _doSimpleSwap(Utils.SimpleData calldata _data) internal {
    //address proxy = IAugustusSwapper(paraswapAddress).getTokenTransferProxy();

    IERC20(_data.fromToken).safeIncreaseAllowance(
        IAugustusSwapper(paraswapAddress).getTokenTransferProxy(), 
        _data.fromAmount
    );

    IParaswap(paraswapAddress).simpleSwap(_data);
  }

  /// @dev _spRecordPendingTx - hash a unique swap ID, and add it to the array of pending TXs, and then emit event
  /// @param _destinationAddressForBTC The BTC address to send BTC to.
  /// @param _btctAmount amount in BTC decimal 8.
  function _spRecordPendingTx(
    string calldata _destinationAddressForBTC,
    uint256 _btctAmount
  ) internal {
    uint256 swapCount = skyPool.getSwapCount();
    //hash TX data for unique ID
    bytes32 ID = keccak256(
        abi.encodePacked(
            BTCT_ADDR, //specific to current chain
            swapCount,
            _destinationAddressForBTC,
            _btctAmount,
            block.timestamp
        )
    );

    spPendingTXs[swapCount] = spPendingTx(
        ID,
        _destinationAddressForBTC,
        msg.sender,
        _btctAmount,
        block.timestamp
    );

    //clean up expired TXs
    spCleanUpOldTXs();

    skyPool.increaseSwapCount();
    skyPool.reduceLimitBTCForSPFlow2(_btctAmount);
  }

}