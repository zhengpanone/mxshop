package router

import (
	"github.com/gin-gonic/gin"
	"order-web/api"
	"order-web/middlewares"
)

func InitOrderRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("order").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("/list", api.GetOrderList) // 订单列表
		GoodsRouter.POST("/create", api.NewOrder)
		GoodsRouter.POST("/:id", api.GetOrder)
	}
}
