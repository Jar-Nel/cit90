package main

import (
	"golang.org/x/crypto/bcrypt"
	//"crypto/sha256"
	"fmt"
	"log"
	//"bytes"
)

func main() {

	pw:="BananaBread12"
	bs,err:=bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err!=nil {
		log.Fatal("GenerateFromPassword error", err)
	}

	fmt.Println(bs)
	fmt.Printf("%x\n",bs)
	p2:="AnotherPW"


	fmt.Println("Compare password stored wiht password entered")
	err=bcrypt.CompareHashAndPassword(bs,[]byte(p2))
	if err!=nil{
		//Exits program if error
		log.Fatalln("Passwords do not match")
	}
	fmt.Println("Passwords Matched")
}