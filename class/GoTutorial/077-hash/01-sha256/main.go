package main

import (
	"crypto/sha256"
	"fmt"
	//"bytes"
)

func main() {
	s1:="Some arbitrary text"
	s2:="Some arbitrary text."

	sum1:=sha256.Sum256([]byte(s1))
	sum11:=sha256.New().Sum([]byte(s1))
	sum2:=sha256.Sum256([]byte(s2))

	//n := bytes.IndexByte(sum1, 0)
	//s := string(sum1[:])

	fmt.Printf("sum1:\t %x\n",sum1)
	fmt.Printf("sum11:\t %x\n",sum11)
	fmt.Printf("sum2:\t %x\n",sum2)
	fmt.Println(sum1==sum2)
}