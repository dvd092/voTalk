package models

import (
	"html/template"
	"time"
)

type Article struct {
	ID         int
	UserExID   int
	CategoryId int
	Title      string
	Plot       template.HTML
	Likes      int
	CreatedAt  time.Time

	// リレーション
	ExUser   UserEx   `gorm:"foreignKey:UserExID"`
	Category Category `gorm:"foreignKey:CategoryId"`
}

func GetArticles() (arts []Article, err error) {
	DB.Order("likes desc").Find(&arts)
	return arts, err
}

func GetArticle(id int) (art Article, err error) {
	err = DB.Where("id = ?", id).First(&art).Error
	if err != nil {
		return Article{}, err
	}

	art.Plot = template.HTML(art.Plot)
	return art, nil
}

func GetArticlesByUser(userId int) (arts []Article, err error) {
	err = DB.Where("user_ex_id = ?", userId).Preload("Category").Find(&arts).Error
	if err != nil {
		return []Article{}, err
	}
	return arts, nil
}

func GetCategoryArticles(id int) (arts []Article, err error) {
	err = DB.Where("category_id = ?", id).Find(&arts).Error
	if err != nil {
		return []Article{}, err
	}
	return arts, nil
}
