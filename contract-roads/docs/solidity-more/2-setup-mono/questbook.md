# Solidity on openquest

Write this code and run the test cases. Let's get a quick grasp of the programming language and syntax.

pragma solidity ^0.8.3;

contract HelloWorld {
string public greet = "Hello World!";
}
pragma defines which version of the compiler to use to compile this program (aka smart contract). Solidity language is rapidly improving and using the correct compiler version is important to avoid code incompatibility.
Like a class, we use contract in solidity
Solidity is a typed language, string is a primitive data type.
When you declare a variable as public in solidity, a getter method with the same name as variable is automatically generated for you.

And that is all, you just wrote a basic Solidity contract!

## Introduction
Let's first create a state variable named num.

uint public num;
We saw earlier saw that each function call is a new transaction. Every transaction has a transaction fees. To write or update a state variable you need to send a transaction & pay the fees.

On the other hand, you can read state variables, for free, without any transaction fee.

## 
