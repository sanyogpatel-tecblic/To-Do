package handlers

import (
	driver "github.com/sanyogpatel-tecblic/To-Do/Driver"
	"github.com/sanyogpatel-tecblic/To-Do/config"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
	}
}
