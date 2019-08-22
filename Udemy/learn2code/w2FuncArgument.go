//https://play.golang.org/p/HvxIq-gpB-r
package main

import (
	"fmt"
)

func main() {
	w2func("Jared Nelson", 42)
	w2func("Friend's Name", 38)
}

func w2func(name string, num int) {
	fmt.Println(name,"is",num)
}
