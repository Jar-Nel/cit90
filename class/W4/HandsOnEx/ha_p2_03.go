package main

import (
	"fmt"
)

func main() {
	defer defertest()
	saysomething()
}

func defertest() {
	defer fmt.Println("this is the last line")
	fmt.Println("This is a defer test.")
}

func saysomething() {
	fmt.Println("Something")
}
