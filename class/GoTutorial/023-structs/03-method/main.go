package main

import f "fmt"

type person struct {
	first string
	last string
	favnum int
	drinks bool
}

//func (receiver) identifer(params) returns {code}
func (p person) speak() {
	f.Printf("I'm a person and my name is %s %s\n", p.first, p.last)
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	f.Printf("I'm a secret agent and my name is %s. %s %s\n", sa.person.last, sa.person.first, sa.person.last)
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

	p1.speak()
	sa2.speak()
	sa1.person.speak()
}
