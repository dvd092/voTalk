package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	// "votalk/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"github.com/joho/godotenv"
	"os"
)

var DB *gorm.DB

var err error

const (
	tableNameExUser  = "ex_users"
	tableNameVwUser  = "vw_users"
	tableNameSession = "sessions"
)

func init() {
	EnvLoad()
	connection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
	DB, err = gorm.Open("mysql", connection)
	if err != nil {
		log.Fatalln(err)
	}

}

func CreateUUID() (uuidobj uuid.UUID) {
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
