package model

// "go.starlark.net/lib/time"

type Task struct {
	ID         int    `json:"id" validate:"required"`
	Tasks      string `json:"tasks" validate:"required"`
	UserID     int    `json:"userid"`
	Statuscode int    `json:"status"`
}

type User struct {
	ID         int    `json:"id"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Statuscode int    `json:"status"`
	Response   string `json:"response"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
