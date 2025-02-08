package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zhengpanone/mxshop/user-web/docs"
)

// SwaggerInit swagger初始化
func SwaggerInit(engine *gin.Engine) {
	// 获取swagger
	swaggerInfo := docs.SwaggerInfo
	// 动态设置swagger
	swaggerInfo.Title = "user-web"
	swaggerInfo.Description = "mxshop-api 用户管理"
	swaggerInfo.Version = "v1.0.0"
	swaggerInfo.Host = "127.0.0.1:18021"
	swaggerInfo.BasePath = ""
	url := ginSwagger.URL("/swagger/doc.json")
	// Serve Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
