package main

import (
	"net/http"
)
// we can directly render tamplate for a partivular page
func Home(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w,"home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w,"about.page.tmpl")
}