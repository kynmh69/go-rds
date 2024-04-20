package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:varchar(255);default:(UUID());primaryKey;not null"`
	Username  string    `json:"username" gorm:"not null;type:varchar(255)"`
	Password  string    `json:"password" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime; not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime; not null"`
}
