package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middleware"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods")
	{
		GoodsRouter.GET("list", api.GetGoodsList)
		GoodsRouter.POST("create", middleware.JWTAuth(), middleware.IsAdmin(), api.NewGoods)

	}
}
