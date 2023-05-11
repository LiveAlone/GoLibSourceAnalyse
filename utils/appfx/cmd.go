package appfx

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd"
	"github.com/spf13/cobra"
)

func CommandProvider() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "utils",
		Short: "utils",
		Long:  "个人项目工具",
	}

	rootCmd.AddCommand(cmd.NewWordCmd())
	rootCmd.AddCommand(cmd.SqlCmd)
	rootCmd.AddCommand(cmd.FileConvert)
	rootCmd.AddCommand(cmd.ApiCmd)

	return rootCmd, nil
}
