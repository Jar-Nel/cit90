package main

import (
	"fmt"
	"io"
	"net/http"
	"log"
	"strconv"
)

func main(){

	http.HandleFunc("/",foo)
	log.Fatal(http.ListenAndServe(":808", nil))
}

func foo(w http.ResponseWriter, r *http.Request){
	xs:=[]string{"John","Jacob","Smith"}
	s:=""
	for i,v:=range xs {
		s+=fmt.Sprintf("<div><span>%v:&nbsp;</span><span>%s</span></div>",i,v)
		s+="<div><span>"+strconv.Itoa(i)+":&nbsp;</span><span>"+v+"</span>";
	}

	hBody:=`
	<html>
	<head>
	</head>
	<body>`+s+`
	</body>
	</html>
	`

	//fmt.Fprintf(w, "Fmt fprintf\n")
	io.WriteString(w, hBody)
}