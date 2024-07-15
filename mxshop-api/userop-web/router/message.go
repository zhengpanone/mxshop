package router

import (
	"github.com/gin-gonic/gin"
	"userop-web/api"
	"userop-web/middlewares"
)

func InitMessageRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("message").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("/list", api.GetMessageList)   // 批量获取留言信息
		GoodsRouter.POST("/create", api.CreateMessage) // 添加留言

	}
}
