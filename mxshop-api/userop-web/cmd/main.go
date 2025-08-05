package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/cmd/run"
	"os"
	"time"
)

//	@title			用户操作服务
//	@description	慕学商城项目，提供商品的查询、创建、更新等功能。
//	@version		1.0
//	@contact.name	zhengpanone
//	@contact.url	http://127.0.0.1:18022/swagger/index.html
//	@host			127.0.0.1:18022
//	@BasePath		/v1/goods

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@tag.name			用户操作管理
//	@tag.description	提供用户操作的增删改查功能

// https://github.com/gphper/ginadmin
func main() {

	// 设置时区
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("设置时区失败:%s\n", err.Error())
		os.Exit(1)
	}
	time.Local = local

	rootCmd := &cobra.Command{Use: "userop-web"}
	rootCmd.AddCommand(run.CmdRun)
	_ = rootCmd.Execute()
}
