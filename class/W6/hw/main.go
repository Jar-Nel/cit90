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
	fmt.Println("Finish counter",counter)
}


func foo(i int){
	fmt.Println(i)

	mut.Lock()
	counter++
	mut.Unlock()

	wg.Done()
}

// add a counter and increment each time foo prints
// print counter at end of program
// use mutex lock
// check there is no race condition with -race flag
