// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"gitee.com/itse/personal-work/app/dao"
	"gitee.com/itse/personal-work/app/model"
	"gitee.com/itse/personal-work/app/service"
)

// Injectors from wire.go:

// BuildInjector 依赖注入
func BuildInjector() (*App, func(), error) {
	db, err := InitGorm()
	if err != nil {
		return nil, nil, err
	}
	userDao := &dao.UserDao{
		DB: db,
	}
	loginModel := &model.LoginModel{
		User: userDao,
	}
	loginService := &service.LoginService{
		LoginModel: loginModel,
	}
	titleDao := &dao.TitleDao{
		DB: db,
	}
	paperModel := &model.PaperModel{
		TitleDao: titleDao,
	}
	paperService := &service.PaperService{
		PaperModel: paperModel,
	}
	app := &App{
		DB:           db,
		Login:        loginService,
		PaperService: paperService,
	}
	return app, func() {
	}, nil
}
