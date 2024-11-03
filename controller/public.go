package controller

import (
	"go-service-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
}

func NewPublicController() *PublicController {
	return &PublicController{}
}

type SendCodeRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// 发送验证码
func (p *PublicController) SendCode(c *gin.Context) {
	var req SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}
