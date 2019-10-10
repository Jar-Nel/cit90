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
	//http.HandleFunc("/about",bar)
	http.ListenAndServe(":808", nil)
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


}

