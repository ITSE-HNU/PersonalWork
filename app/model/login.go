// Package model 逻辑处理层
package model

import (
	"errors"
	"gitee.com/itse/personal-work/app/dao"
	"gitee.com/itse/personal-work/app/schema"
	"github.com/google/wire"
)

// LoginModelSet LoginModel 注入 DI
var LoginModelSet = wire.NewSet(wire.Struct(new(LoginModel), "*"))

// LoginModel 登录操作结构体
type LoginModel struct {
	User *dao.UserDao
}

// Login 登录操作
func (l *LoginModel) Login(username, password string) (*schema.LoginResponse, error) {
	result, err := l.User.Query(username)
	if err != nil {
		return nil, err
	}
	if len(*result) == 0 {
		return nil, errors.New("查无此人")
	}

	if (*result)[0].Password != password {
		return nil, errors.New("密码错误")
	}
	return &schema.LoginResponse{
		Username: username,
		RoleID:   (*result)[0].RoleID,
		Role:     (*result)[0].Role.Name,
	}, nil
}
