package controllers

import (
	// "context"
	// "os"
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
	"votalk/app/models"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)


// loginHandler handles the third-party login process.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	models.EnvLoad()
	gomniauth.SetSecurityKey(os.Getenv("O_AUTH_SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(os.Getenv("O_AUTH_CLIENT_ID"), os.Getenv("O_AUTH_SECRET_KEY"), os.Getenv("REDIRECT_URL")),
	)

	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := "google"

	if action == "login" {
			// ユーザータイプクッキー登録
		cookie := &http.Cookie{
				Name: "user_type", 
				Value: segs[3], 
		}
		http.SetCookie(w, cookie)


		UserTypeCookieValue := segs[3]
		http.SetCookie(w, &http.Cookie{
			Name:  "user_type",
			Value: UserTypeCookieValue,
			Path:  "/"})
	}





	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
			return
		}
		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s: %s", provider, err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
			return
		}
		// get the credentials
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s: %s", provider, err), http.StatusInternalServerError)
			return
		}
		// get the user
		user, err := provider.GetUser(creds)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get user from %s: %s", provider, err), http.StatusInternalServerError)
			return
		}

		cookie, err := r.Cookie("user_type")
		if err != nil {
			log.Fatalln("Cookie:",err)
		}
		user_type := cookie.Value

		if user_type == "viewer" {
			userType := models.UserVw{}
		
			err = models.DB.Where("email = ?", user.Email()).First(&userType).Error
			if err != nil {
				userType.UUID = models.CreateUUID().String()
				userType.Name = user.Name()
				userType.Email = user.Email()
				userType.IsValid = 0
				userType.IsOauth = 1
				
				userType.CreateUser()
				models.DB.Where("email = ?", user.Email()).First(&userType)
				sess := models.Session{
					UUID: userType.UUID,
					Name: userType.Name,
					Email: userType.Email,
					UserId: userType.ID,
					UserType: user_type,
				}
				err = models.DB.Create(&sess).Error
				if err != nil {
					log.Fatalln(err)
				}
				authCookieValue := userType.UUID
				http.SetCookie(w, &http.Cookie{
					Name:  "_cookie",
					Value: authCookieValue,
					Path:  "/"})
			} else {
				sess := models.Session{
					UUID: userType.UUID,
					Name: userType.Name,
					Email: userType.Email,
					UserId: userType.ID,
					UserType: user_type,
				}
				err = models.DB.Create(&sess).Error
				if err != nil {
					log.Fatalln(err)
				}
				authCookieValue := userType.UUID
				http.SetCookie(w, &http.Cookie{
					Name:  "_cookie",
					Value: authCookieValue,
					Path:  "/"})
			}

		} else if user_type == "expert" {
			userType := models.UserEx{}
			err = models.DB.Where("email = ?", user.Email()).First(&userType).Error
			if err != nil {

				userType.UUID = models.CreateUUID().String()
				userType.Name = user.Name()
				userType.Email = user.Email()
				userType.IsValid = 0
				userType.IsOauth = 1
				// ユーザー新規作成作成
				log.Println(userType)
				userType.CreateUser()
				// ユーザー情報取得
				models.DB.Where("email = ?", user.Email()).First(&userType)

				log.Println(&userType)
				sess := models.Session{
					UUID: userType.UUID,
					Name: userType.Name,
					Email: userType.Email,
					UserId: userType.ID,
					UserType: user_type,
				}
				err = models.DB.Create(&sess).Error
				if err != nil {
					log.Fatalln(err)
				}
				authCookieValue := userType.UUID
				http.SetCookie(w, &http.Cookie{
					Name:  "_cookie",
					Value: authCookieValue,
					Path:  "/"})
			} else {
				sess := models.Session{
					UUID: userType.UUID,
					Name: userType.Name,
					Email: userType.Email,
					UserId: userType.ID,
					UserType: user_type,
				}
				err = models.DB.Create(&sess).Error
				if err != nil {
					log.Fatalln(err)
				}
				authCookieValue := userType.UUID
				http.SetCookie(w, &http.Cookie{
					Name:  "_cookie",
					Value: authCookieValue,
					Path:  "/"})
			}
		}

		w.Header().Set("Location", "/articles")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}


