package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/api/controller"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerRouter := router.Group("banner")
	var bannerController = new(controller.BannerController)
	{
		BannerRouter.GET("list", bannerController.ListBanner)
		BannerRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), bannerController.NewBanner)
		BannerRouter.PUT("update/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), bannerController.UpdateBanner)
		BannerRouter.DELETE("delete/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), bannerController.DeleteBanner)

	}
}
