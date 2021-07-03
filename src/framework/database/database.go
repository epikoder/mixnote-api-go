package database

import (
	"os"

	"github.com/mixnote/mixnote-api-go/src/framework/database/connections"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
	"gorm.io/gorm"
)


func DBConnection(connection string) (db *gorm.DB, err error) {
	if connection == "" {
		connection = os.Getenv("DB_CONNECTION")
	}

	switch connection {
	case "mysql":
		db, err = connections.MySql().ConnectDB()
	case "sqlite":
		db, err = connections.SqLite3().ConnectDB()
	case "pgsql":
		db, err = connections.PgSql().ConnectDB()
	default:
		utilities.Console().Fatal("Unknown Database connection, Did you configured it?")
	}

	if err != nil {
		panic("Error connecting to database")
	}

	return
}
