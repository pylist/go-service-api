package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	RDB    *redis.Client
	DB     *gorm.DB
	Logger *zap.Logger
)
