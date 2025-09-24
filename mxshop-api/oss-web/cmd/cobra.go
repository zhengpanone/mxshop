package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/cmd/run"
	"os"
)

var rootCmd = &cobra.Command{Use: "oss-web", Short: "oss-web"}

func init() {
	rootCmd.AddCommand(run.CmdRun)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
