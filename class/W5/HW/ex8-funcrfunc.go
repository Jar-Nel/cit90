package main

import "fmt"

func main () {
	wf1:=wordfactory("Crash")
	fmt.Println(wf1("Override"))
	fmt.Println(wf1("Course"))

	wf2:=wordfactory("Good")
	fmt.Println(wf2("Code"))
}

func wordfactory(w1 string) func(w2 string)string {
	return func(w2 string) string {
		return w1+" "+w2
	}
}