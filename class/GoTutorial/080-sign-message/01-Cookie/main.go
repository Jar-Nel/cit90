package main

import (
	"fmt"
	"crypto/sha256"
	"strings"

)

const secretKey="gibberish"

func main() {
	emailInCookie:="test@test.org"
	s:=emailInCookie+secretKey
	sum:=sha256.Sum256([]byte(s))
	//hash:=string(sum[:])
	hash:=fmt.Sprintf("%x", sum)
	//valueInCookie:=emailInCookie+"|"+hash


	emailTprCookie:="test2@test.org"
	valueFromCookie:=emailTprCookie+"|"+hash

	xs:=strings.Split(valueFromCookie, "|")
	emailFromCookie:= xs[0]
	hashFromCookie:=xs[1]

	s2:=emailFromCookie+secretKey
	sum2:=sha256.Sum256([]byte(s2))
	hash2:=fmt.Sprintf("%x", sum2)

	if hashFromCookie!=hash2{
		fmt.Println("Hash don't match", emailFromCookie)
		fmt.Printf("%s\n",valueFromCookie)
	} else {
		fmt.Print("Match!", emailFromCookie)
	}


}

