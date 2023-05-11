package start

import (
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func MiniPro() {
	app := fx.New(
		fx.Provide(
			zap.NewProduction,
		),
		fx.Invoke(func(shut fx.Shutdowner, log *zap.Logger) {
			log.Info("mini pro")
			err := shut.Shutdown()
			if err != nil {
				fmt.Println("shutdown error", err)
			}
		}),
	)
	app.Run()
}
