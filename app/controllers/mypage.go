package controllers

import (
	"log"
	"net/http"
	"votalk/app/models"
)

func mypage(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		var user interface{}
			if sess.UserType == "viewer" {
				user, err = sess.GetUserBySessionVw()
			if err != nil {
				log.Println(err,user)
			}
			} else if sess.UserType == "expert" {
				user, err = sess.GetUserBySessionEx()
			if err != nil {
				log.Println(err,user)
			}
			}
			data := struct {
				User interface{}
				S    string
			}{
				user,
				sess.UserType,
			}
			generateHTML(w, data, "layout", "private_navbar", "mypage")
		}
	
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		var user interface{}
			if sess.UserType == "viewer" {
				user, err = sess.GetUserBySessionVw()
				if err != nil {
					log.Println(err,user)
				}
			} else if sess.UserType == "expert" {
				user, err = sess.GetUserBySessionEx()
				if err != nil {
					log.Println(err,user)
				}
			}
			models.DB.Delete(user)
			cookie, err := r.Cookie("_cookie")
			if err != nil {
				log.Println(err)
			}
			if err != http.ErrNoCookie {
				session := models.Session{UUID: cookie.Value}
				session.DeleteSessionByUUID()
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
}
