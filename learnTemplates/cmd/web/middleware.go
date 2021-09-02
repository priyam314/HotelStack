package main

import (
	"net/http"
	"fmt"
	"github.com/justinas/nosurf"
)

// when we open a new page, then we want to print on the console some message
// we going to do that using middleware
// all middleware will have somewhat similar format
func WriteToConsole(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// print to the console
		fmt.Println("Hit the page")
		// what is that we have to do after printing to console
		// like opeing a new page or anything
		next.ServeHTTP(w, r)
	})
}

// adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)
	// it sets cookie to know about csrf values
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad load and save the session on every request
func SessionLoad(next http.Handler) http.Handler{
	// provides middleware which automatically load and saves
	// data for current request, and communicate the session token
	// to and from the client in a cookie
	return sessions.LoadAndSave(next)
}
