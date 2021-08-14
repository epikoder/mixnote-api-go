package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Role struct {
		ID        uuid.UUID      `sql:"primary_key;type:uuid;default:uuid_generate_v4()" json:"-"`
		Name      string         `json:"name"`
		CreatedAt time.Time      `json:"-"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `json:"-"`

		Permissions []Permission `gorm:"many2many:role_permissions"`
	}

	Permission struct {
		ID        uuid.UUID      `sql:"primary_key;type:uuid;default:uuid_generate_v4()" json:"-"`
		Name      string         `gorm:"type:char(50)" json:"name"`
		CreatedAt time.Time      `json:"-"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `json:"-"`
	}
)
