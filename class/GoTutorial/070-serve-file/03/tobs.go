package main

import (
	"net/http"
	"log"
)

func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":808",nil))
}


