package repositories

import (
	"github.com/mixnote/mixnote-api-go/src/framework/database"
)

var (
	db, _ = database.DBConnection(nil)
)

