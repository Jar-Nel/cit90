package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}


func main () {
	tr:= truck {
		vehicle:vehicle{
			doors:2,
			color:"Blue",
		},
		fourWheel: true,
	}
	se:=sedan {
		vehicle:vehicle{
			doors:4,
			color:"Red",
		},
		luxury: false,
	}

	fmt.Println(tr)
	fmt.Println(se)

	fmt.Println("Truck doors:",tr.vehicle.doors)
	fmt.Println("Sedan doors:",se.vehicle.doors)	
}

