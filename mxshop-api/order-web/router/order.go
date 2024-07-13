package router

import (
	"github.com/gin-gonic/gin"
	"order-web/api"
	"order-web/middlewares"
)

func InitOrderRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("order")
	{
		GoodsRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdmin(), api.GetOrderList) // 订单列表
		GoodsRouter.POST("/create", middlewares.JWTAuth(), api.NewOrder)
		GoodsRouter.POST("/:id", middlewares.JWTAuth(), api.GetOrder)
	}
}
