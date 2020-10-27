pragma solidity =0.7.0;

import "./interfaces/IERC20.sol";
import "./BurnableToken.sol";
import "./SafeMath.sol";
import "./Burner.sol";
import "./Forwarder.sol";

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

    function deployNewForwarder(address payable _receiver)
        public
        onlyOwner
        returns (address)
    {
        Forwarder f = new Forwarder(_receiver);
        return address(f);
    }

    function forward(address _forwarder, address _token) public {
        Forwarder f = Forwarder(_forwarder);
        uint256 amount = f.forward(_token);
        // forwarder doesn't know about sender....
        //_approveDeposit();
    }

    function changeSwap(address _swap) public onlyOwner {
        swap = _swap;
    }

    function mintLPToken(address _dist, uint256 _amount) public onlyOwner {
        _mint(_dist, _amount);
    }

    function _approveDeposit(
        address _token,
        address _user,
        uint256 _amount,
        address _to
    ) internal {
        // Accepts WBTC/tBTC/renBTC tokens and send out to swap contract
        IERC20(_token).transfer(_to, _amount);
        // Send back BTC-LP tokens to user.
        _mint(_user, _amount);
    }

    // The contract doesn't allow receiving Ether.
    fallback() external {
        revert();
    }
}
