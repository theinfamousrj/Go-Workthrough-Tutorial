package main

import ("fmt"
		"net/http")
	
//port for http	
var port string

//generic handler handles all calls
func genericHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Oh, hey there "+r.URL.Path[1:]+"... I was looking for web.")
}

//single handler to handle all /web calls
func webHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, "+r.URL.Path[1:]+"!")
}

func main() {
	//standard port 8080
	port = ":8080"
	
	//just a little message to let
	//you know things are running
	fmt.Printf("Now serving Hello, web! at 127.0.0.1%v\n", port)
	
	//telling the server which handler
	//handles all /web calls
    http.HandleFunc("/web", webHandler)
    
    //telling the server which handler
	//handles all calls except /web at this point
    http.HandleFunc("/", genericHandler)
    
    //listening on port 8080 to serve pages
    http.ListenAndServe(port, nil) 
}

//Typing http://127.0.0.1:8080/web into your browser
//will produce a page with the greeting 'Hello, web!'