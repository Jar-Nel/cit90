package main

import (
	"net/http"
	"text/template"
	"log"
)

var tpl *template.Template

func init(){
	//parseglob parses everything in a location
	tpl=template.Must(template.ParseGlob("./template/*.gohtml"))
}

func main(){
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/about",about)
	http.HandleFunc("/test",test)
	log.Fatal(http.ListenAndServe(":808", nil))
}

func about(w http.ResponseWriter, r *http.Request){
	data:=struct  {
		FirstName string
		LastName string
	}{
		FirstName:"John",
		LastName:"Doe",
	}
	err:=tpl.ExecuteTemplate(w, "about.gohtml", data)
	if err!=nil {
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}

func test(w http.ResponseWriter, r *http.Request){
	xs:=[]string{"John","Jacob","Smith"}
	data:=struct {
		DataFields []string
    }{
		DataFields: xs,
	}
	err:=tpl.ExecuteTemplate(w, "test.gohtml", data)
	if err!=nil {
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}


