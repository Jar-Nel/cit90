package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const key string ="Keep it secret! Keep it safe!"

func main() {
	//The email
	email:="username@domain.com"
	//Get HMAC Signing for email
	sig,_:=signMessage([]byte(email))
	//convert byte sig to hex string for storing in cookie value
	sigHex:=hex.EncodeToString(sig)
	//create cookie value string
	cookieValue:=fmt.Sprintf("%s|%s",email,sigHex)

	fmt.Println("Email: ", email)
	fmt.Printf("Sig: %s\n",sigHex) 
	fmt.Println("Cookie Value: ", cookieValue)

	//Is the cookie value valid?
	fmt.Printf("ValidCookie: %v\n",validCookie(cookieValue))

	cookieValue=fmt.Sprintf("tamperedemail@domain.com|%s",sigHex)
	fmt.Printf("ValidCookie: %v\n",validCookie(cookieValue))


}

func validCookie(cv string) bool {
	xs:=strings.Split(cv,"|")
	if len(xs)==2 {
		email:=xs[0]
		//Get the hex signature from the cookie value
		sigHex:=xs[1]
		//convert into []byte for verification
		sig,_:=hex.DecodeString(sigHex)
		//check the signature
		b,_:=checkSig([]byte(email), sig)
		return b
	}
	return false
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