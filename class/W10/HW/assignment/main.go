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
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//Three page website using templates and passing in data
	http.HandleFunc("/string",fString) //template pass string
	http.HandleFunc("/slice",fSlice) //template pass slice
	http.HandleFunc("/struct",fStruct) //template pass struct
	log.Fatal(http.ListenAndServe(":808", nil))
}

func fString(w http.ResponseWriter, r *http.Request){
	//Pass a string to a template
	templateErr(w, tpl.ExecuteTemplate(w, "string.gohtml", "This string was passed to the template"))
}

func fSlice(w http.ResponseWriter, r *http.Request){
	//Pass a slice to a template
	xs:=[]string{"John","Jacob","Jingleheimer","Schmidt"}
	templateErr(w,tpl.ExecuteTemplate(w, "slice.gohtml", xs))
}

func fStruct(w http.ResponseWriter, r *http.Request){
	//pass a struct to a template
	xs:=[]string{"James","Jenny","Felix","René"}
	data:=struct {
		Name string
		DataFields []string
    }{
		Name: "Users",
		DataFields: xs,
	}
	templateErr(w,tpl.ExecuteTemplate(w, "struct.gohtml", data))
}

func templateErr(w http.ResponseWriter, err error){
	if err!=nil{
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}


