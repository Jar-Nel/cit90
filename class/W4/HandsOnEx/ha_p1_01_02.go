package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favIceCreamFlavors []string
}

func (p person) printFlav() {
	fmt.Printf("%v %v likes \n",p.firstName, p.lastName)
	for _,v:=range p.favIceCreamFlavors {
		fmt.Printf("  %v\n",v)
	}
}

func main () {
	p1:=person{
		firstName:"John",
		lastName: "Doe",
		favIceCreamFlavors: []string {"Chocolate", "Vanilla", "Strawberry"},
	}
	p2:=person{
		firstName:"Jane",
		lastName: "Smith",
		favIceCreamFlavors: []string {"Rocky Road", "Strawberry Cheesecake"},
	}

	//p1.printFlav()
	//p2.printFlav()
	
	m:=map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k,v:= range m{
		fmt.Println(k)
		v.printFlav()
	}
}

