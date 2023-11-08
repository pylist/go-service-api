package router

import (
	"go-service-api/controller"
	"go-service-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("api")
	api.POST("test", func(c *gin.Context) {
		response.Success(c)
	})
	super := api.Group("super")
	users := super.Group("users")
	users.POST("create", controller.NewUserController().Create)
}
