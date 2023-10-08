package initialize

import (
	"go-service-api/config"

	"github.com/jinzhu/configor"
)

func Config() {
	var cfg = struct {
		DB     *config.DBConfig
		Server *config.ServerConfig
	}{}
	if err := configor.Load(&cfg, "config.toml"); err != nil {
		panic("配置文件加载失败, err:" + err.Error())
	}
	config.Server = cfg.Server
	config.DB = cfg.DB
}
