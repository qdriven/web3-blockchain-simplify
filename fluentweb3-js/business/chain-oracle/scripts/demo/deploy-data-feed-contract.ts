import { ethers } from "hardhat";

async function main() {
  const currentTimestampInSeconds = Math.round(Date.now() / 1000);
  // const unlockTime = currentTimestampInSeconds + 60;

  // const lockedAmount = ethers.utils.parseEther("0.001");
  // testnet contract: 0x0Ec1ab9DfEdcd861Ba54398eD93C91F20d37C009
  const PriceFeedContract = await ethers.getContractFactory("PriceFeedTask");
  const priceFeedTask = await PriceFeedContract.deploy();

  await priceFeedTask.deployed();

  console.log(
    `priceFeedTask deployed to:${priceFeedTask.address}`
  );
  const price = await priceFeedTask.getLatestPrice();
  console.log("the price is ", price.toString());
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
