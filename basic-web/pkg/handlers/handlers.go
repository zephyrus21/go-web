package handlers

import (
	"basic-web/pkg/config"
	"basic-web/pkg/models"
	"basic-web/pkg/renders"
	"net/http"
)

//? repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//! creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

//! sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["name"] = "Piyush"

	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
