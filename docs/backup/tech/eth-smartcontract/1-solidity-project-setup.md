# Solidity Project Setup


## Setup Project from truffle box

```sh
truffle unbox metacoin
```

run:

- truffle compile
- truffle migrate
- truffle test

If everything is OK, the setup is done.

## Smart Contract Project Intro

```sh
├── LICENSE
├── build 
│   └── contracts
│       ├── ConvertLib.json
│       ├── MetaCoin.json
│       └── Migrations.json
├── contracts
│   ├── ConvertLib.sol
│   ├── MetaCoin.sol
│   └── Migrations.sol
├── migrations
│   ├── 1_initial_migration.js
│   └── 2_deploy_contracts.js
├── test
│   ├── TestMetaCoin.sol
│   └── metacoin.js
└── truffle-config.js
```

- build: create by build tools
- contracts: original contract source code
- migrations: migrate script/deployment scripts
- test: unit test script for smart contract

## More Setup 

- truffle more commands

```sh
truffle develop
truffle console
truffle migrate
```

- Ganache: Local Ethereum Network

Launch Ganache:

![img](https://www.trufflesuite.com/img/docs/ganache/quickstart/accounts.png)

- hardhat installation

```sh
npm install --save-dev hardhat
```

- run hardhat tasks
```sh
npx hardhat
```

- run hardhat node

```sh
npm hardhat node
```

## Hardhat

- Hardhat Network
- Mocha as the test runner
- Commands:
```sh
Try running some of the following tasks:
  npx hardhat accounts
  npx hardhat compile
  npx hardhat test
  npx hardhat node
  node scripts/sample-script.js
  npx hardhat help
```
- fork mainnet: 
```sh
npx hardhat node --fork https://eth-mainnet.alchemyapi.io/v2/<key>
```
Or setting:
```sh
networks: {
  hardhat: {
    forking: {
      url: "https://eth-mainnet.alchemyapi.io/v2/<key>"
    }
  }
}
```
- [hardhat-network-docs](https://hardhat.org/hardhat-network/)
- default Hardhat project using:

```sh
npm install --save-dev @nomiclabs/hardhat-waffle ethereum-waffle chai @nomiclabs/hardhat-ethers ethers
```
- truffle and web3j

```sh
npm install --save-dev @nomiclabs/hardhat-truffle5 @nomiclabs/hardhat-web3 web3
```

- hardhat console

```sh
npx hardhat console
```

- Demo project

```sh
https://github.com/nomiclabs/hardhat-hackathon-boilerplate
```