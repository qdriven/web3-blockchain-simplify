package request

import "chains-gotest-backend/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
