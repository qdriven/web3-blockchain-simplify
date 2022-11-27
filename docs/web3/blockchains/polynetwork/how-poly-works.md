# How Polynetwork works

Polynetwork跨链:
1. 统一的跨链协议/区块链底层实现或者SmartContract实现
2. 跨链需要的信息 for replayer:
   1. block header format
   2. serialization/deserialzation methods 
   3. signature verification
   4. merkle tree generation and verification methods for replayer
3. Interface Requirements
   1. Block Header Synchronziation Contract, maintain block header on replayer chain for verifying chain transaction
   2. Cross Chain Management Contract: only one in one chain, for transaction in replayer


## Block Header Synchronization Contract

- SyncGenesisHeader: Synchronizes the relay chain's genesis block header, invoked only once when contract is initialized
- SyncBlockHeader: Consistently synchronizes block cycle change and cross chain transaction block headers from the relay chain

## Cross Chain Interaction

![img](https://github.com/polynetwork/docs/raw/master/poly/resources/ark.png)