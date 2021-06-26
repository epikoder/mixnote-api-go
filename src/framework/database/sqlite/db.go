package sqlite3

import (
	"os"

	app "github.com/mixnote/mixnote-api-go/src"
	"github.com/mixnote/mixnote-api-go/src/framework/database/contracts"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqLite3 struct {
	contracts.IDatabase
	DATABASE_PATH string
}

func New() (sqlite *sqLite3) {
	sqlite = &sqLite3{
		DATABASE_PATH: os.Getenv("DATABASE_PATH"),
	}
	if sqlite.DATABASE_PATH == "" {
		sqlite.DATABASE_PATH = app.DatabasePath()
	}

	return
}

func (sq *sqLite3) ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(sq.DATABASE_PATH))
}
