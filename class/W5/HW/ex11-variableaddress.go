package main

import "fmt"

func main() {
	value:=3000
	fmt.Printf("value of value\t\t\t%v\n",value)
	//show the address of the variable where the value is stored
	fmt.Printf("&value address\t\t\t%v\n",&value)
	//show the type of the variable X
	fmt.Printf("Type of value\t\t\t%T\n",value)
	//show the type of the address for the variable X where the value is stored
	fmt.Printf("Type of value address\t\t%T\n",&value)
	//assign the address of X to another variable Y
	y:=&value
	//show the type of that variable Y
	fmt.Printf("Type of y\t\t\t%T\n",y)
	//dereference that variable Y and assign a new value to it
	*y=42
	//show the value in the original variable X
	fmt.Printf("value of value\t\t\t%v\n",value)
}