package system

import (
	"gitee.com/moorewqk/antcom/server/cores/types"
)

/*
系统配置表
*/
// 配置文件结构体
type System struct {
	Config types.Server `json:"config"`
}
