package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First, Last string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/encode",jsonEncode)
	http.HandleFunc("/decode",bar)
	//http.HandleFunc("/about",bar)
	http.ListenAndServe(":808", nil)
}

func jsonEncode(w http.ResponseWriter, r *http.Request){
	p1:=person{
		First: "Joe",
		Last: "Schmoe",
	}
	p2:=person{
		First: "James",
		Last: "Peachy",
	}

	people:=[]person{p1,p2}

	err:=json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Panicf("error %s",err)
	}

}

func foo(w http.ResponseWriter, r *http.Request){
	p1:=person{
		First: "Joe",
		Last: "Schmoe",
	}
	p2:=person{
		First: "James",
		Last: "Peachy",
	}

	people:=[]person{p1,p2}

	js, err := json.Marshal(p1)
	if err != nil {
		log.Panicf("error %s",err)
	}
	fmt.Fprintf(w,string(js))

	js, err = json.Marshal(people)
	if err != nil {
		log.Panicf("error %s",err)
	}	
	fmt.Fprintf(w,string(js))
	
	err=json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Panicf("error %s",err)
	}

	var p3 person
	json.Unmarshal(js, &p3)
	fmt.Fprintf(w,p3.First)
	//fmt.Println("ok",p3)

}

func bar(w http.ResponseWriter, r *http.Request){
	people2:=[]person{}
	r2, err:=http.Get("http://localhost:808/encode")
	if err!=nil {
		http.Error(w, "Cant read page", 400)
		return
	}
	defer r2.Body.Close();
	err=json.NewDecoder(r2.Body).Decode(&people2)
	if err!=nil{
		http.Error(w, "Couldn't decode",http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "<html><body><h1>%s</h1><h1>%s</h1></body></html>",people2[0].First, people2[0].Last)
}

