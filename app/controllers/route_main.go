package controllers

import (
	"fmt"
	"log"
	"net/http"
	"votalk/app/libs"
)

func top(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, fmt.Sprintf("/%s/index", s.UserType), 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
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
				S    string
			}{
				user,
				s,
			}
			generateHTML(w, data, "layout", "private_navbar", "index")
		} else if s == "expert" {
			user, err := sess.GetUserBySessionEx()
			if err != nil {
				log.Println(err)
			}
			data := struct {
				User interface{}
				S    string
			}{
				user,
				s,
			}
			generateHTML(w, data, "layout", "private_navbar", "index")
		}
		// todos, _ := user.GetTodosByUser()
		// user.Todos = todos

	}
}

/*

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}


func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
  if err!= nil {
    http.Redirect(w,r, "/login",302)
	} else {
		err = r.ParseForm()
		if err!= nil {
      log.Println(err)
	}
	  user, err := sess.GetUserBySession()
    if err!= nil {
      log.Println(err)
    }
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w,r,"/todos",302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
  if err!= nil {
    http.Redirect(w,r, "/login",302)
	} else {
		_, err := sess.GetUserBySession()
		if err!= nil {
      log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err!= nil {
      log.Println(err)
		}
		generateHTML(w,t,"layout","private_navbar","todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err!= nil {
    http.Redirect(w,r, "/login",302)
	} else {
		err = r.ParseForm()
		if err!= nil {
      log.Println(err)
	}
	  user, err := sess.GetUserBySession()
    if err!= nil {
			log.Println(err)
		}
	content := r.PostFormValue("content")
	t:= &models.Todo{ID :id,Content: content,UserID: user.ID}
	if err := t.UpdateTodo(); err != nil {
		log.Println(err)
	}
	http.Redirect(w,r,"/todos",302)
}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int)  {
	sess, err := session(w, r)
	if err!= nil {
    http.Redirect(w,r, "/login",302)
	} else {
		_, err := sess.GetUserBySession()
		if err!= nil {
      log.Println(err)
    }
		t, err := models.GetTodo(id)
		if err!= nil {
      log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w,r,"/todos",302)
	}
}
*/
