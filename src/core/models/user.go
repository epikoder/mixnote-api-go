package models

import (
	"github.com/google/uuid"
	"github.com/mixnote/mixnote-api-go/src/music/model"
	"gorm.io/gorm"
	"time"
)

type (
	User struct {
		ID          uuid.UUID      `sql:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
		FirstName   string         `gorm:"type:varchar(50)" json:"first_name"`
		LastName    string         `gorm:"type:varchar(50)" json:"last_name"`
		Photo       string         `gorm:"type:varchar(255)" json:"photo"`
		Email       string         `gorm:"type:varchar(150);notnull;unique;index" json:"email" valid:"email"`
		PhoneNumber string         `gorm:"type:varchar(20)" json:"phone_number"`
		CreatedAt   time.Time      `json:"-"`
		UpdatedAt   time.Time      `json:"-"`
		DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

		UserBillingInformation UserBillingInformation
		UserCredential         UserCredential
		Artist                 music_models.Artist

		Roles       []Role       `gorm:"many2many:user_roles"`
		Permissions []Permission `gorm:"many2many:user_permissions"`
	}

	UserCredential struct {
		ID               uuid.UUID `sql:"primary_key;type:uuid;default:uuid_generate_v4()" json:"-"`
		UserID           uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()" json:"user_id"`
		Password         string    `gorm:"type:varchar(255)" json:"password"`
		TwoFactorEnabled int       `gorm:"type:varchar(1);default:0" json:"two_factor_enabled"`
		CreatedAt        time.Time `json:"-"`
		UpdatedAt        time.Time `json:"-"`
	}

	UserBillingInformation struct {
		ID           uuid.UUID `sql:"primary_key;type:uuid;default:uuid_generate_v4()" json:"-"`
		UserID       uuid.UUID `sql:"type:uuid;default:uuid_generate_v4()" json:"user_id"`
		StripeID     string    `gorm:"type:varchar(50)" json:"stripe_id"`
		PaystackID   string    `gorm:"type:varchar(50)" json:"paystack_id"`
		BillingEmail string    `gorm:"type:varchar(150);notnull;index" json:"billing_email"`
		CountryCode  int       `gorm:"type:varchar(3);" json:"country_code"`
		PhoneNumber  string    `gorm:"type:varchar(20)" json:"phone_number"`
		State        string    `gorm:"type:varchar(50)" json:"state"`
		ZipCode      int       `gorm:"type:varchar(3)" json:"zip_code"`
		AddressLine1 string    `gorm:"type:varchar(255)" json:"address_line_1"`
		AddressLine2 string    `gorm:"type:varchar(255)" json:"address_line_2"`
		CreatedAt    time.Time `json:"-"`
		UpdatedAt    time.Time `json:"-"`
	}
)

