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
			user, err := sess.GetUserBySessionVw()
			s := libs.GetUTypeFromSess(sess)
			if err != nil {
				log.Println(err)
			}
			arts,err := models.GetArticles()
			models.DB.Preload("ExUser").Find(&arts)
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

	}
}

func article(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		s := libs.GetUTypeFromSess(sess);
			user, err := sess.GetUserBySessionVw()
			if err != nil {
				log.Println(err)
			}
			art, err := models.GetArticle(models.DB,id)
			models.DB.Preload("ExUser").First(&art)
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
		}
}

func likeButton(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// likes増分処理
		articleID := r.FormValue("articleId")
		art := models.Article{}
		record := models.DB.Where("id = ?", articleID).First(&art)
		record.Update("likes", art.Likes + 1)
		// like_num処理
		userID := r.FormValue("userId")
		user := models.UserVw{}
		recordUser := models.DB.Table("vw_users").Where("id = ?", userID).First(&user)
		recordUser.Update("like_num", user.LikeNum - 1)

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{"likes": art.Likes}
		json.NewEncoder(w).Encode(response)
	}
}
