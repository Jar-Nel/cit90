package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Printf("I am named %s\n",p.first)
}

func (s secretAgent) speak() {
	fmt.Printf("I am a secret agent.\n")
}

type human interface {
	speak()
}

func saysomething(h human){
	h.speak()
}

func main() {
	jm:=person{
		first:"Jenny",
	}
	saysomething(jm)

	jb:=secretAgent{
		person: person {first:"James"},
		ltk:true,
	}

	//jb.speak()
	saysomething(jb)
}