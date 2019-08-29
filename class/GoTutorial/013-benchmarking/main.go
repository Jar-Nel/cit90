package main

import f "fmt"

func main() {
	f.Println("Test")
	compLit()
	makeway()
}

func compLit() {
	ii:=[]int{}
	//l:=len(ii)
	c:=cap(ii)
	//f.Println("len",l)
	//f.Println("cap",c)
	for i:=0;i<1000000;i++{
		ii=append(ii, i)
		if c!=cap(ii){
			c=cap(ii)
			//f.Println("new cap:",c)
		}
	}
}

func makeway(){
	ii:=make([]int,0,1000000)
	//l:=len(ii)
	c:=cap(ii)
	//f.Println("len",l)
	//f.Println("cap",c)
	for i:=0;i<1000000;i++{
		ii=append(ii, i)
		if c!=cap(ii){
			c=cap(ii)
			//f.Println("new cap:",c)
		}
	}	
}