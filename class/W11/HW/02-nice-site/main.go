package main

import (
	"net/http"
	//"text/template"
	"html/template"  //text/template with protections against code injection.
	"log"
)

type pageData struct {
	Title string
	Heading string
	Body string
}

var tpl *template.Template

func init(){
	//parseglob parses everything in a location
	tpl=template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.Handle("/file/", http.StripPrefix("/file", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/home",fIndex) 

	log.Fatal(http.ListenAndServe(":808",nil))
}


func fIndex(w http.ResponseWriter, r *http.Request){
	data:=pageData {
		Title: "Home",
		Heading: "Welcome to the site",
		Body: "This is the Home page of the site.",
	}
	templateErr(w, tpl.ExecuteTemplate(w, "site.gohtml", data))
}


func templateErr(w http.ResponseWriter, err error){
	if err!=nil{
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}