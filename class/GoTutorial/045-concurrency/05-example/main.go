package main

import (
	"fmt"
	"sync"
	//"sync/atomic"
	"runtime"
	//"time"
)

var wg sync.WaitGroup
var counter int
var mut = sync.Mutex{}

func main() {
	q:=100
	wg.Add(q)

	for i:=0;i<100;i++ {
		go foo(i)
	}

	wg.Wait()

	fmt.Println("Finish routines", runtime.NumGoroutine())
	// print counter at end of program
	fmt.Println("Finish counter",counter)
}


func foo(i int){
	fmt.Println(i)

	// use mutex lock
	mut.Lock()
	// add a counter and increment each time foo prints
	counter++
	mut.Unlock()

	wg.Done()
}

// check there is no race condition with -race flag
