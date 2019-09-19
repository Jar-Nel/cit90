package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First, Last string
	Drinks map[int]string
}


func main() {
	s:=`{"First":"Joe","Last":"Schmoe","Drinks":{"1":"martini","2":"Gin","3":"Root Beer"}}`

	p1:=person{
		First: "Joe",
		Last: "Schmoe",
		Drinks: map[int]string{1: "martini",2:"Gin",3:"Root Beer"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	p2:=person{}

	err = json.Unmarshal([]byte(s), &p2)
	if err!=nil {
		panic(err)
	}
	fmt.Println(p2)

	var unm = map[string]interface{}{}
	err = json.Unmarshal([]byte(s), &unm)
	if err!=nil {
		panic(err)
	}
	fmt.Println(unm)

	fmt.Println(unm["First"])
	//var d map[int]string
	d:= unm["Drinks"].(map[string]interface{})
	fmt.Println(d["1"])
	fmt.Println(unm["Drinks"])
	fmt.Println(unm["Drinks"].(map[string]interface{})["2"])
}