package jobs

import "testing"

func TestSendPLTToPallet(t *testing.T) {
	job :=&PaletteCrossChainTestJob{}
	job.SendPltToEthJob()
}