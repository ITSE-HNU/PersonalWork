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
