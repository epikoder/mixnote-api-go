package database

import (
	"os"

	"github.com/mixnote/mixnote-api-go/src/framework/database/mysql"
	"github.com/mixnote/mixnote-api-go/src/framework/database/sqlite"
	"gorm.io/gorm"
)


func DBConnection(connection string) (db *gorm.DB, err error) {
	if connection == "" {
		connection = os.Getenv("DB_CONNECTION")
	}

	switch connection {
	case "mysql":
		db, err = mysql.New().ConnectDB()
	default:
		db, err = sqlite3.New().ConnectDB()
	}

	return
}
