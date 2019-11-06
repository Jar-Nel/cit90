package main

import (
	"net/http"
	"io"
//	"fmt"
)

type user struct {
	Firstname string
	Email string
	Password string
}

var db =map[string]user{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/psignup", psignup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/plogin", plogin)
	http.HandleFunc("/logout", logout)
	//http.HandleFunc("/rc", fReadCookie)

	//http.HandleFunc("/process", fProcess)
	http.ListenAndServe(":808",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	//Try to read cookie
	c, err:=r.Cookie("074-session")
	if err!=nil {
		//if error (no cookie read) not logged in
		http.Redirect(w,r,"/login", http.StatusSeeOther)
		return
	}
	//show it.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>"+c.Name+" "+c.Value+"</h1>")
}

func signup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<form action="/psignup" method="POST">
			<input type="text" id="firstname" name="firstname" />
			<input type="text" id="email" name="email" />
			<input type="password" id="pw" name="pw" />
			<input type="submit" value="Submit" />
		</form>
	</body>
	</html>	
	`)
}

func psignup(w http.ResponseWriter, r *http.Request){
	if (r.Method !=http.MethodPost) {
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	fn:=r.FormValue("firstname")
	email:=r.FormValue("email")
	pw:=r.FormValue("pw")

	if fn=="" {
		http.Error(w, "firstname is required", http.StatusBadRequest)
		return
	}	
	if email=="" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}
	if pw=="" {
		http.Error(w, "password is required", http.StatusBadRequest)
		return
	}

	u:=user{
		Firstname: fn,
		Email: email, 
		Password: pw,
	}
	db[email]=u

	c:=&http.Cookie{
		Name: "074-session",
		Value: email,
		Path:"/",
	}

	http.SetCookie(w, c)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>Created Acct for "+fn+"</h1>")
}

func login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Login</title>
	</head>
	<body>
		<form action="/plogin" method="POST">
			<input type="text" id="email" name="email" />
			<input type="text" id="pw" name="pw" />
			<input type="submit" value="Submit" />
		</form>
	</body>
	</html>	
	`)
}

func plogin(w http.ResponseWriter, r *http.Request){
	if (r.Method !=http.MethodPost) {
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	email:=r.FormValue("email")
	pw:=r.FormValue("pw")

	if email=="" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}
	if pw=="" {
		http.Error(w, "password is required", http.StatusBadRequest)
		return
	}

	if db[email].Password==pw {
		c:=&http.Cookie{
			Name: "074-session",
			Value: email,
			Path:"/",
		}
	
		http.SetCookie(w, c)
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<h1>Invalid email pw combo</h1>")
	}

}

func logout(w http.ResponseWriter, r *http.Request){
	c, err:=r.Cookie("074-session")
	if err!=nil {
		//if error (no cookie read) not logged in
		c=&http.Cookie{
			Name:"074-session",
		}
	}
	c.MaxAge=-1
	http.SetCookie(w, c)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>signed out</h1>")
}


