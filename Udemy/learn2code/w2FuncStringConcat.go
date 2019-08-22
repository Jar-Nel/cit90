//https://play.golang.org/p/yvDVb9Wfn7E
package main

import (
	"fmt"
)

func main() {
	fmt.Println(concat("Jared"))
}

func concat(name string) (string) {
	return "Hello, "+name+"!"
}
