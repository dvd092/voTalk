package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"html/template"
	"votalk/app/models"
	"strconv"
)

func articles(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		var user interface{}
		// ユーザータイプ別にセッションから情報取得
		if sess.UserType == "expert" {
			user, err = sess.GetUserBySessionEx()
			if err != nil {
				log.Fatalln(err,user)
			}
		} else if sess.UserType == "viewer" {
			user, err = sess.GetUserBySessionVw()
			if err != nil {
				log.Fatalln(err,user)
			}
		}
			arts,err := models.GetArticles()
			models.DB.Preload("ExUser").Order("likes desc").Find(&arts)
			if err != nil{
				log.Fatalln(err)
			}
			data := struct {
				User interface{}
				S string
				Art []models.Article
			}{
				user,
				sess.UserType,
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
		var user interface{}
		// ユーザータイプ別にセッションから情報取得
		if sess.UserType == "expert" {
			user, err = sess.GetUserBySessionEx()
			if err != nil {
				log.Fatalln(err,user)
			}
		} else if sess.UserType == "viewer" {
			user, err = sess.GetUserBySessionVw()
			if err != nil {
				log.Fatalln(err,user)
			}
		}
			if err != nil {
				log.Println(err)
			}
			art, err := models.GetArticle(id)
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
				sess.UserType,
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

func myArticles(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
		var user interface{}
		// ユーザータイプ別にセッションから情報取得
			user, err = sess.GetUserBySessionEx()
			if err != nil {
				log.Fatalln(err,user)
			}
			arts,err := models.GetArticlesByUser(sess.UserId)
			if err != nil {
				log.Fatalln(err,arts)
			}
			data := struct {
				User interface{}
				S string
				Art []models.Article
			}{
				user,
				sess.UserType,
				arts,
			}
			generateHTML(w, data, "layout", "private_navbar", "my_articles")
		}
}

func newArticles(w http.ResponseWriter, r *http.Request) {
	
	switch r.Method {
	case http.MethodGet:
	sess, err := session(w, r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 302)
	} else {
			user, err := sess.GetUserBySessionEx()
			if err != nil {
				log.Fatalln(err,user)
			}
			data := struct {
				User interface{}
				S string
			}{
				user,
				sess.UserType,
			}
	generateHTML(w, data, "layout", "private_navbar", "article_new")
		}
	case http.MethodPost:
		sess,err := session(w, r)
		if err != nil {
			log.Fatalln(err)
		}
		categoryId,err := strconv.Atoi(r.FormValue("categoryId"))
		art := models.Article{
			Title : r.FormValue("title"),
			Plot : template.HTML(r.FormValue("text")),
			CategoryId : categoryId,
			UserExID : sess.UserId,
			Likes: 0,
	}
	models.DB.Create(&art)
		
	http.Redirect(w, r, "/expert/articles/mine", http.StatusFound)
}
}

func deleteArticle(w http.ResponseWriter, r *http.Request, id int) {
	var art models.Article
	models.DB.Where("id = ?", id).First(&art)
	models.DB.Delete(art)
	http.Redirect(w, r, "/expert/articles/mine", http.StatusFound)
}


func editArticle(w http.ResponseWriter, r *http.Request, id int) {
	switch r.Method{
	case http.MethodGet:
		sess, err := session(w, r)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
				art,err := models.GetArticle(id)
				if err != nil {
					log.Fatalln(err)
				} 
				user, err := sess.GetUserBySessionEx()
				if err != nil {
					log.Fatalln(err,user)
				}
				data := struct {
					User interface{}
					S string
					Art models.Article
				}{
					user,
					sess.UserType,
					art,
				}
		generateHTML(w, data, "layout", "private_navbar", "article_new")
			}
		case http.MethodPost:
			id := r.FormValue("id")
			title := r.FormValue("title")
			categoryId,err := strconv.Atoi(r.FormValue("categoryId"))
				if err != nil {
					log.Printf(err.Error())
				}
			
			Plot := template.HTML(r.FormValue("text"))
			models.DB.Table("articles").Where("id = ?",id).Updates(models.Article{Title: title, Plot: Plot, CategoryId: categoryId})
			
			http.Redirect(w, r, "/expert/articles/mine", http.StatusFound)
	}
	

	
}
