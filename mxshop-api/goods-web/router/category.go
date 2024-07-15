package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middlewares"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	CategoryRouter := router.Group("category")
	{
		CategoryRouter.GET("list", api.GetCategoryList)
		CategoryRouter.POST("create", middlewares.JWTAuth(), middlewares.IsAdmin(), api.NewCategory)
		CategoryRouter.POST("update", middlewares.JWTAuth(), middlewares.IsAdmin(), api.UpdateCategory)

	}
}
