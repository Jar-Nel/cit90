package main

import "fmt"

func main () {

	fmt.Println(foo(2,3,4,5))
	fmt.Println(bar([]int{5,6,7,8}))
}



func foo (ii ...int) int {
	var sum int
	for _,v:=range ii{
		sum+=v
	}
	return sum
}

func bar (ii []int) int {
	var sum int
	for _,v:=range ii{
		sum+=v
	}
	return sum	
}