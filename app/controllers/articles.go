package controllers

import (
	
	"log"
	"net/http"
	"votalk/app/libs"
	"votalk/app/models"
)

func articles(w http.ResponseWriter, r *http.Request) {
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
			arts,err := models.GetArticles()
			if err != nil{
				log.Fatalln(err)
			}
			data := struct {
				User interface{}
				S string
				Art []models.Article
			}{
				user,
				s,
				arts,
			}
			generateHTML(w, data, "layout", "private_navbar", "articles")
		} else if s == "expert" {
			user, err := sess.GetUserBySessionEx()
			if err != nil {
				log.Println(err)
			}
			generateHTML(w, user, "layout", "private_navbar", "articles")
		}
		// todos, _ := user.GetTodosByUser()
		// user.Todos = todos

	}
}

/*
func articles(w http.ResponseWriter, r *http.Request) {
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
				S string
			}{
				user,
				s,
			}
			generateHTML(w, data, "layout", "private_navbar", "articles")
		} else if s == "expert" {
			user, err := sess.GetUserBySessionEx()
			if err != nil {
				log.Println(err)
			}
			generateHTML(w, user, "layout", "private_navbar", "articles")
		}
		// todos, _ := user.GetTodosByUser()
		// user.Todos = todos

	}
}
*/