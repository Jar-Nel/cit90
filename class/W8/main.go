
package main

import (
	"flag"
	"fmt"
	"strconv"
	"math"
)

var miles = flag.Bool("miles", false, "Miles to Kilometers")  
var km = flag.Bool("km", false, "Kilometers to Miles")  
//-h returns help

func main() {
	flag.Parse()
	if len(flag.Args())<1{
		fmt.Println("Incorrect arguments specified, use -h for help")
		return
	}
	if (*miles){
		for _,v := range flag.Args(){
			m,err:=strconv.ParseFloat(v,64)
			if ok:=handleError(err); ok {
				km:=convMilesToKM(m)
				fmt.Printf("%v miles is %v kilometers\n", m, km)
			}
		}
	} else if (*km) {
		for _,v := range flag.Args(){
			km,err:=strconv.ParseFloat(v,64)
			if ok:=handleError(err); ok {
				m:=convKMtoMiles(km)
				fmt.Printf("%v kilometers is %v miles\n", km,m)
			}
		}
	} else {
		fmt.Println("Incorrect arguments specified, use -h for help")
	}
}

func handleError(err error)bool{
	if (err!=nil) {
		fmt.Println("ERROR: ",err)
		return false
	}
	return true
}

func convMilesToKM(mile float64)float64 {
	return math.Round((mile*1.609344)/0.001)*0.001
}
func convKMtoMiles(km float64)float64 {
	return math.Round((km/1.609344)/0.001)*0.001
}



//!-
