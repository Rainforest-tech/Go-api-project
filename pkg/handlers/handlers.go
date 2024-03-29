package handlers

import (
	"net/http"

	"github.com/Rainforest-tech/Go-api-project/pkg/config"
	"github.com/Rainforest-tech/Go-api-project/pkg/models"
	"github.com/Rainforest-tech/Go-api-project/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
