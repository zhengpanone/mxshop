package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/api/controller"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/global"
)

func InitShopCartRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("shop-cart").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	var shopCartApi = &controller.ShopCartApi{}
	{
		GoodsRouter.GET("/list", shopCartApi.GetShopCartList)  // 购物车列表
		GoodsRouter.POST("/create", shopCartApi.NewShopCart)   // 添加商品到购物车
		GoodsRouter.DELETE("/:id", shopCartApi.DeleteShopCart) //删除条目
		GoodsRouter.PATCH("/:id", shopCartApi.UpdateShopCart)  //修改条目
	}
}
