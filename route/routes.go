package route

import (
	"go_demo/controller"
	"go_demo/util"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	CaptchaController := controller.BaseApi{}
	UserController := controller.UserController{}
	groupBase := r.Group("/base")
	groupAuth := r.Group("/api/auth")
	groupUser := r.Group("/api/user")
	groupUser.Use(util.JwtAuthMiddleware)
	{
		groupAuth.POST("/register", controller.Register)
		groupAuth.POST("/login", controller.Login)
		groupAuth.POST("/logout", util.JwtAuthMiddleware, controller.LogOut)
	}
	{
		groupBase.POST("/captcha", CaptchaController.Captcha)
		groupBase.POST("/login", controller.Login)
	}
	{
		groupUser.GET("", UserController.GetPageList)
		groupUser.POST("/add", UserController.Add)
		groupUser.GET("/:id", UserController.Get)
	}

	return r

}
