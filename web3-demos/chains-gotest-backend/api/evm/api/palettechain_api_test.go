package api

import (
	"chains-gotest-backend/api/evm/log"
	"fmt"
	"testing"
)

func TestSendPltToPaletteChain(t *testing.T) {
	request := &PltCrossChainRequest{
		ToAddress: "0xDe6baAE1014755a0C3F3c8Fb929ca6937f695669",
		Amount:    "1",
		ApproveAmount: "10000000000",
		Decimal: 18,
	}

	tx, err := SendPltToPalletChain(request)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(tx.Hash())
}


//TODO: loop to send api