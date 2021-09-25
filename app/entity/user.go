// Package entity 实体（数据表）定义层
package entity

import "gorm.io/gorm"

// User 用户信息表
type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
	RoleID   int    `gorm:"not null"`
	Role     Role   `gorm:"foreignKey:RoleID"`
}
