package main

import (
	"net/http"
	url "net/url"
	//"io"
	"text/template"
	//"html/template"  //text/template with protections against code injection.
	"log"
	"fmt"
)

type user struct {
	Firstname string
	Email string
	Password string
}

var db =map[string]user{}

type pageData struct {
	User user
	LoggedIn bool
	Title string
	Heading string
	Account string
	Body string
}

var tpl *template.Template

/*func init(){
	//parseglob parses everything in a location
	tpl=template.Must(template.ParseGlob("./templates/*.gohtml"))
}*/

func main() {
	//Set up location for static content
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/home", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/psignup", psignup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/plogin", plogin)
	http.HandleFunc("/logout", logout)

	http.HandleFunc("/about",about) 
	http.HandleFunc("/contact",contact) 

	//http.HandleFunc("/rc", fReadCookie)

	fmt.Println("Starting webserver on port 808")
	log.Fatal(http.ListenAndServe(":808",nil))
}

func getUser(w http.ResponseWriter, r *http.Request) (user, bool) {
	var u user
	c, err:=r.Cookie("02-session")
	if err!=nil {
		//if error (no cookie read) not logged in
		http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
		return u, false
	}
	if u, ok:=db[c.Value]; ok {
		return u, true
	}

	http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
	return u, false
}

func index(w http.ResponseWriter, r *http.Request){
	//Try to read cookie
	//c, err:=r.Cookie("02-session")
	//if err!=nil {
		//if error (no cookie read) not logged in
	//	http.Redirect(w,r,"/login", http.StatusSeeOther)
		//return
	//}
	if u, ok:=getUser(w,r); ok {
		//show it.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//btpl:=template.ParseFiles("index.gohtml")
		//tpl.
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
}

func about(w http.ResponseWriter, r *http.Request) {
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

func contact(w http.ResponseWriter, r *http.Request) {
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

func signup(w http.ResponseWriter, r *http.Request){
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
		Name: "02-session",
		Value: email,
		Path:"/",
	}

	http.SetCookie(w, c)

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

func login(w http.ResponseWriter, r *http.Request) {
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

func plogin(w http.ResponseWriter, r *http.Request){
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

func logout(w http.ResponseWriter, r *http.Request){
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