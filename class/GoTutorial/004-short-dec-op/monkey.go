package main

//go run * to execute all files as one program.

import "fmt"

var b="pkg variable"

func main() {
	x:="Jack"
	y:=42
	z:=42.47
	a:=`Backticks	are
	string literals`
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)

	//multi-cursor editing.
	//select next of higligted, control-d
	fmt.Printf("x valueOf: %v, typeOf: %T\n",x,x)
	fmt.Printf("y valueOf: %v, typeOf: %T\n",y,y)
	fmt.Printf("z valueOf: %v, typeOf: %T\n",z,z)
	fmt.Printf("a valueOf: %v, typeOf: %T\n",a,a)
	fmt.Printf("b valueOf: %v, typeOf: %T\n",b,b)

	fmt.Println("Calling foo:",foo())
}