package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model

		ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
		FName string
		LName string
		Email string
		Phone int
	}
)

func New() (u *User) {
	u = &User{}
	return
}

