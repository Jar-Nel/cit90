package utils

import (
	"net/http"
	"strings"
	"encoding/hex"
	"fmt"
)

//ValidCookie returns true if the cookie is valid
func ValidCookie(cv string) bool {
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

//CreateCookie creates a new cookie
func CreateCookie(cv string) *http.Cookie {
	sig,_:=signMessage([]byte(cv))
	sigHex:=hex.EncodeToString(sig)
	cookieValue:=fmt.Sprintf("%s|%s",cv,sigHex)

	c:=&http.Cookie{
		Name: "02-session",
		Value: cookieValue,
		Path:"/",
	}

	return c
}
