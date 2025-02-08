package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/goods-web/api/controller"
	"github.com/zhengpanone/mxshop/goods-web/global"
)

func InitBrandRouter(router *gin.RouterGroup) {
	BrandRouter := router.Group("brand")
	var brandController = new(controller.BrandController)
	{
		BrandRouter.GET("list", brandController.ListBrand)
		BrandRouter.POST("create", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), brandController.NewBrand)
		BrandRouter.PUT("update", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), brandController.UpdateBrand)
		BrandRouter.DELETE("delete", commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey), commonMiddleware.IsAdmin(), brandController.DeleteBrand)
	}
	CategoryBrandRouter := router.Group("category-brand")
	{
		CategoryBrandRouter.GET("list", brandController.CategoryBrandList)     // 类别品牌列表页
		CategoryBrandRouter.DELETE(":id", brandController.DeleteCategoryBrand) // 删除类别品牌
		CategoryBrandRouter.POST("create", brandController.NewCategoryBrand)   //新建类别品牌
		CategoryBrandRouter.PUT(":id", brandController.UpdateCategoryBrand)    //修改类别品牌
		CategoryBrandRouter.GET(":id", brandController.GetCategoryBrandList)   //获取分类的品牌
	}
}
