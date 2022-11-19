package routers

import (
	"gitee.com/moorewqk/antcom/server/cores/routers/system"
	"github.com/gin-gonic/gin"
)

type Routers interface {
	LoadRouter(*gin.RouterGroup)
}

type RouterGroup struct {
	System system.Routers
}

var (
	Group = new(RouterGroup)
)

//type RouterGroup struct {
//	System   system.RouterGroup
//	Example  example.RouterGroup
//	Autocode autocode.RouterGroup
//}
//
//var RouterGroupApp = new(RouterGroup)
