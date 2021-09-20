package model

import "strconv"

// JuniorGenerateSquare 初中生成开方
func JuniorGenerateSquare() string {
	var tmp = randGenerator.Intn(4)
	var num = randGenerator.Intn(100) + 1
	return "(" + juniorOperator[0] + strconv.Itoa(num) + ")" + baseOperator[tmp]
}

// JuniorGeneratePower 初中生成平方
func JuniorGeneratePower() string {
	var tmp = randGenerator.Intn(4)
	var num = randGenerator.Intn(100) + 1
	return "(" + strconv.Itoa(num) + juniorOperator[1] + ")" + baseOperator[tmp]
}
