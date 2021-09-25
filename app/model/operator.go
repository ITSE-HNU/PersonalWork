// Package model 逻辑处理层
package model

// 操作符定义
var (
	baseOperator   = []string{"+", "-", "×", "÷", "()"}
	juniorOperator = []string{"√", "²"}
	highOperator   = []string{"tan", "sin", "cos"}
	finalOperator  = "="
)
