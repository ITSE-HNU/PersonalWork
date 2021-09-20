package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
	RoleID   int    `gorm:"not null"`
	Role     Role
}
