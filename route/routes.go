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
	UserController := controller.UserController{}
	groupUser := r.Group("/api/user")
	{
		groupUser.GET("", UserController.GetPageList)
		groupUser.POST("/add", UserController.Add)
		groupUser.GET("/:id", UserController.Get)
	}

	return r

}
