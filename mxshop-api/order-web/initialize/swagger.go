package initialize

import (
	"common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"order-web/docs"
	"order-web/global"
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
	swaggerInfo.Title = "order-web"
	swaggerInfo.Description = "mxshop-api 订单管理"
	swaggerInfo.Version = "v1.0.0"
	swaggerInfo.Host = fmt.Sprintf("%s:%d", utils.GetIP(), global.ServerConfig.Port)
	swaggerInfo.BasePath = ""
	url := ginSwagger.URL("/swagger/doc.json")
	// Serve Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
