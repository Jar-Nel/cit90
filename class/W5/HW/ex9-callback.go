package main

import "fmt"

func main() {
	s:=func()string {
		return "The operation was a success"
	} 
	f:=func()string{
		return "We totally failed"
	}

	fmt.Println(doSomething(true,s,f))
	fmt.Println(doSomething(false,s,f))
}

func doSomething(done bool, success func()string, failure func()string)string {
	if (done){
		return success()
	}  
	return failure()
}