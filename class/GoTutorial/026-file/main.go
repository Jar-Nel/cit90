package main

import (
	f "fmt"
	"os"
)

func main(){
	file, err :=os.Create("data.txt")
	if err!=nil {
		panic(err)
	}

	defer file.Close()

	for i:=0; i<=10; i++{
		file.WriteString(f.Sprintf("I will not talk in class. %d\n",i))
	}

	
}
