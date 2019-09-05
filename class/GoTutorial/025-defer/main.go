package main

import f "fmt"

func main(){
	//defer last to run.  defers are stacked last in first out.  Change order of execution
	defer foo()
	bar()
}

func foo() {
	f.Println("foo")
}

func bar() {
	f.Println("bar")
}