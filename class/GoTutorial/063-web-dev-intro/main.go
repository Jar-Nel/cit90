package main

import (
	"net/http"
	"io"
	"fmt"
)

//Ket to understand is type handler

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/about",bar)
	http.ListenAndServe(":808", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello World")
}

func bar(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Hello world</h1>")
}