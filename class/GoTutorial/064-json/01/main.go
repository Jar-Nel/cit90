package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type person struct {
	First, Last string
}

func main() {
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
		panic(err)
	}
	fmt.Println(string(js))

	js, err = json.Marshal(people)
	if err != nil {
		log.Panicf(err)
	}	
	fmt.Println(string(js))

}