# Blockchain Intro 区块链介绍

```
A blockchain is a linked list of transactions stored on a network of computers.
```

Blockchain的设计目标：

- Decentralized: Transactions are on a network of computers (nodes).
- Immutable: Transactions cannot be changed once committed. 
- Open: Transactions can be viewed by anyone.

是否完全实现这个目标，以及现实意义是什么？是不是需要根据现实进行调整？都是非常值得讨论的。
Open/Immutable是可以度量的,由于Open的存在，透明度的问题可以帮助解决Immutable的问题，Immutable某种程度上不是完全意义上的Immutable，但是透明度的存在，可以帮助完善.

去中心化本身是个伪命题，没有办法度量，是多少的问题，而不是完全的去中性化.

## Blockchain 如何工作

Each block has:

- A list of transactions
- A hash (a long string of random characters) for the block
- The previous block’s hash (this is how the blocks are linked)

![img](https://global-uploads.webflow.com/617702c73410810254ccd237/619e8c6e3ab36fcb7d00f68c_Blockchain.png)

> First, both Bob and Mary need crypto wallets. These wallets don’t actually store crypto assets. Instead, they store two keys:

    -  A public key links to an address that lets you send and receive transactions. Think of it as your email address.
    - A private key proves that you own the tokens associated with your public address. Think of it as your email password. Since a private key is hard to remember (it’s a long string of random numbers), wallets also give you a 12-24 word seed phrase. You shouldn't share your private key or seed phrase with anyone.

![img](https://global-uploads.webflow.com/617702c73410810254ccd237/619e8c8b3c25b65628d0ae53_Blockchain%20(1).png)

> Bob tells his wallet: "I want to send 1 bitcoin from my public address to Mary's public address." Bob signs this transaction based on his private key. This signature proves that Bob actually owns 1 bitcoin.
Bob's wallet sends the transaction to nodes on the blockchain. These nodes then verify the transaction using Bob’s signature and public key.
A node groups Bob’s transaction with other transactions into a block. It then works with other nodes to add the block to the blockchain.

## 共识机制

To process transactions without middlemen, nodes need to be able to reach consensus themselves. They do this through two popular methods:


![img](https://global-uploads.webflow.com/617702c73410810254ccd237/619e8d156689c4e24716206a_Consensus%20mechanisms.png)

> Proof of work

Nodes called miners compete to solve a math problem using brute force (e.g., rolling a dice thousands of times to get the right number).
The first miner that solves the problem gets to create a block.
Other nodes check if the block is valid. If it is, the miner is rewarded cryptocurrency. If it’s not, the miner wasted their time and energy.
All nodes add the new block to their copy of the blockchain.
Proof of work uses energy because miners compete to solve math problems by building powerful machines that run 24/7.

Proof of stake

Nodes called validators stake some cryptocurrency. A stake is like saying: “I’ll commit this amount of cryptocurrency to win the right to do this transaction.”
Validators with more stake are more likely (but not guaranteed) to be selected to process the transaction and create a block.
Other validators check if the block is valid. If it is, all participating validators earn a transaction fee. If it’s not, the validator that created the block might lose its stake.
All nodes add the new block to their copy of the blockchain.


## Blockchain trilemma

![img](https://global-uploads.webflow.com/617702c73410810254ccd237/620ff3395e52851078b35232_Blockchain%20trilemma.png)

> For example, consider Ethereum and Solana blockchains. As of December 2021:

Solana’s average transaction cost is $0.00025 with 1,000 validator nodes.
Ethereum’s average transaction cost is $5.2 with 10,000+ validator nodes.
The more nodes there are, the less likely it is for a chain to become compromised. This post isn’t about which chain is better - just remember that this trade-off exists.