// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.4;

contract StateChanger {
    uint public num;
        // You need to send a transaction to write to a state variable.

    function set(uint x) public {
        num = x;
    }
    // You can read from a state variable without sending a transaction.

    function get() public view returns(uint) {
        return num;
    }
}
