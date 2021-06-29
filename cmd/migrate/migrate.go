package migrate

import (
	"fmt"
	"log"
	"github.com/mixnote/mixnote-api-go/src/framework/database"
	"github.com/mixnote/mixnote-api-go/src/framework/database/migrations"
)


func Migrate(reset bool) {
	db, err := database.DBConnection("")	
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	migrations.CreateUserTable().Up(db)
	
	db.Migrator().DropTable("user")
	fmt.Println("Migrating user table")
	fmt.Println("Migrated user table")
	
	
	// db.AutoMigrate(&album.Album{})
}
