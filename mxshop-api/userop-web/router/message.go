package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "mxshop-api/common/middleware"
	"userop-web/api"
	"userop-web/global"
)

func InitMessageRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("message").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", api.GetMessageList)   // 批量获取留言信息
		GoodsRouter.POST("/create", api.CreateMessage) // 添加留言

	}
}
