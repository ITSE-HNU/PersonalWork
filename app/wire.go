//go:generate wire
//+build wire inject

package app

import (
	"gitee.com/itse/personal-work/app/config"
	"gitee.com/itse/personal-work/app/dao"
	"gitee.com/itse/personal-work/app/model"
	"gitee.com/itse/personal-work/app/service"
	"github.com/google/wire"
)

func BuildInjector() (*App, func(), error) {
	wire.Build(
		config.InitConfig,
		InitGorm,
		service.WireServiceSet,
		model.WireModelSet,
		dao.WireDaoSet,
		WireAppSet,
	)

	return nil, nil, nil
}
