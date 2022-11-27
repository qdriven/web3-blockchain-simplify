# Simple Block chain

- 区块(Block)
- 

## 区块(Block)

区块链本质上是分布式数据库，但是区别是区块链数据库可以完全公开.
区块链存储信息的是区块(Block). 我们可以构建一个最简单的区块数据结构:
```go 
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}
```

- Data: transaction data
- Timestamp: when block is committed
- Hash: the block hash to identity this block
- PrevBlockHash: the hash of the previous block

Timestamp,PreBlockHash,Hash在Bitcoin中术语区块头.
以下是完整BitCoin区块头的结构:
|Field|Purpose|Updated When|Size(Bytes)|
|-----|-------||-----|-------|
|Version|Block version number|software updated|4|
|hashPrevBlock|256-bit hash of the previous block header|A new block comes in|32|
|hashMerkleRoot|256-bit hash based on all of the transactions in the block|A transaction is accepted|32|
|Time|||4|
|Bits|Current target in compact format|The difficulty is adjusted|4|
|Nonce|32-bit number (starts at 0)|A hash is tried (increments)4|

Block Data:
- Magic No: always 0xD9B4BEF9
- Blocksize: number of bytes following up to end of block
- Blockheader: consists of 6 items
- Transaction counter: positive integer VI = VarInt
- transactions: the (non empty) list of transactions

在简化中，Hash字段就设置为: ***Hash = SHA256(PrevBlockHash + Timestamp + Data)***

```go
func (block *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevBlockHash,block.Data,timestamp},[]byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}
```

- New Block and NewGenesisBlock 

```go 
func NewBlock(data string,prevBlockHash []byte) *Block {
	block:=&Block{time.Now().Unix(),[]byte(data),
		prevBlockHash, []byte{}}
	block.SetHash()
	return  block
}

// creates and returns genesis block
// the previous block is null
func NewGenesisBlock() *Block{
	return NewBlock("Genesis Block",[]byte{})
}
```

- BlockChain structure

```go 
type BlockChain struct {
	blocks [] *Block
}
```

- New Blockchain and Add Block

```go 
func (bc *BlockChain) AddBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks,newBlock)
}

/**
a new block chain
 */
func NewBlockChain() *BlockChain {
	return &BlockChain {[]*Block{NewGenesisBlock()}}
}
```

Then Let's start the new blockchain

```go 
func TestBasic(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
```

## Proof-Of-Work

- 如何确认交易块(Block)
- 出一个新块: 必须要被网络的其他参与者确认和同意
- 工作量证明（proof-of-work）
- 以保持每 10 分钟出 1 个新块的速度。在比特币中，这个工作就是找到一个块的哈希

### Hash计算
![img](http://upload-images.jianshu.io/upload_images/127313-e9b0730b1798704d.png?imageMogr2/auto-orient/strip|imageView2/2/w/1240)


Hash特点:
- 无法从一个哈希值恢复原始数据。也就是说，哈希并不是加密。
- 对于特定的数据，只能有一个哈希，并且这个哈希是唯一的。
- 即使是仅仅改变输入数据中的一个字节，也会导致输出一个完全不同的哈希。

- [bitcoin-hash]（https://en.bitcoin.it/wiki/Block_hashing_algorithm）