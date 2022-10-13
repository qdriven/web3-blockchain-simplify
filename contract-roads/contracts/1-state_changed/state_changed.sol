// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

contract StateChanger {
    uint public num;
    function set(uint x) public {
        num = x;
    }

    function get() public view returns(uint) {
        return num;
    }
}
