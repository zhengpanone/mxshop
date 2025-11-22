package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/cmd/run"
	"os"
)

var rootCmd = &cobra.Command{Use: "userop-web", Short: "userop-web"}

func init() {
	rootCmd.AddCommand(run.CmdRun)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
