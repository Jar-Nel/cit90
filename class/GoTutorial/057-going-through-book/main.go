package main

import (
	"fmt"
	"057/goofy" //go mod init 057
)

var y = 99

func main() {
	fmt.Println(y)
	n:=foo()
	fmt.Println(n)
	fmt.Println(bar())

	fmt.Println(goofy.Add1(21))
}

func foo() int {
	x:=42
	x++
	x*=2
	return x
}