// Package service 服务提供（入口）层
package service

// Current 定义当前用户
type Current struct {
	Username string
	RoleID   int
	Role     string
	IsLogin  bool
}
