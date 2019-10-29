package main

import (
	"net/http"
	"log"
)

func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.Handle("/file", http.FileServer(http.Dir("./static")))
	http.Handle("/file/", http.StripPrefix("/file", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":808",nil))
}


