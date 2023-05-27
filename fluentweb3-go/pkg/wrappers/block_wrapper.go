package wrappers

import (
	"context"
	"errors"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrTargetTimestampAfterLatestBlock = errors.New("target timestamp after latest block")
)

func GetTxSender(tx *types.Transaction) (from common.Address, err error) {
	from, err = types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		from, err = types.Sender(types.HomesteadSigner{}, tx)
	}
	return from, err
}

// Roughly estimate a block number by target timestamp (might be off by a lot)
func EstimateTargetBlocknumber(utcTimestamp int64) int64 {
	// Calculate block-difference from reference block
	referenceBlockNumber := int64(12323940)
	referenceBlockTimestamp := int64(1619546404) // 2021-04-27 19:49:13 +0200 CEST

	secDiff := referenceBlockTimestamp - utcTimestamp
	// fmt.Println("secDiff", secDiff)
	blocksDiff := secDiff / 13
	targetBlock := referenceBlockNumber - blocksDiff
	return targetBlock
}

// GetBlockHeaderAtTimestamp returns the header of the first block at or after the timestamp. If timestamp is after
// latest block, then return latest block. This function is a bit messy, but works. Improvements to this approach are welcome.
func GetFirstBlockHeaderAtOrAfterTime(client *ethclient.Client, targetTime time.Time) (header *types.Header, err error) {
	// Get latest header
	latestBlockHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return header, err
	}

	// Ensure target timestamp is before latest block
	targetTimestamp := targetTime.Unix()
	if uint64(targetTimestamp) > latestBlockHeader.Time {
		return header, ErrTargetTimestampAfterLatestBlock
	}

	// Estimate a target block number
	currentBlockNumber := EstimateTargetBlocknumber(targetTimestamp)

	// If estimation later than latest block, then use latest block as estimation base
	if currentBlockNumber > latestBlockHeader.Number.Int64() {
		currentBlockNumber = latestBlockHeader.Number.Int64()
	}

	// approach the target block from below, to be sure it's the first one at/after the timestamp
	var isNarrowingDownFromBelow = false

	// Ringbuffer for the latest secDiffs, to avoid going in circles when narrowing down
	lastSecDiffs := make([]int64, 7)
	lastSecDiffsIncludes := func(a int64) bool {
		for _, b := range lastSecDiffs {
			if b == a {
				return true
			}
		}
		return false
	}

	// fmt.Printf("Finding start block:\n")
	var secDiff int64
	blockSecAvg := int64(13) // average block time. is adjusted when narrowing down

	for {
		// core.DebugPrintln("Checking block:", currentBlockNumber)
		blockNumber := big.NewInt(currentBlockNumber)
		header, err := client.HeaderByNumber(context.Background(), blockNumber)
		if err != nil {
			return header, err
		}

		secDiff = int64(header.Time) - targetTimestamp

		// fmt.Printf("%d \t blockTime: %d / %v \t secDiff: %5d\n", currentBlockNumber, header.Time, time.Unix(int64(header.Time), 0).UTC(), secDiff)

		// Check if this secDiff was already seen (avoid circular endless loop)
		if lastSecDiffsIncludes(secDiff) && blockSecAvg < 25 {
			blockSecAvg += 1
			// fmt.Println("- Increase blockSecAvg to", blockSecAvg)
		}

		// Pop & add secDiff to array of last values
		lastSecDiffs = lastSecDiffs[1:]
		lastSecDiffs = append(lastSecDiffs, secDiff)
		// core.DebugPrintln("lastSecDiffs:", lastSecDiffs)

		if Abs(secDiff) < 80 || isNarrowingDownFromBelow { // getting close
			if secDiff < 0 {
				// still before wanted startTime. Increase by 1 from here...
				isNarrowingDownFromBelow = true
				currentBlockNumber += 1
				continue
			}

			// Only return if coming block-by-block from below, making sure to take first block after target time
			if isNarrowingDownFromBelow {
				return header, nil
			} else {
				currentBlockNumber -= 1
				continue
			}
		}

		// Try for better block in big steps
		blockDiff := secDiff / blockSecAvg
		currentBlockNumber -= blockDiff
	}
}

// BlockWithTxReceipts contains a single block and receipts for all its transactions (eth does not guarantee that every tx has a receipt)
type BlockWithTxReceipts struct {
	Block      *types.Block
	TxReceipts map[common.Hash]*types.Receipt
}

// GetBlockWithTxReceipts returns a single block with receipts for all transactions
func GetBlockWithTxReceipts(client *ethclient.Client, height int64) (res *BlockWithTxReceipts, err error) {
	res = &BlockWithTxReceipts{}
	res.TxReceipts = make(map[common.Hash]*types.Receipt)

	// Get the block
	res.Block, err = client.BlockByNumber(context.Background(), big.NewInt(height))
	if err != nil {
		return res, err
	}

	// Get receipts for all transactions
	for _, tx := range res.Block.Transactions() {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				// can apparently happen if 0 tx: https://etherscan.io/block/10102170
				continue
			}
			return res, err

		}
		res.TxReceipts[tx.Hash()] = receipt
	}

	return res, nil
}

// GetBlocksWithTxReceipts downloads a range of blocks with tx receipts, and sends each to a channel once it is ready.
// Uses concurrency parallel connections to get data from the eth node fast. 5 seems a good number for a direct IPC connection.
func GetBlocksWithTxReceipts(client *ethclient.Client, blockChan chan<- *BlockWithTxReceipts, startBlock int64, endBlock int64, concurrency int) {
	var blockWorkerWg sync.WaitGroup
	blockHeightChan := make(chan int64, 100) // blockHeight to fetch with receipts

	// Start eth client thread pool
	for w := 1; w <= concurrency; w++ {
		blockWorkerWg.Add(1)

		go func() {
			defer blockWorkerWg.Done()
			for blockHeight := range blockHeightChan {
				res, err := GetBlockWithTxReceipts(client, blockHeight)
				if err != nil {
					log.Println("Error getting block with tx receipts:", err)
					continue
				}
				blockChan <- res
			}
		}()
	}

	// Push blocks into channel, for workers to pick up
	for currentBlockNumber := startBlock; currentBlockNumber <= endBlock; currentBlockNumber++ {
		blockHeightChan <- currentBlockNumber
	}

	// Close worker channel and wait for workers to finish
	close(blockHeightChan)
	blockWorkerWg.Wait()
}
