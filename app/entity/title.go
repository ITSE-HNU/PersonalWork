package entity

import "gorm.io/gorm"

type Title struct {
	gorm.Model
	Content string `gorm:"column:content;type:varchar(255);not null"`
	RoleID  int    `gorm:"column:role_id;type:int;not null"`
}
