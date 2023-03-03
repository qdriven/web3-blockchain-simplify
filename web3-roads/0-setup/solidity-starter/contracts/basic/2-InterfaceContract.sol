// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.7;


interface Regulator {
    function checkValue(uint amount) external returns (bool);
    function loan() external returns (bool);
}

contract Bank is Regulator {
    uint private value;
    address private owner;

    modifier ownerOnly {
        require(owner == msg.sender,"only contract owner to operate");
        _;
    }
    constructor(uint amount) {
        value = amount;
    }
    
    function deposit(uint amount) public ownerOnly {
        value += amount;
    }
    
    function withdraw(uint amount) public ownerOnly {
        if (checkValue(amount)) {
            value -= amount;
        }
    }
    
    function balance() public view returns (uint) {
        return value;
    }
    
    function checkValue(uint amount) public view returns (bool) {
        // Classic mistake in the tutorial value should be above the amount
        return value >= amount;
    }
    
    function loan() public view returns (bool) {
        return value > 0;
    }
}

contract UserBankAccount is Bank(10) {
    string private name;
    uint private age;
    
    function setName(string memory newName) public {
        name = newName;
    }
    
    function getName() public view returns (string memory) {
        return name;
    }
    
    function setAge(uint newAge) public {
        age = newAge;
    }
    
    function getAge() public view returns (uint) {
        return age;
    }
}