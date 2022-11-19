package system

import (
	"gitee.com/moorewqk/antcom/server/cores/types"
)

/*
认证黑名单
*/
type JwtBlacklist struct {
	types.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
