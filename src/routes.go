package main
import "net/http"
type Route struct{
	Name string
	Method string
	Pattern string
	HandlerFunc  http.HandleFunc
}

type Routes []Route

var routes = Routes{
	
}