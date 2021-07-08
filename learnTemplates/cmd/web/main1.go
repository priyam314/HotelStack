package main

import (
	"net/http"
	"github.com/priyam314/HotelStack/learnTemplates/pkg/handler"
)
// here you have used handler.Home not handlers.Home because handler 
// files is handler package now, and all functions, methods are globally
// defined to be used across handler. so you go to root, and used function

func main(){
	http.HandleFunc("/",handler.Home)
	http.HandleFunc("/about",handler.About)
	http.ListenAndServe(":8080",nil)
}
// now you need to do go run cmd/web/*.go
// since other packages are imported and in no need to be executed again
