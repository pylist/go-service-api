package model

type Role struct {
	Model
	Name     string `json:"name" gorm:"comment:角色名称"`
	Remark   string `json:"remark" gorm:"comment:备注"`
	ParentID uint   `json:"parentId" gorm:"comment:父角色ID"`
	MenuList []Menu `json:"menuList" gorm:"many2many:role_menu;"`
	Children []Role `json:"children" gorm:"-"`
}
