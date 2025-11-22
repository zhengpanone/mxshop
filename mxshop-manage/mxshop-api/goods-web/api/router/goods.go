package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/api/controller"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
)

func InitGoodsRouter(router *gin.RouterGroup) {

	var goodsController = new(controller.GoodsController)
	{
		router.POST("getGoodsPageList", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), goodsController.GetGoodsPageList)
		router.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), goodsController.NewGoods)
		router.PATCH("updateStatus/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), goodsController.UpdateStatus)

	}
}
