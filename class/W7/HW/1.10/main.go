//Exercise 1.10: Find a web site that produces a 
//large amount of data. Investigate caching by 
//running fetchall twice in succession to see 
//whether the reported time changes much. Do you
//get the same content each time? Modify fetchall
//to print its output to a ﬁle so it can be examined.

package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"time"
	"strconv"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		fname:="File"+strconv.Itoa(i)+".txt"
		go fetch(url, fname, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, filename string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	nbytes, err := io.Copy(f, resp.Body)
	f.Close()
	resp.Body.Close() // don't leak resources
	
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
