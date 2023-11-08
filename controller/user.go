package controller

import (
	"go-service-api/model"
	"go-service-api/pkg"
	"go-service-api/pkg/response"
	"go-service-api/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserController) Create(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	user := model.User{}
	if err := service.NewUserService().Create(&user); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserController) Login(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	newUser, err := service.NewUserService().Verify(&user)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	token, err := pkg.NewToken().SigningToken(newUser.Username)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.SuccessData(c, map[string]interface{}{
		"token": token,
	})
}

func (u *UserController) Info(c *gin.Context) {
	var user model.User
	if err := service.NewUserService().Info(&user); err != nil {
		response.Failed(c, err.Error())
		return

	}
	response.SuccessData(c, map[string]interface{}{
		"username": user.Username,
	})
}

func (u *UserController) Logout(c *gin.Context) {
	response.Success(c)
}
