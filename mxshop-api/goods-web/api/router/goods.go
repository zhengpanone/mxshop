package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/goods-web/api/controller"
	"github.com/zhengpanone/mxshop/goods-web/global"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods")
	var goodsController = new(controller.GoodsController)
	{
		GoodsRouter.GET("list", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), goodsController.GetGoodsList)
		GoodsRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), goodsController.NewGoods)
		GoodsRouter.PATCH("updateStatus/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), goodsController.UpdateStatus)

	}
}
