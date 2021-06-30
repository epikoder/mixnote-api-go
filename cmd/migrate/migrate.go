package migrate

import (
	"github.com/mixnote/mixnote-api-go/src/framework/database"
	"github.com/mixnote/mixnote-api-go/src/framework/database/migrations"
	"log"
)

func Migrate(reset bool) {
	db, err := database.DBConnection("")
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	migrations.CreateUserTable().Up(db)

	db.Migrator().DropTable("user")

	// db.AutoMigrate(&album.Album{})
}
