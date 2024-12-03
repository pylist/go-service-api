package model

type Option struct {
	Key   string `json:"key" gorm:"primaryKey;not null"`
	Value string `json:"value" gorm:"type:longtext"`
}
