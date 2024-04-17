package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middlewares"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerRouter := router.Group("goods")
	{
		BannerRouter.GET("banner/list", api.ListBanner)
		BannerRouter.POST("banner/create", middlewares.JWTAuth(), middlewares.IsAdmin(), api.NewBanner)
		BannerRouter.PUT("banner/update", middlewares.JWTAuth(), middlewares.IsAdmin(), api.UpdateBanner)
		BannerRouter.DELETE("banner/delete", middlewares.JWTAuth(), middlewares.IsAdmin(), api.DeleteBanner)

	}
}
