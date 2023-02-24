package model

// "go.starlark.net/lib/time"

type Task struct {
	ID     int    `json:"id"`
	Tasks  string `json:"tasks"`
	UserID int    `json:"userid"`
}
