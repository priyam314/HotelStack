package handler

import (
	"net/http"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/render"
)
// we can directly render tamplate for a partivular page
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"about.page.tmpl")
}