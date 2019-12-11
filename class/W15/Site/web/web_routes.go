package web

import (
	"net/http"
	url "net/url"
	//"io"
	"text/template"
	//"html/template"  //text/template with protections against code injection.
	//"fmt"
	"strings"
	//"os"
	utils "main/utils"
)

var tpl *template.Template

//the standard page data structure for use with temnplates
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
	c, err:=r.Cookie(utils.CookieName)
	if err!=nil {
		//if error (no cookie read) not logged in
		//http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
		return u, false
	}
	
	//we have a cookie, check that it is valid (HMAC)
	if (utils.ValidCookie(c.Value)){
		//cookie is valid, get the user
		if u,ok:=utils.ReadUser(strings.Split(c.Value,"|")[0]); ok {
			return u, true
		}
	}

	//http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
	return u, false
}

//Index: Gets the index template and content for the site
func Index(w http.ResponseWriter, r *http.Request, u utils.User){
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


//About: Gets the about content for the site
func About(w http.ResponseWriter, r *http.Request, u utils.User) {
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
	t:=getTemplates("none")  //A value of none loads a generic template to hold the body
	templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
}


//Contact: Gets the Contact template and content for the site
func Contact(w http.ResponseWriter, r *http.Request, u utils.User) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data:=pageData {
		User: u,
		LoggedIn: true,
		Title: "Contact",
		Heading: "Contact Us",
		Body: "",
	}
	t:=getTemplates("contact")
	templateErr(w, t.ExecuteTemplate(w, "site.gohtml", data))
}


//Signup: Gets the Signup template and content for the site
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

//Psignup: Processes the signup form on submit
func Psignup(w http.ResponseWriter, r *http.Request){
	if (r.Method !=http.MethodPost) {
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	//Get user sign up data from the form
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
		http.Error(w, "Error saving user", http.StatusInternalServerError)
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

//Login: gets the login template and login form for the site
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

//Plogin: Process the login form on submit
func Plogin(w http.ResponseWriter, r *http.Request){
	if (r.Method !=http.MethodPost) {
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	loc:=r.URL.Query().Get("l")
	if (len(loc)==0){
		loc="/home"
	}

	//Get email and password from the submitted form
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
			http.Redirect(w,r,loc, http.StatusSeeOther)
			return		
		}

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

//Logout:  Expires cookie and logs out
func Logout(w http.ResponseWriter, r *http.Request){
	c, err:=r.Cookie(utils.CookieName)
	if err!=nil {
		//if error (no cookie read) not logged in
		c=&http.Cookie{
			Name:utils.CookieName,
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

//getTemplates: loads a sub template "on demand" and builds a complete template for the site.
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

//templateErr: error handling if template execution fails.
func templateErr(w http.ResponseWriter, err error){
	if err!=nil{
		http.Error(w, "couldn't load template", http.StatusInternalServerError)
	}
}
