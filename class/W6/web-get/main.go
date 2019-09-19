package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type movie struct {
	Title string `json:"Title"`
	Director string `json:"Director"`
}

const key = "73626b1f"

func main() {
	r, err:=http.Get(fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s",key,"Gladiator"))
	if err!= nil{
		panic(err)
	}
	defer r.Body.Close();

	js, err := ioutil.ReadAll(r.Body)
	if err!= nil{
		panic(err)
	}

	fmt.Println(string(js))
	m:=map[string]interface{}{}
	err=json.Unmarshal(js,&m)

	mov:=movie{}
	err=json.Unmarshal(js,&mov)

	fmt.Println(m["Year"])
	fmt.Println(m["Rated"])
	fmt.Println(mov.Director)

}