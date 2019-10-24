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

func main(){
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//Three page website using templates and passing in data
	http.HandleFunc("/home",fIndex) 
	http.HandleFunc("/about",fAbout) 
	http.HandleFunc("/shop",fShop) 
	http.HandleFunc("/contact",fContact)
	//http.HandleFunc("/struct",fShop)

	log.Fatal(http.ListenAndServe(":808", nil))
}

func fIndex(w http.ResponseWriter, r *http.Request){
	data:=pageData {
		Title: "Home",
		Heading: "Welcome to the site",
		Body: "This is the Home page of the site.",
	}
	templateErr(w, tpl.ExecuteTemplate(w, "site.gohtml", data))
}

func fAbout(w http.ResponseWriter, r *http.Request){
	data:=pageData {
		Title: "About",
		Heading: "About the site",
		Body: "This site served by Go.",
	}
	templateErr(w, tpl.ExecuteTemplate(w, "site.gohtml", data))
}

func fShop(w http.ResponseWriter, r *http.Request){
	data:=pageData {
		Title: "Shop",
		Heading: "Did you want to buy something?",
		Body: "Nothing to buy",
	}
	templateErr(w, tpl.ExecuteTemplate(w, "site.gohtml", data))
}

func fContact(w http.ResponseWriter, r *http.Request){
	data:=pageData {
		Title: "Contact",
		Heading: "Contact Us",
		Body: "This is where I would put contact information if I wanted to hear from you.",
	}
	templateErr(w, tpl.ExecuteTemplate(w, "site.gohtml", data))
}

func templateErr(w http.ResponseWriter, err error){
	if err!=nil{
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}


