package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middlewares"
)

func InitBrandRouter(router *gin.RouterGroup) {
	BrandRouter := router.Group("goods")
	{
		BrandRouter.GET("brand/list", api.ListBrand)
		BrandRouter.POST("brand/create", middlewares.JWTAuth(), middlewares.IsAdmin(), api.NewBrand)
		BrandRouter.PUT("brand/update", middlewares.JWTAuth(), middlewares.IsAdmin(), api.UpdateBrand)
		BrandRouter.DELETE("brand/delete", middlewares.JWTAuth(), middlewares.IsAdmin(), api.DeleteBrand)

	}
}
