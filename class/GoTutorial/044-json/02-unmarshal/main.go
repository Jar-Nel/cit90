package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First, Last string
}

func main() {
	s:= `{"First":"Joe","Last":"Schmoe"}`;
	p1:=person{
		First: "Joe",
		Last: "Schmoe",
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
}