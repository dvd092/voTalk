package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"votalk/config"
	// gorm mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"github.com/joho/godotenv"
	"os"
)

var Db *sql.DB
var DB *gorm.DB

var err error

const (
	tableNameExUser  = "ex_users"
	tableNameVwUser  = "vw_users"
	tableNameSession = "sessions"
)

func init() {
	EnvLoad()
	connection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	DB, err = gorm.Open("mysql", connection)
	if err != nil {
		log.Fatalln(err)
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("接続失敗")
	} else {
		fmt.Println("接続成功")
	}

	//viewユーザーテーブル作成
	cmdVwU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTO_INCREMENT,uuid VARCHAR(100),name VARCHAR(100),email VARCHAR(100),password VARCHAR(100),created_at DATETIME)`, tableNameVwUser)

	_, err = Db.Exec(cmdVwU)
	if err != nil {
		log.Fatalln(err)
	}
	Db.Exec(cmdVwU)

	//expertユーザーテーブル作成
	cmdExU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTO_INCREMENT,uuid VARCHAR(100),name VARCHAR(100),email VARCHAR(100),password VARCHAR(100),created_at DATETIME)`, tableNameExUser)

	_, err = Db.Exec(cmdExU)
	if err != nil {
		log.Fatalln(err)
	}
	Db.Exec(cmdExU)

	// cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTO_INCREMENT,content TEXT, user_id INTEGER, created_at DATETIME)`, tableNameTodo)

	// Db.Exec(cmdT)

	//セッションテーブル作成
	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255),
		user_id INTEGER,
		user_type INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)

}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

// envLoad 環境変数のロード
func EnvLoad() {
	err := godotenv.Load("aws.env")
	if err != nil {
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatalln(err)
		}
	}
}