// Package util 工具封装
package util

import (
	"regexp"
)

// CheckMode 使用正则进行模式检查
func CheckMode(str string) (int, string, error) {
	matched, err := regexp.MatchString("切换为小学", str)
	if err != nil {
		return 0, "", err
	}
	if matched {
		return 1, "小学", nil
	}
	matched, err = regexp.MatchString("切换为初中", str)
	if err != nil {
		return 0, "", err
	}
	if matched {
		return 2, "初中", nil
	}
	matched, err = regexp.MatchString("切换为高中", str)
	if err != nil {
		return 0, "", err
	}
	if matched {
		return 3, "高中", nil
	}
	matched, err = regexp.MatchString("切换为", str)
	if err != nil {
		return 0, "", err
	}
	if matched {
		return 4, "", nil
	}
	return -1, "", nil
}
