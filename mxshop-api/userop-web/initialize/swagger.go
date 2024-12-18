package initialize

import (
	commonUtils "common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"userop-web/docs"
	"userop-web/global"
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
	swaggerInfo.Title = "userop-web"
	swaggerInfo.Description = "mxshop-api 用户操作管理"
	swaggerInfo.Version = "v1.0.0"
	swaggerInfo.Host = fmt.Sprintf("%s:%d", commonUtils.GetIP(), global.ServerConfig.Port)
	swaggerInfo.BasePath = ""
	url := ginSwagger.URL("/swagger/doc.json")
	// Serve Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
