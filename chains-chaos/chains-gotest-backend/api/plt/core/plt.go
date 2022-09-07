package core

import (
	"chains-gotest-backend/api/plt/config"
	"chains-gotest-backend/api/plt/pkg/log"
	"chains-gotest-backend/api/plt/pkg/sdk"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)


func TotalSupply() (succeed bool) {
	expect := 1e9

	totalSupply, err := admcli.PLTTotalSupply("latest")
	if err != nil {
		log.Error(err)
		return
	}

	actual := plt.PrintUPLT(totalSupply)
	if actual != uint64(expect) {
		log.Errorf("totalSupply expect %d actually %d", expect, actual)
		return
	}

	log.Infof("totalSupply %d", utils.UnsafeDiv(totalSupply, plt.OnePLT))

	return true
}

func Decimal() (succeed bool) {
	data, err := admcli.PLTDecimals()
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("decimal %d", data)

	return true
}

func Name() (succeed bool) {
	actual, err := admcli.PLTName()
	if err != nil {
		log.Error(err)
		return
	}

	expect := "Palette Token"
	if actual != expect {
		log.Errorf("contract name expect %s actual %s", expect, actual)
		return
	}

	log.Infof("contract name %s", actual)

	return true
}

func AdminBalance() (succeed bool) {
	balance, err := admcli.BalanceOf(config.AdminAddr, "latest")
	if err != nil {
		log.Error(err)
		return
	}

	actual := plt.PrintUPLT(balance)
	log.Infof("admin %s balance %d", config.AdminAddr.Hex(), actual)

	return true
}

func GovernanceBalance() (succeed bool) {
	owner := common.HexToAddress(native.GovernanceContractAddress)
	balance, err := admcli.BalanceOf(owner, "latest")
	if err != nil {
		log.Error(err)
		return
	}

	actual := plt.PrintUPLT(balance)
	log.Infof("governance %s balance %d", owner.Hex(), actual)

	return true
}

func BalanceOf() (succeed bool) {
	var params struct {
		Owner    string
		BlockNum string
	}

	if err := config.LoadParams("BalanceOf.json", &params); err != nil {
		log.Error(err)
		return
	}

	owner := common.HexToAddress(params.Owner)
	balance, err := admcli.BalanceOf(owner, params.BlockNum)
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("balance %d", utils.UnsafeDiv(balance, plt.OnePLT))

	return true
}

func Transfer() (succeed bool) {
	var params struct {
		From   string
		To     string
		Amount int64
	}

	if err := config.LoadParams("Transfer.json", &params); err != nil {
		log.Error(err)
		return
	}

	key := config.LoadAccount(params.From)
	to := common.HexToAddress(params.To)
	amount := utils.SafeMul(big.NewInt(params.Amount), plt.OnePLT)

	// balance before transfer
	fromBalanceBeforeTrans, err := admcli.BalanceOf(PrivKey2Addr(key), "latest")
	if err != nil {
		log.Error(err)
		return
	}
	toBalanceBeforeTrans, err := admcli.BalanceOf(to, "latest")
	if err != nil {
		log.Error(err)
		return
	}
	if fromBalanceBeforeTrans.Cmp(amount) < 0 {
		log.Errorf("%s balance not enough %d", params.From, utils.UnsafeDiv(fromBalanceBeforeTrans, plt.OnePLT))
		return
	}

	// transfer and waiting for commit
	hash, err := admcli.PLTTransfer(to, amount)
	if err != nil {
		log.Error(err)
		return
	}
	wait(2)
	if err := admcli.DumpEventLog(hash); err != nil {
		log.Error(err)
		return
	}

	// balance after transfer
	fromBalanceAfterTrans, err := admcli.BalanceOf(PrivKey2Addr(key), "latest")
	if err != nil {
		log.Error(err)
		return
	}
	toBalanceAfterTrans, err := admcli.BalanceOf(to, "latest")
	if err != nil {
		log.Error(err)
		return
	}

	// expect sum
	if utils.SafeAdd(toBalanceBeforeTrans, amount).Cmp(toBalanceAfterTrans) != 0 {
		log.Errorf("dst balance before transfer %d, balance after transfer %d, amount %d",
			utils.UnsafeDiv(toBalanceBeforeTrans, plt.OnePLT),
			utils.UnsafeDiv(toBalanceAfterTrans, plt.OnePLT),
			params.Amount,
		)
		return
	}
	if utils.SafeSub(fromBalanceBeforeTrans, amount).Cmp(fromBalanceAfterTrans) != 0 {
		log.Errorf("src balance before transfer %d, balance after transfer %d, amount %d",
			utils.UnsafeDiv(fromBalanceAfterTrans, plt.OnePLT),
			utils.UnsafeDiv(fromBalanceAfterTrans, plt.OnePLT),
			params.Amount,
		)
		return
	}

	return true
}

func Approve() (succeed bool) {
	var params struct {
		Owner   string
		Spender string
		Amount  int
	}

	if err := config.LoadParams("Approve.json", &params); err != nil {
		log.Error(err)
		return
	}

	baseUrl := config.Conf.Nodes[0].RPCAddr()
	key := config.LoadAccount(params.Owner)
	cli := sdk.NewSender(baseUrl, key)

	owner := PrivKey2Addr(key)
	spender := common.HexToAddress(params.Spender)
	amount := plt.MultiPLT(params.Amount)

	// allowance before approve
	allowanceBeforeApprove, err := cli.PLTAllowance(owner, spender, "latest")
	if err != nil {
		log.Error(err)
		return
	}
	hash, err := cli.PLTApprove(spender, amount)
	if err != nil {
		log.Error(err)
		return
	}
	wait(2)
	if err := cli.DumpEventLog(hash); err != nil {
		log.Error(err)
		return
	}

	// allowance after approve
	allowanceAfterApprove, err := cli.PLTAllowance(owner, spender, "latest")
	if err != nil {
		log.Error(err)
		return
	}

	if allowanceAfterApprove.Cmp(amount) != 0 {
		log.Errorf("owner %s, spender %s, allowance before approve %d, allowance after approve %d, amount %d",
			owner.Hex(), spender.Hex(),
			utils.UnsafeDiv(allowanceBeforeApprove, plt.OnePLT),
			utils.UnsafeDiv(allowanceAfterApprove, plt.OnePLT),
			utils.UnsafeDiv(amount, plt.OnePLT))
		return
	}

	return true
}
