package jobs

import (
	"chains-gotest-backend/api/evm/log"
	"github.com/go-resty/resty/v2"

	"time"
)

type PaletteCrossChainTestJob struct {
	url  string
	data interface{}
}

func (job *PaletteCrossChainTestJob) Send() {
	client := resty.New()
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(job.data).Post(job.url)
	if err != nil {
		log.Error("send plt to ETH error")
	}
	if resp != nil {
		log.Info("plt send eth to ETH respose:", resp.String())
	} else {
		log.Error("Send PLT TO ETH Error")
	}
}

func (job *PaletteCrossChainTestJob) SendPltToEthJob() {
	client := resty.New()
	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(`{
  		"amount": "0.0001",
  		"fromAddress": "0xDe6baAE1014755a0C3F3c8Fb929ca6937f695669",
  		"toAddress": "0x344cfc3b8635f72f14200aaf2168d9f75df86fd3"
	}`).Post("http://106.75.232.131:8888/pallet/sendPltToEth")

	if err != nil {
		log.Error("send plt to ETH error")
	}
	if resp != nil {
		log.Info("plt send eth to ETH response:", resp.String())
	} else {
		log.Error("Send PLT TO ETH Error")
	}
}

func (job *PaletteCrossChainTestJob) SendEthToPLTJob() {
	client := resty.New()

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(`{
  		"amount": "0.0001",
		"toAddress": "0xDe6baAE1014755a0C3F3c8Fb929ca6937f695669"
	}`).Post("http://106.75.232.131:8889/chains/sendPltToPalletChain")

	if err != nil {
		log.Error("send ETH to PLT error")
	}
	if resp != nil {
		log.Info("plt send ETH to PLT response:", resp.String())
	} else {
		log.Error("Send ETH TO PLT Error")
	}
}

func InitJob() {
	job := PaletteCrossChainTestJob{}
	for {
		time.Sleep(5 * time.Second)
		job.SendPltToEthJob()
		job.SendEthToPLTJob()
		time.Sleep(300 * time.Second)
	}
}
