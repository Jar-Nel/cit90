package main

import f "fmt"

func main() {
	//Composite Literal
	xs:=[]string{"James","Jenny"}
	f.Println(xs)

	//Composite Literal
	m:=map[string]string {
		"James":"Martini",
		"Jenny":"Manhattan",
	}
	f.Println(m)

	//add value to a map
	m["Jack"]="Jack Daniels"
	f.Println(m)

	//delete from map
	delete(m, "Jack")
	f.Println(m)
}