// contracts/MyNFT.sol
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract DemoNFT is ERC721 {
    constructor() ERC721("DemoNFT","DemoNFT"){
        // console.log("create DEMO NFT")
    }

}