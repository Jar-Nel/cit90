//Exercise 1.3: 
//Experiment to measure the difference in running time
//between our potentially inefﬁcient versions and the 
//one that uses strings.Join.
//(Section1.6 illustrates part of the time package, 
//and Section 11.4 shows how to write benchmark 
//tests for systematic performance evaluation.)

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Echo 1: Start")
	t:=time.Now()
	echo1(os.Args)
	fmt.Println("Echo 1: End: ",time.Since(t).Nanoseconds(),"Nanoseconds")
	fmt.Println("Echo 2: Start")
	t=time.Now()
	echo2(os.Args)
	fmt.Println("Echo 2: End: ",time.Since(t).Nanoseconds(),"Nanoseconds")
	fmt.Println("Echo 3: Start:")
	t=time.Now()
	echo3(os.Args)
	fmt.Println("Echo 3: End: ",time.Since(t).Nanoseconds(),"Nanoseconds")
	
}

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string){
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string){
	fmt.Println(strings.Join(args[1:], " "))
}
