
package main

import (
	// https://godoc.org/fmt
	"fmt"
)

func main() {
	fmt.Println(foo())	
}


// Go function format
// func receiver identifer (parameters) returns {code}
// functions that start with Uppercase letter are public, lowercase letter are private.
// VALUES and TYPE

func foo() string {
	return "Hello, world."
}