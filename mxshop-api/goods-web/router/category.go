package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/global"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	CategoryRouter := router.Group("category")
	var categoryController = new(api.CategoryController)
	{
		CategoryRouter.GET("list", categoryController.GetCategoryList)
		CategoryRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), categoryController.CreateCategory)
		CategoryRouter.POST("update", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), categoryController.UpdateCategory)

	}
}
