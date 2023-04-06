package model

// "go.starlark.net/lib/time"

type Task struct {
	ID         int    `json:"id" validate:"required"`
	Tasks      string `json:"tasks" validate:"required"`
	UserID     int    `json:"userid"`
	Statuscode int    `json:"status"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// type User struct {
// 	ID       int    `json:"id" validate:"required"`
// 	Username string `json:"username" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// 	Email    string `json:"email" validate:"required"`
// }
