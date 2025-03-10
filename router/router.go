package router

import (
	"go-service-api/controller"
	"go-service-api/middleware"
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
	api.GET("test", func(c *gin.Context) {
		c.SetCookie("test_cookie", "test", 3600, "/", ".dalao123.com", false, true)
		response.Success(c)
	})
	auth := api.Group("auth")
	auth.POST("register", controller.NewUserController().Register)
	auth.POST("login", controller.NewUserController().Login)
	auth.GET("codes", controller.NewUserController().Codes)
	auth.POST("logout", middleware.Jwt(), controller.NewUserController().Logout)

	api.GET("getMenus", controller.NewMenuController().GetMenus)

	option := api.Group("option", middleware.Jwt())
	option.GET("getOptions", controller.NewOptionController().GetOptions)
	option.POST("setOptions", controller.NewOptionController().SetOptions)

	super := api.Group("super", middleware.Jwt())
	{
		users := super.Group("user")
		users.POST("uploadAvatar", controller.NewUserController().UploadAvatar)
		users.GET("list", controller.NewUserController().List)
		users.GET("getUserList", controller.NewUserController().List)
		users.POST("createUser", controller.NewUserController().Create)
		users.POST("updateUser", controller.NewUserController().Update)
		users.POST("deleteUser", controller.NewUserController().Delete)

		menu := super.Group("menu")
		menu.POST("createMenu", controller.NewMenuController().Create)
		menu.GET("getMenuList", controller.NewMenuController().GetMenuList)
		menu.POST("updateMenu", controller.NewMenuController().Update)
		menu.POST("deleteMenu", controller.NewMenuController().Delete)

		role := super.Group("role")
		role.POST("setStatus", controller.NewRoleController().SetStatus)
		role.GET("getRoleList", controller.NewRoleController().GetList)
		role.POST("createRole", controller.NewRoleController().Create)
		role.POST("updateRole", controller.NewRoleController().Update)
		role.POST("deleteRole", controller.NewRoleController().Delete)

		dept := super.Group("dept")
		dept.POST("createDept", controller.NewDeptController().CreateDept)
		dept.POST("deleteDept", controller.NewDeptController().DeleteDept)
		dept.POST("updateDept", controller.NewDeptController().UpdateDept)
		dept.GET("getDeptList", controller.NewDeptController().GetDeptList)
	}

	user := api.Group("user", middleware.Jwt())
	user.GET("info", controller.NewUserController().Info)
	user.POST("uploadAvatar", controller.NewUserController().UploadAvatar)
	user.GET("getMenus", controller.NewMenuController().GetMenus)
	user.GET("getPermission", controller.NewUserController().GetPermission)
}
