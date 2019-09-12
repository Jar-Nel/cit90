package main

import "fmt"

func main() {
	doSomething(func() {fmt.Println("Success Function")}, func() {fmt.Println("Failure Function")})
}

func doSomething(success func(), failure func()){
	worked:=true;
	if (worked){
		success()
	} else {
		failure()
	}
}