package initialize

import (
	"go-service-api/common"

	"go.uber.org/zap"
)

func Log() {
	common.Logger, _ = zap.NewProduction()
}
