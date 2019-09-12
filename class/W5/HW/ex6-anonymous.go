package main

import "fmt"

func main(){
	fmt.Println(runanon())
	fmt.Printf("%s\n",func() string{
		return "So is this"
	}())
}

func runanon()string {
  return func()string {
	return "This is a convoluted way to show an anonymous func"
  }()
}