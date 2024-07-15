package router

import (
	"github.com/gin-gonic/gin"
	"userop-web/api"
	"userop-web/middlewares"
)

func InitAddressRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("address").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("/list", api.GetAddressList)   // 查看地址
		GoodsRouter.POST("/create", api.CreateAddress) // 新增地址
		GoodsRouter.DELETE("/:id", api.DeleteAddress)  // 删除地址
		GoodsRouter.PUT("/:id", api.UpdateAddress)     // 修改地址
	}

}
