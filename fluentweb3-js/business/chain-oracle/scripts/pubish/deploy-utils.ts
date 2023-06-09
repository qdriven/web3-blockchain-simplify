import { ethers } from "hardhat";

export async function deployContract(contractName:string, args?:Record<string,any>) {
    const ContractForDeploy = await ethers.getContractFactory(contractName);
    //todo: handler args
    // const linkAddress = "0x779877A7B0D9E8603169DdbD7836e478b4624789"
    // const ownerAddress ="0xBA996c5D1f2131AEca5C7B794874843f081a9721"
    const contractDeployed = await ContractForDeploy.deploy()
  
    await contractDeployed.deployed();
    console.log("deployed address is ", contractDeployed.address);
    return contractDeployed
}

// deployContract("Operator").then((event)=>{
//     console.log(event)
// }).catch((error)=>{
//     console.log(error)
// })