package main

import (
	"fmt"
	"net/http"
	"html/template"
)
// we can directly render tamplate for a partivular page
func Homee(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"homee.page.tmpl")
}
func Aboutt(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"about.page.tmpl")
}
// this function is first of all private
// it can render a template, as argment it takes ResponseWriter and tmpl file
// template.ParsedFiles parses the file
// parsedTemplate.Execute executes it and as an output it returns err
func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parsedTemplate.Execute(w,nil)
	if err!=nil{
		fmt.Println("error parsing template:",err)
		return
	}
}
func main(){
	http.HandleFunc("/",Homee)
	http.HandleFunc("/about",Aboutt)
	http.ListenAndServe(":8080",nil)
}
