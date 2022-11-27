package clients

import (
	"chains-gotest-backend/api/evm/evm/binding/erc20"
	"chains-gotest-backend/api/evm/evm/binding/poly/erc20_abi"
	"chains-gotest-backend/api/evm/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"math/big"
)

// BalanceOf returns the balance of a given coin.
func (c *EvmClient) BalanceOf(erc20Address, fromAddress common.Address) (*big.Int, error) {
	erc20, err := erc20.NewErc20(erc20Address, c.EClient)
	if err != nil {
		return nil, err
	}
	balance, err := erc20.BalanceOf(nil, fromAddress)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (c *EvmClient) DeployERC20(input string) (*types.Transaction, *erc20_abi.ERC20, error) {
	auth := c.MakeAuth()
	erc20Address, tx, erc20Contract, err := erc20_abi.DeployERC20(auth, c.EClient)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	c.WaitTransactionConfirm(tx.Hash())

	log.Info("deploy erc20 contract address is ", erc20Address.String())
	totalSupply, err := erc20Contract.TotalSupply(nil)
	log.Info("deploy erc20 contract total Supply is ", totalSupply)
	return tx, erc20Contract, nil
}

func (c *EvmClient) Approve(erc20Address, spenderAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	erc20Contract, err := erc20.NewErc20(erc20Address, c.EClient)
	if err != nil {
		return nil, err
	}
	auth := c.MakeAuth()
	tx, err := erc20Contract.Approve(auth, spenderAddress, amount)
	if err != nil {
		return nil, err
	}
	c.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (c *EvmClient) ERC20Transfer(erc20Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	newErc20, err := erc20.NewErc20(erc20Address, c.EClient)
	if err != nil {
		return nil, err
	}
	auth := c.MakeAuth()
	tx, err := newErc20.Transfer(auth, to, amount)
	if err != nil {
		return nil, err
	}
	//TODO: wait tx completed
	c.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}
