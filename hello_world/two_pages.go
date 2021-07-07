package main

import (
	"fmt"
	"net/http"
	"errors"
)
const portNumber = ":8080"
var date = 12

// Home is handler for home page
// H in Home is capital for a reason
// this means this function can be used across package and is public
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, fmt.Sprintf("this is the home page and it's %d",date))
}

// About is handler for about page
// About is public method
func About(w http.ResponseWriter, r *http.Request){
	sum := add(2,2)
	fmt.Fprintf(w, fmt.Sprintf("this is about page and 2 + 2 is %d",sum))
}

// Divide is handler for divide page
func Divide(w http.ResponseWriter, r *http.Request){
	val,err := divide(100.0,0)
	if err!=nil{
		fmt.Fprintf(w,"Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f is divided by %f gives %f",100.0,0,val))
}

// divide function two divide two float
// outputs value and error
func divide(x,y float32)(float32,error){
	if y==0{
		err := errors.New("Cannot divide by zero")
		return 0,err
	}
	return x/y,nil
}

// add function add two values and return a integer
// add value is private and can only be used inside this script
// non capital means private and cannot be used across package
func add(x,y int)int{
	return x+y
}
func main(){
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	http.HandleFunc("/divide",Divide)
	http.ListenAndServe(portNumber,nil)
}
