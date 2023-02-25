package controller

import (
	"net/http"

	"github.com/sanyogpatel-tecblic/To-Do/pkg/config"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	r.Context().Deadline()
}
