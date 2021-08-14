package repositories

import (
	"github.com/google/uuid"
	"github.com/mixnote/mixnote-api-go/src/core/models"
)

type userRepository struct {
	User *models.User
}

func UserRepository(user *models.User) (u *userRepository) {
	user.ID = uuid.Nil
	u = &userRepository{
		User: user,
	}
	return u
}

func (u *userRepository) All() {}
func (u *userRepository) FindByID(id string) {
	db.Where("id", id).First(u.User)
}
func (u *userRepository) FindByEmail(email string) (ok bool) {
	db.Preload("UserCredential").Preload("UserBillingInformation").First(u.User, "email = ?", email)
	return u.User.ID != uuid.Nil
}
func (u *userRepository) CreateUser(data *models.User) (ok bool, err error) {
	db.Create(data)
	u.FindByEmail(data.Email)
	return
}
