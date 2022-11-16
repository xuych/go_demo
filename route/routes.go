package route

import (
	"go_demo/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	groupAuth := r.Group("/api/auth")
	{
		groupAuth.POST("/register", controller.Register)
		groupAuth.POST("/login", controller.Login)
	}

	groupUser := r.Group("/api/user")
	{
		groupUser.GET("", controller.List)
	}

	return r

}
