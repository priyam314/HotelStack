package main

import (
	"net/http"
	"fmt"
)

// when we open a new page, then we want to print on the console some message
// we going to do that using middleware
func WriteToConsole(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// print to the console
		fmt.Println("Hit the page")
		// what is that we have to do after printing to console
		// like opeing a new page or anything
		next.ServeHTTP(w, r)
	})
}