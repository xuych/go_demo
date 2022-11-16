package main

import (
	"go_demo/common"
	"go_demo/config"
	"go_demo/route"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.ParseConfig("./config/app.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	//获取初始化的数据库
	db := common.InitDB()
	//延迟关闭数据库
	defer db.Close()

	//创建一个默认的路由引擎
	r := gin.Default()

	//启动路由
	route.CollectRoutes(r)

	//在9090端口启动服务
	panic(r.Run(":" + conf.AppPort))
}
