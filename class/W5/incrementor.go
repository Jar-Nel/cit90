package main

import "fmt"

func main(){
	a:=incrementor();
	fmt.Println(a())
	fmt.Println(a())
}

func incrementor() func() int{
	var x int
	return func() int{
		x++
		return x
	}
}