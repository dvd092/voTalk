package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"votalk/app/models"
	"votalk/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/article/(edit|update|delete|show)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// top
	http.HandleFunc("/", top)
	// auth
	http.HandleFunc("/signup/expert", signup)
	http.HandleFunc("/signup/viewer", signup)
	http.HandleFunc("/login/expert", login)
	http.HandleFunc("/login/viewer", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	// expertページ
	http.HandleFunc("/expert/index", index) //記事一覧
	http.HandleFunc("/expert/mypage", mypage)
	http.HandleFunc("/expert/mypage/edit", mypageEdit)
	http.HandleFunc("/expert/delete", deleteUser)
	http.HandleFunc("/expert/articles/mine", myArticles)
	// expert記事作成
	http.HandleFunc("/expert/article/new", newArticles)
	// viewerページ
	http.HandleFunc("/viewer/index", index)
	http.HandleFunc("/viewer/mypage", mypage)
	http.HandleFunc("/viewer/mypage/edit", mypageEdit)
	http.HandleFunc("/viewer/delete", deleteUser)
	// viewer機能ページ
	// http.HandleFunc("/viewer/matches", index)
	// viewer記事ページ
	http.HandleFunc("/articles", articles)

	// viewerマッチページ
	// http.HandleFunc("/viewer/match/new", index)
	// http.HandleFunc("/viewer/match/save", index)
	// http.HandleFunc("/viewer/match/edit", index)
	// http.HandleFunc("/viewer/match/update", index)

	// common
	//公開記事

	http.HandleFunc("/article/show/", parseURL(article))
	http.HandleFunc("/article/delete/", parseURL(deleteArticle))
	http.HandleFunc("/article/edit/", parseURL(editArticle))
	http.HandleFunc("/like-article", likeButton)
	// 公開討論
	// http.HandleFunc("matches", index)
	// http.HandleFunc("matches/{id}", index)
	// http.HandleFunc("matches/{topic_id}", index) //タグ付けされたexpertのみ討論可能

	return http.ListenAndServe("127.0.0.1:"+config.Config.Port, nil)
}
