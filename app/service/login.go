// Package service 服务提供（入口）层
package service

import (
	"errors"
	"gitee.com/itse/personal-work/app/model"
	"gitee.com/itse/personal-work/app/util"
	"github.com/google/wire"
	"strings"
)

// LoginServiceSet Login DI
var LoginServiceSet = wire.NewSet(wire.Struct(new(LoginService), "*"))

// LoginService 登录结构体
type LoginService struct {
	LoginModel *model.LoginModel
}

// Login 登录入口
func (l *LoginService) Login() (*Current, error) {
	line := util.GetInput()
	res := strings.Split(line, " ")
	if len(res) != 2 {
		return nil, errors.New("input error")
	}
	result, err := l.LoginModel.Login(res[0], res[1])
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return &Current{
		Username: (*result).Username,
		RoleID:   (*result).RoleID,
		Role:     (*result).Role,
		IsLogin:  true,
	}, nil
}
