package main

import "fmt"

func main(){
	xi:=[]int{42,43,44,99,1024}
	s:=foo(xi)
	fmt.Println(s)

	s=bar(1,2,4,8,16,32,64)
	fmt.Println(s)

	//unfurling a slice
	s=bar(xi...)
	fmt.Println(s)
}

func foo(ii []int) int {
	var sum int
	for _,v:=range ii{
		fmt.Println("adding to total:",v)
		sum+=v
	}
	return sum
}

//variatic paramaterr, pass as many ints as you want.  converts to a slice when called.
func bar(ii ...int) int {
	var sum int
	fmt.Printf("%T\n",ii)
	for _,v:=range ii{
		fmt.Println("adding to total:",v)
		sum+=v
	}
	return sum
}