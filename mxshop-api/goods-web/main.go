package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goods-web/global"
	"goods-web/initialize"
	"goods-web/middlewares"
	"goods-web/utils"
	"os"
	"path/filepath"
)

func main() {
	//initialize.InitNacos()
	// 1.初始化Logger
	initialize.InitLogger()
	// 2.初始化配置文件
	initialize.InitConfig()

	// 3.初始化routers
	Router := initialize.Routers()
	// 4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器错误")
		return
	}
	// 5. 初始化srv连接
	initialize.InitSrvConn()

	Router.Use(middlewares.MyLogger()) //注册全局中间件

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("run exe dir is %v", dir)
	//
	currentMod := gin.Mode()
	if currentMod == gin.ReleaseMode {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	zap.S().Debugf("启动服务器，访问地址：http://127.0.0.1:%d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("服务器启动失败：", err.Error())
	}

}
