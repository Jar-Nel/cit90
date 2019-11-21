package main

import (
	"golang.org/x/crypto/bcrypt"
	//"net/http"
	//url "net/url"
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Email string
	Password string
}

var db =map[string]user{}

func readUsers()(map[string]user){
	//try to read users from file.
	userMap:=map[string]user{}

	jsonFile, err :=os.Open("./data/users.json")
	if err!=nil{
		return userMap
	}
	defer jsonFile.Close()
	bv,_:=ioutil.ReadAll(jsonFile)
	//fmt.Println(string(bv[:]))
	json.Unmarshal(bv, &userMap)
	//fmt.Println(users)

	return userMap

}

func readUser(userEmail string)(user, bool) {
	u:=readUsers()[userEmail]
	if u==(user{}) {
		return u,false
	}
	//fmt.Println("|",u,"|")
	return u,true

}

func saveUser(u user)(bool, error) {
	users:=readUsers()
	pwb,_:=bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password=string(pwb[:])
	users[u.Email]=u
	js,_:=json.Marshal(users)
	file, err :=os.Create("./data/users.json")
	if err!=nil{
		fmt.Println("file error, make this http 500")
		return false, err
	}
	defer file.Close()
	file.Write(js)
	return true,nil
}

func checkPW(pwb[] byte, pws string)(string, bool) {
	err:=bcrypt.CompareHashAndPassword(pwb,[]byte(pws))
	if err!=nil{
		return "Passwords do not match", false
	}
	return "Passwords match", true
}
