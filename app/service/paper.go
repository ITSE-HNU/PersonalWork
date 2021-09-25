// Package service 服务提供（入口）层
package service

import (
	"gitee.com/itse/personal-work/app/model"
	"github.com/google/wire"
)

// PaperServiceSet PaperService 注入 DI
var PaperServiceSet = wire.NewSet(wire.Struct(new(PaperService), "*"))

// PaperService 试卷生成入口
type PaperService struct {
	PaperModel *model.PaperModel
}

// GeneratePaper 根据用户 role 生成试卷 入口
func (p *PaperService) GeneratePaper(username string, role int, count int) error {
	verify, err := p.PaperModel.VerifyFactory(role)
	if err != nil {
		return err
	}
	err = verify.GeneratePaper(username, count)
	if err != nil {
		return err
	}
	return nil
}
