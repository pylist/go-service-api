package model

import "gorm.io/datatypes"

type Dept struct {
	Model
	ParentID    uint           `json:"parentId" gorm:"comment:上级部门ID"`
	Name        string         `json:"name" gorm:"comment:部门名称"`
	Description string         `json:"description" gorm:"comment:描述"`
	Status      int            `json:"status" gorm:"comment:部门状态"`
	Sort        int            `json:"sort" gorm:"comment:排序"`
	Leader      string         `json:"leader" gorm:"comment:负责人"`
	Phone       string         `json:"phone" gorm:"comment:联系电话"`
	Email       string         `json:"email" gorm:"comment:邮箱"`
	Remark      string         `json:"remark" gorm:"comment:备注"`
	Meta        datatypes.JSON `json:"meta" gorm:"comment:部门信息"`
	Children    []*Dept        `json:"children" gorm:"-"`
}
