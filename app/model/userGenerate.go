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

type IPaper interface {
	GeneratePaper(count int) error
}

type Primary struct {
}

var baseOperator = []string{"+", "-", "*", "/", "()"}
var juniorOperator = []string{"(square)", "^2"}
var highOperator = []string{"tan", "sin", "cos"}
var finalOperator = "="

func (p *Primary) GeneratePaper(count int) error {
	var ans schema.Paper
	var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		var title string
		var Operand = randGenerator.Intn(4) + 2
		for j := 0; j < Operand-1; j++ {
			var tmp = randGenerator.Intn(4)
			var num = randGenerator.Intn(100) + 1
			fmt.Println(num)
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

type Junior struct {
}

func (j *Junior) GeneratePaper(count int) error {
	return nil
}

type High struct {
}

func (h *High) GeneratePaper(count int) error {
	return nil
}

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
