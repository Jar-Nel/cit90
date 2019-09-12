package main

import "fmt"

func main() {
	x:=42
	fmt.Println(x)
	fmt.Printf("x\t%v\n",x)
	fmt.Printf("&x\t%v\n",&x)
	fmt.Printf("type of x \t\t%T\n",x)
	fmt.Printf("type of &x \t\t%T\n",&x)


	// '&' shows memory location of value
	fmt.Println("x is stored at",&x)

	fmt.Printf("%T\n",x)
	fmt.Printf("%T\n",&x)

	// *int is a TYPE of pointer to an int
	var a *int
	a=&x
	fmt.Println(a)

	//you can DEREFERENCE a pointer with a '*' operator
	fmt.Println("value stored in location: ",*a)
}