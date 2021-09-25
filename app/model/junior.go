// Package model 逻辑处理层
package model

import "strconv"

// JuniorGenerateSquare 初中生成开方
func JuniorGenerateSquare() string {
	var index = randGenerator.Intn(4)
	var operandNumber = randGenerator.Intn(100) + 1
	return juniorOperator[0] + strconv.Itoa(operandNumber) + baseOperator[index]
}

// JuniorGeneratePower 初中生成平方
func JuniorGeneratePower() string {
	var index = randGenerator.Intn(4)
	var operandNumber = randGenerator.Intn(100) + 1
	return strconv.Itoa(operandNumber) + juniorOperator[1] + baseOperator[index]
}
