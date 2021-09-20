package service

import (
	"gitee.com/itse/personal-work/app/model"
	"gitee.com/itse/personal-work/app/util"
	"strconv"
)

func GeneratePaper(role int) error {
	line := util.GetInput()
	count, err := strconv.Atoi(line)
	if err != nil {
		return err
	}
	verify, err := model.VerifyFactory(role)
	if err != nil {
		return err
	}
	err = verify.GeneratePaper(count)
	if err != nil {
		return err
	}
	return nil
}
