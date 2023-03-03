import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { task } from "hardhat/config";
import type { TaskArguments } from "hardhat/types";

import type { Hello } from "../../typechain-types/Hello";
import type { Hello__factory } from "../../typechain-types/factories/Hello__factory";

task("deploy:Greeter")
  .addParam("greeting", "Say hello, be nice")
  .setAction(async function (taskArguments: TaskArguments, { ethers }) {
    const signers: SignerWithAddress[] = await ethers.getSigners();
    const greeterFactory: Hello__factory = <Hello__factory>(
      await ethers.getContractFactory("Greeter")
    );
    const greeter: Hello = <Hello>(
      await greeterFactory.connect(signers[0]).deploy(taskArguments.greeting)
    );
    await greeter.deployed();
    console.log("Hello deployed to: ", greeter.address);
  });
