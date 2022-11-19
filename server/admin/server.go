package admin

import (
	"fmt"
	"gitee.com/moorewqk/antcom/server/cores/g"
	"gitee.com/moorewqk/antcom/server/cores/initialize"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(address int, router *gin.Engine) server {
	addr := fmt.Sprintf(":%d", address)
	s := endless.NewServer(addr, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunServer() {
	router := initialize.NewRouters()
	s := initServer(g.GV_SERVER.System.Addr, router)
	time.Sleep(10 * time.Microsecond)
	//GV_LOG.Infof("启动antcom服务:%d",GV_SERVER.System.Addr)
	err := s.ListenAndServe()
	if err != nil {
		//GV_LOG.Errorf("启动antcom服务:%d,失败",GV_SERVER.System.Addr)
		os.Exit(1)
	}

}
