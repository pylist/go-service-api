package controller

import (
	"go-service-api/global"
	"go-service-api/model"
	"go-service-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type OptionController struct{}

func NewOptionController() *OptionController {
	return &OptionController{}
}

func (o *OptionController) GetOptions(c *gin.Context) {
	options := make([]model.Option, 0)
	global.DB.Find(&options)
	response.SuccessData(c, gin.H{"items": options})
}

func (o *OptionController) SetOptions(c *gin.Context) {
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	for key, value := range req {
		option := model.Option{
			Key:   key,
			Value: value,
		}
		global.DB.Save(&option)
	}
	response.Success(c)
}
