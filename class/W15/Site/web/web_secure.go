package web

import (
	"net/http"
	url "net/url"
	utils "main/utils"
)

type SecureHttpHandlerFunc func(http.ResponseWriter, *http.Request, utils.User)
type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

/*type SecureWebHandler struct {
	Next http.Handler
}

//add ServeHTTP function to SecureWebHandler
func (swh SecureWebHandler) ServeHTTP (w http.ResponseWriter, r * http.Request) {
	//Check User Security
	if _, ok:=getUser(w,r); ok {
		swh.Next.ServeHTTP(w, r)
	} else {
		//redirect to login
		http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
	}
}*/

func SecureWeb(f SecureHttpHandlerFunc)  HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if u, ok:=getUser(w,r); ok {
			f(w, r, u) // original function call
		} else {
			//redirect to login
			http.Redirect(w,r,"/login?l="+url.QueryEscape(r.URL.String()), http.StatusSeeOther)
		}
	}
}
   
