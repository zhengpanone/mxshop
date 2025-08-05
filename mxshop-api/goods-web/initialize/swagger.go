package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/docs"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
)

// SwaggerInit swagger初始化
//
// Parameters:
// - engine: gin.Engine - gin Engine
//
// Returns:
// - nil
func SwaggerInit(engine *gin.Engine) {
	// 获取swagger
	swaggerInfo := docs.SwaggerInfo
	// 动态设置swagger
	swaggerInfo.Title = "goods-web"
	swaggerInfo.Description = "mxshop-api 商品管理"
	swaggerInfo.Version = "v1.0.0"
	swaggerInfo.Host = fmt.Sprintf("%s:%d", commonUtils.GetIP(), global.ServerConfig.Port)
	swaggerInfo.BasePath = ""
	url := ginSwagger.URL("/swagger/doc.json")
	// Serve Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
