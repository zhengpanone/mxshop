package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/api/controller"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/global"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("getAdminList", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey) /*, middleware.IsAdmin(), */, controller.GetAdminUserList)
		UserRouter.POST("pwd_login", controller.PasswordLogin)
		UserRouter.POST("register", controller.Register)
		UserRouter.POST("logout", controller.LogOut)
	}
}
