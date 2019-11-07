package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", fIndex)
	http.HandleFunc("/rc", fReadCookie)

	http.ListenAndServe(":808",nil)
}

func fIndex(w http.ResponseWriter, r *http.Request){
	//Try to read cookie
	c, err:=r.Cookie("01-cookie")
	if err!=nil {
		//if error (no cookie read) write cookie
		c=&http.Cookie{
			Name: "01-cookie",
			Value: "SomeCookieValue",
			Path:"/",
		}
	
	}
	http.SetCookie(w, c)
	//show it.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<h1>Index and Set Cookie</h1>
		<p>A cookie has been set with the value <span style='color:blue;'>`+c.Value+`</span></p>
		<br />
		<p><a href='/rc'>This page reads the cookie.</a></p>
	`)
}

func fReadCookie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	c, err:=r.Cookie("01-cookie")
	if (err!=nil){
		io.WriteString(w, "Error reading cookie")
		return
	}
	io.WriteString(w, `
		<h1>Read Cookie</h1>
		<p>The value read from the cookie is <span style='color:blue;'>`+c.Value+`</span></p>
		<br />
		<p><a href='/'>Go back to Index.</a></p>
	`)

}

