package main

import (
	"github.com/priyam314/HotelStack/learnTemplates/pkg/config"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/handler"
	//"github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// using third pary library called "pat" for mux
func routes(app *config.AppConfig) http.Handler{
	/*
	// Pat 
	mux := pat.New()
	// takes two parameters
	// 1. pattern
	// 2. handler func
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	*/

	// Chi
	// MiddleWare
	// 		middleware is a loosely defined term for any software or service 
	// 		that enables the *part* of a system to communicate and manage data.
	//		it is the software that handles communication between components 
	//		and i/o, so dev can focus on specific purpose of their application.
	//
	// 		middleware helps to proccess some request that comes in your web 
	// 		application and forms some action on it
	mux := chi.NewRouter()

	// gracefully absorbs panics and prints the stackstrace
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about",handler.Repo.About)

	return mux
}
