package router

import (
	"go-service-api/controller"
	"go-service-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Static("/api/static", "./static")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("api")
	api.POST("test", func(c *gin.Context) {
		response.Success(c)
	})
	api.POST("login", controller.NewUserController().Login)
	api.GET("getMenus", controller.NewMenuController().GetMenus)
	super := api.Group("super")
	{
		users := super.Group("users")
		users.POST("create", controller.NewUserController().Create)
		users.POST("delete", controller.NewUserController().Delete)
		users.POST("update", controller.NewUserController().Update)
		users.POST("uploadAvatar", controller.NewUserController().UploadAvatar)
		users.GET("list", controller.NewUserController().List)

		menus := super.Group("menus")
		menus.POST("create", controller.NewMenuController().Create)
		menus.POST("update", controller.NewMenuController().Update)
		menus.GET("list", controller.NewMenuController().GetList)
		menus.POST("delete", controller.NewMenuController().Delete)

		roles := super.Group("roles")
		roles.POST("create", controller.NewRoleController().Create)
		roles.POST("delete", controller.NewRoleController().Delete)
		roles.POST("update", controller.NewRoleController().Update)
		roles.GET("list", controller.NewRoleController().GetList)
	}

	user := api.Group("user")
	user.POST("uploadAvatar", controller.NewUserController().UploadAvatar)
	user.GET("getMenus", controller.NewMenuController().GetMenus)

}
