package common

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Logger *zap.Logger
)
