package service

import (
	"gitee.com/itse/personal-work/app/model"
)

// GeneratePaper 根据用户 role 生成试卷 入口
func GeneratePaper(username string, role int, count int) error {
	verify, err := model.VerifyFactory(role)
	if err != nil {
		return err
	}
	err = verify.GeneratePaper(username, count)
	if err != nil {
		return err
	}
	return nil
}
