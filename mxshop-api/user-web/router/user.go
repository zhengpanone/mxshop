package router

import (
	"github.com/gin-gonic/gin"
	"user-web/api"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("list" /*middlewares.JWTAuth(), middlewares.IsAdmin(), */, api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
		UserRouter.POST("register", api.Register)
	}
}
