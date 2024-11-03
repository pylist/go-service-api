package initialize

import (
	"go-service-api/global"

	"go.uber.org/zap"
)

func Log() {
	global.Logger, _ = zap.NewProduction()
}
