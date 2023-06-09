// SPDX-License-Identifier: SEE LICENSE IN LICENSE

pragma solidity ^0.8.7;
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
//GET Price Feeds
contract PriceFeedTask {
    AggregatorV3Interface internal priceFeed;
    address owner;


    modifier onlyOwner {
        require(msg.sender ==owner,"only owner can send Data Feed" );
        _;
    }

    function getLatestPrice() public view returns(int256){
        return 100;
    }

    function getPriceFeed() public view returns (AggregatorV3Interface) {
        return priceFeed;
    }
    
    function changePriceFeed(address priceFeedAddress) public onlyOwner{
        priceFeed = AggregatorV3Interface(priceFeedAddress);
    }
}