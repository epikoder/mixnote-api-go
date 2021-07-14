package merchants

import "github.com/mixnote/mixnote-api-go/src/core/models"

type IMerchant interface {
	Name() string
	Identifier() string
	Customer(*models.User) (interface{},error)
}