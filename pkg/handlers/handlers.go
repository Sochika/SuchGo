package handlers

import (
	"net/http"

	"github.com/Sochika/SuchGo/pkg/config"
	"github.com/Sochika/SuchGo/pkg/render"
)

var Repo *Respository

type Respository struct {
	App *config.SystemConfig
}

// Create New Repo
func NewRepo(a *config.SystemConfig) *Respository {
	return &Respository{
		App: a,
	}
}

// New Handlers which set repo for the handler
func NewHandlers(r *Respository) {
	Repo = r
}
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
