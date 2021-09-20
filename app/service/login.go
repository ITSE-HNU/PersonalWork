package service

import (
	"errors"
	"gitee.com/itse/personal-work/app/model"
	"gitee.com/itse/personal-work/app/util"
	"github.com/google/wire"
	"strings"
)

var LoginModelSet = wire.NewSet(wire.Struct(new(LoginService), "*"))

type LoginService struct {
	LoginModel *model.LoginModel
}

func (l *LoginService) Login() (*Current, error) {
	line := util.GetInput()
	res := strings.Split(line, " ")
	if len(res) != 2 {
		return nil, errors.New("input error")
	}
	result, err := l.LoginModel.Login(res[0], res[1])
	if err != nil {
		return nil, err
	}

	return &Current{
		Username: (*result).Username,
		RoleID:   (*result).RoleID,
		Role:     (*result).Role,
		IsLogin:  true,
	}, nil
}
