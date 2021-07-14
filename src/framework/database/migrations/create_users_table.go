package migrations

import (
	"github.com/mixnote/mixnote-api-go/src/core/models"
	"gorm.io/gorm"
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
	if err := c.DB.Migrator().AutoMigrate(&models.UserBillingInformation{}); err != nil {
		return err
	}
	return c.DB.Migrator().AutoMigrate(&models.User{})
}

func (c *createUserTable) Down() error {
	if c.DB.Migrator().HasTable("users") {
		if err := c.DB.Migrator().DropTable(&models.UserBillingInformation{}); err != nil {
			return err
		}
		return c.DB.Migrator().DropTable(&models.User{})
	}
	return nil
}
