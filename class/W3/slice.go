package main 
  
import "fmt"
  
func main() { 
  
    s1 := []string{"One String", "Two String", "Red String", "Blue String"} 
	
	// i is optional index
	for i, s:= range s1{
		fmt.Printf("%v: %v\n",i,s)
	}
	for _, s := range s1 {
		fmt.Println(s)
	}
} 