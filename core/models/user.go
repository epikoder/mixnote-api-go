package user

import (
	"github.com/mixnote/mixnote-api-go/core/models/contracts"
	"gorm.io/gorm"
)

type (
	User struct {
		contracts.IUser
		gorm.Model
		FName string
		LName string
		Email string
		Phone int
	}
)

func New() (u *User) {
	u = &User{
	}
	return
}
