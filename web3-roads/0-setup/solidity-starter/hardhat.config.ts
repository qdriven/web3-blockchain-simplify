import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import { config as dotEnvConfig } from "dotenv";
import { resolve } from "path";

// const dotEnvConfigPath:string = process.env.DOTENV_CONFIG_PATH || "./.env"
// dotEnvConfig({path: resolve(__dirname,dotEnvConfigPath)})

// const mnemonic:string | undefined = process.env.MNEMONIC
// if(!mnemonic){
//   throw new Error("please set up you mnemonic in a .env file")
// }

// const infuraApiKey: string | undefined = process.env.INFURA_API_KEY;
// if (!infuraApiKey) {
//   throw new Error("Please set your INFURA_API_KEY in a .env file");
// }

const config: HardhatUserConfig = {
  solidity: "0.8.17",
};

export default config;
