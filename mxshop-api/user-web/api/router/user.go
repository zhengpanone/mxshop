package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"user-web/api/controller"
	"user-web/global"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("list", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey) /*, middleware.IsAdmin(), */, controller.GetUserList)
		UserRouter.POST("pwd_login", controller.PasswordLogin)
		UserRouter.POST("register", controller.Register)
	}
}
