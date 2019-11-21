package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

const key string ="lotta crap"

func main() {
	msg:="this is a message"
	sig,_:=signMessage([]byte(msg))
	s:=hex.EncodeToString(sig)
	fmt.Println("Message: ", msg)
	fmt.Printf("Sig: %x\n",sig) //for HW convert both to compare.
	fmt.Println(s)

	sig2,_:=hex.DecodeString("cf81e1ad23566b543c0d2baaaa80a4105c1a52074bf39b83ec0e2ba75c70e463")
	b, _:=checkSig([]byte(msg), []byte(sig2))
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