package main

import (
	"fmt"
	//"runtime"
	"sync"
	"net/http"
	//"io/ioutil"
	"io"
	"os"
	"strings"
	"time"
)


var wg sync.WaitGroup

func main(){
	urls:=[]string {"https://www.google.com","https://www.fresnocitycollege.edu", "https://my.sjcl.edu","https://www.gov.au/","https://www.agls.gov.au"}
	for _,v:= range urls{
		wg.Add(1)
		go foo(v)
	}

	fmt.Println("Done")
	wg.Wait()
}

func foo(url string) {
	start := time.Now()
	r, err:=http.Get(url)
	if err!=nil {
		fmt.Errorf("Cant read page", err)
		return
	}
	filename:=strings.Split(strings.Split(url,"://")[1],".")[1]+".txt"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(f, r.Body)
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
	defer r.Body.Close()
	
	secs := time.Since(start).Seconds()
	fmt.Println(filename, secs, "seconds")
	wg.Done()
}