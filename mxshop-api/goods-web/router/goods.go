package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middlewares"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods")
	{
		GoodsRouter.GET("goods/list", api.GetGoodsList)
		GoodsRouter.POST("goods/create", middlewares.JWTAuth(), middlewares.IsAdmin(), api.NewGoods)

	}
}
