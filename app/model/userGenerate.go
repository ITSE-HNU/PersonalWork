package model

import (
	"errors"
	"fmt"
	"gitee.com/itse/personal-work/app/schema"
	"gitee.com/itse/personal-work/app/util"
	"math/rand"
	"strconv"
	"time"
)

// IPaper 生成接口 多态实现
type IPaper interface {
	GeneratePaper(count int) error
}

// Primary 小学操作结构体
type Primary struct {
}

// GeneratePaper 小学生成
func (p *Primary) GeneratePaper(count int) error {
	var ans schema.Paper
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		var title string
		var Operand = randGenerator.Intn(4) + 2
		for j := 0; j < Operand-1; j++ {
			var tmp = randGenerator.Intn(4)
			var num = randGenerator.Intn(100) + 1
			title = title + strconv.Itoa(num) + baseOperator[tmp]
		}
		var num = randGenerator.Intn(100) + 1
		title = title + strconv.Itoa(num) + finalOperator
		ans.Topic = append(ans.Topic, schema.Topic{Title: title, ID: i + 1})
	}
	ans.Name = fmt.Sprintf("%v-%v-%v-%v-%v-%v",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
	)
	err := util.SaveTxt(ans)
	if err != nil {
		return err
	}
	return nil
}

// Junior 初中操作结构体
type Junior struct {
}

// GeneratePaper 初中生成
func (j *Junior) GeneratePaper(count int) error {
	return nil
}

// High 高中操作结构体
type High struct {
}

// GeneratePaper 高中生成
func (h *High) GeneratePaper(count int) error {
	return nil
}

// VerifyFactory 动态工厂
func VerifyFactory(role int) (IPaper, error) {
	if role == 0 {
		return nil, errors.New("参数错误")
	}
	if role == 1 {
		return &Primary{}, nil
	}
	if role == 2 {
		return &Junior{}, nil
	}
	return &High{}, nil
}
