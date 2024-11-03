package model

import "gorm.io/datatypes"

type User struct {
	Model
	Status   int            `json:"status" gorm:"type:tinyint;comment:用户状态"` // 0 正常状态， 1删除
	UID      string         `json:"uid" gorm:"uniqueIndex;not null;comment:用户唯一标识"`
	Username string         `json:"username" gorm:"uniqueIndex;not null;comment:用户登录名"`
	Password string         `json:"password" gorm:"not null;comment:用户登录密码"`
	Location string         `json:"location" gorm:"comment:用户位置"`
	NickName string         `json:"nickName" gorm:"comment:用户昵称"`
	Avatar   string         `json:"avatar" gorm:"comment:用户头像"`
	Email    string         `json:"email" gorm:"comment:用户邮箱"`
	Phone    string         `json:"phone" gorm:"comment:用户手机号码"`
	RoleID   uint           `json:"roleId" gorm:"comment:用户角色ID"`
	Role     *Role          `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	DeptID   uint           `json:"deptId" gorm:"comment:用户部门ID"`
	Dept     *Dept          `json:"dept" gorm:"foreignKey:DeptID;references:ID"`
	Meta     datatypes.JSON `json:"meta" gorm:"comment:用户个人信息"`
	Version  int            `json:"version" gorm:"comment:乐观锁"`
}
