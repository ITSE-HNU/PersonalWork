package dao

import (
	"gitee.com/itse/personal-work/app/entity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// UserDaoSet 注入 DI
var UserDaoSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

// UserDao users 表相关的数据库操作
type UserDao struct {
	DB *gorm.DB
}

// RoleResponse 角色返回结构体
type RoleResponse struct {
	ID   int
	Name string
}

// UserDaoResponse dao 层用户查询返回
type UserDaoResponse struct {
	Username string
	Password string
	RoleID   int
	Role     RoleResponse
}

// Query 根据username查询用户
func (u *UserDao) Query(username string) (*[]UserDaoResponse, error) {
	result := new([]UserDaoResponse)
	db := u.DB.Model(&entity.User{})

	db = db.Where("username = ?", username)

	err := db.Joins("Role").Find(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
