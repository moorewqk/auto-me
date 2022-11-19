package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
加载业务路由
*/

func NewRouters() *gin.Engine {
	var Router = gin.Default()

	//获取路由组实例
	//systemRouter := routers.Group.System

	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	return Router
}
