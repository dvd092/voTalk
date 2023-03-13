package models

import (
	"log"
	"time"
	// "fmt"
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
	cmd := `select id, uuid, name, email, password, created_at from ex_users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

func (u *UserEx) CreateSession(s string) (session Session, err error) {
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

func (sess *Session) GetUserBySessionEx() (user UserEx, err error) {
	user = UserEx{}
	err = DB.Where("id = ?", sess.UserId).First(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}
