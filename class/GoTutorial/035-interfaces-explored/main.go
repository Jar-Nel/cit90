package main

import "fmt"

type person struct{
	first, last string
}

type dbPG map[int]person

//type dbMongo map[int]person

//seperate Database routines in method interface 
//for db agnostic apps
func (pg dbPG) access(i int) person {
	return pg[i]
}

func (pg dbPG) set(i int, p person) {
	pg[i]=p
}

type dbMongo map[int]person

func (m dbMongo) access(i int) person {
	return m[i]
}

func (m dbMongo) set(i int, p person) {
	m[i]=p
}




func main (){
	p1:=person {
		first: "Jenny",
		last: "Moneypenny",
	}
	//mPG:=dbPG{}
	mPG:=dbMongo{}

	mPG.set(0,p1)
	fmt.Println(mPG.access(0))
	fmt.Println(mPG)
}