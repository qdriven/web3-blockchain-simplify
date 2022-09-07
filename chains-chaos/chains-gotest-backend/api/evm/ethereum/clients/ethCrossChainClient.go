package clients

import (
	"chains-gotest-backend/api/evm/evm/binding/poly/eccd_abi"
	"chains-gotest-backend/api/evm/evm/binding/poly/eccm_abi"
	"chains-gotest-backend/api/evm/evm/binding/poly/eccmp_abi"
	"chains-gotest-backend/api/evm/evm/binding/poly/erc20_abi"
	"chains-gotest-backend/api/evm/evm/binding/poly/oep4_abi"
	"chains-gotest-backend/api/evm/log"
	"chains-gotest-backend/api/evm/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (c *EvmClient) DeployEthChainDataContract() (common.Address, *eccd_abi.EthCrossChainData, error) {
	auth := c.MakeAuth()
	contractAddress, tx, contract, err := eccd_abi.DeployEthCrossChainData(auth,
		c.EClient)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployEthChainDataContract, err: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployECCMContract(eccdAddress string) (common.Address, *eccm_abi.EthCrossChainManager, error) {
	auth := c.MakeAuth()
	address := common.HexToAddress(eccdAddress)
	contractAddress, tx, contract, err := eccm_abi.DeployEthCrossChainManager(auth,
		c.EClient, address)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployECCMContract, err: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployECCMPContract(eccmAddress string) (common.Address, *eccmp_abi.EthCrossChainManagerProxy, error) {
	auth := c.MakeAuth()
	address := common.HexToAddress(eccmAddress)
	contractAddress, tx, contract, err := eccmp_abi.DeployEthCrossChainManagerProxy(auth,
		c.EClient, address)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployECCMPContract, err: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployERC20Template() (common.Address, *erc20_abi.ERC20Template, error) {
	auth := c.MakeAuth()
	contractAddress, tx, contract, err := erc20_abi.DeployERC20Template(auth,
		c.EClient)
	if err != nil {
		log.Fatal(err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployOEP4(lockProxy string) (common.Address, *oep4_abi.OEP4Template, error) {
	auth := c.MakeAuth()
	lockProxyAddr := common.HexToAddress(lockProxy)
	contractAddress, tx, contract, err := oep4_abi.DeployOEP4Template(auth,
		c.EClient, lockProxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	c.WaitTransactionConfirm(tx.Hash())

	auth = c.MakeAuth()
	tx, err = contract.DeletageToProxy(auth, lockProxyAddr, big.NewInt(1e13))
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to DeletageToProxy: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) PauseCCMP(ccmpAddr common.Address) (common.Hash, error) {
	ccmp, err := eccmp_abi.NewEthCrossChainManagerProxy(ccmpAddr, c.backend)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("new EthCrossChainManagerProxy err: %s", err)
	}

	auth := c.MakeAuth()
	tx, err := ccmp.PauseEthCrossChainManager(auth)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("call ccmp pause err: %s", err)
	}

	c.WaitTransactionConfirm(tx.Hash())
	return tx.Hash(), nil
}

func (c *EvmClient) UnPauseCCMP(ccmpAddr common.Address) (common.Hash, error) {
	ccmp, err := eccmp_abi.NewEthCrossChainManagerProxy(ccmpAddr, c.backend)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("new EthCrossChainManagerProxy err: %s", err)
	}

	auth := c.MakeAuth()
	tx, err := ccmp.UnpauseEthCrossChainManager(auth)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("call ccmp unpause err: %s", err)
	}

	c.WaitTransactionConfirm(tx.Hash())
	return tx.Hash(), nil
}

func (c *EvmClient) UpgradeECCM(newEccmAddr, ccmpAddr common.Address) (common.Hash, error) {
	ccmp, err := eccmp_abi.NewEthCrossChainManagerProxy(ccmpAddr, c.backend)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("new EthCrossChainManagerProxy err: %s", err)
	}

	auth := c.MakeAuth()
	tx, err := ccmp.UpgradeEthCrossChainManager(auth, newEccmAddr)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("call upgradeEthCrossChainManager err: %s", err)
	}

	c.WaitTransactionConfirm(tx.Hash())
	return tx.Hash(), nil
}

func (c *EvmClient) ECCDTransferOwnerShip(eccdAddr, eccmAddr common.Address) (common.Hash, error) {
	eccd, err := eccd_abi.NewEthCrossChainData(eccdAddr, c.backend)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("new EthCrossChainData err: %s", err)
	}

	auth := c.MakeAuth()
	tx, err := eccd.TransferOwnership(auth, eccmAddr)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("call transferOwnerShip err: %s", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return tx.Hash(), nil
}

func (c *EvmClient) ECCMTransferOwnerShip(eccmAddr, ccmpAddr common.Address) (common.Hash, error) {
	eccm, err := eccm_abi.NewEthCrossChainManager(eccmAddr, c.backend)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("new EthCrossChainManager err: %s", err)
	}

	auth := c.MakeAuth()
	tx, err := eccm.TransferOwnership(auth, ccmpAddr)
	if err != nil {
		return utils.EmptyHash, fmt.Errorf("call transferOwnerShip err: %s", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return tx.Hash(), nil
}




