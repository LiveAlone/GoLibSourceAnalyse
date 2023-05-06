package start

import (
	"go.uber.org/fx"
)

func MiniPro() {
	app := fx.New()
	app.Run()
}
