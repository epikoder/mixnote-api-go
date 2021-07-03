package migrations

import (
	"gorm.io/gorm"
	"github.com/mixnote/mixnote-api-go/src/core/models"
)

type createUserTable struct {
	DB *gorm.DB
}

func CreateUserTable(db *gorm.DB) (c *createUserTable) {
	c = &createUserTable{
		DB: db,
	}
	return
}

func (c *createUserTable) Up() error {
	return c.DB.Migrator().CreateTable(&models.User{})
}

func (c *createUserTable) Down() error {
	if c.DB.Migrator().HasTable("users") {
		return c.DB.Migrator().DropTable(&models.User{})
	}
	return nil
}
