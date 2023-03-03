// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.7;
import "./ERC20Interface.sol";

contract MyFirstToken is ERC20 {
    string public constant symbol = "MFT";
    string public constant name = "My First Token";
    uint8 public constant decimals = 18;

    uint256 private constant __totalSupply = 1000;
    mapping(address => uint256) private __balanceOf;
    mapping(address => mapping(address => uint256)) private __allowances;

    constructor() {
        __balanceOf[msg.sender] = __totalSupply;
    }

    function totalSupply() public pure returns (uint256 _totalSupply) {
        _totalSupply = __totalSupply;
    }

    function balanceOf(address _addr) public view returns (uint256 balance) {
        return __balanceOf[_addr];
    }

    function transfer(address _to, uint256 _value)
        public
        returns (bool success)
    {
        if (_value > 0 && _value <= balanceOf(msg.sender)) {
            __balanceOf[msg.sender] -= _value;
            __balanceOf[_to] += _value;
            return true;
        }
        return false;
    }

    function transferFrom(
        address _from,
        address _to,
        uint256 _value
    ) public returns (bool success) {
        if (
            __allowances[_from][msg.sender] > 0 &&
            _value > 0 &&
            __allowances[_from][msg.sender] >= _value &&
            __balanceOf[_from] >= _value
        ) {
            __balanceOf[_from] -= _value;
            __balanceOf[_to] += _value;
            // Missed from the video
            __allowances[_from][msg.sender] -= _value;
            return true;
        }
        return false;
    }

    function approve(address _spender, uint256 _value)
        public
        returns (bool success)
    {
        __allowances[msg.sender][_spender] = _value;
        return true;
    }

    function allowance(address _owner, address _spender)
        public view
        returns (uint256 remaining)
    {
        return __allowances[_owner][_spender];
    }
}
