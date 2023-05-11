package main

import (
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/cmd"
	"github.com/LiveAlone/GoLibSourceAnalyse/utils/common"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Version = "1.0.0"

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
	fx.New(
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			}),
		fx.Provide(
			// 日志注册
			zap.NewNop,
		),
		fx.Invoke(func(shut fx.Shutdowner, log *zap.Logger) {
			common.InitConf()
			err := rootCmd.Execute()
			if err != nil {
				log.Error("rootCmd.Execute error", zap.Error(err))
			}
			err = shut.Shutdown()
			if err != nil {
				log.Error("shutdown error", zap.Error(err))
			}
		}),
	).Run()
}
