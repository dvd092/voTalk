package controllers

import (
	"votalk/app/libs"
	"votalk/app/models"

	// "fmt"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	s := libs.LastUrl(r.URL.String())

	switch r.Method {
	case http.MethodGet:
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, s, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/", 302)
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		//viewer登録
		if s == "viewer" {
			user := models.UserVw{
				Name:     r.PostFormValue("name"),
				Email:    r.PostFormValue("email"),
				Password: r.PostFormValue("password"),
			}

			if err := user.CreateUser(); err != nil {
				log.Println(err)
			}
		}
		//expert登録
		if s == "expert" {
			user := models.UserEx{
				Name:     r.PostFormValue("name"),
				Email:    r.PostFormValue("email"),
				Password: r.PostFormValue("password"),
			}
			if err := user.CreateUser(); err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/", 302)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	s := libs.LastUrl(r.URL.String())
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, s, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	s := r.PostFormValue("lastUrl")

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	//エキスパート処理
	if s == "expert" {
		user, err := models.GetUserByEmailEx(r.PostFormValue("email"), s)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", 302)
		}
		if user.Password == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession(s)
			if err != nil {
				log.Println(err)
			}
			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/articles", 302)
		} else {
			http.Redirect(w, r, "/login/expert", 302)
			log.Println("パスワードが間違っています")
		}
		//ビューワー処理
	} else if s == "viewer" {
		user, err := models.GetUserByEmailVw(r.PostFormValue("email"), s)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login/viewer", 302)
		}
		if user.Password == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession(s)
			if err != nil {
				log.Println(err)
			}

			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/articles", 302)
		} else {
			http.Redirect(w, r, "/login/viewer", 302)
			log.Println("パスワードが間違っています")
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/", 302)
	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}