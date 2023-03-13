package models

import (
	"log"
	"time"
	// "fmt"
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
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		user_type,
		created_at) values (?,?,?,?,?)`

	_, err = Db.Exec(cmd1, CreateUUID(), u.Email, u.ID, s, time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select id, uuid, email, user_id, user_type, created_at from sessions where user_id=? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.Id,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.UserType,
		&session.CreatedAt)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, user_type, created_at from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.Id,
		&sess.UUID,
		&sess.Email,
		&sess.UserId,
		&sess.UserType,
		&sess.CreatedAt)

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
	cmd := `delete from sessions where uuid =?`
	_, err = Db.Exec(cmd, sess.UUID)
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
