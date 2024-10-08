package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/global"
	commonMiddleware "mxshop-api/common/middleware"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods")
	{
		GoodsRouter.GET("list", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), api.GetGoodsList)
		GoodsRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), api.NewGoods)

	}
}
