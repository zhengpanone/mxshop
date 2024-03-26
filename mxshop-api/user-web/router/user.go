package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("ping", api.Ping)
		UserRouter.GET("list", api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
	}
}
