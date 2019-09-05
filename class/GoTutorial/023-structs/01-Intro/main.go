package main

import f "fmt"

type person struct {
	first string
	last string
	dob string
	favnum int
	drinks bool
}


func main() {
	p1:=person{
		first:"James",
		last:"Bond",
		dob:"01/01/1987",
		favnum:7,
		drinks:true,
	}

	f.Println(p1)
	f.Println(p1.first)
	f.Println(p1.last)


	xi:=[]int {42,43,44,99,23423}

	m:=map[int]string{
		7:"James",
		8: "Jenny",
		9: "Jack",
	}

	f.Println(xi)
	f.Println(m)

	type limitedScopePerson struct {
		first string
		last string
		dob string
	}

}