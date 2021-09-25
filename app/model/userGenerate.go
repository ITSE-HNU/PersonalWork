// Package model 逻辑处理层
package model

import (
	"errors"
	"fmt"
	"gitee.com/itse/personal-work/app/dao"
	"gitee.com/itse/personal-work/app/schema"
	"gitee.com/itse/personal-work/app/util"
	"github.com/google/wire"
	"time"
)

// PaperModelSet PaperModel 注入 DI
var PaperModelSet = wire.NewSet(wire.Struct(new(PaperModel), "*"))

// PaperModel 试卷生成操作结构体
type PaperModel struct {
	TitleDao *dao.TitleDao
}

// IPaper 生成接口 多态实现
type IPaper interface {
	GeneratePaper(username string, count int) error
	JudgeContains(content string) (bool, error)
}

// Primary 小学操作结构体
type Primary struct {
	TitleDao *dao.TitleDao
}

// GeneratePaper 小学生成
func (p *Primary) GeneratePaper(username string, count int) error {
	var ans schema.Paper
	for i := 0; i < count; {
		var title string
		var operandCount = randGenerator.Intn(3) + 2
		for j := 0; j < operandCount-1; {
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
		judge, err := p.JudgeContains(title)
		if err != nil {
			return err
		}
		if !judge {
			if err := p.TitleDao.Create(dao.TitleDealParams{
				Content: title,
				RoleID:  1,
			}); err != nil {
				return err
			}
			ans.Topic = append(ans.Topic, schema.Topic{
				ID:    i + 1,
				Title: title,
			})
			i++
		}
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

// JudgeContains 小学查重
func (p *Primary) JudgeContains(content string) (bool, error) {
	res, err := p.TitleDao.Query(dao.TitleDealParams{
		Content: content,
		RoleID:  1,
	})
	if err != nil {
		return false, err
	}
	return len(*res) != 0, nil
}

// Junior 初中操作结构体
type Junior struct {
	TitleDao *dao.TitleDao
}

// GeneratePaper 初中生成
func (j *Junior) GeneratePaper(username string, count int) error {
	var ans schema.Paper
	for i := 0; i < count; {
		var title string
		var operandCount = randGenerator.Intn(3) + 2
		var specialCount = 0
		if operandCount <= 2 {
			specialCount = 1
		} else {
			specialCount = randGenerator.Intn(operandCount-2) + 1
		}
		var index []int
		for k := 0; k < specialCount; {
			tmpIndex := randGenerator.Intn(operandCount - 1)
			if !util.IsContain(index, tmpIndex) {
				index = append(index, tmpIndex)
				k++
			}
		}
		for j := 0; j < operandCount-1; {
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
		judge, err := j.JudgeContains(title)
		if err != nil {
			return err
		}
		if !judge {
			if err := j.TitleDao.Create(dao.TitleDealParams{
				Content: title,
				RoleID:  2,
			}); err != nil {
				return err
			}
			ans.Topic = append(ans.Topic, schema.Topic{
				ID:    i + 1,
				Title: title,
			})
			i++
		}
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

// JudgeContains 初中查重
func (j *Junior) JudgeContains(content string) (bool, error) {
	res, err := j.TitleDao.Query(dao.TitleDealParams{
		Content: content,
		RoleID:  1,
	})
	if err != nil {
		return false, err
	}
	return len(*res) != 0, nil
}

// High 高中操作结构体
type High struct {
	TitleDao *dao.TitleDao
}

// GeneratePaper 高中生成
func (h *High) GeneratePaper(username string, count int) error {
	var ans schema.Paper
	for i := 0; i < count; i++ {
		var title string
		var operandCount = randGenerator.Intn(3) + 2
		var tmpIndex = 0
		if operandCount <= 2 {
			tmpIndex = 1
		} else {
			tmpIndex = randGenerator.Intn(operandCount-2) + 1
		}
		var index []int
		for k := 0; k < tmpIndex; {
			tmp := randGenerator.Intn(operandCount - 1)
			if !util.IsContain(index, tmp) {
				index = append(index, tmp)
				k++
			}
		}
		for j := 0; j < operandCount-1; {
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
		judge, err := h.JudgeContains(title)
		if err != nil {
			return err
		}
		if !judge {
			if err := h.TitleDao.Create(dao.TitleDealParams{
				Content: title,
				RoleID:  3,
			}); err != nil {
				return err
			}
			ans.Topic = append(ans.Topic, schema.Topic{
				ID:    i + 1,
				Title: title,
			})
			i++
		}
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

// JudgeContains 高中查重
func (h *High) JudgeContains(content string) (bool, error) {
	res, err := h.TitleDao.Query(dao.TitleDealParams{
		Content: content,
		RoleID:  2,
	})
	if err != nil {
		return false, err
	}
	return len(*res) != 0, nil
}

// VerifyFactory 动态工厂
func (p *PaperModel) VerifyFactory(role int) (IPaper, error) {
	if role == 0 {
		return nil, errors.New("参数错误")
	}
	if role == 1 {
		return &Primary{
			TitleDao: p.TitleDao,
		}, nil
	}
	if role == 2 {
		return &Junior{
			TitleDao: p.TitleDao,
		}, nil
	}
	return &High{
		TitleDao: p.TitleDao,
	}, nil
}
