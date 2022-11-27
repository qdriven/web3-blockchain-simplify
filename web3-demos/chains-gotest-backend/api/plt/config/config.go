package config

import (
	"chains-gotest-backend/api/plt/pkg/encode"
	"chains-gotest-backend/api/plt/pkg/files"
	"chains-gotest-backend/api/plt/pkg/sdk"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"path"
	"sort"
	"sync"
	"time"
)

const (
	testCaseDir = "cases"
	keystoreDir = "keystore"
	setupDir    = "setup"
)

var (
	Conf, BakConf = new(Config), new(Config)
	AdminKey      *ecdsa.PrivateKey
	AdminAddr     common.Address
)

type Config struct {
	Environment           *Env
	Network               *Network
	DefaultPassphrase     string
	AdminAccount          string
	BaseRewardPool        string
	Accounts              []string
	GasLimit              uint64
	DeployGasLimit        uint64
	BlockPeriod           encode.Duration
	RewardEffectivePeriod int // 区块奖励周期/参数生效周期
	Nodes                 []*Node
	BaseRPCUrl            string
	PalletWalletAddress   string

}

func (c *Config) DeepCopy() *Config {
	cp := new(Config)
	enc, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(enc, cp); err != nil {
		panic(err)
	}
	return cp
}

func (c *Config) AllNodeAddressList() []string {
	list := make([]string, len(c.Nodes))
	for i, node := range c.Nodes {
		list[i] = node.Address
	}
	return list
}

func (c *Config) IpList() []string {
	data := make(map[string]struct{})
	for _, v := range c.Nodes {
		data[v.Host] = struct{}{}
	}

	list := make([]string, 0)
	for host, _ := range data {
		list = append(list, host)
	}

	return list
}

func (c *Config) GenesisNodes() Nodes {
	start := c.Network.NodeIndexStart
	end := start + c.Network.GenesisNodeNumber - 1
	return c.getRangeNodes(start, end)
}

func (c *Config) ValidatorNodes() Nodes {
	genesisStart := c.Network.NodeIndexStart
	genesisEnd := genesisStart + c.Network.GenesisNodeNumber - 1

	start := genesisEnd + 1
	end := start + c.Network.ValidatorsNumber - 1
	return c.getRangeNodes(start, end)
}

func (c *Config) AllNodes() Nodes {
	start := c.Network.NodeIndexStart
	num := c.Network.GenesisNodeNumber + c.Network.ValidatorsNumber
	end := start + num - 1
	return c.getRangeNodes(start, end)
}

func (c *Config) SpareNodes() Nodes {
	end := c.Network.GenesisNodeNumber + c.Network.ValidatorsNumber
	return c.Nodes[end:]
}

func (c *Config) GetNodeByIndex(index int) *Node {
	for _, n := range c.Nodes {
		if n.Index == index {
			return n
		}
	}
	return nil
}

func (c *Config) GetNodeByAddress(address string) *Node {
	for _, n := range c.Nodes {
		if n.Address == address {
			return n
		}
	}
	return nil
}


func (c *Config) getRangeNodes(start, end int) Nodes {
	list := make([]*Node, 0)
	for i := start; i <= end; i++ {
		list = append(list, c.Nodes[i])
	}
	return list
}

type Nodes []*Node

func (n Nodes) Validators() []common.Address {
	list := make([]common.Address, len(n))
	for i, node := range n {
		list[i] = node.NodeAddr()
	}
	return list
}

func (n Nodes) StakeAccounts() []common.Address {
	list := make([]common.Address, len(n))
	for i, node := range n {
		list[i] = node.StakeAddr()
	}
	return list
}

type Node struct {
	Index        int    `json:"Index"`
	Address      string `json:"Address"`
	NodeKey      string `json:"NodeKey"`
	StakeAccount string `json:"StakeAccount"`
	Host         string `json:"Host"`
	RPCPort      string `json:"RPCPort"`
	P2PPort      string `json:"P2PPort"`
	NodeDir      string

	once       sync.Once
	ndpk, sapk *ecdsa.PrivateKey
}

func (n *Node) init() {

	// load node private key
	bz, err := hex.DecodeString(n.NodeKey)
	if err != nil {
		panic(err)
	}
	if pk, err := crypto.ToECDSA(bz); err != nil {
		panic(err)
	} else {
		n.ndpk = pk
	}

	// load node stake account private key
	file := path.Join(Conf.Environment.LocalWorkspace, keystoreDir, n.StakeAccount)
	if bz, err = ioutil.ReadFile(file); err != nil {
		panic(fmt.Sprintf("load keystore err %v", err))
	}
	if ks, err := keystore.DecryptKey(bz, Conf.DefaultPassphrase); err != nil {
		panic(fmt.Sprintf("decrypt key %s err %v", n.StakeAccount, err))
	} else {
		n.sapk = ks.PrivateKey
	}

	// load node workspace
	nodedir := fmt.Sprintf("node%d", n.Index)
	if Conf.Environment.Remote {
		n.NodeDir = path.Join(Conf.Environment.RemoteWorkspace, nodedir)
	} else {
		n.NodeDir = path.Join(Conf.Environment.LocalWorkspace, nodedir)
	}
}

func (n *Node) NodeDirPath() string {
	n.once.Do(n.init)
	return n.NodeDir
}

func (n *Node) PrivateKey() *ecdsa.PrivateKey {
	n.once.Do(n.init)
	return n.ndpk
}


func (n *Node) NodeAddr() common.Address {
	n.once.Do(n.init)
	return common.HexToAddress(n.Address)
}

func (n *Node) StakePrivateKey() *ecdsa.PrivateKey {
	n.once.Do(n.init)
	return n.sapk
}

func (n *Node) StakePrivateKeyString() string {
	return fmt.Sprintf("%x", crypto.FromECDSA(n.StakePrivateKey()))
}

func (n *Node) StakeAddr() common.Address {
	n.once.Do(n.init)
	return common.HexToAddress(n.StakeAccount)
}

func (n *Node) RPCAddr() string {
	n.once.Do(n.init)
	return fmt.Sprintf("http://%s:%s", n.Host, n.RPCPort)
}

type Env struct {
	Remote          bool
	LocalWorkspace  string
	RemoteWorkspace string
	NetworkID       int
	LogLevel        int
	IpList          []string
}

type Network struct {
	NodeIndexStart    int
	GenesisNodeNumber int
	ValidatorsNumber  int
}

func Init(path string) {
	if err := LoadConfig(path, Conf); err != nil {
		panic(err)
	}

	// sort nodes with node index
	sort.Slice(Conf.Nodes, func(i, j int) bool {
		return Conf.Nodes[i].Index < Conf.Nodes[j].Index
	})

	// load nodes privateKey
	sdk.Init(Conf.GasLimit, Conf.DeployGasLimit, time.Duration(Conf.BlockPeriod))

	AdminKey = LoadAccount(Conf.AdminAccount)
	AdminAddr = crypto.PubkeyToAddress(AdminKey.PublicKey)
	BakConf = Conf.DeepCopy()

}

func LoadConfig(filepath string, ins interface{}) error {
	data, err := files.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, ins)
	if err != nil {
		return fmt.Errorf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return nil
}

func LoadParams(fileName string, data interface{}) error {
	filePath := files.FullPath(Conf.Environment.LocalWorkspace, testCaseDir, fileName)
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(bz, data)
}

func LoadAccount(address string) *ecdsa.PrivateKey {
	filepath := files.FullPath(Conf.Environment.LocalWorkspace, keystoreDir, address)
	keyJson, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Errorf("failed to read file: [%v]", err))
	}

	key, err := keystore.DecryptKey(keyJson, Conf.DefaultPassphrase)
	if err != nil {
		panic(fmt.Errorf("failed to decrypt keyjson: [%v]", err))
	}

	return key.PrivateKey
}

func ShellPath(fileName string) string {
	return files.FullPath(Conf.Environment.LocalWorkspace, "", fileName)
}

func GenesisNodeNumber() int {
	filepath := files.FullPath(Conf.Environment.LocalWorkspace, setupDir, "static-nodes.json")
	keyJson, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Errorf("failed to read file: [%v]", err))
	}

	var nodes []string
	if err := json.Unmarshal(keyJson, &nodes); err != nil {
		panic(fmt.Errorf("failed to unmarshal static-nodes.json: [%v]", err))
	}

	return len(nodes)
}
