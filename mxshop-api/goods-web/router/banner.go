package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middleware"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerRouter := router.Group("banner")
	{
		BannerRouter.GET("list", api.ListBanner)
		BannerRouter.POST("create", middleware.JWTAuth(), middleware.IsAdmin(), api.NewBanner)
		BannerRouter.PUT("update/:id", middleware.JWTAuth(), middleware.IsAdmin(), api.UpdateBanner)
		BannerRouter.DELETE("delete/:id", middleware.JWTAuth(), middleware.IsAdmin(), api.DeleteBanner)

	}
}
