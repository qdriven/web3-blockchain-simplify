package core

import (
	"chains-gotest-backend/api/plt/config"
	"chains-gotest-backend/api/plt/pkg/log"
	"chains-gotest-backend/api/plt/pkg/sdk"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"math/big"
	"strconv"
	"github.com/ethereum/go-ethereum/common"
)

func BlockNumber() bool {
	blockNumber := admcli.GetBlockNumber()
	log.Infof("current block number %d", blockNumber)
	return true
}

func Nonce() (succeed bool) {
	var params struct {
		Address string
	}
	if err := config.LoadParams("GetNonce.json", &params); err != nil {
		log.Error(err)
		return
	}

	nonce := admcli.GetNonce(params.Address)
	log.Infof("%s nonce is %d", params.Address, nonce)
	return true
}

// 检查数据一致性(重要):
// 轮询N个节点，比较其查询所得的lastRewardBlock是否一致。非验证节点同步速度可能会慢上几个块.
func Consistency() (succeed bool) {
	var params struct {
		NodeIndex []int
	}
	if err := config.LoadParams("Consistency.json", &params); err != nil {
		log.Error(err)
		return
	}

	nodes := make([]*config.Node, 0)
	for _, v := range params.NodeIndex {
		node := config.Conf.Nodes[v]
		nodes = append(nodes, node)
	}

	clients := make([]*sdk.Client, len(nodes))
	for i, node := range nodes {
		clients[i] = sdk.NewSender(node.RPCAddr(), config.AdminKey)
	}

	currentBlkNo := clients[0].GetBlockNumber() - 10

	var i, blkNo uint64 = 0, 10
	for i = currentBlkNo - blkNo; i < currentBlkNo; i++ {
		lastRdBlk := big.NewInt(0)
		lastRdProposer := common.Address{}
		queryBlkHex := "0x" + strconv.FormatInt(int64(i), 16)
		for j, client := range clients {
			rdBlk, err := client.GetRewardRecordBlock(queryBlkHex)
			if err != nil {
				log.Error(err)
				return
			}
			rdProp, err := client.GetLatestRewardProposer(queryBlkHex)
			if err != nil {
				log.Error(err)
				return
			}

			if j == 0 {
				lastRdBlk = rdBlk
				lastRdProposer = rdProp
				continue
			}
			if lastRdBlk.Cmp(rdBlk) != 0 {
				log.Errorf("%s query result %d, %s query result %d", clients[0].Url(), lastRdBlk.Uint64(), clients[i].Url(), rdBlk.Uint64())
			}
			if lastRdProposer != rdProp {
				log.Errorf("%s query result %s, %s query result %s", clients[0].Url(), lastRdProposer.Hex(), clients[i].Url(), rdProp.Hex())
			}
		}
		log.Infof("last reward block %d, last reward proposer %s", lastRdBlk.Uint64(), lastRdProposer.Hex())
	}

	return true
}

// 准备测试需要的一定量PLT
func Deposit() (succeed bool) {
	var params struct {
		Amount float64
	}
	if err := config.LoadParams("Deposit.json", &params); err != nil {
		log.Error(err)
		return
	}

	amount := plt.MultiFloatPLT(params.Amount)
	accounts := config.Conf.Accounts
	hashList := make([]common.Hash, len(accounts))
	for i, acc := range accounts {
		to := common.HexToAddress(acc)
		hash, err := admcli.PLTTransfer(to, amount.BigInt())
		if err != nil {
			log.Errorf("failed to deposit to %s, amount %f, err: %v", to.Hex(), plt.PrintFPLT(amount), err)
			return
		}
		hashList[i] = hash
	}

	wait(2)

	if err := DumpHashList(hashList, "deposit"); err != nil {
		return
	}

	return true
}

func RemoteBuild() (succeed bool) {
	ExecRemoteBuild()
	return true
}

func RemoteSetup() (succeed bool) {
	ExecRemoteSetup()
	return true
}
