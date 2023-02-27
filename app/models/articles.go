package models

import (
	// "encoding/json"
	// "fmt"
	"log"
	"time"
	// "net/http"
	// gorm mysql
)

type Article struct {
	ID int
	UserExID int
	CategoryId int
	Title string
	Plot string
	Likes int
	CreatedAt	time.Time

	// リレーション
	ExUser UserEx `gorm:"foreignKey:UserExID"`
}


func GetArticles() (arts []Article, err error) {
	cmd := `select id, user_ex_id, categories_id, title, plot, likes, created_at from articles`
	rows, err := Db.Query(cmd)
		if err!= nil {
      log.Fatalln(err)
    }
		for rows.Next() {
			var art Article
			err = rows.Scan(&art.ID, &art.UserExID, &art.CategoryId, &art.Title, &art.Plot, &art.Likes, &art.CreatedAt)
		if err!= nil {
      log.Fatalln(err)
		}
		arts = append(arts,art)
	}
	rows.Close()
	return arts, err
}

func GetArticle(id int)  (art Article, err error) {
	err = DB.Where("id = ?", id).First(&art).Error
	if err != nil {
		return Article{},err
	}
	return art,nil
}

func GetArticlesByUser(userId int) (arts []Article, err error) {
	err = DB.Where("user_ex_id = ?", userId).Find(&arts).Error
	if err != nil {
		return []Article{},err
	}
	return arts,nil
}