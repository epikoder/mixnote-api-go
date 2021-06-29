package migrations

import (
	// "github.com/mixnote/mixnote-api-go/src/core/models"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/migrator"
)

type createUserTable struct {

}

func CreateUserTable() (c *createUserTable){
	return
}

func (createUserTable) Up(db *gorm.DB) {
	// m := db.Migrator();// m.CreateTable(&models.User{})
	// m.AddColumn("", "name")
	fmt.Println(db)
}

func (createUserTable) Down(m *migrator.Migrator) {

}