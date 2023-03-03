import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/dist/src/signer-with-address";
import { ethers } from "hardhat";

import type { Signers } from "../../types";
import { shouldBehaveLikeHello } from "./hello.behavior";
import { deployHelloFixture } from "./hello.fixture";

describe("Unit tests", function () {
  before(async function () {
    this.signers = {} as Signers;

    const signers: SignerWithAddress[] = await ethers.getSigners();
    this.signers.admin = signers[0];

    this.loadFixture = loadFixture;
  });

  describe("Hello", function () {
    beforeEach(async function () {
      const { hello } = await this.loadFixture(deployHelloFixture);
      this.hello = hello;
    });

    shouldBehaveLikeHello();
  });
});
