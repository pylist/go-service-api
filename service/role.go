package service

import (
	"go-service-api/common"
	"go-service-api/model"
)

type RoleService struct {
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

// Create 创建角色
func (r *RoleService) Create(role *model.Role) error {
	if err := common.DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除角色
func (r *RoleService) Delete(id uint) error {
	if err := common.DB.Delete(&model.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新角色
func (r *RoleService) Update(role *model.Role) error {
	if err := common.DB.Save(&role).Error; err != nil {
		return err
	}
	return nil
}
