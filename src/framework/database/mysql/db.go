package mysql

import (
	"fmt"
	"github.com/mixnote/mixnote-api-go/src/framework/database/contracts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type mySql struct {
	contracts.IDatabase
	DATABASE_URL string
	DB_HOST      string
	DB_PORT      int
	DB_USERNAME  string
	DB_PASSWORD  string
	DB_NAME      string
}

func New() (mysql *mySql) {
	mysql = &mySql{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_USERNAME:  os.Getenv("DB_USERNAME"),
		DB_PASSWORD:  os.Getenv("DB_PASSWORD"),
		DB_NAME:      os.Getenv("DB_NAME"),
	}

	port, err := (func() (int, error) {
		port := os.Getenv("DB_PORT")
		if port != "" {
			return strconv.Atoi(port)
		}
		return 3306, nil
	})()

	if err != nil {
		panic("Invalid database port")
	}
	mysql.DB_PORT = port
	return
}

func (m *mySql) ConnectDB() (*gorm.DB, error) {
	dns := (func() string {
		if m.DATABASE_URL != "" {
			return m.DATABASE_URL
		}
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", m.DB_USERNAME, m.DB_PASSWORD, m.DB_HOST, m.DB_NAME)
	})()
	return gorm.Open(mysql.Open(dns))
}
