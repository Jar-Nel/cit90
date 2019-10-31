package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", fIndex)
	http.HandleFunc("/process", fProcess)
	http.ListenAndServe(":808",nil)
}

func fIndex(w http.ResponseWriter, r *http.Request){
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
		<form action="/process" method="POST">
			<input type="text" id="first" name="first" />
			<input type="submit" value="Submit" />
		</form>
	</body>
	</html>	
	`)
}

func fProcess(w http.ResponseWriter, r *http.Request){
	if (r.Method !=http.MethodPost) {
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	f:=r.FormValue("first")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>"+f+"</h1>")
}