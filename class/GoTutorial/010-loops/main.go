package main

import f "fmt"

func main() {
	loop1()
	loop2()
	loop3()
}

func loop1() {
	x:=0
	for {
		f.Println(x)
		x++
		if x==123{
			break
		}
	}
}

func loop2() {
	//for init;condition;post {}
	for i:=123; i>=0; i-- {
		f.Println(i)
	}
}

func loop3() {
	i:=0
	for i<100 {
		f.Println(i)
		i++
	}
}