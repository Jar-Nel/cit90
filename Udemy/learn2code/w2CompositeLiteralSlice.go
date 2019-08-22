//https://play.golang.org/p/sb1PYDF1Cnc
package main

import (
	"fmt"
)

func main() {
	sli:=[]int{2,4,6,8}
	fmt.Println("print the slice:",sli)
	fmt.Println("print a value by index:",sli[2])
	fmt.Println("print all values using a for range:")
	for i:=range sli {
		fmt.Println(sli[i])
	}
}

