import { ethers } from "hardhat";

import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { EmissionAppContract__factory } from "../types/factories/EmissionAppContract__factory";
import { EmissionAppContract } from "../types/EmissionAppContract";


describe("CDPAllInOneContract",function(){
    before(async function(){ //load all fixtures
        const signers: SignerWithAddress[] = await ethers.getSigners();
        const admin: SignerWithAddress = signers[0];
        this.singers = signers
        this.admin = admin
        const eacFactory: EmissionAppContract__factory = <EmissionAppContract__factory>await ethers.getContractFactory("EmissionAppContract");
        const eac: EmissionAppContract = <EmissionAppContract>await eacFactory.connect(admin).deploy();
        await eac.deployed();
        this.eac = eac
    });

    it("start to test EmissionAppContract",async function(){
            console.log("this is admin address:",this.admin.address)
            console.log("this is contract address:",this.eac.address)
    });
    it("add data to eac contract", async function(){
        const emisiionData = "{\"k1\":\"v\"}";
        this.eac.connect("0xB581C9264f59BF0289fA76D61B2D0746dCE3C30D")
        const result =this.eac.addEmissionData(emisiionData,)
        console.log(result)
        const retrievedData = this.eac.getEmissionData()
        console.log(retrievedData)
    })
   
})