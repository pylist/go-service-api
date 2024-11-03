package controller

import (
	"go-service-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
}

func NewSystemController() SystemController {
	return SystemController{}
}

func (s *SystemController) GetDeptList(c *gin.Context) {
	response.Success(c)
}
