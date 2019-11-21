package main

//Convert User to JSON.
//Save user to file.
//read users from file.

import (
	"net/http"
	"log"
	"fmt"
	//"os"
	//"encoding/json"
	//"io/ioutil"
	//"strings"

	//Encryption
	//"golang.org/x/crypto/bcrypt"
	//"crypto/hmac"
	//"crypto/sha256"
	//"encoding/hex"
)


func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/home", index)

	http.HandleFunc("/signup", signup)
	//Process SignUp
	http.HandleFunc("/psignup", psignup)

	http.HandleFunc("/login", login)
	//Process Login
	http.HandleFunc("/plogin", plogin)

	http.HandleFunc("/logout", logout)

	http.HandleFunc("/about",about) 
	http.HandleFunc("/contact",contact) 

	fmt.Println("Starting webserver on port 808")
	log.Fatal(http.ListenAndServe(":808",nil))
}

