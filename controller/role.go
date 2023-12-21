package controller

import (
	"go-service-api/common"
	"go-service-api/model"
	"go-service-api/pkg/orm"
	"go-service-api/pkg/response"
	"go-service-api/service"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
}

func NewRoleController() *RoleController {
	return &RoleController{}
}

// 创建角色
func (r *RoleController) Create(c *gin.Context) {
	var req model.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	role := req
	if err := service.NewRoleService().Create(&role); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 删除角色
func (r *RoleController) Delete(c *gin.Context) {
	var req model.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err := service.NewRoleService().Delete(req.ID); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 更新角色
func (r *RoleController) Update(c *gin.Context) {
	var req model.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	if err := service.NewRoleService().Update(&req); err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c)
}

// 获取角色列表
func (r *RoleController) GetList(c *gin.Context) {
	var roles []model.Role
	tx := common.DB.Model(&model.Role{})
	orm.Paginate(c)(tx).Find(&roles)
	response.SuccessData(c, gin.H{
		"list": roles,
	})
}
