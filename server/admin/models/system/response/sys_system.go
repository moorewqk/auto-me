package response

import "gitee.com/moorewqk/antcom/server/cores/types"

type SysConfigResponse struct {
	Config types.Server `json:"config"`
}
