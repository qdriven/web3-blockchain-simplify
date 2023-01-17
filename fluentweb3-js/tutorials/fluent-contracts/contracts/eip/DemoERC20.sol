// contracts/MyNFT.sol
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract DemoERC20 is ERC20 {
    constructor() ERC20("DEMO-ERC20","DERC"){
        _mint(msg.sender,1000);
    }
}