// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.4;

contract StateChanger {
    uint public num;
    function set(uint x) public {
        num = x;
    }

    function get() public view returns(uint) {
        return num;
    }
}
