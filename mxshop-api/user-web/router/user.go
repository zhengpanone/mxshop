package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "mxshop-api/common/middleware"
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
