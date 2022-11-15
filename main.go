package main

import (
	"go_demo/common"
	"go_demo/route"

	"github.com/gin-gonic/gin"
)

func main() {

	//获取初始化的数据库
	db := common.InitDB()
	//延迟关闭数据库
	defer db.Close()

	//创建一个默认的路由引擎
	r := gin.Default()

	//启动路由
	route.CollectRoutes(r)

	//在9090端口启动服务
	panic(r.Run(":9090"))
}
