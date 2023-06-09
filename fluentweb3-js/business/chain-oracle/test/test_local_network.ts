import { ethers } from "hardhat"

describe("local network functionalities", function () {
    it("get network id",async function(){
        const network = await ethers.provider.getNetwork()
        console.log(network)
    })
    it("get block information",async function() {
        const blockInform = await ethers.provider.getBlockNumber();
        console.log(blockInform)
        const addresses = await ethers.getSigners();
        console.log(addresses[0])
        const provider = ethers.getDefaultProvider();
        console.log(provider);

        const addressBalanace = await ethers.provider.getBalance(addresses[0].address);
        console.log(addressBalanace.toString());
    })
})