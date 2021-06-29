package database

import (
	"os"
	"gorm.io/gorm"
	"github.com/mixnote/mixnote-api-go/src/framework/database/connections"
)


func DBConnection(connection string) (db *gorm.DB, err error) {
	if connection == "" {
		connection = os.Getenv("DB_CONNECTION")
	}

	switch connection {
	case "mysql":
		db, err = connections.MySql().ConnectDB()
	default:
		db, err = connections.SqLite3().ConnectDB()
	}

	if err != nil {
		panic("Error connecting to database")
	}
	return
}
