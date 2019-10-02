package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	f, err := os.Create("test.txt")
	if err!=nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "hello world")
	//Fprint uses interface writer, so pass in something with writer
	fmt.Fprintf(os.Stdout,"\nHello world\n")

	//can also use http response writer as output.
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Hello world again")
	})
	http.ListenAndServe(":8080", nil)


}

func booga (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello world again")
}