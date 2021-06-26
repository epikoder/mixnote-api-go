package main

import (
	"fmt"
	"log"

	app "github.com/mixnote/mixnote-api-go/src"
	"github.com/mixnote/mixnote-api-go/src/core/models/user"
	"github.com/mixnote/mixnote-api-go/src/framework/database"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	app.RegisterBindings()
	DB, err = database.DBConnection("")
	if err != nil {
		log.Fatalln("Error connecting to database")
	}
}

func main() {
	user := user.New()
	fmt.Println(&user)
}