package model

import "time"

type Role struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Status      int       `json:"status" gorm:"comment:角色状态"` // 0 正常状态， 1删除
	Name        string    `json:"name" gorm:"comment:角色名称"`
	Remark      string    `json:"remark" gorm:"comment:备注"`
	ParentID    uint      `json:"parentId" gorm:"comment:父角色ID"`
	MenuList    []Menu    `json:"menuList" gorm:"many2many:role_menu;"`
	Children    []Role    `json:"children" gorm:"-"`
	OrderNo     int       `json:"orderNo" gorm:"comment:排序"`
	Description string    `json:"description" gorm:"comment:描述"`
}
