package controllers

import (
	"expert/app/libs"
	"expert/app/models"
	// "fmt"
	"log"
	"net/http"
	// "net/url"
)

func signup(w http.ResponseWriter, r *http.Request) {
	s := libs.LastUrl(r.URL.String())

	switch r.Method {
	case http.MethodGet:
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, s, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	case http.MethodPost:
		log.Println(s)
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		log.Println(s)
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

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
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
	http.Redirect(w, r, "/login", 302)
	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}
