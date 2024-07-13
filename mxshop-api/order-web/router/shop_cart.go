package router

import (
	"github.com/gin-gonic/gin"
	"order-web/api"
	"order-web/middlewares"
)

func InitShopCartRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("shop-cart", middlewares.JWTAuth())
	{
		GoodsRouter.GET("/list", api.GetShopCartList)  // 购物车列表
		GoodsRouter.POST("/create", api.NewShopCart)   // 添加商品到购物车
		GoodsRouter.DELETE("/:id", api.DeleteShopCart) //删除条目
		GoodsRouter.PATCH("/:id", api.UpdateShopCart)  //修改条目
	}
}
