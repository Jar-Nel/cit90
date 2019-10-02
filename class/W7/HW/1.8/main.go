//Exercise 1.8: 
//Modify fetch to add the preﬁx http:// 
//to each argument URL if it is missing. 
//You might want to use strings.HasPrefix.


package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url,"http://") || strings.HasPrefix(url,"https://")){
			url="http://"+url	
		} 
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_,err=io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//!-
