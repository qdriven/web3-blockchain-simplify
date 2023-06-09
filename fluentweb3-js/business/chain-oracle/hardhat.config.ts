import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@chainlink/hardhat-chainlink";
import { networkSettings } from "./chain.config";
const commonCompilerSettings = {
  optimizer: {
    enabled: true,
    runs: 1_000,
  },
}
const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.4.11",
        settings: commonCompilerSettings,
      }, { version: "0.4.24" ,settings: commonCompilerSettings},
      {
        version: "0.5.16",settings: commonCompilerSettings,
      },
      { version: "0.6.0",settings: commonCompilerSettings, },
      { version: "0.7.6",settings: commonCompilerSettings, },
      { version: "0.8.10",settings: commonCompilerSettings, },
      { version: "0.8.18" ,settings: commonCompilerSettings,}
    ]
  },
  networks: networkSettings
};

export default config;
