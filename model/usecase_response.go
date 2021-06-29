package model

type Actor struct {
	ID          int64       `json:"id"`
	Username    string      `json:"username"`
	RoleLabel   string      `json:"role_label"`
	LastLoginAt interface{} `json:"last_login_at"`
}
