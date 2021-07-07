package  main

import (
	"net/http"
	"fmt"
)

// ServeMux(or just Mux) - is HTTP request multiplexer.
// it matches the url of each incoming request against a list of registered
// patterns and calls the handler for the pattern that most closely matches
// the URL.
// mux is used for routing

//       ServeMux(type)
//              |
// Methods __________
//         |        |
//     HandleFunc  Handle

// for e.g you want to execute a function named "hello" whenever a user visits
// www.example.com and a function named info whenever a user visits 
// www.example.com/info. you can use HandleFunc or Handle method from ServeMux
// multipplexer in golang is like multiplexer in hardware which multiply some
// inputs into some outputs

// http.HandleFunc method works as my given simple multiplexer

// ListenAndServe - starts an HTTP server with a given address and handler
// the handler is usually nil, which means to use DefaultServeMux
// Handle and HandleFunc add handlers to DefaultServeMux

// DefaultServeMux is just a predefined http.ServeMux

// http.Handle calls DefaultServeMux internally

// http.ListenAndServe returns an error
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello world")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("Bytes written:" , n)
	})

	http.ListenAndServe(":8080",nil)
}
