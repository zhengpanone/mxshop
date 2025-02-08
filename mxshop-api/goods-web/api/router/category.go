package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/goods-web/api/controller"
	"github.com/zhengpanone/mxshop/goods-web/global"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	CategoryRouter := router.Group("category")
	var categoryController = new(controller.CategoryController)
	{
		CategoryRouter.GET("list", categoryController.GetCategoryList)
		CategoryRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), categoryController.CreateCategory)
		CategoryRouter.POST("update", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), categoryController.UpdateCategory)

	}
}
