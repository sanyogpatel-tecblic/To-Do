package model

// "go.starlark.net/lib/time"

type Task struct {
	ID     int    `json:"id"`
	Tasks  string `json:"tasks"`
	UserID int    `json:"userid"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
