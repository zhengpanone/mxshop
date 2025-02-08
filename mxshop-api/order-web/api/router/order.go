package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/order-web/api/controller"
	"github.com/zhengpanone/mxshop/order-web/global"
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
