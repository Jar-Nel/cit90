package main

import f "fmt"

func main(){
	for x:=0;x<7;x++{
		cond1(x)
	}
	for x:=0;x<7;x++{
		cond2(x)
	}

	cond3()
}

func cond1(x int) {
	if x==4{
		f.Printf("c1: x=%v four\n",x);
	} else if x==5 {
		f.Printf("c1: x=%v five\n",x)
	} else if x==6 {
		f.Printf("c1: x=%v six\n",x)
	} else {
		f.Printf("c1: x=%v not that\n",x)
	}
}

func cond2(x int) {
	//default no fall through.
	//explicit fallthrough keyword.
	switch x {
	case 4:
		f.Printf("c2: x=%v four\n",x)
		fallthrough
	case 5:
		f.Printf("c2: x=%v five\n",x)
	case 6: 
		f.Printf("c2: x=%v six\n",x)
	default: f.Printf("c2: x=%v not that either\n",x)
	}
}

func cond3(){
	x:=true
	switch {
	case (3==2):
		f.Println("c3: not true")
	case (3==3):
		f.Println("c3: 3==3 true")
	case x:
		f.Printf("c3: x is %v\n",x)
	default: 
		f.Println("c3: None of these happened")
	}
}

/*func cond4(x int){
	switch x {
		
	}
}*/