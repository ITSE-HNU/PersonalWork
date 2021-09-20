package schema

// LoginParams 登录参数
type LoginParams struct {
	Username string
	Password string
}

type LoginResponse struct {
	Username string
	RoleID   int
	Role     string
}
