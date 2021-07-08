package main

import (
	"net/http"
	"log"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/handler"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/config"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/render"
)
// here you have used handler.Home not handlers.Home because handler 
// files is handler package now, and all functions, methods are globally
// defined to be used across handler. so you go to root, and used function

func main(){
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
	}

	app.AppCache = tc
	app.UseCache = false
	
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/",handler.Repo.Home)
	http.HandleFunc("/about",handler.Repo.About)
	http.ListenAndServe(":8080",nil)
}
// now you need to do go run cmd/web/*.go
// since other packages are imported and in no need to be executed again
