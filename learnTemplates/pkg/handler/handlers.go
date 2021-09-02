package handler

import (
	"net/http"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/render"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/config"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/models"
)
// we can directly render tamplate for a partivular page


var Repo *Repository

type Repository struct{
	App *config.AppConfig
}
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App:a, 
	}
}
func NewHandlers(r *Repository){
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	// gives ip address
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w,"home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	// business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w,"about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}