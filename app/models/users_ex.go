package models

import (
	"log"
	"time"
)

type UserEx struct {
	ID        int `gorm:"primaryKey"`
	UUID      string
	Name      string
	Email     string
	Password  string
	IsValid   int
	IsOauth   int
	CreatedAt time.Time
}

// UserExテーブル名を設定する
func (UserEx) TableName() string {
	return "ex_users"
}

//エキスパート登録
func (u *UserEx) CreateUser() (err error) {
	err = DB.Create(&u).Error
	if err != nil {
		log.Fatalln(err.Error())
	}
	return err
}

func GetUserByEmailEx(email string, s string) (user UserEx, err error) {
	user = UserEx{}
	err = DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Println(err)
	}
	return user, err
}

func (u *UserEx) CreateSession(s string) (session Session, err error) {
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

func (sess *Session) GetUserBySessionEx() (user UserEx, err error) {
	user = UserEx{}
	err = DB.Where("id = ?", sess.UserId).First(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}
