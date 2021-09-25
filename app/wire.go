//go:generate wire
//+build wireinject

// Package app 应用封装层
package app

import (
	"gitee.com/itse/personal-work/app/dao"
	"gitee.com/itse/personal-work/app/model"
	"gitee.com/itse/personal-work/app/service"
	"github.com/google/wire"
)

// BuildInjector 依赖注入
func BuildInjector() (*App, func(), error) {
	wire.Build(
		InitGorm,
		service.WireServiceSet,
		model.WireModelSet,
		dao.WireDaoSet,
		WireAppSet,
	)

	return nil, nil, nil
}
