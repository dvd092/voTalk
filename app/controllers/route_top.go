package controllers

import (
	"fmt"
	"log"
	"net/http"
	"votalk/app/libs"
)

func top(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, fmt.Sprintf("/%s/mypage", s.UserType), 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		if s := libs.SecondLastUrl(r.URL.String()); s == "viewer" {
			user, err := sess.GetUserBySessionVw()
			if err != nil {
				log.Println(err)
			}
			data := struct {
				User interface{}
				S    string
			}{
				user,
				s,
			}
			generateHTML(w, data, "layout", "private_navbar", "index")
		} else if s == "expert" {
			user, err := sess.GetUserBySessionEx()
			if err != nil {
				log.Println(err)
			}
			data := struct {
				User interface{}
				S    string
			}{
				user,
				s,
			}
			generateHTML(w, data, "layout", "private_navbar", "index")
		}
	}
}
