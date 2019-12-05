package main

//Convert User to JSON.
//Save user to file.
//read users from file.

import (
	"net/http"
	"log"
	"fmt"
	web "main/web"
)


func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.HandleFunc("/", secureRoute)

	http.HandleFunc("/sec", web.SecureRoute)
	http.HandleFunc("/sec/", web.SecureRoute)

	http.HandleFunc("/home", web.SecureRoute)
	http.HandleFunc("/index", web.SecureRoute)

	http.HandleFunc("/signup", web.Signup)
	//Process SignUp
	http.HandleFunc("/psignup", web.Psignup)

	http.HandleFunc("/login", web.Login)
	//Process Login
	http.HandleFunc("/plogin", web.Plogin)

	http.HandleFunc("/logout", web.Logout)

	http.HandleFunc("/about",web.About) 
	http.HandleFunc("/contact",web.Contact) 

	fmt.Println("Starting webserver on port 808")
	log.Fatal(http.ListenAndServe(":808",nil))
}

