package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/middlewares"
	"order-web/api"
)

func InitShopCartRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("shop-cart").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("/list", api.GetShopCartList)  // 购物车列表
		GoodsRouter.POST("/create", api.NewShopCart)   // 添加商品到购物车
		GoodsRouter.DELETE("/:id", api.DeleteShopCart) //删除条目
		GoodsRouter.PATCH("/:id", api.UpdateShopCart)  //修改条目
	}
}
