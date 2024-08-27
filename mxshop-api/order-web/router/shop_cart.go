package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"order-web/api"
	"order-web/global"
)

func InitShopCartRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("shop-cart").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", api.GetShopCartList)  // 购物车列表
		GoodsRouter.POST("/create", api.NewShopCart)   // 添加商品到购物车
		GoodsRouter.DELETE("/:id", api.DeleteShopCart) //删除条目
		GoodsRouter.PATCH("/:id", api.UpdateShopCart)  //修改条目
	}
}
