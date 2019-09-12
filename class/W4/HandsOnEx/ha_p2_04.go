package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func (p person) speak() {
	fmt.Printf("Hi! My name is %v, my age is %v\n",p.firstName, p.age)
}

func main () {
	p1:=person{
		firstName:"John",
		lastName: "Doe",
		age: 22,
	}

	p1.speak()
	
}

