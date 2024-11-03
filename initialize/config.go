package initialize

import (
	"go-service-api/config"

	"github.com/jinzhu/configor"
)

func Config() {
	cfg := struct {
		DB     *config.DBConfig
		Server *config.ServerConfig
		Redis  *config.RedisConfig
	}{}
	if err := configor.Load(&cfg, "config.toml"); err != nil {
		panic("配置文件加载失败, err:" + err.Error())
	}
	config.Server = cfg.Server
	config.DB = cfg.DB
	config.Redis = cfg.Redis
}
