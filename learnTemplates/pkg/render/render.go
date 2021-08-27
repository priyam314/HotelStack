package render

import (
	"fmt"
	"net/http"
	"html/template"
	"path/filepath"
	"log"
	"bytes"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/config"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/models"
)

// Whole Story
// CreateTemplateCache() create templates and save it in cache
// CreateTemplateCache()
// ...1) save all *.page.tmpl in a variable
// ...2) looping over every page
// ......1) find the name of file
// ......2) allocate new template with name of file
// ......3) add any functions if connected to it 
// ......4) and then parse the file
// ......5) now check for all *.layout.tmpl files in directory
// ......6) if there is any layout template then create a new template 
// ......   using ParseGlob
// ......7) at last add the name of file to map with its value to ts variable
// RenderTemplate()
// ...1) it first instatiates CreateTemplateCache()
// ...2) now it has the map which has all the templates 
// ...3) now search for the page in map
// ...4) execute tha value found from map and write to buffer(here buf)
// ...5) now buf write to w which is http.ResponseWriter and then to html page

// FuncMap is the type of map defining mapping from names to function
// each function must have either a single return value or two return values
// of which the second has type error
// for now we will not talk lot about FuncMap and Funcs, but later in course
var function = template.FuncMap{

}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}
// it can render a template, as argment it takes ResponseWriter and tmpl file
// template.ParsedFiles parses the file
// parsedTemplate.Execute executes it and as an output it returns err
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){

	var tc map[string]*template.Template

	if app.UseCache{
		tc = app.AppCache
	}else{
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("Could not get template from template cache")
	}

	// Initiating a buffer
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	// its actually like *template.Template.Execute
	// Execute applies a parsed template to the specified data objet, and 
	// writes the output to w io.Writer, here to buf
	_ = t.Execute(buf, td)

	// WriteTo writes data to w until the buffer is drained or an error occurs
	_, err := buf.WriteTo(w)
	if err!=nil{
		fmt.Println("error writing template to browser",err)
	}
}

// Package template implements data-driven templates for 
// generating textual output. 
func CreateTemplateCache()(map[string]*template.Template,error){
	myCache := map[string]*template.Template{}

	// searchng for all the pages in a directory
	pages, err := filepath.Glob("./template/*.page.tmpl")
	if err!=nil{
		return myCache,err
	}

	for _,page := range pages{
		// Base returns the last element of path
		name := filepath.Base(page)

		// new allocates a new, undefined template with the given name
		// returns *Template
		// Funcs method defines a function that you want in that template 
		// "function" is basicallly template.FuncsMap{ inside it has map}
		// map[string]interface{}
		// ParseFile parses the certain page 
		// ts is of type *template.Template
		ts, err := template.New(name).Funcs(function).ParseFiles(page)
		if err!=nil{
			return myCache, err
		}

		// searching for *.layout.tmpl pages in template directory
		matches, err := filepath.Glob("./template/*.layout.tmpl")
		if err!=nil{
			return myCache,err
		}

		if len(matches)>0{
			// ParseGlob creates a new Template and parses the template 
			// definitons from files identified by the pattern
			// the returned template will have the (base) name and (parsed) 
			// contents of the first file matched by the pattern
			// ParseGlob is equivalent to calling ParseFiles with the list of 
			// files matched by the pattern
			ts, err = ts.ParseGlob("./template/*.layout.tmpl")
			if err!=nil{
				return myCache, err
			}
		}

		// adding name of file from *.page.tmpl as key
		// and the *template.Template of all the files which match the pattern
		myCache[name] = ts
	}
	return myCache,nil
}
