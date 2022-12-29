package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("key")
	store = sessions.NewCookieStore(key)
)

func AuthenticatedCookie(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(w, "the cake is a lie")
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/cookie", AuthenticatedCookie)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)

	http.ListenAndServe(":9090",nil)
}
