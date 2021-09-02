package models

// suppose while rendering an about page or hoke page
// you want to do some business logic, and send those values to html page
// data is going to have some data type
// so here is TemplateData, which takes what you need
// now you dont know how any data you want to send over render template
// so you use map
// when developing forms , you need to have some security tokem
// its CSRFToken

type TemplateData struct{
	StringMap   map[string]string
	IntMap 		map[string]int
	FloatMap 	map[string]float32
	Data 		map[string]interface{}
	// cross site request forgery token, it nothig more but when you build a web
	// page there is a form on it, hidden field in that form which contains 
	// long random sequence of numbers, it changes everytime when you go o it
	// it seems like suitable place of middleware
	CSRFToken   String
	Flash 		string
	Warning		string
	Error       string
}