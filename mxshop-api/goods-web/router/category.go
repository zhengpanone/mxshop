package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middlewares"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	CategoryRouter := router.Group("goods")
	{
		CategoryRouter.GET("category/list", api.GetCategoryList)
		CategoryRouter.POST("category/create", middlewares.JWTAuth(), middlewares.IsAdmin(), api.NewCategory)
		CategoryRouter.POST("category/update", middlewares.JWTAuth(), middlewares.IsAdmin(), api.UpdateCategory)

	}
}
