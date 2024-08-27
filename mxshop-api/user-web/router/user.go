package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"user-web/api"
	"user-web/global"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("list", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey) /*, middleware.IsAdmin(), */, api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
		UserRouter.POST("register", api.Register)
	}
}
