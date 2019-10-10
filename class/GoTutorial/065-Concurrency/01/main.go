package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t",runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t",runtime.NumGoroutine())

	//wait for sync to be done
	wg.Wait()
}

func foo(){
	for i:=0;i<10;i++ {
		fmt.Println("foo: ",i)
	}
	wg.Done()
}

func bar(){
	for i:=0;i<10;i++ {
		fmt.Println("bar: ",i)
	}
}