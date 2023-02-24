package models

import (
	// "encoding/json"
	// "fmt"
	"log"
	"time"
	// "net/http"
	// gorm mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Article struct {
	ID int
	ExId string
	CategoryId int
	Title string
	Plot string
	Likes int
	CreatedAt	time.Time
}

// func (u *User) CreateTodo(content string) (err error) {
// 	cmd := `insert into todos (content, user_id, created_at) values(?, ?, ?)`

// 	_, err = Db.Exec(cmd, content, u.ID, time.Now())
// 	if err!= nil {
//     log.Fatalln(err) 
//   }
// 	return err
// }

// func GetTodo(id int) (art Article, err error) {
// 	cmd := `select id, content, user_id, created_at from articles where id =?`
// 	art = Article{}

// 	err = Db.QueryRow(cmd, id).Scan(&art.ID, &art.Content, &art.UserID, &art.CreatedAt)
// 	return art ,err
// }

func GetArticles() (arts []Article, err error) {
	cmd := `select id, ex_id, categories_id, title, plot, likes, created_at from articles`
	rows, err := Db.Query(cmd)
		if err!= nil {
      log.Fatalln(err)
    }
		for rows.Next() {
			var art Article
			err = rows.Scan(&art.ID, &art.ExId, &art.CategoryId, &art.Title, &art.Plot, &art.Likes, &art.CreatedAt)
		if err!= nil {
      log.Fatalln(err)
		}
		arts = append(arts,art)
	}
	rows.Close()
	return arts, err
}

func GetArticle(id int)  (gorm *gorm.DB) {
	result := db.Where("id = ?", id).Find(&Article{})
	return result
}

/*
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id =?`
  rows, err := Db.Query(cmd, u.ID)
  if err!= nil {
      log.Fatalln(err)
	}
	for rows.Next() {
    var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err!= nil {
      log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
  return todos, err
}

func (t *Todo)UpdateTodo() (err error) {
	cmd := `update todos set content =?, user_id =? where id =?`
  _, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err!= nil {
      log.Fatalln(err)
  }
  return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todos where id =?`
  _, err = Db.Exec(cmd, t.ID)
  if err!= nil {
      log.Fatalln(err)
  }
	return err
}
*/