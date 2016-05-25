package main

import (
	"phalgo-sample/Routes"
	"github.com/wenzhenxi/phalgo"
)

func main() {
	//初始化ECHO路由
	phalgo.NewEcho()
	//初始化配置文件
	phalgo.NewConfig("Config", "conf")
	//初始化数据库连接
	phalgo.NewDB("dbDefault")
	// Routes 载入路由
	routes.GetRoutes()
	//开启服务
	phalgo.RunFasthttp(phalgo.Config.GetString("system.port"))
}
