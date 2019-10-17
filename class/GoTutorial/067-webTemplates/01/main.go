package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"net/http"
	"log"
	"strconv"
	"strings"
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
	<body>##NAMES##
	</body>
	</html>
	`
	file, err := os.Open("./template/index.html")
    if err != nil {
		log.Print(err)
		return
    }
	defer file.Close()
	b, err := ioutil.ReadAll(file)
    if err != nil {
		log.Print(err)
		return
	}
	
	hBody=strings.Replace(string(b), "##NAMES##",s, -1)

	//fmt.Fprintf(w, "Fmt fprintf\n")
	io.WriteString(w, hBody)
}