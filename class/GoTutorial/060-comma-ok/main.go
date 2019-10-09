package main

import "fmt"



func main() {
	m:=map[int]string {1:"Jenny", 2:"James"}
	fmt.Println(m);

	v, ok := m[1];
	fmt.Println(v, ok)

	if v, ok :=m[2]; ok {  //limits scope of ok to this line when on one line.
		fmt.Println(v)
	}

	var x = 42
	fmt.Printf("%T\t\t%v",x,x)
	//fmt.Printf("%T\t\t%v",x.(string),x)

}

/*
v, ok = m[key]
v, ok = x.(T)
v, ok = <-ch
*/