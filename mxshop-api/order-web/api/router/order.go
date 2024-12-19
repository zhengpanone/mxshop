package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"order-web/api/controller"
	"order-web/global"
)

func InitOrderRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("order").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	var orderApi = &controller.OrderApi{}
	{
		GoodsRouter.GET("/list", orderApi.GetOrderList)  // 订单列表
		GoodsRouter.POST("/create", orderApi.NewOrder)   // 新建订单
		GoodsRouter.GET("/:id", orderApi.GetOrderDetail) // 查询订单详情
	}
	PayRouter := router.Group("pay")
	var payApi = &controller.PayApi{}
	{
		PayRouter.POST("alipay/notify", payApi.AliPayNotify)
	}
}
