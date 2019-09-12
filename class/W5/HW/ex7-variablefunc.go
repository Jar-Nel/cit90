package main

import "fmt"

func main() {
	s:=square
	n:=4
	fmt.Printf("Square of %v is %v\n",n,s(n))

	pwr:=func(x int, y int) int {
		res:=x
		for i:=1;i<y;i++{
			res=res*x
		}
		return res
	}

	fmt.Printf("2 to the power of 3 is %v\n",pwr(2,3))
}

func square(a int)int{
	return a*a
}