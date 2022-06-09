// SPDX-License-Identifier: AGPL-3.0
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

import "./interfaces/ISkyBridgeV1.sol";
import "./interfaces/ISwapRewards.sol";
import "./interfaces/IBurnableToken.sol";
import "./interfaces/IParams.sol";
import "./interfaces/IAugustusSwapper.sol";
import "./interfaces/IParaswap.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/IWETH.sol";
import "./libraries/Utils.sol";
import "./libraries/SafeERC20.sol";

contract SkyPool is Ownable, ReentrancyGuard {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;
    ISkyBridgeV1 public skyBridge;
    ISwapRewards public sw;
    address public BTCT_ADDR;
    IBurnableToken public lpToken;
    IParams public ip;
    address public paraswapAddress;
    //keep track of ether in tokens[][]
    address constant ETHER = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    address public immutable wETH;

    constructor(
        address _skyBridge,
        address _swapRewards,
        address _btct,
        address _lpToken,
        address _params,
        address _weth
    ) {
        skyBridge = ISkyBridgeV1(_skyBridge);
        sw = ISwapRewards(_swapRewards);
        BTCT_ADDR = _btct;
        lpToken = IBurnableToken(_lpToken);
        ip = IParams(_params);
        wETH = _weth;
    }

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
    ) external onlyOwner returns (bool) {
        require(skyBridge.isWhiteListed(_destToken), "14"); // //_destToken is not whitelisted
        require(
            _destToken != address(0),
            "15" //_destToken should not be address(0)
        );
        address _feesToken = address(0);
        if (_totalSwapped > 0) {
            sw.pullRewards(_destToken, _to, _totalSwapped);
            skyBridge.swap(address(0), BTCT_ADDR, _totalSwapped);
        } else {
            _feesToken = (_destToken == address(lpToken))
                ? address(lpToken)
                : BTCT_ADDR;
        }
        skyBridge.rewardsCollection(_feesToken, _rewardsAmount);
        skyBridge.addUsedTxs(_redeemedFloatTxIds);
        skyBridge.safeTransfer(_destToken, _to, _amount);
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
    ) external onlyOwner returns (bool) {
        require(
            skyBridge.isWhiteListed(_destToken),
            "_destTokne is not whiteListed"
        );
        require(
            _destToken != address(0),
            "_destToken should not be address(0)"
        );
        address _feesToken = address(0);
        if (_totalSwapped > 0) {
            skyBridge.swap(address(0), BTCT_ADDR, _totalSwapped);
        } else {
            _feesToken = (_destToken == address(lpToken))
                ? address(lpToken)
                : BTCT_ADDR;
        }
        skyBridge.rewardsCollection(_feesToken, _rewardsAmount);
        skyBridge.addUsedTxs(_redeemedFloatTxIds);
        for (uint256 i = 0; i < _addressesAndAmounts.length; i++) {
            skyBridge.safeTransfer(
                _destToken,
                address(uint160(uint256(_addressesAndAmounts[i]))),
                uint256(uint96(bytes12(_addressesAndAmounts[i])))
            );
        }
        return true;
    }

    /// @dev collectSwapFeesForBTC collects fees in the case of swap BTCT to BTC.
    /// @param _incomingAmount The spent amount. (BTCT)
    /// @param _rewardsAmount The fees that should be paid.
    function collectSwapFeesForBTC(
        uint256 _incomingAmount,
        uint256 _rewardsAmount,
        address[] memory _spenders,
        uint256[] memory _swapAmounts,
        bool _isUpdatelimitBTCForSPFlow2
    ) external onlyOwner returns (bool) {
        address _feesToken = BTCT_ADDR;
        if (_incomingAmount > 0) {
            uint256 swapAmount = _incomingAmount.sub(_rewardsAmount);
            sw.pullRewardsMulti(address(0), _spenders, _swapAmounts);
            skyBridge.swap(BTCT_ADDR, address(0), swapAmount);
        } else if (_incomingAmount == 0) {
            _feesToken = address(0);
        }
        skyBridge.rewardsCollection(_feesToken, _rewardsAmount);
        if (_isUpdatelimitBTCForSPFlow2) {
            skyBridge.updateLimitBTCForSPFlow2();
        }
        return true;
    }

    /// @dev recordIncomingFloat mints LP token.
    /// @param _token The address of target token.
    /// @param _addressesAndAmountOfFloat The address of recipient and amount.
    /// @param _txid The txids which is for recording.
    function recordIncomingFloat(
        address _token,
        bytes32 _addressesAndAmountOfFloat,
        bytes32 _txid
    ) external onlyOwner returns (bool) {
        require(skyBridge.priceCheck(), "Invalid LPT price");
        require(skyBridge.isWhiteListed(_token), "16"); //_token is invalid
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
    ) external onlyOwner returns (bool) {
        require(skyBridge.priceCheck(), "Invalid LPT price");
        require(skyBridge.isWhiteListed(_token), "16"); // _token is invalid
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

    /// @dev _issueLPTokensForFloat
    /// @param _token The address of target token.
    /// @param _transaction The recevier address and amount.
    /// @param _txid The txid which is for recording.
    function _issueLPTokensForFloat(
        address _token,
        bytes32 _transaction,
        bytes32 _txid
    ) internal returns (bool) {
        require(!skyBridge.isTxUsed(_txid), "06"); //"The txid is already used");
        require(_transaction != 0x0, "07"); //"The transaction is not valid");
        // Define target address which is recorded on the tx data (20 bytes)
        // Define amountOfFloat which is recorded top on tx data (12 bytes)
        (address to, uint256 amountOfFloat) = _splitToValues(_transaction);
        // Calculate the amount of LP token
        uint256 nowPrice = skyBridge.getCurrentPriceLP();
        uint256 amountOfLP = amountOfFloat.mul(skyBridge.getLpDivisor()).div(
            nowPrice
        );
        // Send LP tokens to LP
        lpToken.mint(to, amountOfLP);
        // Add float amount
        skyBridge.addFloat(_token, amountOfFloat);
        skyBridge.addUsedTx(_txid);
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
        require(!skyBridge.isTxUsed(_txid), "06"); //"The txid is already used");
        require(_transaction != 0x0, "07"); //"The transaction is not valid");
        // Define target address which is recorded on the tx data (20bytes)
        // Define amountLP which is recorded top on tx data (12bytes)
        (address to, uint256 amountOfLP) = _splitToValues(_transaction);
        // Calculate the amount of LP token
        uint256 nowPrice = skyBridge.getCurrentPriceLP();
        // Calculate the amountOfFloat
        uint256 amountOfFloat = amountOfLP.mul(nowPrice).div(
            skyBridge.getLpDivisor()
        );
        uint256 withdrawalFees = amountOfFloat.mul(ip.withdrawalFeeBPS()).div(
            10000
        );
        require(
            amountOfFloat.sub(withdrawalFees) >= _minerFee,
            "09" //"Error: amountOfFloat.sub(withdrawalFees) < _minerFee"
        );
        uint256 withdrawal = amountOfFloat.sub(withdrawalFees).sub(_minerFee);
        (uint256 reserveA, uint256 reserveB) = skyBridge.getFloatReserve(
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
        skyBridge.rewardsCollection(_token, withdrawalFees);
        // Remove float amount
        skyBridge.removeFloat(_token, amountOfFloat);
        // Add txid for recording.
        skyBridge.addUsedTx(_txid);
        // BTCT transfer if token address is BTCT_ADDR
        if (_token == BTCT_ADDR) {
            // _minerFee should be zero
            skyBridge.safeTransfer(_token, to, withdrawal);
        }
        // Burn LP tokens
        require(lpToken.burn(amountOfLP));
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

        skyBridge.swap(address(0), BTCT_ADDR, _totalSwapped);
        skyBridge.addTokenBalance(BTCT_ADDR, _to, _totalSwapped);
        skyBridge.rewardsCollection(address(0), _rewardsAmount);
        skyBridge.addUsedTxs(_usedTxIds);
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
        skyBridge.swap(address(0), BTCT_ADDR, _totalSwapped);

        skyBridge.rewardsCollection(address(0), _rewardsAmount);

        skyBridge.addUsedTxs(_usedTxIds);
        skyBridge.addMultiTokenBalance(BTCT_ADDR, _addressesAndAmounts);

        return true;
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
            skyBridge.checkTokenBalance(
                _data.fromToken,
                _data.beneficiary,
                _data.fromAmount
            ),
            "Balance insufficient"
        );

        require(_data.fromToken == BTCT_ADDR, "fromToken != BTCT_ADDR");

        skyBridge.removeTokenBalance(
            _data.fromToken,
            _data.beneficiary,
            _data.fromAmount
        );

        _doSimpleSwap(_data); //no received amount, tokens to go user's wallet
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
            skyBridge.checkTokenBalance(fromToken, msg.sender, _amountIn),
            "Balance insufficient"
        );

        require(fromToken == BTCT_ADDR, "fromToken != BTCT_ADDR");
        require(endToken != ETHER, "Use path wBTC -> wETH");

        uint256 preSwapBalance = IERC20(endToken).balanceOf(address(this));
        skyBridge.removeTokenBalance(fromToken, msg.sender, _amountIn);

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

        require(receivedAmount >= _amountOutMin, "Received < minimum");

        skyBridge.addTokenBalance(endToken, msg.sender, receivedAmount);

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
            skyBridge.checkTokenBalance(fromToken, msg.sender, _amountIn),
            "Balance insufficient"
        );
        require(fromToken != ETHER, "Use path wETH -> wBTC");
        require(endToken == BTCT_ADDR, "swap => BTCT");

        uint256 preSwapBalance = IERC20(endToken).balanceOf(address(this));

        skyBridge.removeTokenBalance(fromToken, msg.sender, _amountIn);

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

        require(receivedAmount >= _amountOutMin, "Received < minimum");
        require(
            receivedAmount >= ip.minimumSwapAmountForWBTC(),
            "receivedAmount < minimumSwapAmountForWBTC"
        );

        skyBridge.spRecordPendingTx(_destinationAddressForBTC, receivedAmount);

        return receivedAmount;
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
            skyBridge.checkTokenBalance(
                _data.fromToken,
                msg.sender,
                _data.fromAmount
            ),
            "Balance insufficient"
        );

        uint256 preSwapBalance = IERC20(_data.toToken).balanceOf(address(this));

        skyBridge.removeTokenBalance(
            _data.fromToken,
            msg.sender,
            _data.fromAmount
        );

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

        skyBridge.spRecordPendingTx(_destinationAddressForBTC, receivedAmount);

        return receivedAmount;
    }

    // user functions

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

            uint256 received = IERC20(_token).balanceOf(address(this)).sub(
                initBalance
            );

            skyBridge.addTokenBalance(_token, msg.sender, received);
        } else {
            require(msg.value > 0);
            //swap to wETH tokens - contract now holds wETH instead of ether
            IWETH(wETH).deposit{value: msg.value}();
            skyBridge.addTokenBalance(wETH, msg.sender, msg.value);
        }
    }

    /// @dev redeemEther for skypools - swap wETH for ether and send to user's wallet
    /// @param _amount amount to withdraw
    function redeemEther(uint256 _amount) external nonReentrant {
        require(skyBridge.checkTokenBalance(wETH, msg.sender, _amount));

        IWETH(wETH).withdraw(_amount);
        skyBridge.removeTokenBalance(wETH, msg.sender, _amount);

        address payable sender = payable(msg.sender);

        (
            bool success, /*bytes memory data*/

        ) = sender.call{value: _amount}("");

        require(success, "receiver rejected ETH transfer");
    }

    // fallback function
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
        require(
            skyBridge.checkTokenBalance(_token, msg.sender, _amount),
            "Insufficient Balance"
        );
        skyBridge.removeTokenBalance(_token, msg.sender, _amount);
    }

    /// @dev recordUTXOSweepMinerFee reduces float amount by collected miner fees.
    /// @param _minerFee The miner fees of BTC transaction.
    /// @param _txid The txid which is for recording.
    function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid)
        public
        onlyOwner
        returns (bool)
    {
        require(!skyBridge.isTxUsed(_txid), "The txid is already used");
        skyBridge.removeFloat(address(0), _minerFee);
        skyBridge.addUsedTx(_txid);
        return true;
    }
}
