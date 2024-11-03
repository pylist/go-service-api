package controller

import (
	"go-service-api/global"
	"go-service-api/model"
	"go-service-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type DeptController struct {
}

func NewDeptController() *DeptController {
	return &DeptController{}
}

// 新增部门
func (d *DeptController) CreateDept(c *gin.Context) {
	var req model.Dept
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	if err := global.DB.Create(&req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c)
}

// 删除部门
func (d *DeptController) DeleteDept(c *gin.Context) {
	var req model.Dept
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	if err := global.DB.Delete(&req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c)
}

// 更新部门
func (d *DeptController) UpdateDept(c *gin.Context) {
	var req model.Dept
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	if err := global.DB.Debug().Save(&req).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c)
}

// 获取部门列表
func (d *DeptController) GetDeptList(c *gin.Context) {
	var list []model.Dept
	if err := global.DB.Find(&list).Error; err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	// var result []model.Dept
	// for _, v := range list {
	// 	if v.ParentID == 0 {
	// 		result = append(result, v)
	// 	}
	// }
	// for i, _ := range result {
	// 	for j, _ := range list {
	// 		if result[i].ID == list[j].ParentID {
	// 			result[i].Children = append(result[i].Children, &list[j])
	// 		}
	// 	}
	// }
	response.SuccessData(c, gin.H{"list": list})
}
