package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middleware"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	CategoryRouter := router.Group("category")
	{
		CategoryRouter.GET("list", api.GetCategoryList)
		CategoryRouter.POST("create", middleware.JWTAuth(), middleware.IsAdmin(), api.NewCategory)
		CategoryRouter.POST("update", middleware.JWTAuth(), middleware.IsAdmin(), api.UpdateCategory)

	}
}
