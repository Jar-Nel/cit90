package main

import f "fmt"

func main(){
	//composite data type
	//compisite literal  x:=type(value)
	x:= []int {42,43,44,45,1024}
	//slice is like an array.  slice is dynamic, arrays are not
	//slice stores values of same type

	f.Println(x)
	f.Printf("Type type of x is %T\n",x)

	//index, value.  foreach xs in x 
	for i,xs:=range x{
		f.Println(i,xs)
	}

	for _,xs:=range x{
		f.Println(xs)
	}

	for i:=range x {
		f.Println(i,x[i])
	}

	//slicing a slice
	f.Println(x)
	f.Println(x[0])
	f.Println(x[2:])
	f.Println(x[1:4])

	y:=[]int {20,21,22,23,24,25,26,27,28,29}
	f.Println(y[2:4])
	f.Println(y[2:])
	f.Println(y[:4])

	//appending to a slice
	x = append(x, 2048, 4096)
	f.Println(x)

	z:=[]int {35,36,37}
	y = append(y,33,34)
	y = append(y, z...)
	f.Println(y)

	//delete from slice
	y = y[5:10]
	f.Println(y)

	z=[]int {50,51,52,53,54,55,56,57,58,59}
	z = append(z[0:2], z[8:]...)
	f.Println(z)

	//make slice
	f.Println("len",len(z))
	f.Println("cap",cap(z))

	//Slice is a data structure made up of three values
	//len length
	//cap capacity
	//pointer to array

	//If you know what the capacity will be, define it with make
	h:= make([]int, 10, 20);
	f.Println("len",len(h))
	f.Println("cap",cap(h))
	for i,v:=range h {
		f.Println(i,v)
	}
	h[2] = 12;
	f.Println(h)

	//can't do
	//h[11]=42
	//can do
	h=append(h,42);
	f.Println("len",len(h))
	f.Println("cap",cap(h))
	f.Println(h)
}