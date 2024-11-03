package service

import (
	"go-service-api/global"
	"go-service-api/model"
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

// 创建菜单
func (m *MenuService) Create(menu *model.Menu) error {
	if err := global.DB.Create(&menu).Error; err != nil {
		return err
	}
	return nil
}

// 更新菜单
func (m *MenuService) Update(menu *model.Menu) error {
	if err := global.DB.Save(&menu).Error; err != nil {
		return err
	}
	return nil
}

// 删除菜单
func (m *MenuService) Delete(id uint) error {
	if err := global.DB.Delete(&model.Menu{}, id).Error; err != nil {
		return err
	}
	return nil
}
