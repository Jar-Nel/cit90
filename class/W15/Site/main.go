package main

//Convert User to JSON.
//Save user to file.
//read users from file.

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	web "main/web"
	//utils "main/utils"
)

func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//Set up secure routes
	secureRoutes:=map[string]web.SecureHttpHandlerFunc{}
	secureRoutes["/home"]=web.Index
	secureRoutes["/index"]=web.Index
	secureRoutes["/about"]=web.About
	secureRoutes["/contact"]=web.Contact

	for k,v:= range secureRoutes {
		http.HandleFunc(k,web.SecureWeb(v))
		if !strings.HasSuffix(k, "/"){
			k+="/"
			http.HandleFunc(k,web.SecureWeb(v))
		}
	}


	//Set up anon routes
	routes:=map[string]web.HttpHandlerFunc{}
	//Sign up routes, signup form and process signup form
	routes["/signup"]=web.Signup
	routes["/psignup"]=web.Psignup
	//Log in routes, login form and process login form
	routes["/login"]=web.Login
	routes["/plogin"]=web.Plogin
	//log out route
	routes["/logout"]=web.Logout

	for k,v:= range routes {
		http.HandleFunc(k,v)
		if !strings.HasSuffix(k, "/"){
			k+="/"
			http.HandleFunc(k,v)
		}
	}


	fmt.Println("Starting webserver on port 808")
	log.Fatal(http.ListenAndServe(":808",nil))
}

