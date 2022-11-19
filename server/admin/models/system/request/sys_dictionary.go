package request

import (
	"gitee.com/moorewqk/antcom/server/cores/models/system"
	"gitee.com/moorewqk/antcom/server/utils/request"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
