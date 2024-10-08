package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/global"
	commonMiddleware "mxshop-api/common/middleware"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerRouter := router.Group("banner")
	{
		BannerRouter.GET("list", api.ListBanner)
		BannerRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), api.NewBanner)
		BannerRouter.PUT("update/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), api.UpdateBanner)
		BannerRouter.DELETE("delete/:id", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), api.DeleteBanner)

	}
}
