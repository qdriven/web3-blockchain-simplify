import { task } from "hardhat/config";
import { Signer } from "@ethersproject/abstract-signer";


task("accounts","print the list of accounts", async (_taskArgs,hre)=>{
    const accounts: Signer[] = await hre.ethers.getSigners();
    for (const account of accounts) {
        console.log(await account.getAddress());
      }
});