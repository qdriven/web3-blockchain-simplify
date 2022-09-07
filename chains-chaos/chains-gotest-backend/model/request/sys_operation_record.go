package request

import "chains-gotest-backend/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
