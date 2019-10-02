//Exercise 1.11: 
//Try fetchall with longer argument lists, such as
//samples from the top million web sites available 
//at alexa.com. How does the program behave if a web
//site just doesn’t respond? 

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
	ch <- fmt.Sprintf("%.2fs  %7d  %s  %s", secs, nbytes, filename, url)
}

//!-
