package models

import (
	// "log"
	"time"
	// "fmt"
)

type Session struct {
	Id        int
	UUID      string
	Email     string
	UserId    int
	UserType  string
	CreatedAt time.Time
}

/*
//エキスパート登録
func (u *UserEx) CreateUser() (err error) {
	cmd := `insert into ex_users (
	uuid,
	name,
	email,
	password,
	created_at) values (?, ?, ?, ?, ?)`

		_, err = Db.Exec(cmd, createUUID(),u.Name,u.Email,Encrypt(u.Password),time.Now())

		if err != nil {
			log.Fatalln(err)
		}
		return err
}
//ビューワー登録
func (u *UserVw) CreateUser() (err error) {
	cmd := `insert into vw_users (
	uuid,
	name,
	email,
	password,
	created_at) values (?, ?, ?, ?, ?)`

		_, err = Db.Exec(cmd, createUUID(),u.Name,u.Email,Encrypt(u.Password),time.Now())

		if err != nil {
			log.Fatalln(err)
		}
		return err
}

// func GetUser(id int) (user UserEx, err error) {
// 	user = UserEx{}
// 	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
// 	err = Db.QueryRow(cmd, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
// 	return user,err
// }

// func (u *UserEx) UpdateUser() (err error) {
// 	cmd := `update users set name =?, email =? where id =?`
//   _, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
// 	if err!= nil {
//     log.Fatalln(err)
// 	}
// 	return err
// }

// func (u *UserEx) DeleteUser() (err error) {
// 	cmd := `delete from users where id =?`
//   _, err = Db.Exec(cmd, u.ID)
//   if err!= nil {
//     log.Fatalln(err)
//   }
// 	return err
// }

func GetUserByEmail(email string, s string) (user UserVw, err error) {
	user = UserVw{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd,email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
    &user.Email,
		&user.Password,
    &user.CreatedAt)

	return user,err
}

func (u *UserVw) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?,?,?,?)`

		_, err = Db.Exec(cmd1, createUUID(),u.Email,u.ID,time.Now())

    if err!= nil {
      log.Fatalln(err)
    }

		cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id=? and email = ?`

		err = Db.QueryRow(cmd2,u.ID,u.Email).Scan(
			&session.Id,
			&session.UUID,
			&session.Email,
			&session.UserId,
			&session.CreatedAt)

    return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`

	err = Db.QueryRow(cmd,sess.UUID).Scan(
		&sess.Id,
		&sess.UUID,
		&sess.Email,
		&sess.UserId,
		&sess.CreatedAt)

		if err != nil {
			valid = false
			return
		}

		if sess.Id != 0 {
			valid = true
		}

		return valid,err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid =?`
  _, err = Db.Exec(cmd,sess.UUID)
  if err!= nil {
    log.Fatalln(err)
  }

	return err
}

func (sess *Session) GetUserBySession() (user UserEx, err error) {
	user = UserEx{}
	cmd := `select id, uuid, name, email, created_at from users where id = ?`
	err = Db.QueryRow(cmd,sess.UserId).Scan(
		&user.ID,
    &user.UUID,
    &user.Name,
    &user.Email,
    &user.CreatedAt)

		return user, err
}
*/
