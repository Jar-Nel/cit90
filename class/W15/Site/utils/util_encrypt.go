package utils

import (
	//Encryption
	//"golang.org/x/crypto/bcrypt"
	"crypto/hmac"
	"crypto/sha256"
	//"encoding/hex"
	"fmt"	
)

const key string ="Keep it secret! Keep it safe!"

/* #region HMAC */

func signMessage(msg []byte)([]byte, error){
	h:=hmac.New(sha256.New, []byte(key))
	_, err:=h.Write(msg)
	if err!=nil{
		return nil, fmt.Errorf("Error in signMessage %w",err)
	}
	sig:=h.Sum(nil)
	return sig, nil
}

func checkSig(msg, sig []byte)(bool, error) {
	newSig, err:=signMessage(msg)
	if err!=nil {
		return false, fmt.Errorf("Error in checkSig : No match")
	}
	same:=hmac.Equal(newSig, sig)
	return same, nil
}

/* #endregion */