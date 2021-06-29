package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model

		ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
		FirstName string `json:"first_name"`
		LastName string
		Email string
		Phone int
	}
)

func New_User() (u *User) {
	u = &User{}
	return
}

