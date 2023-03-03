// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.7;

contract Transaction {
    
    event SenderLogger(address);
    event ValueLogger(uint);
    
    address private owner;
    
    modifier isOwner {
        require(owner == msg.sender,"only owner");
        _;
    }
    
    modifier validValue {
        assert(msg.value >= 1 ether );
        _;
    }
    
    constructor(){
        owner = msg.sender;
    }
    
    function emmitEvent() public payable isOwner validValue {
        emit SenderLogger(msg.sender);
        emit ValueLogger(msg.value);
    }
}
