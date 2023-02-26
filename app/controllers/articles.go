package controllers

import (
	"log"
	"net/http"
	"votalk/app/libs"
	"votalk/app/models"
	"encoding/json"

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
	}
}

func article(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		if s := libs.GetUTypeFromSess(sess); s == "viewer" {
			user, err := sess.GetUserBySessionVw()
			if err != nil {
				log.Println(err)
			}
			art, err := models.GetArticle(models.DB,id)
			if err != nil{
				log.Fatalln(err)
			}
			data := struct {
				User interface{}
				S string
				Art models.Article
			}{
				user,
				s,
				art,
			}
			generateHTML(w, data, "layout", "private_navbar", "article")
		} else if s == "expert" {
			user, err := sess.GetUserBySessionEx()
			if err != nil {
				log.Println(err)
			}
			generateHTML(w, user, "layout", "private_navbar", "articles")
		}
	
	}
}

func likeButton(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		articleID := r.FormValue("articleId")
		art := models.Article{}
		record := models.DB.Where("id = ?", articleID).First(&art)
		record.Update("likes", art.Likes + 1)

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"likes": art.Likes}
		json.NewEncoder(w).Encode(response)
	}
}
