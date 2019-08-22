// https://play.golang.org/p/hqan3UUy65i
package main

import (
	"fmt"
)


func main() {
	m:=map[string]string {  //map (key value)
		"author": "Fancy Bookwriter",
		"title":"Fancy Book Title",  
	}
	
	fmt.Println("print one of the entries: ",m["author"])
	fmt.Println("print the whole map: ")
	for key, value := range m {
        fmt.Println(key, "=", value)
    }
}



