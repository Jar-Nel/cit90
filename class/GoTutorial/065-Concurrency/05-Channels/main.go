package main

import (
	"fmt"
)

func main() {
	//ch:=make(chan string)
	//go fooCh("Me:",ch)
	//fmt.Println(<-ch)

	c:=make(chan int)
	num:=0
	for i:=0;i<100;i++{
		done:=make(chan bool)
		go func(cha chan<-int) {
			num++
			c<-num
			done<-true
		}(c)
		close(done)
		<-done  //wait for func to end
		fmt.Println("done",i)
		func (cha chan int)  {
			//for v:= range cha{
				fmt.Println(<-cha)
			
		}(c)
	}
	

	//for v:=range c{
	//	fmt.Println(v)
	//}

	//close(c)

	//for i:=0;i<100;i++{
	//	fmt.Println(<-c)
	//}


	close(c)

}

func fooCh(me string, ch chan<- string) {
	ch<-me+" From FooCH"
}