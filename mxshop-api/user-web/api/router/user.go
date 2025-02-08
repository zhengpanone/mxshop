package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/user-web/api/controller"
	"github.com/zhengpanone/mxshop/user-web/global"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("list", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey) /*, middleware.IsAdmin(), */, controller.GetUserList)
		UserRouter.POST("pwd_login", controller.PasswordLogin)
		UserRouter.POST("register", controller.Register)
	}
}
