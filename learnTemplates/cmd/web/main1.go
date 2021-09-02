package main

import (
	"net/http"
	"log"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/handler"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/config"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/render"
	"github.com/alexedwards/scs/v2"
	"time"
)
// here you have used handler.Home not handlers.Home because handler 
// files is handler package now, and all functions, methods are globally
// defined to be used across handler. so you go to root, and used function

// now its available in package
var app config.AppConfig
var sessions *scs.SessionManager

func main(){
	// change this to true, when in production
	app.InProduction = false

	sessions = scs.New()
	sessions.Lifetime = 24*time.Hour
	// if you want to forget the cookies once someone closes the window, 
	// set to false
	sessions.Cookie.Persist = true
	// how strict you wanna be to what site cookie applies to
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	// if you want to secure your cookie, then do true
	// we did false, because we use localhost, it ain't secure
	sessions.Cookie.Secure = app.InProduction

	app.Session = sessions
	tc, err := render.CreateTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
	}
	portNumber := ":8080"
	app.AppCache = tc
	app.UseCache = false
	
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	// not good to have routes in main func
	// with complexity, it would become cumbersome
	// http.HandleFunc("/",handler.Repo.Home)
	// http.HandleFunc("/about",handler.Repo.About)
	// http.ListenAndServe(":8080",nil)
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
// now you need to do go run cmd/web/*.go
// since other packages are imported and in no need to be executed again
