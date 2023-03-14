package models

import (
	"log"
	"time"
)

type UserVw struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	LikeNum   int
	IsValid   int
	IsOauth   int
	CreatedAt time.Time
}

type Session struct {
	Id        int
	UUID      string
	Name      string
	Email     string
	UserId    int
	UserType  string
	CreatedAt time.Time
}

func (UserVw) TableName() string {
	return "vw_users"
}

//ビューワー登録
func (u *UserVw) CreateUser() (err error) {
	err = DB.Create(&u).Error
	if err != nil {
		log.Fatalln(err.Error())
	}
	return err
}

func GetUserByEmailVw(email string, s string) (user UserVw, err error) {
	user = UserVw{}
	err = DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}

func (u *UserVw) CreateSession(s string) (session Session, err error) {
	// セッション作成
	session = Session{
		UUID: CreateUUID().String(),
		Email: u.Email,
		UserId: u.ID,
		UserType: s,
		CreatedAt: time.Now(),
	}
	err = DB.Create(&session).Error
		if err != nil {
		log.Println(err)
	}
	// セッション取得
	err = DB.Where("email = ?", u.Email).First(&session).Error

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	err = DB.Where("uuid = ?", sess.UUID).Find(&sess).Error
	if err != nil {
		valid = false
		return
	}

	if sess.Id != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	err = DB.Where("uuid = ?", sess.UUID).Unscoped().Delete(&sess).Error
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (sess *Session) GetUserBySessionVw() (user UserVw, err error) {
	user = UserVw{}
	err = DB.Where("id = ?", sess.UserId).First(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}
