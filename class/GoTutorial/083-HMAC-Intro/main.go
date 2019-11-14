package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
)

const key string ="lotta crap"

func main() {
	msg:="this is a message"
	sig,_:=signMessage([]byte(msg))
	fmt.Println("Message: ", msg)
	fmt.Printf("Sig: %x\n",sig)

	b, _:=checkSig([]byte(msg), sig)
	fmt.Printf("%v\n",b)

	b, _=checkSig([]byte("Changed message"), sig)
	fmt.Printf("%v\n",b)


}

func signMessage(msg []byte)([]byte, error){
	h:=hmac.New(sha256.New, []byte(key))
	_, err:=h.Write(msg)
	if err!=nil{
		return nil, fmt.Errorf("ERROR!!!!! %w",err)
	}
	sig:=h.Sum(nil)
	return sig, nil
}

func checkSig(msg, sig []byte)(bool, error) {
	newSig, err:=signMessage(msg)
	if err!=nil {
		return false, fmt.Errorf("No match")
	}
	same:=hmac.Equal(newSig, sig)
	return same, nil
}