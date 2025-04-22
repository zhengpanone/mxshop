package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhengpanone/mxshop/oss-web/cmd/run"
	"os"
	"time"
)

//	@title			oss服务
//	@description	慕学商城项目，提供订单的查询、创建、更新等功能。
//	@version		1.0
//	@contact.name	zhengpanone
//	@contact.url	http://127.0.0.1:18022/swagger/index.html
//	@host			127.0.0.1:18022
//	@BasePath		/v1/oss

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @tag.name			oss管理
// @tag.description	提供oss的增删改查功能
func main() {
	// 设置时区
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("设置时区失败:%s\n", err.Error())
		os.Exit(1)
	}
	time.Local = local

	rootCmd := &cobra.Command{Use: "oss-web"}
	rootCmd.AddCommand(run.CmdRun)
	_ = rootCmd.Execute()
}
