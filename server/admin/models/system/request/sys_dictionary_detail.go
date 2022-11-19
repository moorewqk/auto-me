package request

import (
	"gitee.com/moorewqk/antcom/server/cores/models/system"
	"gitee.com/moorewqk/antcom/server/utils/request"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
