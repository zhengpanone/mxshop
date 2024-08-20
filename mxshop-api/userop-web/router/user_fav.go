package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "mxshop-api/common/middleware"
	"userop-web/api"
	"userop-web/global"
)

func InitUserFavRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("userfavs").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", api.GetFavList)      // 过滤收藏信息
		GoodsRouter.POST("/create", api.AddUserFav)   // 添加收藏
		GoodsRouter.GET("/:id", api.GetUserFavDetail) // 获取用户是否收藏
		GoodsRouter.DELETE("/:id", api.DeleteUserFav) // 取消收藏
	}

}
