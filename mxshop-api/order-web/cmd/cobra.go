package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/cmd/run"
	"os"
)

var rootCmd = &cobra.Command{Use: "order-web", Short: "order-web"}

func init() {
	rootCmd.AddCommand(run.CmdRun)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
