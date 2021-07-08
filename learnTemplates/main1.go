package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	http.ListenAndServe(":8080",nil)
}
