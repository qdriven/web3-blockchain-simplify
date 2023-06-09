import { ethers } from "hardhat";

async function deploy_aggerator_contract(){
    const Aggerator = await ethers.getContractFactory("MockV3Aggregator");
    const aggerator = await Aggerator.deploy("18","100000000000000000000");
    await aggerator.deployed();
    console.log("Aggerator deployed to:", aggerator.address);     
}

async function deploy_linktoken() {
    
}

async function main() {
    await deploy_aggerator_contract()
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });