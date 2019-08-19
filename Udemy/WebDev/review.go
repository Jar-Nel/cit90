package main

import "fmt"

type person struct {
	fname string  //lowercase is not visible outside package
	Lname string  //uppercase is visible outside package
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {  //func (target) appends function to object.
	fmt.Println(p.fname, p.Lname, `says "Good morning, James."`)
}

func (sa secretAgent) speak() {  //func (target) appends function to object.
	fmt.Println(sa.fname, sa.Lname, `says "Shaken, not stirred."`)
}

type human interface {  //polymorphism p and sa have speak so they will be this type
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	//var x int   //Long variable declaration
	//x=7

	x:=7  //Short Variable declaration operator.  
		  //Declares variable when assigned, figures out type.
	fmt.Println(x)

	xi:=[]int{2,4,6,8} //slice of int (list of data of type int)
	fmt.Println(xi)

	m:=map[string]int {  //map (key value)
		"Todd": 45,
		"Job":42,  // need trailing comma.
	}

	p1:=person{
		"Miss",
		"Moneypenny",
	}  //if give values of fields in order do not need to provide field name.

	p1.speak()

	sa1:=secretAgent {
		person {
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)
}