import { ethers } from "ethers";
import { expect } from 'chai';

describe("Balances", () => {
    it("should have a balance", async () => {
        const provider = new ethers.JsonRpcProvider("http://localhost:8545");
        const balance = await provider.getBalance(`0x${process.env.WALLET_ADDRESS}`);
        expect(balance).to.be.gt(0);
    })
})