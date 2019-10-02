package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rNum := rand.Intn(3)
	fmt.Println(rNum)

	for i:=0; i<10; i++{
		switch rand.Intn(5) {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		default:
			fmt.Println("Number too high")
		}
	}
}