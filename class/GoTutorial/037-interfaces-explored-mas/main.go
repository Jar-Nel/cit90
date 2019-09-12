package main

import "fmt"

type person struct{
	first, last string
}

//dbPG and dbMongo is also of type storer interface, because they implement the functions
type storer interface {
	access(int)person
	set(int,person)
}

func getter(s storer, i int) person{
	return s.access(i)
}

func setter(s storer, i int, p person){
	s.set(i,p)
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
	//Map way
	//mPG:=dbPG{}
	mPG:=dbMongo{}

	mPG.set(0,p1)
	fmt.Println(mPG.access(0))
	fmt.Println(mPG)

	//Interface way
	setter(mPG, 0, p1)
	fmt.Println(getter(mPG,0))

}