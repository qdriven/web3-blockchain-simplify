import { ethers } from "hardhat";

async function deployContract(contractName:string, args?:Record<string,any>) {
    const ContractForDeploy = await ethers.getContractFactory(contractName);
    //todo: handler args
    const contractDeployed = await ContractForDeploy.deploy();
  
    await contractDeployed.deployed();
    console.log("deployed address is ", contractDeployed.address);
    return contractDeployed
}

