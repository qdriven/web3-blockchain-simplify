package model

import (
	"chains-gotest-backend/global"
)

type JwtBlacklist struct {
	global.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
