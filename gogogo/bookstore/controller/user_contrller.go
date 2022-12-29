package controller

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/repository"
	"gogogo/bookstore/utils"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := repository.IsLogin(r)
	if flag {
		GetPageBookByPrice(w, r)
		log.Println("没有登录")
	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		log.Println(username, password)
		user, _ := repository.CheckUsernameAndPassword(username, password)
		if user.Id > 0 {
			uuid := utils.CreateUUID()
			sess := &model.Session{
				SessionId: uuid,
				Username:  user.Username,
				UserId:    user.Id,
			}
			repository.AddSession(sess)
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或者密码不正确")
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		repository.DelSession(cookieValue)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	GetPageBookByPrice(w, r)
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := repository.CheckUsername(username)
	if user.Id > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在")
	} else {
		repository.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("usernmae")
	user, _ := repository.CheckUsername(username)
	if user.Id > 0 {
		w.Write([]byte("用户名已经存在"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))
	}
}
