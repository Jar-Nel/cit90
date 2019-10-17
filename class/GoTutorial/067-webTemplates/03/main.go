package main

import (
	"net/http"
	//"text/template"
	"html/template"  //text/template with protections against code injection.
	"log"
)

var tpl *template.Template

func init(){
	//parseglob parses everything in a location
	tpl=template.Must(template.ParseGlob("./template/*.gohtml"))
}

func main(){
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/about",about) //template pass string
	http.HandleFunc("/slice",slice) //template pass slice
	http.HandleFunc("/test",test) //template pass struct
	log.Fatal(http.ListenAndServe(":808", nil))
}

func about(w http.ResponseWriter, r *http.Request){

	err:=tpl.ExecuteTemplate(w, "about.gohtml", "Passed Name")
	if err!=nil {
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}

func slice(w http.ResponseWriter, r *http.Request){
	xs:=[]string{"John","<script>alert('test')</script>Jacob","Smith"}
	templateErr(w,tpl.ExecuteTemplate(w, "slice.gohtml", xs))
}

func test(w http.ResponseWriter, r *http.Request){
	xs:=[]string{"John","Jacob","Smith"}
	data:=struct {
		Name string
		DataFields []string
    }{
		Name: "Users",
		DataFields: xs,
	}
	templateErr(w,tpl.ExecuteTemplate(w, "test.gohtml", data))
}

func templateErr(w http.ResponseWriter, err error){
	if err!=nil{
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}


