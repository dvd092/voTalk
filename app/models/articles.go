package models

import (
	// "encoding/json"
	// "fmt"
	// "log"
	"time"
	"html/template"
	// "net/http"
	// gorm mysql
)

type Article struct {
	ID int
	UserExID int
	CategoryId int
	Title string
	Plot template.HTML
	Likes int
	CreatedAt	time.Time

	// リレーション
	ExUser UserEx `gorm:"foreignKey:UserExID"`
}


func GetArticles() (arts []Article, err error) {
	DB.Order("likes desc").Find(&arts)
	return arts, err
}

func GetArticle(id int)  (art Article, err error) {
	err = DB.Where("id = ?", id).First(&art).Error
	if err != nil {
		return Article{},err
	}

	art.Plot = template.HTML(art.Plot)
	return art,nil
}

func GetArticlesByUser(userId int) (arts []Article, err error) {
	err = DB.Where("user_ex_id = ?", userId).Find(&arts).Error
	if err != nil {
		return []Article{},err
	}
	return arts,nil
}