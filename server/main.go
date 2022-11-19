package main

import (
	"gitee.com/moorewqk/antcom/server/cores"
	"gitee.com/moorewqk/antcom/server/cores/g"
	"gitee.com/moorewqk/antcom/server/cores/initialize"
)

var (
	configPath = "config.yml"
)

func main() {
	//初始化配置:  加载配置文件,初始化GV_SERVER全局对象
	g.GV_VP = cores.NewConfig(configPath)
	g.GV_LOG = &g.ZapLogger{cores.NewZapLogger()}
	g.GV_DB = initialize.NewMySQL()
	//初始化表
	//mg := initialize.GromManger{g.GV_DB}
	//mg.MigrateTable()
	cores.RunServer()

}
