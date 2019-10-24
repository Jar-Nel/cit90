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
}

var tpl *template.Template

func init() {
	tpl=template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	log.Fatal(http.ListenAndServe(":808",nil))

}

func index(w http.ResponseWriter, r *http.Request){
	data:=pageData {
		Title: "Index",
		Heading: "Welcome to the site",
	}
	templateErr(w, tpl.ExecuteTemplate(w, "template.gohtml", data))
}

func about(w http.ResponseWriter, r *http.Request){
	data:=pageData {
		Title: "About",
		Heading: "About The Site",
	}
	templateErr(w, tpl.ExecuteTemplate(w, "template.gohtml", data))
	
}

func templateErr(w http.ResponseWriter, err error){
	if err!=nil{
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}
