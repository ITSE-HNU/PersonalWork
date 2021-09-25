// Package util 工具封装
package util

// IsContain 去重 util
func IsContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
