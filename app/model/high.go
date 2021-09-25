// Package model 逻辑处理层
package model

import "strconv"

// HighGenerateWithTrigonometric 高中生成三角函数
func HighGenerateWithTrigonometric() string {
	var baseOperatorIndex = randGenerator.Intn(4)
	var operandNumber = randGenerator.Intn(13) * 15
	var trigonometricOperatorIndex = randGenerator.Intn(3)
	if trigonometricOperatorIndex == 0 && operandNumber%6 == 0 {
		return highOperator[trigonometricOperatorIndex+1] + strconv.Itoa(operandNumber) + baseOperator[baseOperatorIndex]
	}
	return highOperator[trigonometricOperatorIndex] + strconv.Itoa(operandNumber) + baseOperator[baseOperatorIndex]
}
