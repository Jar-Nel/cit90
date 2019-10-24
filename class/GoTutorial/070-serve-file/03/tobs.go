package main

import (
	"net/http"
	//"text/template"
	//"html/template"  //text/template with protections against code injection.
	"log"
	//"os"
)

func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":808",nil))
}


