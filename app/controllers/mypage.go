package controllers

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/sessions"
	"votalk/app/models"
	// "votalk/app/libs"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func mypage(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	session, _ := store.Get(r, "edit_success")
	flashMessages := session.Flashes()
	session.Save(r, w)
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
				Flash interface{}
			}{
				user,
				sess.UserType,
				flashMessages,
			}
			generateHTML(w, data, "layout", "private_navbar", "mypage")
		}
	
}

func mypageEdit(w http.ResponseWriter, r *http.Request) {

	userType := r.FormValue("userType")
	userId := r.FormValue("userId")
	email := r.FormValue("new-email")
	name := r.FormValue("new-name")

	if userType == "viewer" {
		if email == "" {
			err := models.DB.Table("vw_users").Where("id = ?", userId).Update("name", name).Error
			if err != nil { 
				log.Println(err.Error())
			}
		}else {
			err := models.DB.Table("vw_users").Where("id = ?", userId).Update("email", email).Error
		if err != nil { 
			log.Println(err.Error())
		}
		}
	} else if userType == "expert" {
		if email == "" {
			err := models.DB.Table("ex_users").Where("id = ?", userId).Update("name", name).Error
			if err != nil { 
				log.Println(err.Error())
			}
		}else {
			err := models.DB.Table("ex_users").Where("id = ?", userId).Update("email", email).Error
		if err != nil {
			log.Println(err.Error())
		}
		}
		
	}

	session, _ := store.Get(r, "edit_success")
	session.AddFlash("変更が保存されました。")
	session.Save(r, w)

	http.Redirect(w, r, fmt.Sprintf("/%s/mypage", userType), http.StatusFound)	
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
