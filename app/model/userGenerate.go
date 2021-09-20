package model

import (
	"errors"
	"fmt"
	"gitee.com/itse/personal-work/app/schema"
	"gitee.com/itse/personal-work/app/util"
	"time"
)

// IPaper 生成接口 多态实现
type IPaper interface {
	GeneratePaper(username string, count int) error
}

// Primary 小学操作结构体
type Primary struct {
}

// GeneratePaper 小学生成
func (p *Primary) GeneratePaper(username string, count int) error {
	var ans schema.Paper
	for i := 0; i < count; i++ {
		var title string
		var Operand = randGenerator.Intn(4) + 2
		for j := 0; j < Operand-1; {
			if randGenerator.Intn(5) != 4 {
				res := BaseGenerateCommon()
				title = title + res
				j = j + 1
				continue
			}
			res := BaseGenerateWithBrackets()
			title = title + res
			j = j + 2
		}
		res := BaseGenerateFinal()
		title = title + res
		ans.Topic = append(ans.Topic, schema.Topic{
			ID:    i + 1,
			Title: title,
		})
	}
	ans.Name = fmt.Sprintf("%v-%v-%v-%v-%v-%v",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
	)
	err := util.SaveTxt(username, ans)
	if err != nil {
		return err
	}
	return nil
}

// Junior 初中操作结构体
type Junior struct {
}

// GeneratePaper 初中生成
func (j *Junior) GeneratePaper(username string, count int) error {
	var ans schema.Paper
	for i := 0; i < count; i++ {
		var title string
		var Operand = randGenerator.Intn(4) + 2
		var tmpNum = 0
		if Operand <= 2 {
			tmpNum = 1
		} else {
			tmpNum = randGenerator.Intn(Operand-2) + 1
		}
		var index []int
		for k := 0; k < tmpNum; {
			tmp := randGenerator.Intn(Operand - 1)
			if !util.IsContain(index, tmp) {
				index = append(index, tmp)
				k++
			}
		}
		for j := 0; j < Operand-1; {
			if util.IsContain(index, j) {
				if randGenerator.Intn(10)%2 == 0 {
					res := JuniorGenerateSquare()
					title = title + res
					j = j + 1
					continue
				}
				res := JuniorGeneratePower()
				title = title + res
				j = j + 1
				continue
			}
			if randGenerator.Intn(5) != 4 {
				res := BaseGenerateCommon()
				title = title + res
				j = j + 1
				continue
			}
			if !util.IsContain(index, j+1) {
				res := BaseGenerateWithBrackets()
				title = title + res
				j = j + 2
			}
		}
		res := BaseGenerateFinal()
		title = title + res
		ans.Topic = append(ans.Topic, schema.Topic{
			ID:    i + 1,
			Title: title,
		})
	}
	ans.Name = fmt.Sprintf("%v-%v-%v-%v-%v-%v",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
	)
	err := util.SaveTxt(username, ans)
	if err != nil {
		return err
	}
	return nil
}

// High 高中操作结构体
type High struct {
}

// GeneratePaper 高中生成
func (h *High) GeneratePaper(username string, count int) error {
	var ans schema.Paper
	for i := 0; i < count; i++ {
		var title string
		var Operand = randGenerator.Intn(4) + 2
		var tmpNum = 0
		if Operand <= 2 {
			tmpNum = 1
		} else {
			tmpNum = randGenerator.Intn(Operand-2) + 1
		}
		var index []int
		for k := 0; k < tmpNum; {
			tmp := randGenerator.Intn(Operand - 1)
			if !util.IsContain(index, tmp) {
				index = append(index, tmp)
				k++
			}
		}
		for j := 0; j < Operand-1; {
			if util.IsContain(index, j) {
				res := HighGenerateWithTrigonometric()
				title = title + res
				j = j + 1
				continue
			}
			if randGenerator.Intn(5) != 4 {
				res := BaseGenerateCommon()
				title = title + res
				j = j + 1
				continue
			}
			if !util.IsContain(index, j+1) {
				res := BaseGenerateWithBrackets()
				title = title + res
				j = j + 2
			}
		}
		res := BaseGenerateFinal()
		title = title + res
		ans.Topic = append(ans.Topic, schema.Topic{
			ID:    i + 1,
			Title: title,
		})
	}
	ans.Name = fmt.Sprintf("%v-%v-%v-%v-%v-%v",
		time.Now().Year(),
		int(time.Now().Month()),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
	)
	err := util.SaveTxt(username, ans)
	if err != nil {
		return err
	}
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
