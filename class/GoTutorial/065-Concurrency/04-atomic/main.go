package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	launchRoutines()
	wg.Wait()
	fmt.Println("Finish routines", runtime.NumGoroutine())
	fmt.Println("Finish counter",counter)
}

func launchRoutines(){
	const gs=100
	wg.Add(gs)

	for i:=0;i<gs;i++ {
		go func() {
			time.Sleep(time.Second)

			runtime.Gosched()
			
			counter=atomic.AddInt64(&counter, 1)

			wg.Done()
		}()
		fmt.Println("Working routines",runtime.NumGoroutine())
	}

	wg.Wait()

	/*
	Waitgroups
	Mutex
	Atomic

	All old.  New way is channels
	*/
}
