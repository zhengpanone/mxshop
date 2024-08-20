package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middleware"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	CategoryRouter := router.Group("category")
	var categoryController = new(api.CategoryController)
	{
		CategoryRouter.GET("list", api.GetCategoryList)
		CategoryRouter.POST("create", middleware.JWTAuth(), middleware.IsAdmin(), categoryController.CreateCategory)
		CategoryRouter.POST("update", middleware.JWTAuth(), middleware.IsAdmin(), api.UpdateCategory)

	}
}
