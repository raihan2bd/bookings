package handler

import (
	"net/http"

	"github.com/raihan2bd/bookings/pkg/config"
	"github.com/raihan2bd/bookings/pkg/models"
	"github.com/raihan2bd/bookings/pkg/render"
)

// Repo the reposity use by the handlers
var Repo *Repoository

// Repository is the repository type
type Repoository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repoository {
	return &Repoository{
		App: a,
	}
}

// NewHandler sets the Repository for the handler
func NewHandler(r *Repoository) {
	Repo = r
}

// Home page handler
func (m *Repoository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About page handler
func (m *Repoository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
