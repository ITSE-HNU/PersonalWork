// Package schema 参数封装（数据表映射）
package schema

// LoginParams 登录参数
type LoginParams struct {
	Username string
	Password string
}

// LoginResponse 登录返回参数
type LoginResponse struct {
	Username string
	RoleID   int
	Role     string
}
