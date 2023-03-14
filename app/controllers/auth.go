package controllers

import (
	"fmt"
	"log"
	"net/http"
	"votalk/app/libs"
	"votalk/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	s := libs.LastUrl(r.URL.String())

	switch r.Method {
	case http.MethodGet:
		_, err := session(w, r)
		// 登録失敗メッセージ
		session, _ := store.Get(r, "sign_up_failed")
		flashMessages := session.Flashes()
		session.Save(r, w)
		data := struct {
			S     string
			Flash interface{}
		}{
			s,
			flashMessages,
		}
		if err != nil {
			generateHTML(w, data, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/", 302)
		}

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//viewer登録
		if s == "viewer" {
			user := models.UserVw{
				UUID: models.CreateUUID().String(),
				Name:     r.PostFormValue("name"),
				Email:    r.PostFormValue("email"),
				Password: models.Encrypt(r.PostFormValue("password")),
				IsValid: 1,
				LikeNum: 1,
			}

			user_vw := models.UserVw{}

			err = models.DB.Where("email = ?", user.Email).First(&user_vw).Error
			if err == nil {
				session, _ := store.Get(r, "sign_up_failed")
				session.AddFlash("メールアドレスはすでに登録されています")
				session.Save(r, w)
				http.Redirect(w, r, "/signup/viewer", 302)
				return
			}
			err = models.DB.Where("name = ?", user.Name).First(&user_vw).Error
			if err == nil {
				session, _ := store.Get(r, "sign_up_failed")
				session.AddFlash("名前はすでに登録されています")
				session.Save(r, w)
				http.Redirect(w, r, "/signup/viewer", 302)
				return
			}

			if err := user.CreateUser(); err != nil {
				log.Println(err)
			}
		}
		//expert登録
		if s == "expert" {
			user := models.UserEx{
				Name:     r.PostFormValue("name"),
				Email:    r.PostFormValue("email"),
				Password: models.Encrypt(r.PostFormValue("password")),
				IsValid: 1,
			}

			err = models.DB.Where("email = ?", user.Email).First(user).Error
			if err == nil {
				session, _ := store.Get(r, "sign_up_failed")
				session.AddFlash("メールアドレスはすでに登録されています")
				session.Save(r, w)
				http.Redirect(w, r, "/signup/expert", 302)
			}
			err = models.DB.Where("name = ?", user.Name).First(user).Error
			if err == nil {
				session, _ := store.Get(r, "sign_up_failed")
				session.AddFlash("名前はすでに登録されています")
				session.Save(r, w)
				http.Redirect(w, r, "/signup/expert", 302)
			}

			if err := user.CreateUser(); err != nil {
				log.Println(err)
			}
		}
		url := fmt.Sprintf("/login/%s",s)
		http.Redirect(w, r, url, 302)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	s := libs.LastUrl(r.URL.String())
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, s, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	s := r.PostFormValue("lastUrl")

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	//エキスパート処理
	if s == "expert" {
		user, err := models.GetUserByEmailEx(r.PostFormValue("email"), s)
		if user.IsValid == 0 {
			log.Println("有効なアカウントではありません")
			http.Redirect(w, r, "/login/expert", 302)
		}

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", 302)
		}
		if user.Password == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession(s)
			if err != nil {
				log.Println(err)
			}
			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/articles", 302)
		} else {
			http.Redirect(w, r, "/login/expert", 302)
			log.Println("パスワードが間違っています")
		}
		//ビューワー処理
	} else if s == "viewer" {
		user, err := models.GetUserByEmailVw(r.PostFormValue("email"), s)
		if user.IsValid == 0 {
			log.Println("有効なアカウントではありません")
			http.Redirect(w, r, "/login/viewer", 302)
		}
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login/viewer", 302)
		}
		if user.Password == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession(s)
			if err != nil {
				log.Println(err)
			}

			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/articles", 302)
		} else {
			http.Redirect(w, r, "/login/viewer", 302)
			log.Println("パスワードが間違っています")
		}
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
	http.Redirect(w, r, "/", 302)
	w.Write([]byte("Old cookie deleted. Logged out!\n"))
}
