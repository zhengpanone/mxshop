package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"order-web/api/controller"
	"order-web/global"
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
