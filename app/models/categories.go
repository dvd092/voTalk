package models

import "log"

type Category struct {
	ID   int
	Name string
}

// 基本データ追加
func InsertCategory() error {

	var categories = []Category{}
	categories = []Category{
		{ID: 4, Name: "スポーツ"},
		{ID: 5, Name: "健康"},
		{ID: 6, Name: "サイエンス"},
		{ID: 7, Name: "政治"},
		{ID: 8, Name: "自然環境"},
		{ID: 9, Name: "社会"},
		{ID: 10, Name: "アート"},
		{ID: 11, Name: "音楽"},
		{ID: 12, Name: "映画"},
		{ID: 13, Name: "テレビ番組"},
		{ID: 14, Name: "ゲーム"},
		{ID: 15, Name: "スポーツニュース"},
		{ID: 16, Name: "スポーツイベント"},
		{ID: 17, Name: "ダイエット"},
		{ID: 18, Name: "病気"},
		{ID: 19, Name: "医療"},
		{ID: 20, Name: "宇宙"},
		{ID: 21, Name: "物理学"},
		{ID: 22, Name: "化学"},
		{ID: 23, Name: "生物学"},
		{ID: 24, Name: "歴史"},
		{ID: 25, Name: "国際情勢"},
		{ID: 26, Name: "国内情勢"},
		{ID: 27, Name: "法律"},
		{ID: 28, Name: "キャリア"},
		{ID: 29, Name: "教育"},
		{ID: 30, Name: "宗教"},
	}
	log.Println(&categories)
	for _, v := range categories {
		DB.Table("categories").Create(v)
	}
	return nil
}

// 全データ取得
func AllCategories() []Category {
	categories := []Category{}
	err := DB.Find(&categories).Error
	if err != nil {
		log.Printf(err.Error())
	}
	return categories
}
