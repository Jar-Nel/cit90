package main

import (
	f "fmt"
	//s "./stuff" //it is in a different folder
	s "main/stuff"
)

func main(){
	f.Println(s.Foo())
	f.Println("exiting")
}