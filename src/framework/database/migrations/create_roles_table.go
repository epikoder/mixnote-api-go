package migrations

import (
	"github.com/mixnote/mixnote-api-go/src/core/models"
	"gorm.io/gorm"
)


type createRoleTable struct {
	DB *gorm.DB
}

func CreateRoleTable(db *gorm.DB) (c *createRoleTable) {
	c = &createRoleTable{
		DB: db,
	}
	return
}

func (c *createRoleTable) Up() error {
	return c.DB.Migrator().CreateTable(&models.Role{})
}

func (c *createRoleTable) Down() error {
	if c.DB.Migrator().HasTable("roles") {
		return c.DB.Migrator().DropTable(&models.Role{})
	}
	return nil
}