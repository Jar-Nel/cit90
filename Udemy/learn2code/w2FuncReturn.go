//https://play.golang.org/p/gZeNieIzXCO
package main

import (
	"fmt"
)

func main() {
	fmt.Println(mult(3,4))
}

func mult(a, b int) (int){
	return a*b
}
