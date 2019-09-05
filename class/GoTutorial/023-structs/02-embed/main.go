package main

import f "fmt"

type person struct {
	first string
	last string
	favnum int
	drinks bool
}

type secretAgent struct {
	person
	licenseToKill bool
}


func main(){
	p1:= person{
		first:"John",
		last:"Doe",
		favnum:13,
		drinks:false,
	}

	f.Println(p1)

	sa1:=secretAgent {
		person: p1,
		licenseToKill: false,
	}

	sa2:=secretAgent {
		person: person {
			first:"James",
			last:"Bond",
			favnum: 7,
			drinks: true,
		},
		licenseToKill: true,
	}

	f.Println(sa1)
	f.Println(sa2)
}
