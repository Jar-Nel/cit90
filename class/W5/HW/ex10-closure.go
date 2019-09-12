package main

import "fmt"

func main() {
	bytwo:=adder();
	fmt.Println(bytwo())
	fmt.Println(bytwo())
	fmt.Println(bytwo())
	fmt.Println(bytwo())
	fmt.Println(bytwo())
	fmt.Println(bytwo())
}

func adder() func() int {
	num:=0
	return func() int{
		num+=2
		return num
	}
}