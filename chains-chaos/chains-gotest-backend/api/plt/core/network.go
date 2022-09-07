package core

import (
	"chains-gotest-backend/api/plt/config"
	"time"
)

// --------------------------------
// genesis network
// --------------------------------
func ResetGenesisNetwork() (succeed bool) {
	if !StopGenesisNetwork() {
		return
	}
	if !ClearGenesisNetwork() {
		return
	}
	if !InitGenesisNetwork() {
		return
	}
	if !StartGenesisNetwork() {
		return
	}
	return true
}

func ReStartGenesisNetwork() (succeed bool) {
	if !StopGenesisNetwork() {
		return
	}
	if !StartGenesisNetwork() {
		return
	}
	return true
}

func InitGenesisNetwork() (succeed bool) {
	nodes := config.Conf.GenesisNodes()
	for _, node := range nodes {
		ExecInitNode(node)
	}

	return true
}

func StartGenesisNetwork() (succeed bool) {
	nodes := config.Conf.GenesisNodes()
	for _, node := range nodes {
		ExecStartNode(node)
	}

	ExecGrep()
	return true
}

func StopGenesisNetwork() (succeed bool) {
	nodes := config.Conf.GenesisNodes()
	for _, node := range nodes {
		ExecStopNode(node)
	}

	ExecGrep()
	return true
}

func ClearGenesisNetwork() (succeed bool) {
	nodes := config.Conf.GenesisNodes()
	for _, node := range nodes {
		ExecClearNode(node)
	}

	return true
}

// --------------------------------
// validators network
// --------------------------------
func ResetValidatorNetwork() (succeed bool) {
	if !StopValidatorNetwork() {
		return
	}
	if !ClearValidatorNetwork() {
		return
	}
	if !InitValidatorNetwork() {
		return
	}
	if !StartValidatorNetwork() {
		return
	}
	return true
}

func ReStartValidatorNetwork() (succeed bool) {
	if !StopValidatorNetwork() {
		return
	}
	if !StartValidatorNetwork() {
		return
	}
	return true
}

func InitValidatorNetwork() (succeed bool) {
	nodes := config.Conf.ValidatorNodes()
	for _, node := range nodes {
		ExecInitNode(node)
	}

	return true
}

func StartValidatorNetwork() (succeed bool) {
	nodes := config.Conf.ValidatorNodes()
	for _, node := range nodes {
		ExecStartNode(node)
	}

	ExecGrep()
	return true
}

func StopValidatorNetwork() (succeed bool) {
	nodes := config.Conf.ValidatorNodes()
	for _, node := range nodes {
		ExecStopNode(node)
	}

	ExecGrep()
	return true
}

func ClearValidatorNetwork() (succeed bool) {
	nodes := config.Conf.ValidatorNodes()
	for _, node := range nodes {
		ExecClearNode(node)
	}

	return true
}

// --------------------------------
// all network
// --------------------------------
func ReStartAllNetwork() (succeed bool) {
	if !StopAllNetwork() {
		return
	}
	time.Sleep(3 * time.Second)
	if !StartAllNetwork() {
		return
	}
	return true
}

func InitAllNetwork() (succeed bool) {
	if !InitGenesisNetwork() {
		return
	}
	if !InitValidatorNetwork() {
		return
	}
	return true
}

func StartAllNetwork() (succeed bool) {
	nodes := config.Conf.AllNodes()
	for _, node := range nodes {
		ExecStartNode(node)
	}

	ExecGrep()
	return true
}

func StopAllNetwork() (succeed bool) {
	nodes := config.Conf.AllNodes()
	for _, node := range nodes {
		ExecStopNode(node)
	}

	ExecGrep()
	return true
}

func ClearAllNetwork() (succeed bool) {
	nodes := config.Conf.AllNodes()
	for _, node := range nodes {
		ExecClearNode(node)
	}

	return true
}

func ResetAllNetwork() (succeed bool) {
	if !StopAllNetwork() {
		return
	}
	if !ClearAllNetwork() {
		return
	}
	if !InitAllNetwork() {
		return
	}
	return StartAllNetwork()
}

//// --------------------------------
//// start sync node
//// --------------------------------
//func StartSyncNode() bool {
//	shell.Exec(shStartSyncNode)
//	wait(3)
//
//	url := "http://127.0.0.1:22010"
//	client := sdk.NewSender(url, config.AdminKey)
//
//	// check current block number
//	for i := 0; i < 3; i++ {
//		currentBlock := client.GetBlockNumber()
//		wait(1)
//		log.Infof("sync node check current block %d", currentBlock)
//	}
//
//	// check admin nonce
//	nonce := client.GetNonce(config.AdminAddr.Hex())
//	if nonce == 0 {
//		log.Error("sync node check admin nonce failed")
//		return false
//	}
//	log.Infof("sync node check admin nonce %d", nonce)
//
//	// check PLT total supply
//	totalSupply, err := client.PLTTotalSupply("latest")
//	if err != nil {
//		log.Errorf("sync node check total supply failed, [%v]", err)
//		return false
//	} else {
//		log.Infof("sync node check total supply %d", utils.UnsafeDiv(totalSupply, plt.OnePLT))
//	}
//
//	return true
//}
//
//func StopSyncNode() bool {
//	shell.Exec(shStopSyncNode)
//	return true
//}
