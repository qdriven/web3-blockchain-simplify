# Contract Operations

- Compiling
- Testing
- Deploying
- Verifying
- Task and Scripts

## Compile Contract

- Solidity code to Binary
- After Compile, then the contract could be deployed
- compile command
```shell
npx hardhat compile
hardhat compile
```
- configurations

```js
module.exports = {
  solidity: {
    version: "0.8.9",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
};
```
additional options: evmVersion

- Testing

```shell
npx hardhat test
```
