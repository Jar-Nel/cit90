package main

import "fmt"

//STACK
// moo
// bar
// foo
// main

//HEAP

func main() {
	s:=foo()
	fmt.Println(s)
}

func foo() string {
	s:=bar()
	return fmt.Sprintf("from foo - %s",s)
}

func bar() string {
	s:=moo()
	return fmt.Sprintf("from bar - %s", s)
}

func moo() string {
	return "from moo"
}