package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "mxshop-api/common/middleware"
	"order-web/api"
	"order-web/global"
)

func InitOrderRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("order").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", api.GetOrderList)  // 订单列表
		GoodsRouter.POST("/create", api.NewOrder)   // 新建订单
		GoodsRouter.GET("/:id", api.GetOrderDetail) // 查询订单详情
	}
	PayRouter := router.Group("pay")
	{
		PayRouter.POST("alipay/notify", api.AliPayNotify)
	}
}
