package main

import (
	"fmt"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd/file_convert"
	"math"

	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmd.WordCmd)
	rootCmd.AddCommand(cmd.SqlCmd)
	rootCmd.AddCommand(file_convert.FileConvert)
}

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "utils",
	Long:  "个人项目工具",
}

func main() {
	//cmd.InitConf()
	//err := rootCmd.Execute()
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(math.MaxInt8)
}
