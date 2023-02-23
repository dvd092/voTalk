package controllers

import (
	"fmt"
	"html/template"
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
	templates.ExecuteTemplate(w, "layout", data)
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

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

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
	http.HandleFunc("/expert/index", index)//記事一覧
	http.HandleFunc("/expert/mypage", mypage)
	http.HandleFunc("/expert/articles", index)
	// viewerページ
	http.HandleFunc("/viewer/index", index)
	http.HandleFunc("/viewer/mypage", mypage)
		// viewer機能ページ
		http.HandleFunc("/viewer/matches", index)
		http.HandleFunc("/viewer/", index)
			// viewerマッチページ
			http.HandleFunc("/viewer/match/new", index)
			http.HandleFunc("/viewer/match/save", index)
			http.HandleFunc("/viewer/match/edit", index)
			http.HandleFunc("/viewer/match/update", index)

	// common
		//公開記事
		http.HandleFunc("/viewer/articles", articles)
		// http.HandleFunc("article/{id}", article)
		// 公開討論
		http.HandleFunc("matches", index)
		http.HandleFunc("matches/{id}", index)
		http.HandleFunc("matches/{topic_id}", index)//タグ付けされたexpertのみ討論可能

	// http.HandleFunc("/todos/new", todoNew)
	// http.HandleFunc("/todos/save",todoSave)
	// http.HandleFunc("/todos/edit/",parseURL(todoEdit))
	// http.HandleFunc("/todos/update/",parseURL(todoUpdate))
	// http.HandleFunc("/todos/delete/",parseURL(todoDelete))
	return http.ListenAndServe("127.0.0.1:"+config.Config.Port, nil)
}
