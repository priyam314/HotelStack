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
	CSRFToken   string
	Flash 		string
	Warning		string
	Error       string
}