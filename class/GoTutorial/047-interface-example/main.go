package main

import "fmt"

type hotdog int
type hamburger string

func (h hotdog) foo() {  //hotdog.foo()
	// do something
}

func (ha hamburger) foo() {  //hamburger.foo()
	// do something
}

type vendor interface {  //vendor.foo()
	foo()
}

func main() {
	var v vendor = hotdog(42)
	var ht hotdog = 33
	var ha hamburger="Food"

	fmt.Println(v)
	fmt.Println(ht)
	fmt.Println(ha)
}