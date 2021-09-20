package model

import "strconv"

// HighGenerateWithTrigonometric 高中生成三角函数
func HighGenerateWithTrigonometric() string {
	var tmp = randGenerator.Intn(4)
	var num = randGenerator.Intn(13)
	var index = randGenerator.Intn(3)
	if index == 0 && num%6 == 0 {
		return highOperator[index+1] + strconv.Itoa(num*15) + baseOperator[tmp]
	}
	return highOperator[index] + strconv.Itoa(num*15) + baseOperator[tmp]
}
