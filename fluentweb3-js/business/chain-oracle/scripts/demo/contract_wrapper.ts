import { ethers } from "hardhat";

const contract_abis = {
    "PriceFeedTask":{
        "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "priceFeedAddress",
                  "type": "address"
                }
              ],
              "name": "changePriceFeed",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getLatestPrice",
              "outputs": [
                {
                  "internalType": "int256",
                  "name": "",
                  "type": "int256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getPriceFeed",
              "outputs": [
                {
                  "internalType": "contract AggregatorV3Interface",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ],
    }
}

async function loadContract() {
   const localProvider = new ethers.providers.JsonRpcProvider("http://127.0.0.1:8545");
   const priceTaskContract = new ethers.Contract("0x5FbDB2315678afecb367f032d93F642f64180aa3", 
            contract_abis.PriceFeedTask.abi, localProvider);
//    await ethers.getContractAt("PriceFeedTask", 
//    "0x5FbDB2315678afecb367f032d93F642f64180aa3");
   console.log(priceTaskContract.address);
   const latestPrice = await priceTaskContract.getLatestPrice()
   console.log(latestPrice.toString())
}

loadContract().catch((error)=>{
    console.error(error);    
})

