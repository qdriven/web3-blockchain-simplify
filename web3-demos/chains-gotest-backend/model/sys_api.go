package model

import (
	"chains-gotest-backend/global"
)

type SysApi struct {
	global.BaseModel
	Path        string `json:"path" gorm:"comment:api路径"`                    // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`           // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`                 // api组
	Method      string `json:"method" gorm:"default:POST" gorm:"comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}
