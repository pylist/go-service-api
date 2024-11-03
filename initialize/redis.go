package initialize

import (
	"go-service-api/config"
	"go-service-api/global"

	"github.com/redis/go-redis/v9"
)

func Redis() {
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
}
