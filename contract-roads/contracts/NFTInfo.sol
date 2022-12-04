// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Base64.sol";

contract COTNFT is ERC721Enumerable, Ownable {
    using Counters for Counters.Counter;

    struct NFTInfo{
        string  background;
        string  wing;
        string  body;
        string  hat;
        string  eye;
        string  sdg;
        string imageUrl;
        uint256 cot;
        bool neutralized;

    }

    Counters.Counter private _tokenIds;
    
    uint256 public totalSupply = 5000000;

    mapping(uint256 => NFTInfo) public tokenInfoMap;

    constructor() ERC721("DCCMCOTNFT", "COTNFT") {
        _tokenIds._value = 10000;
    }

    // for opensea
    function contractURI() public view returns (string memory) {
        // return "https://metadata-url.com/my-metadata";
        // Project name
        // location
        // registry
        // methodology

        abi.encodePacked(
                    "data:application/json;base64,",
                    Base64.encode(
                        bytes(
                            abi.encodePacked(
                                '{"name":"', "DCCMCOTNFT",'", "description":"DCCMCOTNFT", "attributes":"", "image":"',imageURI,'"}'
                            )
                        )
                    )
                );
    }


    function split(uint256 tokenId, NFTInfo[] memory array) public {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "split caller is not owner nor approved");
        require(array.length != 0, "array length zero");

        NFTInfo splitNft = tokenInfoMap[tokenId];
        require(splitNft.neutralized == false, "token has neutralized");

        require(splitNft.cot > 0, "cot amount error");

        uint256 cotTotal = 0;

        for (uint i = 0; i < array.length; i++) {
            cotTotal += array[i].cot;
        }

        require(cotTotal <= splitNft.cot, "cot amount error");

        for (uint i = 0; i < array.length; i++) {
            // mint
            uint256 newTokenId = _tokenIds.current();
            _safeMint(_msgSender(), newTokenId);

            // set cot
            uint256 info = array[i];
            tokenInfoMap[newTokenId] = info;

            // id increment
            _tokenIds.increment();
        }

        uint256 newCot = splitNft.cot - cotTotal;
        splitNft.cot = newCot;
    }

    function neutralized(uint256 tokenId, string memory imageUrl) public {

        require(_isApprovedOrOwner(_msgSender(), tokenId), "neutralized caller is not owner nor approved");
        NFTInfo nft = tokenInfoMap[tokenId];
        nft.neutralized = true;
        nft.imageUrl = imageUrl
    }

    function mint(address to, NFTInfo memory nftInfo) public onlyOwner {
        // check total supply
        uint256 newTokenId = _tokenIds.current();
        _safeMint(to, newTokenId);
        tokenInfoMap[newTokenId] = nftInfo;
        _tokenIds.increment();
    }

    function burn(uint256 tokenId) public {
        require(_isApprovedOrOwner(_msgSender(), tokenId), "burn caller is not owner nor approved");
        _burn(tokenId);
    }


    // -------------
    function getNFTInfo(uint256 tokenId) public view returns (NFTInfo memory nftInfo) {
        nftInfo = tokenInfoMap[tokenId];
    }

}
