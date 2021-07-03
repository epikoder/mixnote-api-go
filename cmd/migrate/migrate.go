package migrate

import (
	"os"
	"github.com/mixnote/mixnote-api-go/cmd/helper"
	"github.com/mixnote/mixnote-api-go/src/framework/database"
	"github.com/mixnote/mixnote-api-go/src/framework/database/migrations"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
)

var migrations_ map[string]migrations.IMigrations
var argHelper helper.ArgHelper = helper.New()

var reset bool = argHelper.Switch("-F")
var conn string = argHelper.Option("conn")
var model string = argHelper.Option("model")

func init() {
	db, _ := database.DBConnection(conn)
	migrations_ = map[string]migrations.IMigrations{
		"user": migrations.CreateUserTable(db),
		"role": migrations.CreateRoleTable(db),
	}	
}

func Migrate() {
	if reset {
		resetDB()
		os.Exit(0)
	}

	if model != "" {
		m_ := migrations_[model]
		if m_ == nil {
			utilities.Console().Fatal("Model not found")
		}

		err := m_.Down()
		if err != nil {
			utilities.Console().Error("Unable to drop table for model " + model)
			utilities.Console().Fatal(err)
		}
		
		utilities.Console().Warn("Migrating table for model " + model)
		err = m_.Up()
		if err != nil {
			utilities.Console().Error("Unable to migrate table for model " + model)
			utilities.Console().Fatal(err)
		}
		utilities.Console().Success("Migrated table for model " + model)
		os.Exit(0)
	}
}

func resetDB() {
	for key, m_ := range migrations_ {
		err := m_.Down()
		if err != nil {
			utilities.Console().Error("Unable to drop table for model " + key)
			utilities.Console().Fatal(err)
		}
	}
	utilities.Console().Warn("Dropped All tables")


	for key, m_ := range migrations_ {
		utilities.Console().Warn("Migrating table for model " + key)
		err := m_.Up()
		if err != nil {
			utilities.Console().Error("Unable to migrate table for model " + key)
			utilities.Console().Fatal(err)
		}
		utilities.Console().Success("Migrated table for model " + key)
	}
}
