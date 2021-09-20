package model

import (
	"errors"
	"fmt"
	"gitee.com/itse/personal-work/app/dao"
	"gitee.com/itse/personal-work/app/schema"
	"github.com/google/wire"
)

var LoginModelSet = wire.NewSet(wire.Struct(new(LoginModel), "*"))

type LoginModel struct {
	User *dao.UserDao
}

func (l *LoginModel) Login(username, password string) (*schema.LoginResponse, error) {
	fmt.Println(username)
	fmt.Println(password)
	result, err := l.User.Query(username)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(*result))
	if len(*result) == 0 {
		return nil, errors.New("查无此人")
	}

	if (*result)[0].Password != password {
		return nil, errors.New("密码错误")
	}
	fmt.Println("登录成功")
	return &schema.LoginResponse{
		Username: username,
		RoleID:   (*result)[0].RoleID,
		Role:     (*result)[0].Role.Name,
	}, nil
}
