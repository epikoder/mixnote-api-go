package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Role struct {
		ID        uuid.UUID `sql:"primary_key;type:uuid;default:uuid_generate_v4()"`
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt

		Permissions []Permission	`gorm:"many2many:role_permissions"`
	}

	Permission struct {
		ID        uuid.UUID `sql:"primary_key;type:uuid;default:uuid_generate_v4()"`
		Name      string    `gorm:"type:char(50)"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
	}
)
