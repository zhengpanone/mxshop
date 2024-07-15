package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middlewares"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods")
	{
		GoodsRouter.GET("list", api.GetGoodsList)
		GoodsRouter.POST("create", middlewares.JWTAuth(), middlewares.IsAdmin(), api.NewGoods)

	}
}
