package main 
  
import "fmt"
  
func main() { 
  
    m:=map[string]string {  //map (key value)
		"Stephen King": "The Dark Tower",
		"John Scalzi":"Old Man's War",  
	}

	for k, v := range m {
        fmt.Println("Key:", k, "Value:", v)
    }
} 