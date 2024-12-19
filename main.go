package main

import (
	"fmt"
	"go-service-api/config"
	"go-service-api/global"
	"go-service-api/initialize"
	"go-service-api/middleware"
	"go-service-api/router"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	initialize.Config()
	initialize.Mysql()
	initialize.Log()
	r := gin.New()
	r.Use(middleware.Logger(global.Logger, time.RFC3339, true), middleware.Recovery(global.Logger, true))
	r.Use(middleware.Cors())
	router.Router(r)
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	r.Run(addr)
}
