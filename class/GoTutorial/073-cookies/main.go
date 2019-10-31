package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", fIndex)
	http.HandleFunc("/rc", fReadCookie)

	//http.HandleFunc("/process", fProcess)
	http.ListenAndServe(":808",nil)
}

func fIndex(w http.ResponseWriter, r *http.Request){
	//Try to read cookie
	c, err:=r.Cookie("073-cookie")
	if err!=nil {
		//if error (no cookie read) write cookie
		c=&http.Cookie{
			Name: "073-cookie",
			Value: "JamesBondUUID",
			Path:"/",
		}
	
	}
	http.SetCookie(w, c)
	//show it.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>"+c.Name+" "+c.Value+"</h1>")
}

func fReadCookie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	c, err:=r.Cookie("073-cookie")
	if (err!=nil){
		io.WriteString(w, "Error reading cookie")
	}
	io.WriteString(w, "<h1>"+c.Name+" "+c.Value+"</h1>")

}
