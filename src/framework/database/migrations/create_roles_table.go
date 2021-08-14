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

func (c *createRoleTable) Up() (err error) {
	if err = c.DB.Migrator().AutoMigrate(&models.Role{}); err != nil {
		return
	}

	if err = c.DB.Migrator().AutoMigrate(&models.Permission{}); err != nil {
		return
	}
	return
}

func (c *createRoleTable) Down() error {
	if c.DB.Migrator().HasTable("roles") {
		c.DB.Migrator().DropTable("user_permissions")
		c.DB.Migrator().DropTable("user_roles")
		c.DB.Migrator().DropTable("role_permissions")
		c.DB.Migrator().DropTable(&models.Permission{})
		return c.DB.Migrator().DropTable(&models.Role{})
	}
	return nil
}
