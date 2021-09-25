// Package entity 实体（数据表）定义层
package entity

import "gorm.io/gorm"

// Title 题目存储表 用于查重
type Title struct {
	gorm.Model
	Content string `gorm:"column:content;type:varchar(255);not null"`
	RoleID  int    `gorm:"column:role_id;type:int;not null"`
}
