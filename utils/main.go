package main

import (
	"log"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmd.WordCmd)
	rootCmd.AddCommand(cmd.SqlCmd)
	rootCmd.AddCommand(cmd.FileConvert)
	rootCmd.AddCommand(cmd.ApiCmd)
}

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "utils",
	Long:  "个人项目工具",
}

func main() {
	cmd.InitConf()
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
