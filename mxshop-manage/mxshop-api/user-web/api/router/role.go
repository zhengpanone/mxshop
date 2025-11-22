package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/api/controller"
)

func InitRoleRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("role")
	{
		UserRouter.POST("addRole" /*commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey),*/, controller.CreateRole)

	}
}
