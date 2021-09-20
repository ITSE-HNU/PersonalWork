package model

import "strconv"

// BaseGenerateWithBrackets 带有括号的题目
func BaseGenerateWithBrackets() string {
	var tmp = randGenerator.Intn(4)
	var num = randGenerator.Intn(100) + 1
	var temp = randGenerator.Intn(4)
	var tempNum = randGenerator.Intn(100) + 1
	return "(" + strconv.Itoa(num) + baseOperator[tmp] + strconv.Itoa(tempNum) + ")" + baseOperator[temp]
}

// BaseGenerateCommon 普通题目
func BaseGenerateCommon() string {
	var tmp = randGenerator.Intn(4)
	var num = randGenerator.Intn(100) + 1
	return strconv.Itoa(num) + baseOperator[tmp]
}

// BaseGenerateFinal 最后
func BaseGenerateFinal() string {
	return strconv.Itoa(randGenerator.Intn(100)+1) + finalOperator
}
