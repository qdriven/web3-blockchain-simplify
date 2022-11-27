package request

import "chains-gotest-backend/model"

type SysDictionaryDetailSearch struct {
	model.SysDictionaryDetail
	PageInfo
}
