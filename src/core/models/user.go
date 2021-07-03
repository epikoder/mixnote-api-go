package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/mixnote/mixnote-api-go/src/music/model"
)

type (
	User struct {
		ID		uuid.UUID `sql:"primary_key;type:uuid;default:uuid_generate_v4()"`
		FirstName string    `gorm:"type:char(50)"`
		LastName  string	`gorm:"type:char(50)"`
		Email     string	`gorm:"type:char(150)"`
		Phone     int		`gorm:"type:int(18)"`
		CreatedAt	time.Time
		UpdatedAt	time.Time
		DeletedAt	gorm.DeletedAt	`gorm:"index"`

		Artist    music_models.Artist
	}
)

func (u *User) BeforeCreate(*gorm.DB) error {
	return nil
}

func New_User() (u *User) {
	return
}

