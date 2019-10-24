package main

import (
	"io"
	"net/http"
	//"text/template"
	//"html/template"  //text/template with protections against code injection.
	"log"
	//"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/tob", tob)
	log.Fatal(http.ListenAndServe(":808",nil))
}

func index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<html>
		<head>
			<style>
				img {
					border-radius:10px;
					max-height:90vh;
					max-width:90vw;
				}
			</style>
		</head>
		<body>
		<p>Image here:</p>
			<img src="/tob" />
		</body>
		</html>
	`)
}

func tob(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "01-Boo.jpg")
}
