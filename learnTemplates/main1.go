package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func Homee(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"homee.page.tmpl")
}
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
	http.ListenAndServe(":8080",nil)
}
