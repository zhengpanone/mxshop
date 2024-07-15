package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middlewares"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerRouter := router.Group("banner")
	{
		BannerRouter.GET("list", api.ListBanner)
		BannerRouter.POST("create", middlewares.JWTAuth(), middlewares.IsAdmin(), api.NewBanner)
		BannerRouter.PUT("update/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), api.UpdateBanner)
		BannerRouter.DELETE("delete/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), api.DeleteBanner)

	}
}
