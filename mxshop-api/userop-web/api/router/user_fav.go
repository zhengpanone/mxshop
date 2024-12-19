package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"userop-web/api/controller"
	"userop-web/global"
)

func InitUserFavRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("userfavs").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", controller.GetFavList)      // 过滤收藏信息
		GoodsRouter.POST("/create", controller.AddUserFav)   // 添加收藏
		GoodsRouter.GET("/:id", controller.GetUserFavDetail) // 获取用户是否收藏
		GoodsRouter.DELETE("/:id", controller.DeleteUserFav) // 取消收藏
	}

}
