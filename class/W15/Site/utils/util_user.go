package utils

import (
	"golang.org/x/crypto/bcrypt"
	//"net/http"
	//url "net/url"
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//User user struct
type User struct {
	Name string
	Email string
	Password string
}

var db =map[string]User{}

//readUsers: Reads users from a json file and puts them in a map. 
func readUsers()(map[string]User){
	//try to read users from file.
	userMap:=map[string]User{}

	jsonFile, err :=os.Open("./data/users.json")
	if err!=nil{
		return userMap
	}
	defer jsonFile.Close()
	bv,_:=ioutil.ReadAll(jsonFile)
	json.Unmarshal(bv, &userMap)

	return userMap
}

//ReadUser: reads users from JSON file and returns matching user (if there is one)
func ReadUser(userEmail string)(User, bool) {
	u:=readUsers()[userEmail]
	if u==(User{}) {
		return u,false
	}
	return u,true
}

//SaveUser: saves user to json file.
func SaveUser(u User)(bool, error) {
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

//CheckPW Checks PW
func CheckPW(pwb[] byte, pws string)(string, bool) {
	err:=bcrypt.CompareHashAndPassword(pwb,[]byte(pws))
	if err!=nil{
		return "Passwords do not match", false
	}
	return "Passwords match", true
}
