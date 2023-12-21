package service

import (
	"go-service-api/common"
	"go-service-api/model"
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

// 创建菜单
func (m *MenuService) Create(menu *model.Menu) error {
	if err := common.DB.Create(&menu).Error; err != nil {
		return err
	}
	return nil
}

// 更新菜单
func (m *MenuService) Update(menu *model.Menu) error {
	if err := common.DB.Debug().Save(&menu).Error; err != nil {
		return err
	}
	return nil
}

// 删除菜单
func (m *MenuService) Delete(id uint) error {
	if err := common.DB.Delete(&model.Menu{}, id).Error; err != nil {
		return err
	}
	return nil
}
