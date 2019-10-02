//Create 2 types 
//  1. person
//  2. secret agent
//give each a speak method
//Create interface human as speak method
//create func saysomething(human)
//use it all.

package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	codename string
}

type mathProfessor struct {
	first string
	num1 int
	num2 int
}

func (p person) speak() {
	fmt.Printf("The name is %s. %s %s\n",p.last, p.first, p.last)
}

func (s secretAgent) speak() {
	fmt.Printf("My code name is %s.\n", s.codename)
}

func (m mathProfessor) speak() {
	fmt.Printf("%s says: %v plus %v is %v\n",m.first, m.num1, m.num2, (m.num1+m.num2))
}

type human interface {
	speak()
}

func saysomething(h human){
	h.speak()
}

func main() {
	jb:=person{
		first:"James",
		last: "Bond",
	}
	saysomething(jb)

	jbsa:=secretAgent{
		person: jb,
		codename: "007",
	}

	saysomething(jbsa)

	var mp human = mathProfessor{ first: "Professor John", num1: 12, num2:33,}
	saysomething(mp)

	//dynamic assignment
	vs:=[]interface{}{"this a string", 42, 44.56, person{first:"john", last:"doe"}}
	for _, v := range vs {
      fmt.Printf("%v %T\n", v, v)
    } 
}