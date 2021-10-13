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
	if err := c.DB.Migrator().AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := c.DB.Migrator().AutoMigrate(&models.UserCredential{}); err != nil {
		return err
	}

	if err := c.DB.Migrator().AutoMigrate(&models.UserActivities{}); err != nil {
		return err
	}
	if err := c.DB.Migrator().AutoMigrate(&models.AccessLog{}); err != nil {
		return err
	}
	if err := c.DB.Migrator().AutoMigrate(&models.Transaction{}); err != nil {
		return err
	}
	return c.DB.Migrator().AutoMigrate(&models.UserBillingInformation{})
}

func (c *createUserTable) Down() error {
	if c.DB.Migrator().HasTable("user_credentials") {
		if err := c.DB.Migrator().DropTable(&models.UserCredential{}); err != nil {
			return err
		}
	}

	if c.DB.Migrator().HasTable("user_billing_informations") {
		if err := c.DB.Migrator().DropTable(&models.UserBillingInformation{}); err != nil {
			return err
		}
	}

	if c.DB.Migrator().HasTable("user_activities") {
		if err := c.DB.Migrator().DropTable(&models.UserActivities{}); err != nil {
			return err
		}
	}

	if c.DB.Migrator().HasTable("users") {
		return c.DB.Migrator().DropTable(&models.User{})
	}
	return nil
}
