package router

import (
	"github.com/gin-gonic/gin"
	"userop-web/api"
	"userop-web/middlewares"
)

func InitUserFavRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("userfavs").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("/list", api.GetFavList)      // 过滤收藏信息
		GoodsRouter.POST("/create", api.AddUserFav)   // 添加收藏
		GoodsRouter.GET("/:id", api.GetUserFavDetail) // 获取用户是否收藏
		GoodsRouter.DELETE("/:id", api.DeleteUserFav) // 取消收藏
	}

}
