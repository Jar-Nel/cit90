package main

import (
	f "fmt"
	"os"
	"bufio"
)

func main(){
	toddFile()

}

func toddFile(){
	file, err := os.Open("data.txt")
	if err!=nil {
		//panic(err)
		f.Fprintln(os.Stderr, "Reading input: ",err)
	}	

	s:=bufio.NewScanner(file)
	s.Split(bufio.ScanWords)
	var count int
	m:=map[string]int{}
	for s.Scan(){
		t:=s.Text()
		f.Println(t)
		m[s.Text()]++
		count++
	}
	f.Println(count)
	

	
}

func appendTXTFile(){
	file, err := os.OpenFile("data.txt", os.O_APPEND, 0666)
	if err!=nil {
		panic(err)
	}

	defer file.Close()

	for i:=0; i<=10; i++{
		file.WriteString(f.Sprintf("I will not Dance in class. %d\n",i))
	}

}

func readFileFMT() {
	file, err :=os.Open("data.txt")
	if err!=nil {
		panic(err)
	}

	defer file.Close()
	
	var a string
	var itemCount int
	itemCount, err = f.Fscanf(file, "%s", &a)
	if err!=nil {
		panic(err)
	}

	f.Println(itemCount, err)
	for itemCount > 0 && err == nil {
		f.Println(a)
		itemCount, err = f.Fscanf(file, "%s",&a)
		if err!=nil {
			panic(err)
		}
	}	
}
