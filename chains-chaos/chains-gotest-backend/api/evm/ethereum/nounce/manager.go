package nounce

import (
	"chains-gotest-backend/api/evm/log"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"time"
)

const ClearNonceInterval = 10 * time.Minute

type NonceManager struct {
	addressNonce map[common.Address]uint64
	ethClient    *ethclient.Client
	lock         sync.RWMutex
}

func NewNonceManager(ethClient *ethclient.Client) *NonceManager {
	nonceManager := &NonceManager{
		addressNonce: make(map[common.Address]uint64),
		ethClient:    ethClient,
	}
	go nonceManager.clearNonce()
	return nonceManager
}

// return account nonce, and than nonce++
func (this *NonceManager) GetAddressNonce(address common.Address) uint64 {
	this.lock.Lock()
	defer this.lock.Unlock()

	// return a new point
	nonce, ok := this.addressNonce[address]
	if !ok {
		// get nonce from eth network
		uintNonce, err := this.ethClient.PendingNonceAt(context.Background(), address)
		if err != nil {
			log.Infof("GetAddressNonce: cannot get account %s nonce, err: %s, set it to nil!",
				address, err)
		}
		this.addressNonce[address] = uintNonce
		nonce = uintNonce
	}
	// increase record
	this.addressNonce[address]++
	return nonce
}

func (this *NonceManager) DecreaseAddressNonce(address common.Address) {
	this.lock.Lock()
	defer this.lock.Unlock()

	nonce, ok := this.addressNonce[address]
	if ok && nonce > 0 {
		this.addressNonce[address]--
	}
}

// clear nonce per
func (this *NonceManager) clearNonce() {
	for {
		<-time.After(ClearNonceInterval)
		this.lock.Lock()
		for addr := range this.addressNonce {
			delete(this.addressNonce, addr)
		}
		this.lock.Unlock()
		//log.Infof("clearNonce: clear all cache nonce")
	}
}
