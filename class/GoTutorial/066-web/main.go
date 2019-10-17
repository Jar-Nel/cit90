package main

import (
	"net/http"
	"fmt"
	"sort"
	"log"
	//"github.com/valyala/fasthttp"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/index",serveIndex)
	http.HandleFunc("/about",foo)
	log.Fatal(http.ListenAndServe(":808", nil))
}

func foo(w http.ResponseWriter, r *http.Request){
	m:= map[string]string {
		"Method":r.Method,
		"URL":r.URL.String(),
		"Proto":r.Proto,
		"URI":r.RequestURI,
		"Referrer":r.Referer(),
		"Host":r.Host,
	}

	keys:=[]string{}
	for k:= range m{
		keys = append(keys, k)
	}

	sort.Strings(keys)

	//sorted
	for _,k := range keys{
		fmt.Fprintf(w,"%s: %s\n",k,m[k])
	}
	
	//Unsorted
	//for k,v:=range m{
	//	fmt.Fprintf(w,"%s: %s\n",k,v)
	//}

	fmt.Fprintf(w,"\tHeaders:\n")
	for k,v :=range r.Header {
		fmt.Fprintf(w, "\t\t%s - %s\n", k, v)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}