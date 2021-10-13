package repositories

import (
	"github.com/google/uuid"
	"github.com/mixnote/mixnote-api-go/src/core/models"
)

type userRepository struct {
	User *models.User
}

var (
	userCredential         = "UserCredential"
	userBillingInformation = "UserBillingInformation"
	activities             = "Activities"
)

func UserRepository(user *models.User) (u *userRepository) {
	user.ID = uuid.Nil
	u = &userRepository{
		User: user,
	}
	return u
}

func (u *userRepository) All() {}
func (u *userRepository) FindByID(id string) (ok bool) {
	db.Preload(userCredential).
		Preload(userBillingInformation).
		Preload(activities).
		First(u.User, "id = ?", id)
	return u.User.ID != uuid.Nil
}

func (u *userRepository) FindByEmail(email string) (ok bool) {
	db.Preload(userCredential).
		Preload(userBillingInformation).
		Preload(activities).
		First(u.User, "email = ?", email)
	return u.User.ID != uuid.Nil
}

func (u *userRepository) CreateUser(data *models.User) (ok bool) {
	db.Create(data)
	return u.FindByEmail(data.Email)
}
