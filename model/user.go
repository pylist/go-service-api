package model

import "gorm.io/datatypes"

type User struct {
	Model
	UID      string `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Location string `json:"location"`
	RoleID   string `json:"roleId"`
	Meta     datatypes.JSON	`json:"meta"`
}
