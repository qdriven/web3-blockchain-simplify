
import { EmissionAppContract } from '../types/EmissionAppContract';
import {ethers} from "hardhat"
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { EmissionAppContract__factory } from '../types/factories/EmissionAppContract__factory';



export async function deployEmissionAppContract(): Promise<{ eac: EmissionAppContract }> {
    const signers: SignerWithAddress[] = await ethers.getSigners();
    const admin: SignerWithAddress = signers[0];
  
    const eacFactory: EmissionAppContract__factory = <EmissionAppContract__factory>await ethers.getContractFactory("EmissionAppContract");
    const eac: EmissionAppContract = <EmissionAppContract>await eacFactory.connect(admin).deploy();
    await eac.deployed();
  
    return { eac };
  }