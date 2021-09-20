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

// Query 根据username查询用户
func (u *UserDao) Query(username string) (*[]entity.User, error) {
	result := new([]entity.User)
	db := u.DB.Model(&entity.User{})

	db = db.Where("username = ?", username)

	err := db.Find(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
