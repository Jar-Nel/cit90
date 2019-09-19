package main

import "fmt"

type person struct {
	first string
}

func (p person) foo() {
	fmt.Printf("Person foo here %s\n", p)
}

func bar(t talker) {
	fmt.Println(t)
}


type talker interface {
	foo()
}

func (p *person) fooToo() {
	fmt.Printf("Person fooToo here %s\n", p)
}

type talkerToo interface {
	fooToo()
}

func barToo(t talker) {
	fmt.Println(t)
}

//type talkerToo interf

func main() {
  p1:= person{
	  first: "MyName",
  }

  bar(p1);

  p1.foo()

  barToo(p1)

  p1.fooToo()
}