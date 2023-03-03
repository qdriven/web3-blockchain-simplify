// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.7;

contract DataTypes {
    
    bool public myBool = false;
    
    int8 public myInt = -128;
    uint8 public myUInt = 255;
    
    string public myString;
    uint8[] public myStringArr;

    bytes public myValue;
    bytes1 public myBytes1; 
    bytes32 public myBytes32;
    
//    fixed256x8 myFixed = 1; // 255.0
//    ufixed myFixed = 1;

    enum Action {ADD, REMOVE, UPDATE}
    
    Action public myAction = Action.ADD;
    
    address public myAddress;
    
    function assignAddress() public {
        myAddress = msg.sender;
        myAddress.balance;
        myAddress.transfer(10);
    }
    
    uint[] public myIntArr = [1,2,3];
    
    function arrFunc() public {
        myIntArr.push(1);
        myIntArr.length;
        myIntArr[0];
    }
    
    uint[10] public myFixedArr;
    
    struct Account {
        uint balance;
        uint dailyLimit;
    }
    
    Account public myAccount;
    
    function structFunc() public {
        myAccount.balance = 100;
    }
    
    mapping(address => Account) public _accounts;
    
    function () public payable  {
        _accounts[msg.sender].balance += msg.value;
    }
    
    function getBalance() public view returns (uint) {
        return _accounts[msg.sender].balance;
    }
}
