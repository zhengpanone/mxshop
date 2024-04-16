package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods")
	{
		GoodsRouter.GET("goods/list" /*middlewares.JWTAuth(), middlewares.IsAdmin(), */, api.GetGoodsList)

	}
}
