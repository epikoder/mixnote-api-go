package migrate

import (
	"os"
	"strconv"

	"github.com/mixnote/mixnote-api-go/cmd/helpers"
	"github.com/mixnote/mixnote-api-go/src/framework/database"
	"github.com/mixnote/mixnote-api-go/src/framework/database/migrations"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
	"gorm.io/gorm"
)

var migrations_ map[string]migrations.IMigrations
var cli = helpers.Cli()

var bReset bool = cli.Switch("-F")
var conn string = cli.Option("conn")
var bRollback bool = cli.Switch("-R")
var bCreate bool = cli.Switch("-C")

var model string = cli.Option("model")
var db *gorm.DB

func Migrate() {
	//Initialize
	db, _ = database.DBConnection(conn)
	migrations_ = map[string]migrations.IMigrations{
		"users": migrations.CreateUserTable(db),
		"roles": migrations.CreateRoleTable(db),
		"artists": migrations.CreateArtistTable(db),
		"albums": migrations.CreateAlbumTable(db),
		"songs": migrations.CreateSongTable(db),
	}
	///
	
	if model != "" {
		m_ := migrations_[model]
		if m_ == nil {
			utilities.Console.Fatal("Model not found")
		}

		if bRollback && db.Migrator().HasTable(model) {
			utilities.Console.Warn("Rollback:: Table " + model)
			if err := m_.Down(); err != nil {
				utilities.Console.Error("Unable to drop table for model " + model)
				utilities.Console.Fatal(err)
			} 
			utilities.Console.Success("Table " + model + " status: " + strconv.FormatBool(db.Migrator().HasTable(model)))
		}		

		if db.Migrator().HasTable(model) {
			utilities.Console.Warn("Nothing to migrate. Use '--model=%s -R -C' -R:rollback -C:create to reset table", model)
			os.Exit(0)
		}

		if bCreate && !db.Migrator().HasTable(model) || !bCreate && !bRollback {
			utilities.Console.Warn("Migrating:: Table " + model)
			if err := m_.Up(); err != nil {
				utilities.Console.Error("Unable to migrate table for model " + model)
				utilities.Console.Fatal(err)
			}
			utilities.Console.Success("Migrated:: Table " + model)
		}
		os.Exit(0)
	}

	if bReset {
		dropTables()
		createTables()
		os.Exit(0)
	}

	if bRollback {
		dropTables()
		os.Exit(0)
	}

	createTables()
}

func dropTables() {
	for _, m_ := range migrations_ {
		if err := m_.Down(); err != nil {
			utilities.Console.Fatal(err)
		}
	}
	utilities.Console.Log("Dropped all tables successfully")
}

func createTables() {
	mi := 0
	for key, m_ := range migrations_ {
		if db.Migrator().HasTable(key) {
			utilities.Console.Log("Skipping %s", key)
			continue
		}
		utilities.Console.Warn("Migrating:: table for model " + key)
		if err := m_.Up(); err != nil {
			utilities.Console.Fatal(err)
		}
		utilities.Console.Success("Migrated:: table for model " + key)
		mi++
	}
	if mi == 0 {
		utilities.Console.Warn("Nothing to migrate")
	}
}
