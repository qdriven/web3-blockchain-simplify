import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/dist/src/signer-with-address";
import { ethers } from "hardhat";

import type { Hello } from "../../../typechain-types/Hello";
import type { Hello__factory } from "../../../typechain-types/factories/Hello__factory";

export async function deployHelloFixture(): Promise<{ hello: Hello }> {
  const signers: SignerWithAddress[] = await ethers.getSigners();
  const admin: SignerWithAddress = signers[0];

  const greeting: string = "Hello, world!";
  const greeterFactory: Hello__factory = <Hello__factory>(
    await ethers.getContractFactory("Hello")
  );
  const hello: Hello = <Hello>(
    await greeterFactory.connect(admin).deploy(greeting)
  );
  await hello.deployed();

  return { hello };
}
