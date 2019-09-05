package main

import f "fmt"

func main() {
	m:=map[string]int{
		"James":32,
		"Jenny":24,
	}
	f.Println(m)
	//Access by key
	f.Println(m["James"])
	f.Println(m["Jenny"])

	v, ok := m["James"]
	f.Println(v)
	f.Println(ok)

	//Two variables.  V has the value, ok has a boolean on if the map exists
	v, ok = m["Jack"]
	f.Println(v)
	f.Println(ok)

	xi:= []int{42,43,44,45}
	for i,v:=range xi{
		f.Println("slice",i,v)
	}

	for k,v:=range m{
		f.Println("map",k,v)
	}

	//Create your own type
	type flavor string
	xf :=[]flavor{"chocolate","vanilla","strawberry"}
	for i,v:=range xf {
		f.Println("flavors",i,v)
	}

	type orders map[int][]flavor

	o:=orders{
		1: []flavor{"chocolate"},
		2: []flavor{"vanilla", "chocolate"},
		3: []flavor{"strawberry", "mint chip"},
	}

	for k,v := range o {
		f.Printf("Order %d had ", k)
		for _, vv:= range v{
			f.Printf("%s ",vv)
		}
		f.Printf("\n")
	}

	//loop init, condition, post
	for i:=0; i<10; i++{
		f.Printf("Loop %d \n",i)
		v, ok := o[i]
		if(ok) {
			f.Println(i,v)
		}
	}

	//note: order is guaranteed in a slice,  it is not guaranteed in a map.  map is for association.
}