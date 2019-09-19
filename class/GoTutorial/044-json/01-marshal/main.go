package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First, Last string
}

func main() {
	p1:=person{
		First: "Joe",
		Last: "Schmoe",
	}

	js, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}