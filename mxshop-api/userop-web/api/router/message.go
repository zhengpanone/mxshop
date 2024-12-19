package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"userop-web/api/controller"
	"userop-web/global"
)

func InitMessageRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("message").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", controller.GetMessageList)   // 批量获取留言信息
		GoodsRouter.POST("/create", controller.CreateMessage) // 添加留言

	}
}
