// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;
contract FunctionCases{
    uint256 public number = 5;
    // 函数类型
    // function (<parameter types>) {internal|external} [pure|view|payable] [returns (<return types>)]
    // 默认function
    function add() internal{
        number = number + 1;
    }

    function add_two(uint256 n1) external 
            returns(uint256 added_num){
         add();
         number = number+n1;
         return number;
    }

    function add_all() public returns(uint256 added_num) {
        add();
        return this.add_two(1);
    }

    function pure_func(uint256 temp) public pure returns(uint256 new_temp){
        return temp*2;
    }
    
    function minus() internal {
        number = number - 1;
    }
    // 合约内的函数可以调用内部函数
    function minusCall() external {
        minus();
    }

    // payable: 需要支付
    function minusPayable() external payable returns(uint256 balance) {
        minus();    
        balance = address(this).balance;
    }
}
