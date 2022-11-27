package model

import (
	"chains-gotest-backend/global"
)

type PMetrics struct {
	global.BaseModel
	TotalCount         uint64
	TotalExecutionTime int
	AvgTps             string

}
