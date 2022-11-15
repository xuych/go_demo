package route

import (
	"go_demo/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	groupAuth := r.Group("/api")
	//注册
	groupAuth.POST("/register", controller.Register)
	//登录
	groupAuth.POST("/login", controller.Login)

	return r

}
