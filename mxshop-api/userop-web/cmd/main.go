package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhengpanone/mxshop/userop-web/cmd/run"
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

	rootCmd := &cobra.Command{Use: "userop-web"}
	rootCmd.AddCommand(run.CmdRun)
	_ = rootCmd.Execute()
}
