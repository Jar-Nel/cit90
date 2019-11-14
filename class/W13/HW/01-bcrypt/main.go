package main

import (
	"golang.org/x/crypto/bcrypt"
	//"crypto/sha256"
	"fmt"
	"log"
	//"bytes"
)

func main() {

	pw:="pw123"
	bs,err:=bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err!=nil {
		log.Fatal("GenerateFromPassword error", err)
	}
	fmt.Println(string(bs[:]))
	pw2:="pw123"
	pw3:="Notpw123"


	fmt.Println("Compare password stored with password entered")
	fmt.Println("pw & pw2:")
	msg, ok:=checkPW(bs, pw2)
	fmt.Printf("\t ok: %v\t msg: %v\n", ok, msg)
	msg, ok=checkPW(bs, pw3)
	fmt.Println("pw & pw3:")
	fmt.Printf("\t ok: %v\t msg: %v\n", ok, msg)
}

func checkPW(pwb[] byte, pws string)(string, bool) {
	err:=bcrypt.CompareHashAndPassword(pwb,[]byte(pws))
	if err!=nil{
		return "Passwords do not match", false
	}
	return "Passwords match", true
}