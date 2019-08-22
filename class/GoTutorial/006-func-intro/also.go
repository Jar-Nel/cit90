package main

import "fmt"
//functions
//func receiver identifer (parameters) return(s) {code}

func foo() string {
	return "from foo"
}

func bar() string{
	return "from bar"
}

func chip(num int) int {
	return num+1
}

func dale(x int) string {
	return fmt.Sprintf("from dale %d",x)
}

func mickey()(string, int) {
	return "from Mickey", 45
}

func minnie() (string, int, bool){
	return "from Minnie", 50, true
}