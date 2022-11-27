package utils

import (
	"chains-gotest-backend/api/evm/log"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
		"strings"
)

type ExecPlan struct {
	From   string
	To     string
	It     string
	Amount string
	Action string
	Wal    string
	Symbol string
}

func ReadExcelPlan(excelPath, sheetName string) []ExecPlan {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		log.Info(err)
		return nil
	}
	result := make([]ExecPlan, 0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Info(err)
		return nil
	}
	for i := 1; i < len(rows); i++ {
		plan := &ExecPlan{
			From:   strings.TrimSpace(rows[i][0]),
			To:     strings.TrimSpace(rows[i][1]),
			Wal:    strings.TrimSpace(rows[i][2]),
			Amount: strings.TrimSpace(rows[i][3]),
			Symbol: strings.TrimSpace(rows[i][4]),
			Action: strings.TrimSpace(rows[i][5]),
			It:     strings.TrimSpace(rows[i][6]),
		}
		result = append(result, *plan)
	}
	return result
}
