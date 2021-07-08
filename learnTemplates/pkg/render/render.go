package render

import (
	"fmt"
	"net/http"
	"html/template"
)
// it can render a template, as argment it takes ResponseWriter and tmpl file
// template.ParsedFiles parses the file
// parsedTemplate.Execute executes it and as an output it returns err
func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parsedTemplate.Execute(w,nil)
	if err!=nil{
		fmt.Println("error parsing template:",err)
		return
	}
}