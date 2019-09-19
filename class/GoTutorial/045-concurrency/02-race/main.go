package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t",runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)

	counter:=0
	const gs=100

	var wg sync.WaitGroup
	wg.Add(gs)

	m:=sync.Mutex{}

	for i:=0;i<gs;i++ {
		go func() {
			time.Sleep(time.Second)

			//Yeilds Processor
			runtime.Gosched()
			
			//hold variable
			m.Lock()
			counter++
			m.Unlock()
			
			wg.Done()
		}()
		fmt.Println("Working routines",runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Finish routines", runtime.NumGoroutine())
	fmt.Println("Finish counter",counter)



}

