pragma solidity =0.7.0;

import "./interfaces/IERC20.sol";
import "./BurnableToken.sol";
import "./Ownable.sol";
import "./SafeMath.sol";
import "./Burner.sol";

// The TSS address should have ownership of this contract.
contract FloatMinter is BurnableToken {
    using SafeMath for uint256;
    address public swap;
    Burner public burner;
    mapping(address => mapping(address => uint256)) private balances;

    event ApproveDeposit(
        address indexed token,
        address indexed _user,
        uint256 amount
    );

    function initialize(address _swap) public {
        require(swap == address(0));
        _initialize("BTC LP Token", "BTC-LP", 18, 0, true);
        swap = _swap;
        burner = new Burner(address(this));
    }

    function changeSwap(address _swap) public onlyOwner {
        swap = _swap;
    }

    function mintLPToken(address _dist, uint256 _amount) public onlyOwner {
        _mint(_dist, _amount);
    }

    function approveDeposit(
        address _token,
        address _user,
        uint256 _amount
    ) public onlyOwner {
        _approveDeposit(_token, _user, _amount);
        // Accepts WBTC/tBTC/renBTC tokens and send out to swap contract
        IERC20(_token).transfer(swap, _amount);
        // Send back BTC-LP tokens to user.
        _mint(_user, _amount);
    }

    function _approveDeposit(
        address _token,
        address _user,
        uint256 _amount
    ) internal {
        balances[_token][_user] = balances[_token][_user].add(_amount);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
