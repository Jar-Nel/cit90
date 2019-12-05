package web

import (
	"net/http"
	url "net/url"
	"io"
	"text/template"
	//"html/template"  //text/template with protections against code injection.
	"fmt"
	"strings"
	"os"
	utils "main/utils"
)

var tpl *template.Template

type pageData struct {
	User utils.User
	LoggedIn bool
	Title string
	Heading string
	Account string
	Body string
}

/*func init(){
	//parseglob parses everything in a location
	tpl=template.Must(template.ParseGlob("./templates/*.gohtml"))
}*/

//Security Routine.  Gets user if exists.
func getUser(w http.ResponseWriter, r *http.Request) (utils.User, bool) {
	var u utils.User
	c, err:=r.Cookie("02-session")
	if err!=nil {
		//if error (no cookie read) not logged in
		http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
		return u, false
	}
	
	//we have a cookie, check that it is valid (HMAC)
	if (utils.ValidCookie(c.Value)){
		//cookie is valid, get the user
		if u,ok:=utils.ReadUser(strings.Split(c.Value,"|")[0]); ok {
			return u, true
		}
	}

	http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
	return u, false
}

//look at maps to functions one for secure, one for unsecure, default to static dir
//consider adding group membership
//also add session mgmt to user.

func SecureRoute(w http.ResponseWriter, r *http.Request){
	fmt.Println(strings.ReplaceAll(r.URL.EscapedPath(), "/", ""))
	if u, ok:=getUser(w,r); ok {
		switch(strings.ReplaceAll(r.URL.EscapedPath(), "/", "")){
		case "index":
			fallthrough
		case "home":
			Index(w, r, u)
		default:
			fmt.Println("./static"+r.URL.EscapedPath())
			http.ServeFile(w,r,"./static"+r.URL.EscapedPath())

		}
		return
	}
	fmt.Println("./static"+r.URL.EscapedPath())
	fileReq, err :=os.Open("./static"+r.URL.EscapedPath())
	if err!=nil {
		fmt.Println(err.Error())
		//http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer fileReq.Close()
	io.Copy(w,fileReq)
}

func Index(w http.ResponseWriter, r *http.Request, u utils.User){
	//show it.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data:=pageData {
		User: u,
		LoggedIn: true,
		Title: "Home",
		Heading: "Welcome to the site.",
		Body: "",
	}
	t:=getTemplates("index")
	templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
}

func About(w http.ResponseWriter, r *http.Request) {
	if u, ok:=getUser(w,r); ok {
		//show it.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data:=pageData {
			User: u,
			LoggedIn: true,
			Title: "About",
			Heading: "About this site",
			Body: `
				<p>This site is served by Go.</p>
			`,
		}
		t:=getTemplates("none")
		templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if u, ok:=getUser(w,r); ok {
		//show it.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data:=pageData {
			User: u,
			LoggedIn: true,
			Title: "Contact",
			Heading: "Contact Us",
			Body: `
				<p>If we wanted to hear from you, this is where you would find out how to contact us.</p>
			`,
		}
		t:=getTemplates("none")
		templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
	}
}

func Signup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data:=pageData {
		Title: "SignUp",
		LoggedIn: false,
		Heading: "Sign Up",
		Body: "",
	}
	t:=getTemplates("signup")
	templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
}

//process sign up.
func Psignup(w http.ResponseWriter, r *http.Request){
	if (r.Method !=http.MethodPost) {
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	fn:=r.FormValue("firstname")
	email:=r.FormValue("email")
	pw:=r.FormValue("pw")

	if fn=="" {
		http.Error(w, "Name is required", http.StatusBadRequest)
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

	if _,ok:=utils.ReadUser(email); ok {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return
	}
		
	u:=utils.User{
		Name: fn,
		Email: email, 
		Password: pw,
	}
	
	ok, err:=utils.SaveUser(u)
	if err!=nil{
		http.Error(w, "Error saving user", http.StatusBadRequest)
	}
	if ok {
		http.SetCookie(w, utils.CreateCookie(email))

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data:=pageData {
			Title: "SignUp",
			LoggedIn: false,
			Heading: "Sign Up",
			Account: "",
			Body: `
				<h1>Congratulations!</h1>
				<h1>Your account has been created</h1>
				<br />
				<p><a href="/home">Take me to the site!</a></p>
			`,
		}
		t:=getTemplates("none")
		templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))	
	}
	//Failed to create user
	http.Error(w, "Error Saving user", http.StatusBadRequest)

}

func Login(w http.ResponseWriter, r *http.Request) {
	loc:=r.URL.Query().Get("l")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data:=pageData {
		Title: "Login",
		LoggedIn: false,
		Heading: "Login",
		Account: "",
		Body: url.QueryEscape(loc),
	}
	t:=getTemplates("login")
	templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
}

/*func login(w http.ResponseWriter, r *http.Request){
	loc:=r.URL.Query().Get("l")
	
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
		<form action="/plogin?l=`+url.QueryEscape(loc)+`" method="POST">
			<input type="text" id="email" name="email" />
			<input type="text" id="pw" name="pw" />
			<input type="submit" value="Submit" />
		</form>
	</body>
	</html>	
	`)
}*/

func Plogin(w http.ResponseWriter, r *http.Request){
	if (r.Method !=http.MethodPost) {
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	loc:=r.URL.Query().Get("l")
	if (len(loc)==0){
		loc="/home"
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

	if u, ok:=utils.ReadUser(email); ok{
		if _, ok:=utils.CheckPW([]byte(u.Password), pw); ok{
			http.SetCookie(w, utils.CreateCookie(email))
			
			//sig,_:=signMessage([]byte(email))
			//sigHex:=hex.EncodeToString(sig)
			//cookieValue:=fmt.Sprintf("%s|%s",email,sigHex)
			//fmt.Println("logged in:",msg)
			//c:=&http.Cookie{
			//	Name: "02-session",
			//	Value: cookieValue,
			//	Path:"/",
			//}

			//http.SetCookie(w, c)
			http.Redirect(w,r,loc, http.StatusSeeOther)
			return		
		}

	}

	/*
	if db[email].Password==pw {
		c:=&http.Cookie{
			Name: "02-session",
			Value: email,
			Path:"/",
		}
	
		http.SetCookie(w, c)
		http.Redirect(w,r,loc, http.StatusSeeOther)
		return
	} 
	*/

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data:=pageData {
		Title: "Login",
		LoggedIn: false,
		Heading: "Login",
		Account: "",
		Body: `
			<h1>Invalid email and password combo</h1>
			<br />
			<p><a href="/login?l=`+loc+`">Return to login screen.</a></p>
		`,
	}
	t:=getTemplates("none")
	templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
}

func Logout(w http.ResponseWriter, r *http.Request){
	c, err:=r.Cookie("02-session")
	if err!=nil {
		//if error (no cookie read) not logged in
		c=&http.Cookie{
			Name:"02-session",
		}
	}
	c.MaxAge=-1
	http.SetCookie(w, c)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data:=pageData {
		Title: "Logout",
		LoggedIn: false,
		Heading: "Logout",
		Account: "",
		Body: `
			<h1>You have been signed out</h1>
			<br />
			<p><a href="/home">Sign In</a></p>
		`,
	}
	t:=getTemplates("none")
	templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
}

func getTemplates(tplContent string) *template.Template {
	files:=[]string {
		"./templates/content/"+tplContent+".gohtml",
		"./templates/body.gohtml",
		"./templates/footer.gohtml",
		"./templates/header.gohtml",
		"./templates/site.gohtml",
	}
	t,_:=template.ParseFiles(files...)
	return t
}

func templateErr(w http.ResponseWriter, err error){
	if err!=nil{
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}
