package main

import f "fmt"

func main() {
	jb:= []string{"James","Bond","Martini","Shaken, not stirred"}
	jm:= []string{"Jenny", "Moneypenny", "Manhattan", "Women do it better"}
	xs:= [][]string{jb,jm}
	
	f.Println(xs)
	f.Println(xs[0])
	f.Println(xs[0][0])
	f.Println(xs[0][1])
	f.Println(xs[1])
	f.Println(xs[1][0])
	f.Println(xs[1][1])
}