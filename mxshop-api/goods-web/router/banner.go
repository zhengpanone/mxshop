package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/global"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerRouter := router.Group("banner")
	var bannerController = new(api.BannerController)
	{
		BannerRouter.GET("list", bannerController.ListBanner)
		BannerRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), bannerController.NewBanner)
		BannerRouter.PUT("update/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), bannerController.UpdateBanner)
		BannerRouter.DELETE("delete/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), bannerController.DeleteBanner)

	}
}
