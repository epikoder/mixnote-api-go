package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/mixnote/mixnote-api-go/src/music/model"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uuid.UUID `sql:"primary_key;type:uuid;default:uuid_generate_v4()"`
		FirstName string    `gorm:"type:char(50)"`
		LastName  string    `gorm:"type:char(50)"`
		Email     string    `gorm:"type:char(150)"`
		Phone     int       `gorm:"type:int(18)"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`

		UserBillingInformation UserBillingInformation
		Artist                 music_models.Artist

		Roles       []Role       `gorm:"many2many:user_roles"`
		Permissions []Permission `gorm:"many2many:user_permissions"`
	}

	UserBillingInformation struct {
		ID           uuid.UUID `sql:"primary_key;type:uuid;default:uuid_generate_v4()"`
		UserID       uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()"`
		StripeID     string    `gorm:"type:char(50)"`
		PaystackID   string    `gorm:"type:char(50)"`
		BillingEmail string    `gorm:"type:char(150);not_null;"`
		CountryCode  string    `gorm:"type:char(3)"`
		PhoneNumber  int       `gorm:"type:int(15)"`
		State        string    `gorm:"type:char(50)"`
		ZipCode      string    `gorm:"type:int(8)"`
		AddressLine1 string    `gorm:"type:char(255)"`
		AddressLine2 string    `gorm:"type:char(255)"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
		DeletedAt    gorm.DeletedAt `gorm:"index"`
	}
)

func NewUser() (u *User) {
	return
}
