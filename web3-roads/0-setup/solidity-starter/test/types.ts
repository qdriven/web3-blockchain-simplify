import { Hello } from "../typechain-types";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/dist/src/signer-with-address";

type Fixture<T> = () => Promise<T>;

declare module "mocha" {
  export interface Context {
    hello: Hello;
    loadFixture: <T>(fixture: Fixture<T>) => Promise<T>;
    signers: Signers;
  }
}

export interface Signers {
  admin: SignerWithAddress;
}