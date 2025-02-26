package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/userop-web/api/controller"
	"github.com/zhengpanone/mxshop/userop-web/global"
)

func InitMessageRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("message").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", controller.GetMessageList)   // 批量获取留言信息
		GoodsRouter.POST("/create", controller.CreateMessage) // 添加留言

	}
}
