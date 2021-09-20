package model

// 操作符定义
var (
	baseOperator   = []string{"+", "-", "*", "/", "()"}
	juniorOperator = []string{"(square)", "^2"}
	highOperator   = []string{"tan", "sin", "cos"}
	finalOperator  = "="
)
