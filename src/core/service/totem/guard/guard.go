package guard

import "github.com/mixnote/mixnote-api-go/src/core/models"

var (
	user *models.User
)

func SetUser(u *models.User) {
	user = u
}

func User() *models.User {
	return user
}

