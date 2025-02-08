package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/oss-web/cmd/run"
	"github.com/zhengpanone/mxshop/oss-web/global"
	"os"
	"time"
)

func main() {
	// 设置时区
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("设置时区失败:%s\n", err.Error())
		os.Exit(1)
	}
	time.Local = local

	rootCmd := &cobra.Command{Use: "goods-web"}
	rootCmd.AddCommand(run.CmdRun)
	_ = rootCmd.Execute()
}

// registerMiddleware 注册中间件
func registerMiddleware(r *gin.Engine) {
	// 打印日志 、异常保护
	r.Use(commonMiddleware.GinLogger(global.Logger), commonMiddleware.GinRecovery(global.Logger, true))
}
