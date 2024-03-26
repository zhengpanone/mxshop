package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
	"mxshop-api/user-web/middlewares"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("ping", api.Ping)
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdmin(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
	}
}
